# Nome do binário gerado
APP_NAME = gerador-de-token

# Diretório principal da aplicação
SRC_DIR = .

# Variáveis do Go
GO = go
BUILD_DIR = ./bin
VERSION = 1.0.0

# Variáveis de compilação adicionais (opcional)
LDFLAGS = -X "main.Version=$(VERSION)"

# Comandos do Makefile
.PHONY: all build run clean test install

# Regra principal (compilar)
all: build

# Compilar para o arquivo binário
build:
	@echo "Compilando o programa..."
	$(GO) build -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/$(APP_NAME) $(SRC_DIR)

# Executar o programa após a compilação
run: build
	@echo "Executando o programa..."
	$(BUILD_DIR)/$(APP_NAME)

# Limpar artefatos gerados
clean:
	@echo "Limpando os arquivos compilados..."
	rm -rf $(BUILD_DIR)

# Rodar os testes
test:
	@echo "Executando testes..."
	$(GO) test ./...

# Instalar o binário no sistema (em $GOPATH/bin ou /usr/local/bin)
install: build
	@echo "Instalando o programa no sistema..."
	install -m 0755 $(BUILD_DIR)/$(APP_NAME) /usr/local/bin/$(APP_NAME)

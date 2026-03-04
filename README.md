# 🌐 Monitoramento de Sites em Go

Aplicação de linha de comando desenvolvida em Go para monitorar a disponibilidade de sites através de requisições HTTP.

## 🚀 Funcionalidades

* Monitoramento periódico de sites
* Leitura de lista de sites a partir de arquivo `.txt`
* Registro de logs com data e hora
* Exibição do histórico de monitoramentos

## 📂 Estrutura do projeto

* `main.go` → Código principal da aplicação
* `sites.txt` → Lista de sites monitorados
* `logs.txt` → Arquivo gerado com os registros de monitoramento

## ⚙️ Como configurar

Antes de executar o projeto, adicione no arquivo `sites.txt` os links que deseja monitorar.

Exemplo de conteúdo do `sites.txt`:

```
https://www.carlos0101.xyz
https://financas.carlos0101.xyz
https://birdgame.carlos0101.xyz
```

Cada site deve estar em uma linha separada.

## ▶ Como executar

No terminal, dentro da pasta do projeto:

```bash
go run main.go
```


Projeto criado para prática dos fundamentos da linguagem Go, incluindo requisições HTTP, loops, condicionais e manipulação de arquivos.

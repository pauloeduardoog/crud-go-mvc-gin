
# CRUD em Go com Gin e MVC

Este é um projeto de exemplo de uma API RESTful construída com a linguagem **Go (Golang)**, usando o framework **Gin Gonic** e o padrão de arquitetura **MVC (Model-View-Controller)**.

A aplicação simula um CRUD (Create, Read, Update, Delete) de livros, inicialmente com dados em memória, e posteriormente com suporte a banco de dados PostgreSQL.

---

## 🚀 Tecnologias utilizadas

- [Go](https://golang.org)
- [Gin Gonic](https://github.com/gin-gonic/gin)
- [PostgreSQL (opcional)](https://www.postgresql.org/)
- [Docker & Docker Compose (opcional)](https://www.docker.com/)
- Git

---

## 📁 Estrutura de pastas

```
crud-go-mvc/
├── controllers/    # Funções para lidar com requisições HTTP
├── models/         # Estruturas e lógica de dados
├── routes/         # Definições de rotas da API
├── main.go         # Ponto de entrada da aplicação
└── go.mod          # Gerenciamento de dependências
```

---

## 📌 Endpoints disponíveis

| Método | Rota             | Descrição                  |
|--------|------------------|----------------------------|
| GET    | /books           | Listar todos os livros     |
| GET    | /books/:id       | Obter um livro por ID      |
| POST   | /books           | Adicionar um novo livro    |
| PUT    | /books/:id       | Atualizar um livro         |
| DELETE | /books/:id       | Remover um livro           |

---

## ▶️ Como rodar o projeto (modo local - sem Docker)

```bash
# Clone o repositório
git clone https://github.com/seuusuario/crud-go-mvc.git
cd crud-go-mvc

# Instale as dependências
go mod tidy

# Rode o projeto
go run main.go
```

Acesse em: [http://localhost:8080/books](http://localhost:8080/books)

---

## 🐳 Rodando com Docker + PostgreSQL

1. Configure suas variáveis de ambiente no `docker-compose.yml`
2. Execute:

```bash
docker compose up --build
```

A aplicação estará disponível em [http://localhost:8080](http://localhost:8080)

---

## 📄 Licença

Este projeto está licenciado sob a [MIT License](LICENSE).

---

Desenvolvido por Paulo Eduardo de Oliveira Gonçalves 🚀

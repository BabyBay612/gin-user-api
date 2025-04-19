# 基于Gin的用户管理API

这是一个使用Gin](https://github.com/gin-gonic/gin)框架开发的用户管理API，基于Go语言实现，支持基本的 **增删改查（CRUD）** 操作，适合作为入门练习或中小型项目接口模板。

---

## 功能列表

- 创建用户（POST `/users`）
- 查询所有用户（GET `/users`）
- 查询单个用户（GET `/users/:id`）
- 更新用户信息（PUT `/users/:id`）
- 删除用户（DELETE `/users/:id`）

---

## 接口文档

### 创建用户

- **URL**：`POST /users`
- **请求体**（JSON）：

```json
{
  "name": "Tom",
  "age": 22
}
```

- **成功响应**：

```json
{
  "id": 1,
  "name": "Tom",
  "age": 22
}
```

---

### 获取所有用户

- **URL**：`GET /users`
- **响应示例**：

```json
[
  {
    "id": 1,
    "name": "Tom",
    "age": 22
  },
  {
    "id": 2,
    "name": "Jerry",
    "age": 20
  }
]
```

---

### 查询单个用户

- **URL**：`GET /users/:id`
- **成功响应**：

```json
{
  "id": 1,
  "name": "Tom",
  "age": 22
}
```

- **失败响应**：

```json
{
  "error": "User not found"
}
```

---

### 更新用户

- **URL**：`PUT /users/:id`
- **请求体**：

```json
{
  "name": "Tommy",
  "age": 23
}
```

- **响应**：

```json
{
  "id": 1,
  "name": "Tommy",
  "age": 23
}
```

---

### 删除用户

- **URL**：`DELETE /users/:id`
- **成功响应**：

```json
{
  "message": "User deleted"
}
```

- **失败响应**：

```json
{
  "error": "User not found"
}
```

---

## 启动方式

```bash
go run main.go
```

默认监听端口：

```
http://localhost:8080
```

---

## 作者

> GitHub: [@BabyBay612](https://github.com/BabyBay612)  
> 欢迎 Star、Fork、Issue！


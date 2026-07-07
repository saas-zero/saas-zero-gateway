# SaaS-Zero Gateway

API 网关 — go-zero gateway 纯 HTTP 透传代理。统一入口 `:18080`。

## 路由表

| 上游服务 | 路由前缀 | 目标端口 |
|---|---|---|
| Auth API | `/oauth/*` | `:18081` |
| Basedata API | `/system/*`, `/init/*` | `:18083` |
基于zero构建的多租户微服务版本  

网关不处理业务逻辑，只做请求转发。所有 JWT 认证、Casbin 权限检查由后端服务自行完成。
地址：https://github.com/saas-zero/saas-zero-gateway  

## 配置

`etc/gateway.yaml` 中的 `Upstreams` 定义路由映射：

```yaml
Upstreams:
  - Name: authservice.api
    Http:
      Target: 127.0.0.1:18081
    Mappings:
      - Method: POST
        Path: /oauth/login
      - Method: GET
        Path: /oauth/verify
      # ...

  - Name: basedata-api
    Http:
      Target: 127.0.0.1:18083
    Mappings:
      - Method: POST
        Path: /system/user/create
      # ...
      - Method: POST
        Path: /init/all
```

## 启动

```bash
# 确保 etcd 已启动
go run ./apps/saas-zero-gateway
# 或进入目录
cd apps/saas-zero-gateway
go run gateway.go -f etc/gateway.yaml
```

> 网关依赖 etcd 服务发现，需等所有后端服务就绪后再启动。

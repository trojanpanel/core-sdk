# 该项目先处于测试阶段

# Core SDK

Trojan Panel Core SDK

# 注意

1. Trojan Panel Core version >= 2.0.5

2. 在调用API前需要在Redis中添加以下数据

| key                    | 说明      |
|------------------------|---------|
| `trojan-panel:jwt-key` | jwt的key |
| `trojan-panel:token`   | 认证      |

使用`github.com/golang-jwt/jwt`生成token，例如：

```
type AccountVo struct {
	Id         uint      `json:"id"`
	Username   string    `json:"username"`
	RoleId     uint      `json:"roleId"`
	Deleted    uint      `json:"deleted"`
	Roles      []string  `json:"roles"`
}

type MyClaims struct {
	AccountVo vo.AccountVo `json:"accountVo"`
	jwt.StandardClaims
}

// TokenExpireDuration 过期时间默认2小时
const TokenExpireDuration = time.Hour * 2

// GenToken 生成Token
func GenToken(accountVo vo.AccountVo, jwtKey string) (string, error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		// 自定义字段
		AccountVo: accountVo,
		StandardClaims: jwt.StandardClaims{
			// 过期时间
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			// 签发人
			Issuer: "trojan-panel",
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	if err != nil {
		return "", err
	}
	return token.SignedString(jwtKey)
}
```

# Community

- Telegram Channel: [Trojan Panel](https://t.me/TrojanPanel)
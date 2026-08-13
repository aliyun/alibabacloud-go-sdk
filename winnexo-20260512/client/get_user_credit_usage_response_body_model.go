// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserCreditUsageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetUserCreditUsageResponseBody
	GetCode() *string
	SetCreditLimit(v string) *GetUserCreditUsageResponseBody
	GetCreditLimit() *string
	SetMessage(v string) *GetUserCreditUsageResponseBody
	GetMessage() *string
	SetRemainingCredits(v string) *GetUserCreditUsageResponseBody
	GetRemainingCredits() *string
	SetRequestId(v string) *GetUserCreditUsageResponseBody
	GetRequestId() *string
	SetShadowCreditLimit(v string) *GetUserCreditUsageResponseBody
	GetShadowCreditLimit() *string
	SetShadowRemainingCredits(v string) *GetUserCreditUsageResponseBody
	GetShadowRemainingCredits() *string
	SetShadowUsedCredits(v string) *GetUserCreditUsageResponseBody
	GetShadowUsedCredits() *string
	SetTenantId(v int64) *GetUserCreditUsageResponseBody
	GetTenantId() *int64
	SetUsedCredits(v string) *GetUserCreditUsageResponseBody
	GetUsedCredits() *string
	SetUserId(v int64) *GetUserCreditUsageResponseBody
	GetUserId() *int64
}

type GetUserCreditUsageResponseBody struct {
	// 业务状态码：成功为 200，失败为后端错误码（ERR.	- / InvalidParameter.*）
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 用户 credit 限额
	//
	// example:
	//
	// string_value
	CreditLimit *string `json:"creditLimit,omitempty" xml:"creditLimit,omitempty"`
	// 错误描述，成功时为空
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 剩余 credit（实时，来自 Redis）
	//
	// example:
	//
	// string_value
	RemainingCredits *string `json:"remainingCredits,omitempty" xml:"remainingCredits,omitempty"`
	// 请求追踪 ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 影子 credit 限额
	//
	// example:
	//
	// string_value
	ShadowCreditLimit *string `json:"shadowCreditLimit,omitempty" xml:"shadowCreditLimit,omitempty"`
	// 剩余影子 credit（实时，来自 Redis）
	//
	// example:
	//
	// string_value
	ShadowRemainingCredits *string `json:"shadowRemainingCredits,omitempty" xml:"shadowRemainingCredits,omitempty"`
	// 已消耗影子 credit（实时，来自 Redis）
	//
	// example:
	//
	// string_value
	ShadowUsedCredits *string `json:"shadowUsedCredits,omitempty" xml:"shadowUsedCredits,omitempty"`
	// 租户ID
	//
	// example:
	//
	// 10000
	TenantId *int64 `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// 已消耗 credit（实时，来自 Redis）
	//
	// example:
	//
	// string_value
	UsedCredits *string `json:"usedCredits,omitempty" xml:"usedCredits,omitempty"`
	// 用户ID
	//
	// example:
	//
	// 1
	UserId *int64 `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetUserCreditUsageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserCreditUsageResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserCreditUsageResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetUserCreditUsageResponseBody) GetCreditLimit() *string {
	return s.CreditLimit
}

func (s *GetUserCreditUsageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetUserCreditUsageResponseBody) GetRemainingCredits() *string {
	return s.RemainingCredits
}

func (s *GetUserCreditUsageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserCreditUsageResponseBody) GetShadowCreditLimit() *string {
	return s.ShadowCreditLimit
}

func (s *GetUserCreditUsageResponseBody) GetShadowRemainingCredits() *string {
	return s.ShadowRemainingCredits
}

func (s *GetUserCreditUsageResponseBody) GetShadowUsedCredits() *string {
	return s.ShadowUsedCredits
}

func (s *GetUserCreditUsageResponseBody) GetTenantId() *int64 {
	return s.TenantId
}

func (s *GetUserCreditUsageResponseBody) GetUsedCredits() *string {
	return s.UsedCredits
}

func (s *GetUserCreditUsageResponseBody) GetUserId() *int64 {
	return s.UserId
}

func (s *GetUserCreditUsageResponseBody) SetCode(v string) *GetUserCreditUsageResponseBody {
	s.Code = &v
	return s
}

func (s *GetUserCreditUsageResponseBody) SetCreditLimit(v string) *GetUserCreditUsageResponseBody {
	s.CreditLimit = &v
	return s
}

func (s *GetUserCreditUsageResponseBody) SetMessage(v string) *GetUserCreditUsageResponseBody {
	s.Message = &v
	return s
}

func (s *GetUserCreditUsageResponseBody) SetRemainingCredits(v string) *GetUserCreditUsageResponseBody {
	s.RemainingCredits = &v
	return s
}

func (s *GetUserCreditUsageResponseBody) SetRequestId(v string) *GetUserCreditUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserCreditUsageResponseBody) SetShadowCreditLimit(v string) *GetUserCreditUsageResponseBody {
	s.ShadowCreditLimit = &v
	return s
}

func (s *GetUserCreditUsageResponseBody) SetShadowRemainingCredits(v string) *GetUserCreditUsageResponseBody {
	s.ShadowRemainingCredits = &v
	return s
}

func (s *GetUserCreditUsageResponseBody) SetShadowUsedCredits(v string) *GetUserCreditUsageResponseBody {
	s.ShadowUsedCredits = &v
	return s
}

func (s *GetUserCreditUsageResponseBody) SetTenantId(v int64) *GetUserCreditUsageResponseBody {
	s.TenantId = &v
	return s
}

func (s *GetUserCreditUsageResponseBody) SetUsedCredits(v string) *GetUserCreditUsageResponseBody {
	s.UsedCredits = &v
	return s
}

func (s *GetUserCreditUsageResponseBody) SetUserId(v int64) *GetUserCreditUsageResponseBody {
	s.UserId = &v
	return s
}

func (s *GetUserCreditUsageResponseBody) Validate() error {
	return dara.Validate(s)
}

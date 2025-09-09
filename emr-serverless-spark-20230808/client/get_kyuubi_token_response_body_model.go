// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKyuubiTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetKyuubiTokenResponseBodyData) *GetKyuubiTokenResponseBody
	GetData() *GetKyuubiTokenResponseBodyData
	SetRequestId(v string) *GetKyuubiTokenResponseBody
	GetRequestId() *string
}

type GetKyuubiTokenResponseBody struct {
	Data *GetKyuubiTokenResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetKyuubiTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetKyuubiTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetKyuubiTokenResponseBody) GetData() *GetKyuubiTokenResponseBodyData {
	return s.Data
}

func (s *GetKyuubiTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetKyuubiTokenResponseBody) SetData(v *GetKyuubiTokenResponseBodyData) *GetKyuubiTokenResponseBody {
	s.Data = v
	return s
}

func (s *GetKyuubiTokenResponseBody) SetRequestId(v string) *GetKyuubiTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetKyuubiTokenResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetKyuubiTokenResponseBodyData struct {
	AutoExpireConfiguration *GetKyuubiTokenResponseBodyDataAutoExpireConfiguration `json:"autoExpireConfiguration,omitempty" xml:"autoExpireConfiguration,omitempty" type:"Struct"`
	// example:
	//
	// 1749456094000
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// admin
	CreatedBy *string `json:"createdBy,omitempty" xml:"createdBy,omitempty"`
	// example:
	//
	// 1753932319390
	ExpireTime *int64 `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	// example:
	//
	// 1749456098000
	LastUsedTime *int64    `json:"lastUsedTime,omitempty" xml:"lastUsedTime,omitempty"`
	MemberArns   []*string `json:"memberArns,omitempty" xml:"memberArns,omitempty" type:"Repeated"`
	// example:
	//
	// dev_serverless_spark
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// dxj**********wfg
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
	// Token ID。
	//
	// example:
	//
	// tk-zpi0*****hdv4y
	TokenId *string `json:"tokenId,omitempty" xml:"tokenId,omitempty"`
}

func (s GetKyuubiTokenResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetKyuubiTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetKyuubiTokenResponseBodyData) GetAutoExpireConfiguration() *GetKyuubiTokenResponseBodyDataAutoExpireConfiguration {
	return s.AutoExpireConfiguration
}

func (s *GetKyuubiTokenResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetKyuubiTokenResponseBodyData) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *GetKyuubiTokenResponseBodyData) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *GetKyuubiTokenResponseBodyData) GetLastUsedTime() *int64 {
	return s.LastUsedTime
}

func (s *GetKyuubiTokenResponseBodyData) GetMemberArns() []*string {
	return s.MemberArns
}

func (s *GetKyuubiTokenResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetKyuubiTokenResponseBodyData) GetToken() *string {
	return s.Token
}

func (s *GetKyuubiTokenResponseBodyData) GetTokenId() *string {
	return s.TokenId
}

func (s *GetKyuubiTokenResponseBodyData) SetAutoExpireConfiguration(v *GetKyuubiTokenResponseBodyDataAutoExpireConfiguration) *GetKyuubiTokenResponseBodyData {
	s.AutoExpireConfiguration = v
	return s
}

func (s *GetKyuubiTokenResponseBodyData) SetCreateTime(v int64) *GetKyuubiTokenResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetKyuubiTokenResponseBodyData) SetCreatedBy(v string) *GetKyuubiTokenResponseBodyData {
	s.CreatedBy = &v
	return s
}

func (s *GetKyuubiTokenResponseBodyData) SetExpireTime(v int64) *GetKyuubiTokenResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *GetKyuubiTokenResponseBodyData) SetLastUsedTime(v int64) *GetKyuubiTokenResponseBodyData {
	s.LastUsedTime = &v
	return s
}

func (s *GetKyuubiTokenResponseBodyData) SetMemberArns(v []*string) *GetKyuubiTokenResponseBodyData {
	s.MemberArns = v
	return s
}

func (s *GetKyuubiTokenResponseBodyData) SetName(v string) *GetKyuubiTokenResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetKyuubiTokenResponseBodyData) SetToken(v string) *GetKyuubiTokenResponseBodyData {
	s.Token = &v
	return s
}

func (s *GetKyuubiTokenResponseBodyData) SetTokenId(v string) *GetKyuubiTokenResponseBodyData {
	s.TokenId = &v
	return s
}

func (s *GetKyuubiTokenResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetKyuubiTokenResponseBodyDataAutoExpireConfiguration struct {
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// example:
	//
	// 365
	ExpireDays *int32 `json:"expireDays,omitempty" xml:"expireDays,omitempty"`
}

func (s GetKyuubiTokenResponseBodyDataAutoExpireConfiguration) String() string {
	return dara.Prettify(s)
}

func (s GetKyuubiTokenResponseBodyDataAutoExpireConfiguration) GoString() string {
	return s.String()
}

func (s *GetKyuubiTokenResponseBodyDataAutoExpireConfiguration) GetEnable() *bool {
	return s.Enable
}

func (s *GetKyuubiTokenResponseBodyDataAutoExpireConfiguration) GetExpireDays() *int32 {
	return s.ExpireDays
}

func (s *GetKyuubiTokenResponseBodyDataAutoExpireConfiguration) SetEnable(v bool) *GetKyuubiTokenResponseBodyDataAutoExpireConfiguration {
	s.Enable = &v
	return s
}

func (s *GetKyuubiTokenResponseBodyDataAutoExpireConfiguration) SetExpireDays(v int32) *GetKyuubiTokenResponseBodyDataAutoExpireConfiguration {
	s.ExpireDays = &v
	return s
}

func (s *GetKyuubiTokenResponseBodyDataAutoExpireConfiguration) Validate() error {
	return dara.Validate(s)
}

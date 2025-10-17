// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLivyComputeTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetLivyComputeTokenResponseBody
	GetCode() *string
	SetData(v *GetLivyComputeTokenResponseBodyData) *GetLivyComputeTokenResponseBody
	GetData() *GetLivyComputeTokenResponseBodyData
	SetMessage(v string) *GetLivyComputeTokenResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetLivyComputeTokenResponseBody
	GetRequestId() *string
}

type GetLivyComputeTokenResponseBody struct {
	// example:
	//
	// 1000000
	Code *string                              `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetLivyComputeTokenResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 484D9DDA-300D-525E-AF7A-0CCCA5C64A7A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetLivyComputeTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLivyComputeTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetLivyComputeTokenResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetLivyComputeTokenResponseBody) GetData() *GetLivyComputeTokenResponseBodyData {
	return s.Data
}

func (s *GetLivyComputeTokenResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetLivyComputeTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLivyComputeTokenResponseBody) SetCode(v string) *GetLivyComputeTokenResponseBody {
	s.Code = &v
	return s
}

func (s *GetLivyComputeTokenResponseBody) SetData(v *GetLivyComputeTokenResponseBodyData) *GetLivyComputeTokenResponseBody {
	s.Data = v
	return s
}

func (s *GetLivyComputeTokenResponseBody) SetMessage(v string) *GetLivyComputeTokenResponseBody {
	s.Message = &v
	return s
}

func (s *GetLivyComputeTokenResponseBody) SetRequestId(v string) *GetLivyComputeTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLivyComputeTokenResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetLivyComputeTokenResponseBodyData struct {
	AutoExpireConfiguration *GetLivyComputeTokenResponseBodyDataAutoExpireConfiguration `json:"autoExpireConfiguration,omitempty" xml:"autoExpireConfiguration,omitempty" type:"Struct"`
	// example:
	//
	// 1749456094000
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// alice
	CreatedBy *string `json:"createdBy,omitempty" xml:"createdBy,omitempty"`
	// example:
	//
	// 1749457994000
	ExpireTime *int64 `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	// example:
	//
	// 1749456098000
	LastUsedTime *int64 `json:"lastUsedTime,omitempty" xml:"lastUsedTime,omitempty"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// d25561157a635bb
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
	// Token ID。
	//
	// example:
	//
	// lctk-xxxxxxxxxx
	TokenId *string `json:"tokenId,omitempty" xml:"tokenId,omitempty"`
}

func (s GetLivyComputeTokenResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetLivyComputeTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetLivyComputeTokenResponseBodyData) GetAutoExpireConfiguration() *GetLivyComputeTokenResponseBodyDataAutoExpireConfiguration {
	return s.AutoExpireConfiguration
}

func (s *GetLivyComputeTokenResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetLivyComputeTokenResponseBodyData) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *GetLivyComputeTokenResponseBodyData) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *GetLivyComputeTokenResponseBodyData) GetLastUsedTime() *int64 {
	return s.LastUsedTime
}

func (s *GetLivyComputeTokenResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetLivyComputeTokenResponseBodyData) GetToken() *string {
	return s.Token
}

func (s *GetLivyComputeTokenResponseBodyData) GetTokenId() *string {
	return s.TokenId
}

func (s *GetLivyComputeTokenResponseBodyData) SetAutoExpireConfiguration(v *GetLivyComputeTokenResponseBodyDataAutoExpireConfiguration) *GetLivyComputeTokenResponseBodyData {
	s.AutoExpireConfiguration = v
	return s
}

func (s *GetLivyComputeTokenResponseBodyData) SetCreateTime(v int64) *GetLivyComputeTokenResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetLivyComputeTokenResponseBodyData) SetCreatedBy(v string) *GetLivyComputeTokenResponseBodyData {
	s.CreatedBy = &v
	return s
}

func (s *GetLivyComputeTokenResponseBodyData) SetExpireTime(v int64) *GetLivyComputeTokenResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *GetLivyComputeTokenResponseBodyData) SetLastUsedTime(v int64) *GetLivyComputeTokenResponseBodyData {
	s.LastUsedTime = &v
	return s
}

func (s *GetLivyComputeTokenResponseBodyData) SetName(v string) *GetLivyComputeTokenResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetLivyComputeTokenResponseBodyData) SetToken(v string) *GetLivyComputeTokenResponseBodyData {
	s.Token = &v
	return s
}

func (s *GetLivyComputeTokenResponseBodyData) SetTokenId(v string) *GetLivyComputeTokenResponseBodyData {
	s.TokenId = &v
	return s
}

func (s *GetLivyComputeTokenResponseBodyData) Validate() error {
	if s.AutoExpireConfiguration != nil {
		if err := s.AutoExpireConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetLivyComputeTokenResponseBodyDataAutoExpireConfiguration struct {
	// example:
	//
	// false
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// example:
	//
	// 7
	ExpireDays *int32 `json:"expireDays,omitempty" xml:"expireDays,omitempty"`
}

func (s GetLivyComputeTokenResponseBodyDataAutoExpireConfiguration) String() string {
	return dara.Prettify(s)
}

func (s GetLivyComputeTokenResponseBodyDataAutoExpireConfiguration) GoString() string {
	return s.String()
}

func (s *GetLivyComputeTokenResponseBodyDataAutoExpireConfiguration) GetEnable() *bool {
	return s.Enable
}

func (s *GetLivyComputeTokenResponseBodyDataAutoExpireConfiguration) GetExpireDays() *int32 {
	return s.ExpireDays
}

func (s *GetLivyComputeTokenResponseBodyDataAutoExpireConfiguration) SetEnable(v bool) *GetLivyComputeTokenResponseBodyDataAutoExpireConfiguration {
	s.Enable = &v
	return s
}

func (s *GetLivyComputeTokenResponseBodyDataAutoExpireConfiguration) SetExpireDays(v int32) *GetLivyComputeTokenResponseBodyDataAutoExpireConfiguration {
	s.ExpireDays = &v
	return s
}

func (s *GetLivyComputeTokenResponseBodyDataAutoExpireConfiguration) Validate() error {
	return dara.Validate(s)
}

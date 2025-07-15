// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTrafficSpecialControlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSecurityToken(v string) *AddTrafficSpecialControlRequest
	GetSecurityToken() *string
	SetSpecialKey(v string) *AddTrafficSpecialControlRequest
	GetSpecialKey() *string
	SetSpecialType(v string) *AddTrafficSpecialControlRequest
	GetSpecialType() *string
	SetTrafficControlId(v string) *AddTrafficSpecialControlRequest
	GetTrafficControlId() *string
	SetTrafficValue(v int32) *AddTrafficSpecialControlRequest
	GetTrafficValue() *int32
}

type AddTrafficSpecialControlRequest struct {
	// The security token included in the WebSocket request header. The system uses this token to authenticate the request.
	//
	// example:
	//
	// fa876ffb-caab-4f0a-93b3-3409f2fa5199
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The ID of the app or Alibaba Cloud account. Specify this parameter based on the value of the **SpecialType*	- parameter. You can view your account ID on the [Account Management](https://account.console.aliyun.com/?spm=a2c4g.11186623.2.15.3f053654YpMPwo#/secure) page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3382463
	SpecialKey *string `json:"SpecialKey,omitempty" xml:"SpecialKey,omitempty"`
	// The type of the special throttling policy. Valid values:
	//
	// 	- **APP**
	//
	// 	- **USER**
	//
	// This parameter is required.
	//
	// example:
	//
	// APP
	SpecialType *string `json:"SpecialType,omitempty" xml:"SpecialType,omitempty"`
	// The ID of the specified throttling policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// tf123456
	TrafficControlId *string `json:"TrafficControlId,omitempty" xml:"TrafficControlId,omitempty"`
	// The special throttling value.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	TrafficValue *int32 `json:"TrafficValue,omitempty" xml:"TrafficValue,omitempty"`
}

func (s AddTrafficSpecialControlRequest) String() string {
	return dara.Prettify(s)
}

func (s AddTrafficSpecialControlRequest) GoString() string {
	return s.String()
}

func (s *AddTrafficSpecialControlRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *AddTrafficSpecialControlRequest) GetSpecialKey() *string {
	return s.SpecialKey
}

func (s *AddTrafficSpecialControlRequest) GetSpecialType() *string {
	return s.SpecialType
}

func (s *AddTrafficSpecialControlRequest) GetTrafficControlId() *string {
	return s.TrafficControlId
}

func (s *AddTrafficSpecialControlRequest) GetTrafficValue() *int32 {
	return s.TrafficValue
}

func (s *AddTrafficSpecialControlRequest) SetSecurityToken(v string) *AddTrafficSpecialControlRequest {
	s.SecurityToken = &v
	return s
}

func (s *AddTrafficSpecialControlRequest) SetSpecialKey(v string) *AddTrafficSpecialControlRequest {
	s.SpecialKey = &v
	return s
}

func (s *AddTrafficSpecialControlRequest) SetSpecialType(v string) *AddTrafficSpecialControlRequest {
	s.SpecialType = &v
	return s
}

func (s *AddTrafficSpecialControlRequest) SetTrafficControlId(v string) *AddTrafficSpecialControlRequest {
	s.TrafficControlId = &v
	return s
}

func (s *AddTrafficSpecialControlRequest) SetTrafficValue(v int32) *AddTrafficSpecialControlRequest {
	s.TrafficValue = &v
	return s
}

func (s *AddTrafficSpecialControlRequest) Validate() error {
	return dara.Validate(s)
}

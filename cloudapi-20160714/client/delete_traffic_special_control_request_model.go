// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTrafficSpecialControlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSecurityToken(v string) *DeleteTrafficSpecialControlRequest
	GetSecurityToken() *string
	SetSpecialKey(v string) *DeleteTrafficSpecialControlRequest
	GetSpecialKey() *string
	SetSpecialType(v string) *DeleteTrafficSpecialControlRequest
	GetSpecialType() *string
	SetTrafficControlId(v string) *DeleteTrafficSpecialControlRequest
	GetTrafficControlId() *string
}

type DeleteTrafficSpecialControlRequest struct {
	// The security token included in the WebSocket request header. The system uses this token to authenticate the request.
	//
	// example:
	//
	// 7c51b234-48d3-44e1-9b36-e2ddccc738e3
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The ID of the app or Alibaba Cloud account. You can view your account ID on the [Account Management](https://account.console.aliyun.com/?spm=a2c4g.11186623.2.15.343130a8sDi8cO#/secure) page.
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
	// The ID of the throttling policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// tf123456
	TrafficControlId *string `json:"TrafficControlId,omitempty" xml:"TrafficControlId,omitempty"`
}

func (s DeleteTrafficSpecialControlRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTrafficSpecialControlRequest) GoString() string {
	return s.String()
}

func (s *DeleteTrafficSpecialControlRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DeleteTrafficSpecialControlRequest) GetSpecialKey() *string {
	return s.SpecialKey
}

func (s *DeleteTrafficSpecialControlRequest) GetSpecialType() *string {
	return s.SpecialType
}

func (s *DeleteTrafficSpecialControlRequest) GetTrafficControlId() *string {
	return s.TrafficControlId
}

func (s *DeleteTrafficSpecialControlRequest) SetSecurityToken(v string) *DeleteTrafficSpecialControlRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteTrafficSpecialControlRequest) SetSpecialKey(v string) *DeleteTrafficSpecialControlRequest {
	s.SpecialKey = &v
	return s
}

func (s *DeleteTrafficSpecialControlRequest) SetSpecialType(v string) *DeleteTrafficSpecialControlRequest {
	s.SpecialType = &v
	return s
}

func (s *DeleteTrafficSpecialControlRequest) SetTrafficControlId(v string) *DeleteTrafficSpecialControlRequest {
	s.TrafficControlId = &v
	return s
}

func (s *DeleteTrafficSpecialControlRequest) Validate() error {
	return dara.Validate(s)
}

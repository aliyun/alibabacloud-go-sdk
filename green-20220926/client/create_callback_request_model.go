// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCallbackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCryptType(v string) *CreateCallbackRequest
	GetCryptType() *string
	SetName(v string) *CreateCallbackRequest
	GetName() *string
	SetRegionId(v string) *CreateCallbackRequest
	GetRegionId() *string
	SetScope(v string) *CreateCallbackRequest
	GetScope() *string
	SetUrl(v string) *CreateCallbackRequest
	GetUrl() *string
}

type CreateCallbackRequest struct {
	// Encryption algorithm.
	//
	// example:
	//
	// SHA256
	CryptType *string `json:"CryptType,omitempty" xml:"CryptType,omitempty"`
	// Plan name.
	//
	// example:
	//
	// 消息通知1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Review result.
	//
	// example:
	//
	// all
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// Callback URL.
	//
	// example:
	//
	// https://console.aliyun.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CreateCallbackRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCallbackRequest) GoString() string {
	return s.String()
}

func (s *CreateCallbackRequest) GetCryptType() *string {
	return s.CryptType
}

func (s *CreateCallbackRequest) GetName() *string {
	return s.Name
}

func (s *CreateCallbackRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateCallbackRequest) GetScope() *string {
	return s.Scope
}

func (s *CreateCallbackRequest) GetUrl() *string {
	return s.Url
}

func (s *CreateCallbackRequest) SetCryptType(v string) *CreateCallbackRequest {
	s.CryptType = &v
	return s
}

func (s *CreateCallbackRequest) SetName(v string) *CreateCallbackRequest {
	s.Name = &v
	return s
}

func (s *CreateCallbackRequest) SetRegionId(v string) *CreateCallbackRequest {
	s.RegionId = &v
	return s
}

func (s *CreateCallbackRequest) SetScope(v string) *CreateCallbackRequest {
	s.Scope = &v
	return s
}

func (s *CreateCallbackRequest) SetUrl(v string) *CreateCallbackRequest {
	s.Url = &v
	return s
}

func (s *CreateCallbackRequest) Validate() error {
	return dara.Validate(s)
}

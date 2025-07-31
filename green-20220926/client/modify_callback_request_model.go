// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCallbackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCryptType(v string) *ModifyCallbackRequest
	GetCryptType() *string
	SetId(v int64) *ModifyCallbackRequest
	GetId() *int64
	SetName(v string) *ModifyCallbackRequest
	GetName() *string
	SetRegionId(v string) *ModifyCallbackRequest
	GetRegionId() *string
	SetScope(v string) *ModifyCallbackRequest
	GetScope() *string
	SetUrl(v string) *ModifyCallbackRequest
	GetUrl() *string
}

type ModifyCallbackRequest struct {
	// example:
	//
	// SHA256
	CryptType *string `json:"CryptType,omitempty" xml:"CryptType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 112
	Id   *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// all
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// example:
	//
	// https://www.aliyuncs.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ModifyCallbackRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCallbackRequest) GoString() string {
	return s.String()
}

func (s *ModifyCallbackRequest) GetCryptType() *string {
	return s.CryptType
}

func (s *ModifyCallbackRequest) GetId() *int64 {
	return s.Id
}

func (s *ModifyCallbackRequest) GetName() *string {
	return s.Name
}

func (s *ModifyCallbackRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyCallbackRequest) GetScope() *string {
	return s.Scope
}

func (s *ModifyCallbackRequest) GetUrl() *string {
	return s.Url
}

func (s *ModifyCallbackRequest) SetCryptType(v string) *ModifyCallbackRequest {
	s.CryptType = &v
	return s
}

func (s *ModifyCallbackRequest) SetId(v int64) *ModifyCallbackRequest {
	s.Id = &v
	return s
}

func (s *ModifyCallbackRequest) SetName(v string) *ModifyCallbackRequest {
	s.Name = &v
	return s
}

func (s *ModifyCallbackRequest) SetRegionId(v string) *ModifyCallbackRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyCallbackRequest) SetScope(v string) *ModifyCallbackRequest {
	s.Scope = &v
	return s
}

func (s *ModifyCallbackRequest) SetUrl(v string) *ModifyCallbackRequest {
	s.Url = &v
	return s
}

func (s *ModifyCallbackRequest) Validate() error {
	return dara.Validate(s)
}

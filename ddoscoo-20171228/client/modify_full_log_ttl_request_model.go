// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFullLogTtlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ModifyFullLogTtlRequest
	GetLang() *string
	SetResourceGroupId(v string) *ModifyFullLogTtlRequest
	GetResourceGroupId() *string
	SetSourceIp(v string) *ModifyFullLogTtlRequest
	GetSourceIp() *string
	SetTtl(v int32) *ModifyFullLogTtlRequest
	GetTtl() *int32
}

type ModifyFullLogTtlRequest struct {
	// example:
	//
	// cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// xx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
}

func (s ModifyFullLogTtlRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyFullLogTtlRequest) GoString() string {
	return s.String()
}

func (s *ModifyFullLogTtlRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifyFullLogTtlRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyFullLogTtlRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *ModifyFullLogTtlRequest) GetTtl() *int32 {
	return s.Ttl
}

func (s *ModifyFullLogTtlRequest) SetLang(v string) *ModifyFullLogTtlRequest {
	s.Lang = &v
	return s
}

func (s *ModifyFullLogTtlRequest) SetResourceGroupId(v string) *ModifyFullLogTtlRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyFullLogTtlRequest) SetSourceIp(v string) *ModifyFullLogTtlRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyFullLogTtlRequest) SetTtl(v int32) *ModifyFullLogTtlRequest {
	s.Ttl = &v
	return s
}

func (s *ModifyFullLogTtlRequest) Validate() error {
	return dara.Validate(s)
}

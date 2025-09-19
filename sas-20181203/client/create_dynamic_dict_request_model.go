// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDynamicDictRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOverride(v bool) *CreateDynamicDictRequest
	GetOverride() *bool
	SetSourceIp(v string) *CreateDynamicDictRequest
	GetSourceIp() *string
}

type CreateDynamicDictRequest struct {
	// Specifies whether to overwrite existing data. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	Override *bool `json:"Override,omitempty" xml:"Override,omitempty"`
	// The source IP address.
	//
	// example:
	//
	// 123.103.9.***
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s CreateDynamicDictRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDynamicDictRequest) GoString() string {
	return s.String()
}

func (s *CreateDynamicDictRequest) GetOverride() *bool {
	return s.Override
}

func (s *CreateDynamicDictRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *CreateDynamicDictRequest) SetOverride(v bool) *CreateDynamicDictRequest {
	s.Override = &v
	return s
}

func (s *CreateDynamicDictRequest) SetSourceIp(v string) *CreateDynamicDictRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateDynamicDictRequest) Validate() error {
	return dara.Validate(s)
}

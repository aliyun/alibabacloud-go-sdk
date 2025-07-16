// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDomainsByLogConfigIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v string) *ListDomainsByLogConfigIdRequest
	GetConfigId() *string
}

type ListDomainsByLogConfigIdRequest struct {
	// The ID of the custom configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
}

func (s ListDomainsByLogConfigIdRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDomainsByLogConfigIdRequest) GoString() string {
	return s.String()
}

func (s *ListDomainsByLogConfigIdRequest) GetConfigId() *string {
	return s.ConfigId
}

func (s *ListDomainsByLogConfigIdRequest) SetConfigId(v string) *ListDomainsByLogConfigIdRequest {
	s.ConfigId = &v
	return s
}

func (s *ListDomainsByLogConfigIdRequest) Validate() error {
	return dara.Validate(s)
}

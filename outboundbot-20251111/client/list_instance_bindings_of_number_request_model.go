// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceBindingsOfNumberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNumber(v string) *ListInstanceBindingsOfNumberRequest
	GetNumber() *string
	SetTenantId(v string) *ListInstanceBindingsOfNumberRequest
	GetTenantId() *string
}

type ListInstanceBindingsOfNumberRequest struct {
	// The bound number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 057100001234
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 1234567890
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ListInstanceBindingsOfNumberRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceBindingsOfNumberRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceBindingsOfNumberRequest) GetNumber() *string {
	return s.Number
}

func (s *ListInstanceBindingsOfNumberRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListInstanceBindingsOfNumberRequest) SetNumber(v string) *ListInstanceBindingsOfNumberRequest {
	s.Number = &v
	return s
}

func (s *ListInstanceBindingsOfNumberRequest) SetTenantId(v string) *ListInstanceBindingsOfNumberRequest {
	s.TenantId = &v
	return s
}

func (s *ListInstanceBindingsOfNumberRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTestPreModelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomerModuleId(v int32) *TestPreModelRequest
	GetCustomerModuleId() *int32
}

type TestPreModelRequest struct {
	// Customer model ID
	//
	// example:
	//
	// 456
	CustomerModuleId *int32 `json:"CustomerModuleId,omitempty" xml:"CustomerModuleId,omitempty"`
}

func (s TestPreModelRequest) String() string {
	return dara.Prettify(s)
}

func (s TestPreModelRequest) GoString() string {
	return s.String()
}

func (s *TestPreModelRequest) GetCustomerModuleId() *int32 {
	return s.CustomerModuleId
}

func (s *TestPreModelRequest) SetCustomerModuleId(v int32) *TestPreModelRequest {
	s.CustomerModuleId = &v
	return s
}

func (s *TestPreModelRequest) Validate() error {
	return dara.Validate(s)
}

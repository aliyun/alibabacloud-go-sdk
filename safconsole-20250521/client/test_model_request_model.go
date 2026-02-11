// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTestModelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomerModuleId(v int32) *TestModelRequest
	GetCustomerModuleId() *int32
}

type TestModelRequest struct {
	// Customer model ID
	//
	// example:
	//
	// 456
	CustomerModuleId *int32 `json:"CustomerModuleId,omitempty" xml:"CustomerModuleId,omitempty"`
}

func (s TestModelRequest) String() string {
	return dara.Prettify(s)
}

func (s TestModelRequest) GoString() string {
	return s.String()
}

func (s *TestModelRequest) GetCustomerModuleId() *int32 {
	return s.CustomerModuleId
}

func (s *TestModelRequest) SetCustomerModuleId(v int32) *TestModelRequest {
	s.CustomerModuleId = &v
	return s
}

func (s *TestModelRequest) Validate() error {
	return dara.Validate(s)
}

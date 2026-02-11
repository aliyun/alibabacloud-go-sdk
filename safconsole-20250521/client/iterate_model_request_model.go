// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIterateModelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomerModuleId(v int32) *IterateModelRequest
	GetCustomerModuleId() *int32
}

type IterateModelRequest struct {
	// Customer model ID
	//
	// example:
	//
	// 456
	CustomerModuleId *int32 `json:"CustomerModuleId,omitempty" xml:"CustomerModuleId,omitempty"`
}

func (s IterateModelRequest) String() string {
	return dara.Prettify(s)
}

func (s IterateModelRequest) GoString() string {
	return s.String()
}

func (s *IterateModelRequest) GetCustomerModuleId() *int32 {
	return s.CustomerModuleId
}

func (s *IterateModelRequest) SetCustomerModuleId(v int32) *IterateModelRequest {
	s.CustomerModuleId = &v
	return s
}

func (s *IterateModelRequest) Validate() error {
	return dara.Validate(s)
}

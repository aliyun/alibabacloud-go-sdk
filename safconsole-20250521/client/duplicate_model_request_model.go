// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDuplicateModelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomerModuleId(v int32) *DuplicateModelRequest
	GetCustomerModuleId() *int32
}

type DuplicateModelRequest struct {
	// Customer model ID
	//
	// example:
	//
	// 456
	CustomerModuleId *int32 `json:"CustomerModuleId,omitempty" xml:"CustomerModuleId,omitempty"`
}

func (s DuplicateModelRequest) String() string {
	return dara.Prettify(s)
}

func (s DuplicateModelRequest) GoString() string {
	return s.String()
}

func (s *DuplicateModelRequest) GetCustomerModuleId() *int32 {
	return s.CustomerModuleId
}

func (s *DuplicateModelRequest) SetCustomerModuleId(v int32) *DuplicateModelRequest {
	s.CustomerModuleId = &v
	return s
}

func (s *DuplicateModelRequest) Validate() error {
	return dara.Validate(s)
}

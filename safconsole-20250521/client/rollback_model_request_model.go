// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackModelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomerModuleId(v int32) *RollbackModelRequest
	GetCustomerModuleId() *int32
}

type RollbackModelRequest struct {
	// Customer model ID
	//
	// example:
	//
	// 456
	CustomerModuleId *int32 `json:"CustomerModuleId,omitempty" xml:"CustomerModuleId,omitempty"`
}

func (s RollbackModelRequest) String() string {
	return dara.Prettify(s)
}

func (s RollbackModelRequest) GoString() string {
	return s.String()
}

func (s *RollbackModelRequest) GetCustomerModuleId() *int32 {
	return s.CustomerModuleId
}

func (s *RollbackModelRequest) SetCustomerModuleId(v int32) *RollbackModelRequest {
	s.CustomerModuleId = &v
	return s
}

func (s *RollbackModelRequest) Validate() error {
	return dara.Validate(s)
}

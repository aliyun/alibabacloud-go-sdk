// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteModelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomerModuleId(v int32) *DeleteModelRequest
	GetCustomerModuleId() *int32
}

type DeleteModelRequest struct {
	// Customer model ID
	//
	// example:
	//
	// 1
	CustomerModuleId *int32 `json:"CustomerModuleId,omitempty" xml:"CustomerModuleId,omitempty"`
}

func (s DeleteModelRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteModelRequest) GoString() string {
	return s.String()
}

func (s *DeleteModelRequest) GetCustomerModuleId() *int32 {
	return s.CustomerModuleId
}

func (s *DeleteModelRequest) SetCustomerModuleId(v int32) *DeleteModelRequest {
	s.CustomerModuleId = &v
	return s
}

func (s *DeleteModelRequest) Validate() error {
	return dara.Validate(s)
}

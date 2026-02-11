// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPrepublishModelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomerModuleId(v int32) *PrepublishModelRequest
	GetCustomerModuleId() *int32
}

type PrepublishModelRequest struct {
	// Customer model ID
	//
	// example:
	//
	// 456
	CustomerModuleId *int32 `json:"CustomerModuleId,omitempty" xml:"CustomerModuleId,omitempty"`
}

func (s PrepublishModelRequest) String() string {
	return dara.Prettify(s)
}

func (s PrepublishModelRequest) GoString() string {
	return s.String()
}

func (s *PrepublishModelRequest) GetCustomerModuleId() *int32 {
	return s.CustomerModuleId
}

func (s *PrepublishModelRequest) SetCustomerModuleId(v int32) *PrepublishModelRequest {
	s.CustomerModuleId = &v
	return s
}

func (s *PrepublishModelRequest) Validate() error {
	return dara.Validate(s)
}

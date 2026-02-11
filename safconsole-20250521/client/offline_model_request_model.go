// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflineModelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomerModuleId(v int32) *OfflineModelRequest
	GetCustomerModuleId() *int32
}

type OfflineModelRequest struct {
	// Customer model ID
	//
	// example:
	//
	// 456
	CustomerModuleId *int32 `json:"CustomerModuleId,omitempty" xml:"CustomerModuleId,omitempty"`
}

func (s OfflineModelRequest) String() string {
	return dara.Prettify(s)
}

func (s OfflineModelRequest) GoString() string {
	return s.String()
}

func (s *OfflineModelRequest) GetCustomerModuleId() *int32 {
	return s.CustomerModuleId
}

func (s *OfflineModelRequest) SetCustomerModuleId(v int32) *OfflineModelRequest {
	s.CustomerModuleId = &v
	return s
}

func (s *OfflineModelRequest) Validate() error {
	return dara.Validate(s)
}

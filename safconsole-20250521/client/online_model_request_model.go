// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnlineModelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomerModuleId(v int32) *OnlineModelRequest
	GetCustomerModuleId() *int32
}

type OnlineModelRequest struct {
	// example:
	//
	// 456
	CustomerModuleId *int32 `json:"CustomerModuleId,omitempty" xml:"CustomerModuleId,omitempty"`
}

func (s OnlineModelRequest) String() string {
	return dara.Prettify(s)
}

func (s OnlineModelRequest) GoString() string {
	return s.String()
}

func (s *OnlineModelRequest) GetCustomerModuleId() *int32 {
	return s.CustomerModuleId
}

func (s *OnlineModelRequest) SetCustomerModuleId(v int32) *OnlineModelRequest {
	s.CustomerModuleId = &v
	return s
}

func (s *OnlineModelRequest) Validate() error {
	return dara.Validate(s)
}

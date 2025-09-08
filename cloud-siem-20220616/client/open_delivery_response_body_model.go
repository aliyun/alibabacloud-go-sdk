// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenDeliveryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *OpenDeliveryResponseBody
	GetData() *bool
	SetRequestId(v string) *OpenDeliveryResponseBody
	GetRequestId() *string
}

type OpenDeliveryResponseBody struct {
	// Indicates whether the log delivery feature is enabled. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 15FD134E-D69B-51E8-B052-73F97BD8****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenDeliveryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenDeliveryResponseBody) GoString() string {
	return s.String()
}

func (s *OpenDeliveryResponseBody) GetData() *bool {
	return s.Data
}

func (s *OpenDeliveryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenDeliveryResponseBody) SetData(v bool) *OpenDeliveryResponseBody {
	s.Data = &v
	return s
}

func (s *OpenDeliveryResponseBody) SetRequestId(v string) *OpenDeliveryResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenDeliveryResponseBody) Validate() error {
	return dara.Validate(s)
}

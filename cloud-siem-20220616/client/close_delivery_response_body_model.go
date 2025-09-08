// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseDeliveryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *CloseDeliveryResponseBody
	GetData() *bool
	SetRequestId(v string) *CloseDeliveryResponseBody
	GetRequestId() *string
}

type CloseDeliveryResponseBody struct {
	// Indicates whether the threat analysis feature was disabled. Valid values:
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
	// F375A043-4F5B-55F2-A564-CC47FFC6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloseDeliveryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloseDeliveryResponseBody) GoString() string {
	return s.String()
}

func (s *CloseDeliveryResponseBody) GetData() *bool {
	return s.Data
}

func (s *CloseDeliveryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloseDeliveryResponseBody) SetData(v bool) *CloseDeliveryResponseBody {
	s.Data = &v
	return s
}

func (s *CloseDeliveryResponseBody) SetRequestId(v string) *CloseDeliveryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloseDeliveryResponseBody) Validate() error {
	return dara.Validate(s)
}

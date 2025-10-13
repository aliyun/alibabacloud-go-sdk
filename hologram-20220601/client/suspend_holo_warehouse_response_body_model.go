// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendHoloWarehouseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *SuspendHoloWarehouseResponseBody
	GetData() *string
	SetRequestId(v string) *SuspendHoloWarehouseResponseBody
	GetRequestId() *string
}

type SuspendHoloWarehouseResponseBody struct {
	// Indicates whether the request was successful. Valid values: true and false.
	//
	// example:
	//
	// true
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2A8DEF6E-067E-5DB0-BAE1-2894266E6C6A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SuspendHoloWarehouseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SuspendHoloWarehouseResponseBody) GoString() string {
	return s.String()
}

func (s *SuspendHoloWarehouseResponseBody) GetData() *string {
	return s.Data
}

func (s *SuspendHoloWarehouseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SuspendHoloWarehouseResponseBody) SetData(v string) *SuspendHoloWarehouseResponseBody {
	s.Data = &v
	return s
}

func (s *SuspendHoloWarehouseResponseBody) SetRequestId(v string) *SuspendHoloWarehouseResponseBody {
	s.RequestId = &v
	return s
}

func (s *SuspendHoloWarehouseResponseBody) Validate() error {
	return dara.Validate(s)
}

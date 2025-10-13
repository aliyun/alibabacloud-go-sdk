// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebalanceHoloWarehouseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *RebalanceHoloWarehouseResponseBody
	GetData() *string
	SetRequestId(v string) *RebalanceHoloWarehouseResponseBody
	GetRequestId() *string
}

type RebalanceHoloWarehouseResponseBody struct {
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
	// C0EA5844-AB00-5653-8711-CD9FD1798412
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RebalanceHoloWarehouseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RebalanceHoloWarehouseResponseBody) GoString() string {
	return s.String()
}

func (s *RebalanceHoloWarehouseResponseBody) GetData() *string {
	return s.Data
}

func (s *RebalanceHoloWarehouseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RebalanceHoloWarehouseResponseBody) SetData(v string) *RebalanceHoloWarehouseResponseBody {
	s.Data = &v
	return s
}

func (s *RebalanceHoloWarehouseResponseBody) SetRequestId(v string) *RebalanceHoloWarehouseResponseBody {
	s.RequestId = &v
	return s
}

func (s *RebalanceHoloWarehouseResponseBody) Validate() error {
	return dara.Validate(s)
}

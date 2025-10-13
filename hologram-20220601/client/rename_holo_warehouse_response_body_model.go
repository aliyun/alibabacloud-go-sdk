// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenameHoloWarehouseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *RenameHoloWarehouseResponseBody
	GetData() *bool
	SetRequestId(v string) *RenameHoloWarehouseResponseBody
	GetRequestId() *string
}

type RenameHoloWarehouseResponseBody struct {
	// Indicates whether the request was successful. Valid values: true and false.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2A8DEF6E-067E-5DB0-BAE1-2894266E6C6A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RenameHoloWarehouseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RenameHoloWarehouseResponseBody) GoString() string {
	return s.String()
}

func (s *RenameHoloWarehouseResponseBody) GetData() *bool {
	return s.Data
}

func (s *RenameHoloWarehouseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RenameHoloWarehouseResponseBody) SetData(v bool) *RenameHoloWarehouseResponseBody {
	s.Data = &v
	return s
}

func (s *RenameHoloWarehouseResponseBody) SetRequestId(v string) *RenameHoloWarehouseResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenameHoloWarehouseResponseBody) Validate() error {
	return dara.Validate(s)
}

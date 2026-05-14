// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableWarehouseAutoScaleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DisableWarehouseAutoScaleResponseBody
	GetData() *bool
	SetRequestId(v string) *DisableWarehouseAutoScaleResponseBody
	GetRequestId() *string
}

type DisableWarehouseAutoScaleResponseBody struct {
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 819A7F0F-2951-540F-BD94-6A41ECF0281F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DisableWarehouseAutoScaleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableWarehouseAutoScaleResponseBody) GoString() string {
	return s.String()
}

func (s *DisableWarehouseAutoScaleResponseBody) GetData() *bool {
	return s.Data
}

func (s *DisableWarehouseAutoScaleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableWarehouseAutoScaleResponseBody) SetData(v bool) *DisableWarehouseAutoScaleResponseBody {
	s.Data = &v
	return s
}

func (s *DisableWarehouseAutoScaleResponseBody) SetRequestId(v string) *DisableWarehouseAutoScaleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableWarehouseAutoScaleResponseBody) Validate() error {
	return dara.Validate(s)
}

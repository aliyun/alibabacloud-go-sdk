// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDrdsInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateDrdsInstanceResponseBodyData) *CreateDrdsInstanceResponseBody
	GetData() *CreateDrdsInstanceResponseBodyData
	SetRequestId(v string) *CreateDrdsInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDrdsInstanceResponseBody
	GetSuccess() *bool
}

type CreateDrdsInstanceResponseBody struct {
	// Indicates the details of the result.
	Data *CreateDrdsInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Indicates the ID of the request.
	//
	// example:
	//
	// D99D4644-A70D-49A3-B8B4-767ACC50SE2R
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDrdsInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDrdsInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDrdsInstanceResponseBody) GetData() *CreateDrdsInstanceResponseBodyData {
	return s.Data
}

func (s *CreateDrdsInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDrdsInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDrdsInstanceResponseBody) SetData(v *CreateDrdsInstanceResponseBodyData) *CreateDrdsInstanceResponseBody {
	s.Data = v
	return s
}

func (s *CreateDrdsInstanceResponseBody) SetRequestId(v string) *CreateDrdsInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDrdsInstanceResponseBody) SetSuccess(v bool) *CreateDrdsInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDrdsInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateDrdsInstanceResponseBodyData struct {
	// Indicates the ID of the instance.
	DrdsInstanceIdList *CreateDrdsInstanceResponseBodyDataDrdsInstanceIdList `json:"DrdsInstanceIdList,omitempty" xml:"DrdsInstanceIdList,omitempty" type:"Struct"`
	// Indicates the ID of the order.
	//
	// example:
	//
	// 111111111111111
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CreateDrdsInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateDrdsInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateDrdsInstanceResponseBodyData) GetDrdsInstanceIdList() *CreateDrdsInstanceResponseBodyDataDrdsInstanceIdList {
	return s.DrdsInstanceIdList
}

func (s *CreateDrdsInstanceResponseBodyData) GetOrderId() *int64 {
	return s.OrderId
}

func (s *CreateDrdsInstanceResponseBodyData) SetDrdsInstanceIdList(v *CreateDrdsInstanceResponseBodyDataDrdsInstanceIdList) *CreateDrdsInstanceResponseBodyData {
	s.DrdsInstanceIdList = v
	return s
}

func (s *CreateDrdsInstanceResponseBodyData) SetOrderId(v int64) *CreateDrdsInstanceResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *CreateDrdsInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type CreateDrdsInstanceResponseBodyDataDrdsInstanceIdList struct {
	DrdsInstanceIdList []*string `json:"drdsInstanceIdList,omitempty" xml:"drdsInstanceIdList,omitempty" type:"Repeated"`
}

func (s CreateDrdsInstanceResponseBodyDataDrdsInstanceIdList) String() string {
	return dara.Prettify(s)
}

func (s CreateDrdsInstanceResponseBodyDataDrdsInstanceIdList) GoString() string {
	return s.String()
}

func (s *CreateDrdsInstanceResponseBodyDataDrdsInstanceIdList) GetDrdsInstanceIdList() []*string {
	return s.DrdsInstanceIdList
}

func (s *CreateDrdsInstanceResponseBodyDataDrdsInstanceIdList) SetDrdsInstanceIdList(v []*string) *CreateDrdsInstanceResponseBodyDataDrdsInstanceIdList {
	s.DrdsInstanceIdList = v
	return s
}

func (s *CreateDrdsInstanceResponseBodyDataDrdsInstanceIdList) Validate() error {
	return dara.Validate(s)
}

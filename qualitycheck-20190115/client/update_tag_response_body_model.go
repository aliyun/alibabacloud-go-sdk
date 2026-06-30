// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateTagResponseBody
	GetCode() *string
	SetData(v *UpdateTagResponseBodyData) *UpdateTagResponseBody
	GetData() *UpdateTagResponseBodyData
	SetMessage(v string) *UpdateTagResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateTagResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateTagResponseBody
	GetSuccess() *bool
}

type UpdateTagResponseBody struct {
	// The result code. A value of **200*	- indicates success. Other values indicate failure. You can use this field to determine the cause of the failure.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned result.
	Data *UpdateTagResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned when an error occurs.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 3CEA0495-341B-4482-9AD9-8191EF4***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values: true: The call was successful. false: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTagResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTagResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateTagResponseBody) GetData() *UpdateTagResponseBodyData {
	return s.Data
}

func (s *UpdateTagResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTagResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateTagResponseBody) SetCode(v string) *UpdateTagResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateTagResponseBody) SetData(v *UpdateTagResponseBodyData) *UpdateTagResponseBody {
	s.Data = v
	return s
}

func (s *UpdateTagResponseBody) SetMessage(v string) *UpdateTagResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateTagResponseBody) SetRequestId(v string) *UpdateTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTagResponseBody) SetSuccess(v bool) *UpdateTagResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateTagResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateTagResponseBodyData struct {
	// The time when the label node was created.
	//
	// example:
	//
	// 1748428991000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The label description.
	//
	// example:
	//
	// 涵盖退款、退货、维修、咨询等售后链路
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the label node was last modified.
	//
	// example:
	//
	// 1748431368000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The label name.
	//
	// example:
	//
	// 售后服务
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the parent label node.
	//
	// example:
	//
	// -1
	ParentTagId *int64 `json:"ParentTagId,omitempty" xml:"ParentTagId,omitempty"`
	// The label ID.
	//
	// example:
	//
	// 128
	TagId *int64 `json:"TagId,omitempty" xml:"TagId,omitempty"`
}

func (s UpdateTagResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateTagResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateTagResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *UpdateTagResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *UpdateTagResponseBodyData) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *UpdateTagResponseBodyData) GetName() *string {
	return s.Name
}

func (s *UpdateTagResponseBodyData) GetParentTagId() *int64 {
	return s.ParentTagId
}

func (s *UpdateTagResponseBodyData) GetTagId() *int64 {
	return s.TagId
}

func (s *UpdateTagResponseBodyData) SetCreateTime(v int64) *UpdateTagResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *UpdateTagResponseBodyData) SetDescription(v string) *UpdateTagResponseBodyData {
	s.Description = &v
	return s
}

func (s *UpdateTagResponseBodyData) SetModifyTime(v int64) *UpdateTagResponseBodyData {
	s.ModifyTime = &v
	return s
}

func (s *UpdateTagResponseBodyData) SetName(v string) *UpdateTagResponseBodyData {
	s.Name = &v
	return s
}

func (s *UpdateTagResponseBodyData) SetParentTagId(v int64) *UpdateTagResponseBodyData {
	s.ParentTagId = &v
	return s
}

func (s *UpdateTagResponseBodyData) SetTagId(v int64) *UpdateTagResponseBodyData {
	s.TagId = &v
	return s
}

func (s *UpdateTagResponseBodyData) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetTagResponseBody
	GetCode() *string
	SetData(v *GetTagResponseBodyData) *GetTagResponseBody
	GetData() *GetTagResponseBodyData
	SetMessage(v string) *GetTagResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTagResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTagResponseBody
	GetSuccess() *bool
}

type GetTagResponseBody struct {
	// The result code. A value of 200 indicates success. Other values indicate failure. You can use this field to determine the cause of a failure.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned result.
	Data *GetTagResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message, if any.
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
	// Indicates whether the call was successful. Valid values:
	//
	// - true: The call was successful.
	//
	// - false: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTagResponseBody) GoString() string {
	return s.String()
}

func (s *GetTagResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTagResponseBody) GetData() *GetTagResponseBodyData {
	return s.Data
}

func (s *GetTagResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTagResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTagResponseBody) SetCode(v string) *GetTagResponseBody {
	s.Code = &v
	return s
}

func (s *GetTagResponseBody) SetData(v *GetTagResponseBodyData) *GetTagResponseBody {
	s.Data = v
	return s
}

func (s *GetTagResponseBody) SetMessage(v string) *GetTagResponseBody {
	s.Message = &v
	return s
}

func (s *GetTagResponseBody) SetRequestId(v string) *GetTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTagResponseBody) SetSuccess(v bool) *GetTagResponseBody {
	s.Success = &v
	return s
}

func (s *GetTagResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTagResponseBodyData struct {
	// The number of direct child nodes.
	//
	// example:
	//
	// 5
	ChildCount *int32 `json:"ChildCount,omitempty" xml:"ChildCount,omitempty"`
	// The time when the label was created.
	//
	// example:
	//
	// 1748428991000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The label description.
	//
	// example:
	//
	// 用于归集售后服务相关的所有意图与 FAQ
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The level of the current node.
	//
	// example:
	//
	// 2
	Level *int32 `json:"Level,omitempty" xml:"Level,omitempty"`
	// The time when the label was last modified.
	//
	// example:
	//
	// 1748428991000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The label name.
	//
	// example:
	//
	// 售后问题
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

func (s GetTagResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTagResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTagResponseBodyData) GetChildCount() *int32 {
	return s.ChildCount
}

func (s *GetTagResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetTagResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetTagResponseBodyData) GetLevel() *int32 {
	return s.Level
}

func (s *GetTagResponseBodyData) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *GetTagResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetTagResponseBodyData) GetParentTagId() *int64 {
	return s.ParentTagId
}

func (s *GetTagResponseBodyData) GetTagId() *int64 {
	return s.TagId
}

func (s *GetTagResponseBodyData) SetChildCount(v int32) *GetTagResponseBodyData {
	s.ChildCount = &v
	return s
}

func (s *GetTagResponseBodyData) SetCreateTime(v int64) *GetTagResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetTagResponseBodyData) SetDescription(v string) *GetTagResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetTagResponseBodyData) SetLevel(v int32) *GetTagResponseBodyData {
	s.Level = &v
	return s
}

func (s *GetTagResponseBodyData) SetModifyTime(v int64) *GetTagResponseBodyData {
	s.ModifyTime = &v
	return s
}

func (s *GetTagResponseBodyData) SetName(v string) *GetTagResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetTagResponseBodyData) SetParentTagId(v int64) *GetTagResponseBodyData {
	s.ParentTagId = &v
	return s
}

func (s *GetTagResponseBodyData) SetTagId(v int64) *GetTagResponseBodyData {
	s.TagId = &v
	return s
}

func (s *GetTagResponseBodyData) Validate() error {
	return dara.Validate(s)
}

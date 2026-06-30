// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateTagResponseBody
	GetCode() *string
	SetData(v *CreateTagResponseBodyData) *CreateTagResponseBody
	GetData() *CreateTagResponseBodyData
	SetMessage(v string) *CreateTagResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateTagResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateTagResponseBody
	GetSuccess() *bool
}

type CreateTagResponseBody struct {
	// The result code. A value of **200*	- indicates success. Other values indicate failure. You can use this field to determine the cause of failure.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned result.
	Data *CreateTagResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 106C6CA0-282D-4AF7-85F0-D2D24***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values: true: The call was successful. false: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTagResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTagResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateTagResponseBody) GetData() *CreateTagResponseBodyData {
	return s.Data
}

func (s *CreateTagResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTagResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateTagResponseBody) SetCode(v string) *CreateTagResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTagResponseBody) SetData(v *CreateTagResponseBodyData) *CreateTagResponseBody {
	s.Data = v
	return s
}

func (s *CreateTagResponseBody) SetMessage(v string) *CreateTagResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTagResponseBody) SetRequestId(v string) *CreateTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTagResponseBody) SetSuccess(v bool) *CreateTagResponseBody {
	s.Success = &v
	return s
}

func (s *CreateTagResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateTagResponseBodyData struct {
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
	// 用于归集售后服务相关的所有意图与 FAQ
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the label node was last modified.
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
	// The parent label node ID.
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

func (s CreateTagResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateTagResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateTagResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *CreateTagResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *CreateTagResponseBodyData) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *CreateTagResponseBodyData) GetName() *string {
	return s.Name
}

func (s *CreateTagResponseBodyData) GetParentTagId() *int64 {
	return s.ParentTagId
}

func (s *CreateTagResponseBodyData) GetTagId() *int64 {
	return s.TagId
}

func (s *CreateTagResponseBodyData) SetCreateTime(v int64) *CreateTagResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *CreateTagResponseBodyData) SetDescription(v string) *CreateTagResponseBodyData {
	s.Description = &v
	return s
}

func (s *CreateTagResponseBodyData) SetModifyTime(v int64) *CreateTagResponseBodyData {
	s.ModifyTime = &v
	return s
}

func (s *CreateTagResponseBodyData) SetName(v string) *CreateTagResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateTagResponseBodyData) SetParentTagId(v int64) *CreateTagResponseBodyData {
	s.ParentTagId = &v
	return s
}

func (s *CreateTagResponseBodyData) SetTagId(v int64) *CreateTagResponseBodyData {
	s.TagId = &v
	return s
}

func (s *CreateTagResponseBodyData) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTaskCallChatsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *TaskCallChatsResponseBody
	GetCode() *int64
	SetMessage(v string) *TaskCallChatsResponseBody
	GetMessage() *string
	SetModel(v []*TaskCallChatsResponseBodyModel) *TaskCallChatsResponseBody
	GetModel() []*TaskCallChatsResponseBodyModel
	SetRequestId(v string) *TaskCallChatsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TaskCallChatsResponseBody
	GetSuccess() *bool
	SetTimestamp(v int64) *TaskCallChatsResponseBody
	GetTimestamp() *int64
}

type TaskCallChatsResponseBody struct {
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值
	Message *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   []*TaskCallChatsResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Repeated"`
	// example:
	//
	// 8EFC6D10-307B-1ECA-A8C6-7CBDF776AAD2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 1683440860035
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s TaskCallChatsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TaskCallChatsResponseBody) GoString() string {
	return s.String()
}

func (s *TaskCallChatsResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *TaskCallChatsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TaskCallChatsResponseBody) GetModel() []*TaskCallChatsResponseBodyModel {
	return s.Model
}

func (s *TaskCallChatsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TaskCallChatsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TaskCallChatsResponseBody) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *TaskCallChatsResponseBody) SetCode(v int64) *TaskCallChatsResponseBody {
	s.Code = &v
	return s
}

func (s *TaskCallChatsResponseBody) SetMessage(v string) *TaskCallChatsResponseBody {
	s.Message = &v
	return s
}

func (s *TaskCallChatsResponseBody) SetModel(v []*TaskCallChatsResponseBodyModel) *TaskCallChatsResponseBody {
	s.Model = v
	return s
}

func (s *TaskCallChatsResponseBody) SetRequestId(v string) *TaskCallChatsResponseBody {
	s.RequestId = &v
	return s
}

func (s *TaskCallChatsResponseBody) SetSuccess(v bool) *TaskCallChatsResponseBody {
	s.Success = &v
	return s
}

func (s *TaskCallChatsResponseBody) SetTimestamp(v int64) *TaskCallChatsResponseBody {
	s.Timestamp = &v
	return s
}

func (s *TaskCallChatsResponseBody) Validate() error {
	if s.Model != nil {
		for _, item := range s.Model {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type TaskCallChatsResponseBodyModel struct {
	// 说话内容
	//
	// example:
	//
	// 示例值示例值
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// 说话时间
	//
	// example:
	//
	// 2022-01-13 14:56:46.604
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 说话号码
	//
	// example:
	//
	// 138*****265
	FromNumber *string `json:"FromNumber,omitempty" xml:"FromNumber,omitempty"`
}

func (s TaskCallChatsResponseBodyModel) String() string {
	return dara.Prettify(s)
}

func (s TaskCallChatsResponseBodyModel) GoString() string {
	return s.String()
}

func (s *TaskCallChatsResponseBodyModel) GetContent() *string {
	return s.Content
}

func (s *TaskCallChatsResponseBodyModel) GetCreateTime() *string {
	return s.CreateTime
}

func (s *TaskCallChatsResponseBodyModel) GetFromNumber() *string {
	return s.FromNumber
}

func (s *TaskCallChatsResponseBodyModel) SetContent(v string) *TaskCallChatsResponseBodyModel {
	s.Content = &v
	return s
}

func (s *TaskCallChatsResponseBodyModel) SetCreateTime(v string) *TaskCallChatsResponseBodyModel {
	s.CreateTime = &v
	return s
}

func (s *TaskCallChatsResponseBodyModel) SetFromNumber(v string) *TaskCallChatsResponseBodyModel {
	s.FromNumber = &v
	return s
}

func (s *TaskCallChatsResponseBodyModel) Validate() error {
	return dara.Validate(s)
}

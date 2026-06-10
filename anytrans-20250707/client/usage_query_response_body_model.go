// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUsageQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UsageQueryResponseBody
	GetCode() *string
	SetData(v []*UsageQueryResponseBodyData) *UsageQueryResponseBody
	GetData() []*UsageQueryResponseBodyData
	SetHttpStatusCode(v string) *UsageQueryResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *UsageQueryResponseBody
	GetMessage() *string
	SetRequestId(v string) *UsageQueryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UsageQueryResponseBody
	GetSuccess() *bool
}

type UsageQueryResponseBody struct {
	// example:
	//
	// success
	Code *string                       `json:"code,omitempty" xml:"code,omitempty"`
	Data []*UsageQueryResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 42FF90E5-5D40-5797-AAF6-8A4D837CCCD5
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UsageQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UsageQueryResponseBody) GoString() string {
	return s.String()
}

func (s *UsageQueryResponseBody) GetCode() *string {
	return s.Code
}

func (s *UsageQueryResponseBody) GetData() []*UsageQueryResponseBodyData {
	return s.Data
}

func (s *UsageQueryResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *UsageQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UsageQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UsageQueryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UsageQueryResponseBody) SetCode(v string) *UsageQueryResponseBody {
	s.Code = &v
	return s
}

func (s *UsageQueryResponseBody) SetData(v []*UsageQueryResponseBodyData) *UsageQueryResponseBody {
	s.Data = v
	return s
}

func (s *UsageQueryResponseBody) SetHttpStatusCode(v string) *UsageQueryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UsageQueryResponseBody) SetMessage(v string) *UsageQueryResponseBody {
	s.Message = &v
	return s
}

func (s *UsageQueryResponseBody) SetRequestId(v string) *UsageQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *UsageQueryResponseBody) SetSuccess(v bool) *UsageQueryResponseBody {
	s.Success = &v
	return s
}

func (s *UsageQueryResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UsageQueryResponseBodyData struct {
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 533
	InputTokens *int64 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	// example:
	//
	// 38
	OutputTokens *int64 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	// example:
	//
	// d3a2397bc2c14ab4a2e40a4f5b46241b
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// 535
	TotalTokens *int64 `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
	UpdateTime  *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s UsageQueryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UsageQueryResponseBodyData) GoString() string {
	return s.String()
}

func (s *UsageQueryResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *UsageQueryResponseBodyData) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *UsageQueryResponseBodyData) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *UsageQueryResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *UsageQueryResponseBodyData) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *UsageQueryResponseBodyData) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *UsageQueryResponseBodyData) SetCreateTime(v int64) *UsageQueryResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *UsageQueryResponseBodyData) SetInputTokens(v int64) *UsageQueryResponseBodyData {
	s.InputTokens = &v
	return s
}

func (s *UsageQueryResponseBodyData) SetOutputTokens(v int64) *UsageQueryResponseBodyData {
	s.OutputTokens = &v
	return s
}

func (s *UsageQueryResponseBodyData) SetTaskId(v string) *UsageQueryResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *UsageQueryResponseBodyData) SetTotalTokens(v int64) *UsageQueryResponseBodyData {
	s.TotalTokens = &v
	return s
}

func (s *UsageQueryResponseBodyData) SetUpdateTime(v int64) *UsageQueryResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *UsageQueryResponseBodyData) Validate() error {
	return dara.Validate(s)
}

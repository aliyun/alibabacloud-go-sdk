// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStyleLearningResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListStyleLearningResultResponseBody
	GetCode() *string
	SetCurrent(v int32) *ListStyleLearningResultResponseBody
	GetCurrent() *int32
	SetData(v []*ListStyleLearningResultResponseBodyData) *ListStyleLearningResultResponseBody
	GetData() []*ListStyleLearningResultResponseBodyData
	SetHttpStatusCode(v int32) *ListStyleLearningResultResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListStyleLearningResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListStyleLearningResultResponseBody
	GetRequestId() *string
	SetSize(v int32) *ListStyleLearningResultResponseBody
	GetSize() *int32
	SetSuccess(v bool) *ListStyleLearningResultResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *ListStyleLearningResultResponseBody
	GetTotal() *int32
}

type ListStyleLearningResultResponseBody struct {
	// example:
	//
	// NoData
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1
	Current *int32                                     `json:"Current,omitempty" xml:"Current,omitempty"`
	Data    []*ListStyleLearningResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 100
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListStyleLearningResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListStyleLearningResultResponseBody) GoString() string {
	return s.String()
}

func (s *ListStyleLearningResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListStyleLearningResultResponseBody) GetCurrent() *int32 {
	return s.Current
}

func (s *ListStyleLearningResultResponseBody) GetData() []*ListStyleLearningResultResponseBodyData {
	return s.Data
}

func (s *ListStyleLearningResultResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListStyleLearningResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListStyleLearningResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListStyleLearningResultResponseBody) GetSize() *int32 {
	return s.Size
}

func (s *ListStyleLearningResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListStyleLearningResultResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListStyleLearningResultResponseBody) SetCode(v string) *ListStyleLearningResultResponseBody {
	s.Code = &v
	return s
}

func (s *ListStyleLearningResultResponseBody) SetCurrent(v int32) *ListStyleLearningResultResponseBody {
	s.Current = &v
	return s
}

func (s *ListStyleLearningResultResponseBody) SetData(v []*ListStyleLearningResultResponseBodyData) *ListStyleLearningResultResponseBody {
	s.Data = v
	return s
}

func (s *ListStyleLearningResultResponseBody) SetHttpStatusCode(v int32) *ListStyleLearningResultResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListStyleLearningResultResponseBody) SetMessage(v string) *ListStyleLearningResultResponseBody {
	s.Message = &v
	return s
}

func (s *ListStyleLearningResultResponseBody) SetRequestId(v string) *ListStyleLearningResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListStyleLearningResultResponseBody) SetSize(v int32) *ListStyleLearningResultResponseBody {
	s.Size = &v
	return s
}

func (s *ListStyleLearningResultResponseBody) SetSuccess(v bool) *ListStyleLearningResultResponseBody {
	s.Success = &v
	return s
}

func (s *ListStyleLearningResultResponseBody) SetTotal(v int32) *ListStyleLearningResultResponseBody {
	s.Total = &v
	return s
}

func (s *ListStyleLearningResultResponseBody) Validate() error {
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

type ListStyleLearningResultResponseBodyData struct {
	// example:
	//
	// AIGC 生成的内容
	AigcResult *string `json:"AigcResult,omitempty" xml:"AigcResult,omitempty"`
	// example:
	//
	// 70
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 用户修订后内容
	RewriteResult *string `json:"RewriteResult,omitempty" xml:"RewriteResult,omitempty"`
	// example:
	//
	// 文体风格名称
	StyleName *string `json:"StyleName,omitempty" xml:"StyleName,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ListStyleLearningResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListStyleLearningResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListStyleLearningResultResponseBodyData) GetAigcResult() *string {
	return s.AigcResult
}

func (s *ListStyleLearningResultResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *ListStyleLearningResultResponseBodyData) GetRewriteResult() *string {
	return s.RewriteResult
}

func (s *ListStyleLearningResultResponseBodyData) GetStyleName() *string {
	return s.StyleName
}

func (s *ListStyleLearningResultResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *ListStyleLearningResultResponseBodyData) SetAigcResult(v string) *ListStyleLearningResultResponseBodyData {
	s.AigcResult = &v
	return s
}

func (s *ListStyleLearningResultResponseBodyData) SetId(v int64) *ListStyleLearningResultResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListStyleLearningResultResponseBodyData) SetRewriteResult(v string) *ListStyleLearningResultResponseBodyData {
	s.RewriteResult = &v
	return s
}

func (s *ListStyleLearningResultResponseBodyData) SetStyleName(v string) *ListStyleLearningResultResponseBodyData {
	s.StyleName = &v
	return s
}

func (s *ListStyleLearningResultResponseBodyData) SetTaskId(v string) *ListStyleLearningResultResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ListStyleLearningResultResponseBodyData) Validate() error {
	return dara.Validate(s)
}

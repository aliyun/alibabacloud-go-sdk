// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocClusterTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDocClusterTaskResponseBody
	GetCode() *string
	SetData(v *GetDocClusterTaskResponseBodyData) *GetDocClusterTaskResponseBody
	GetData() *GetDocClusterTaskResponseBodyData
	SetHttpStatusCode(v int32) *GetDocClusterTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetDocClusterTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDocClusterTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDocClusterTaskResponseBody
	GetSuccess() *bool
}

type GetDocClusterTaskResponseBody struct {
	// example:
	//
	// NoData
	Code *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetDocClusterTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDocClusterTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDocClusterTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetDocClusterTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDocClusterTaskResponseBody) GetData() *GetDocClusterTaskResponseBodyData {
	return s.Data
}

func (s *GetDocClusterTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetDocClusterTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDocClusterTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDocClusterTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDocClusterTaskResponseBody) SetCode(v string) *GetDocClusterTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetDocClusterTaskResponseBody) SetData(v *GetDocClusterTaskResponseBodyData) *GetDocClusterTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetDocClusterTaskResponseBody) SetHttpStatusCode(v int32) *GetDocClusterTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDocClusterTaskResponseBody) SetMessage(v string) *GetDocClusterTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetDocClusterTaskResponseBody) SetRequestId(v string) *GetDocClusterTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDocClusterTaskResponseBody) SetSuccess(v bool) *GetDocClusterTaskResponseBody {
	s.Success = &v
	return s
}

func (s *GetDocClusterTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDocClusterTaskResponseBodyData struct {
	// example:
	//
	// 错误信息
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// PENDING
	Status *string                                    `json:"Status,omitempty" xml:"Status,omitempty"`
	Topics []*GetDocClusterTaskResponseBodyDataTopics `json:"Topics,omitempty" xml:"Topics,omitempty" type:"Repeated"`
}

func (s GetDocClusterTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDocClusterTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDocClusterTaskResponseBodyData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetDocClusterTaskResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetDocClusterTaskResponseBodyData) GetTopics() []*GetDocClusterTaskResponseBodyDataTopics {
	return s.Topics
}

func (s *GetDocClusterTaskResponseBodyData) SetErrorMessage(v string) *GetDocClusterTaskResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *GetDocClusterTaskResponseBodyData) SetStatus(v string) *GetDocClusterTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetDocClusterTaskResponseBodyData) SetTopics(v []*GetDocClusterTaskResponseBodyDataTopics) *GetDocClusterTaskResponseBodyData {
	s.Topics = v
	return s
}

func (s *GetDocClusterTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetDocClusterTaskResponseBodyDataTopics struct {
	DocIds []*string `json:"DocIds,omitempty" xml:"DocIds,omitempty" type:"Repeated"`
	// example:
	//
	// 聚类主题摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// example:
	//
	// 聚类主题名
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetDocClusterTaskResponseBodyDataTopics) String() string {
	return dara.Prettify(s)
}

func (s GetDocClusterTaskResponseBodyDataTopics) GoString() string {
	return s.String()
}

func (s *GetDocClusterTaskResponseBodyDataTopics) GetDocIds() []*string {
	return s.DocIds
}

func (s *GetDocClusterTaskResponseBodyDataTopics) GetSummary() *string {
	return s.Summary
}

func (s *GetDocClusterTaskResponseBodyDataTopics) GetTitle() *string {
	return s.Title
}

func (s *GetDocClusterTaskResponseBodyDataTopics) SetDocIds(v []*string) *GetDocClusterTaskResponseBodyDataTopics {
	s.DocIds = v
	return s
}

func (s *GetDocClusterTaskResponseBodyDataTopics) SetSummary(v string) *GetDocClusterTaskResponseBodyDataTopics {
	s.Summary = &v
	return s
}

func (s *GetDocClusterTaskResponseBodyDataTopics) SetTitle(v string) *GetDocClusterTaskResponseBodyDataTopics {
	s.Title = &v
	return s
}

func (s *GetDocClusterTaskResponseBodyDataTopics) Validate() error {
	return dara.Validate(s)
}

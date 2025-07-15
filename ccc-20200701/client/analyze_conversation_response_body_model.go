// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAnalyzeConversationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AnalyzeConversationResponseBody
	GetCode() *string
	SetData(v string) *AnalyzeConversationResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *AnalyzeConversationResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AnalyzeConversationResponseBody
	GetMessage() *string
	SetRequestId(v string) *AnalyzeConversationResponseBody
	GetRequestId() *string
}

type AnalyzeConversationResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9F766284-F103-4298-8EC5-19F9F9BE5522
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AnalyzeConversationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AnalyzeConversationResponseBody) GoString() string {
	return s.String()
}

func (s *AnalyzeConversationResponseBody) GetCode() *string {
	return s.Code
}

func (s *AnalyzeConversationResponseBody) GetData() *string {
	return s.Data
}

func (s *AnalyzeConversationResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AnalyzeConversationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AnalyzeConversationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AnalyzeConversationResponseBody) SetCode(v string) *AnalyzeConversationResponseBody {
	s.Code = &v
	return s
}

func (s *AnalyzeConversationResponseBody) SetData(v string) *AnalyzeConversationResponseBody {
	s.Data = &v
	return s
}

func (s *AnalyzeConversationResponseBody) SetHttpStatusCode(v int32) *AnalyzeConversationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AnalyzeConversationResponseBody) SetMessage(v string) *AnalyzeConversationResponseBody {
	s.Message = &v
	return s
}

func (s *AnalyzeConversationResponseBody) SetRequestId(v string) *AnalyzeConversationResponseBody {
	s.RequestId = &v
	return s
}

func (s *AnalyzeConversationResponseBody) Validate() error {
	return dara.Validate(s)
}

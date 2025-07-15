// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartEditContactFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StartEditContactFlowResponseBody
	GetCode() *string
	SetData(v string) *StartEditContactFlowResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *StartEditContactFlowResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *StartEditContactFlowResponseBody
	GetMessage() *string
	SetRequestId(v string) *StartEditContactFlowResponseBody
	GetRequestId() *string
}

type StartEditContactFlowResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// b3114362-9062-46c7-82dc-ae55ac168b2e
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// BC2733CE-C470-564A-8C11-9DC02468823A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartEditContactFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartEditContactFlowResponseBody) GoString() string {
	return s.String()
}

func (s *StartEditContactFlowResponseBody) GetCode() *string {
	return s.Code
}

func (s *StartEditContactFlowResponseBody) GetData() *string {
	return s.Data
}

func (s *StartEditContactFlowResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *StartEditContactFlowResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StartEditContactFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartEditContactFlowResponseBody) SetCode(v string) *StartEditContactFlowResponseBody {
	s.Code = &v
	return s
}

func (s *StartEditContactFlowResponseBody) SetData(v string) *StartEditContactFlowResponseBody {
	s.Data = &v
	return s
}

func (s *StartEditContactFlowResponseBody) SetHttpStatusCode(v int32) *StartEditContactFlowResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StartEditContactFlowResponseBody) SetMessage(v string) *StartEditContactFlowResponseBody {
	s.Message = &v
	return s
}

func (s *StartEditContactFlowResponseBody) SetRequestId(v string) *StartEditContactFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartEditContactFlowResponseBody) Validate() error {
	return dara.Validate(s)
}

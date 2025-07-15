// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCommitContactFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CommitContactFlowResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *CommitContactFlowResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CommitContactFlowResponseBody
	GetMessage() *string
	SetRequestId(v string) *CommitContactFlowResponseBody
	GetRequestId() *string
}

type CommitContactFlowResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 937617D5-01E9-5A39-B52D-15D5C143260C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CommitContactFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CommitContactFlowResponseBody) GoString() string {
	return s.String()
}

func (s *CommitContactFlowResponseBody) GetCode() *string {
	return s.Code
}

func (s *CommitContactFlowResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CommitContactFlowResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CommitContactFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CommitContactFlowResponseBody) SetCode(v string) *CommitContactFlowResponseBody {
	s.Code = &v
	return s
}

func (s *CommitContactFlowResponseBody) SetHttpStatusCode(v int32) *CommitContactFlowResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CommitContactFlowResponseBody) SetMessage(v string) *CommitContactFlowResponseBody {
	s.Message = &v
	return s
}

func (s *CommitContactFlowResponseBody) SetRequestId(v string) *CommitContactFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *CommitContactFlowResponseBody) Validate() error {
	return dara.Validate(s)
}

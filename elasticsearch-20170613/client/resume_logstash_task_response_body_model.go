// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeLogstashTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ResumeLogstashTaskResponseBody
	GetCode() *string
	SetMessage(v string) *ResumeLogstashTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *ResumeLogstashTaskResponseBody
	GetRequestId() *string
	SetResult(v bool) *ResumeLogstashTaskResponseBody
	GetResult() *bool
}

type ResumeLogstashTaskResponseBody struct {
	// The error code. This parameter is not returned if the call is successful.
	//
	// example:
	//
	// InstanceNotFound
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message. This parameter is not returned if the call is successful.
	//
	// example:
	//
	// The specified cluster does not exist. Check the cluster status and try again.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0FA05123-745C-42FD-A69B-AFF48EF9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result. Valid values:
	//
	// - true: The task is resumed.
	//
	// - false: The task failed to be resumed.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s ResumeLogstashTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResumeLogstashTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ResumeLogstashTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *ResumeLogstashTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ResumeLogstashTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResumeLogstashTaskResponseBody) GetResult() *bool {
	return s.Result
}

func (s *ResumeLogstashTaskResponseBody) SetCode(v string) *ResumeLogstashTaskResponseBody {
	s.Code = &v
	return s
}

func (s *ResumeLogstashTaskResponseBody) SetMessage(v string) *ResumeLogstashTaskResponseBody {
	s.Message = &v
	return s
}

func (s *ResumeLogstashTaskResponseBody) SetRequestId(v string) *ResumeLogstashTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResumeLogstashTaskResponseBody) SetResult(v bool) *ResumeLogstashTaskResponseBody {
	s.Result = &v
	return s
}

func (s *ResumeLogstashTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

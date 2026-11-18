// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitVideoRemakeJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *SubmitVideoRemakeJobResponseBody
	GetErrorCode() *string
	SetJobId(v string) *SubmitVideoRemakeJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *SubmitVideoRemakeJobResponseBody
	GetRequestId() *string
}

type SubmitVideoRemakeJobResponseBody struct {
	// The error code.
	//
	// example:
	//
	// InvalidParameter
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The job ID.
	//
	// example:
	//
	// 0123456789abcdef0123456789abcdef
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 12345678-1234-4321-ABCD-123456789012
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitVideoRemakeJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitVideoRemakeJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitVideoRemakeJobResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SubmitVideoRemakeJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *SubmitVideoRemakeJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitVideoRemakeJobResponseBody) SetErrorCode(v string) *SubmitVideoRemakeJobResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SubmitVideoRemakeJobResponseBody) SetJobId(v string) *SubmitVideoRemakeJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitVideoRemakeJobResponseBody) SetRequestId(v string) *SubmitVideoRemakeJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitVideoRemakeJobResponseBody) Validate() error {
	return dara.Validate(s)
}

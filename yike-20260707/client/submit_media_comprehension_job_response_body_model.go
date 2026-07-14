// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitMediaComprehensionJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *SubmitMediaComprehensionJobResponseBody
	GetErrorCode() *string
	SetJobId(v string) *SubmitMediaComprehensionJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *SubmitMediaComprehensionJobResponseBody
	GetRequestId() *string
}

type SubmitMediaComprehensionJobResponseBody struct {
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitMediaComprehensionJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitMediaComprehensionJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitMediaComprehensionJobResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SubmitMediaComprehensionJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *SubmitMediaComprehensionJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitMediaComprehensionJobResponseBody) SetErrorCode(v string) *SubmitMediaComprehensionJobResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SubmitMediaComprehensionJobResponseBody) SetJobId(v string) *SubmitMediaComprehensionJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitMediaComprehensionJobResponseBody) SetRequestId(v string) *SubmitMediaComprehensionJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitMediaComprehensionJobResponseBody) Validate() error {
	return dara.Validate(s)
}

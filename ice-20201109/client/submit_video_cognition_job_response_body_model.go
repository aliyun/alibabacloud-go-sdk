// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitVideoCognitionJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *SubmitVideoCognitionJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *SubmitVideoCognitionJobResponseBody
	GetRequestId() *string
}

type SubmitVideoCognitionJobResponseBody struct {
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitVideoCognitionJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitVideoCognitionJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitVideoCognitionJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *SubmitVideoCognitionJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitVideoCognitionJobResponseBody) SetJobId(v string) *SubmitVideoCognitionJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitVideoCognitionJobResponseBody) SetRequestId(v string) *SubmitVideoCognitionJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitVideoCognitionJobResponseBody) Validate() error {
	return dara.Validate(s)
}

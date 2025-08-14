// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDynamicImageJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *SubmitDynamicImageJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *SubmitDynamicImageJobResponseBody
	GetRequestId() *string
}

type SubmitDynamicImageJobResponseBody struct {
	// The job ID.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitDynamicImageJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitDynamicImageJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitDynamicImageJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *SubmitDynamicImageJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitDynamicImageJobResponseBody) SetJobId(v string) *SubmitDynamicImageJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitDynamicImageJobResponseBody) SetRequestId(v string) *SubmitDynamicImageJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitDynamicImageJobResponseBody) Validate() error {
	return dara.Validate(s)
}

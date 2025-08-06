// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAIImageJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *SubmitAIImageJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *SubmitAIImageJobResponseBody
	GetRequestId() *string
}

type SubmitAIImageJobResponseBody struct {
	// The ID of the image AI processing job.
	//
	// example:
	//
	// cf08a2c6e11e*****de1711b738b9067
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 218A6807-A21E-43*****54-C0512880B0B0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitAIImageJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIImageJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitAIImageJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *SubmitAIImageJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitAIImageJobResponseBody) SetJobId(v string) *SubmitAIImageJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitAIImageJobResponseBody) SetRequestId(v string) *SubmitAIImageJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitAIImageJobResponseBody) Validate() error {
	return dara.Validate(s)
}

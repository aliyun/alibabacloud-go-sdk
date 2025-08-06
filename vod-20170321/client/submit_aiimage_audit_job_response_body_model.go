// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAIImageAuditJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *SubmitAIImageAuditJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *SubmitAIImageAuditJobResponseBody
	GetRequestId() *string
}

type SubmitAIImageAuditJobResponseBody struct {
	// The ID of the image review job.
	//
	// example:
	//
	// b1aa3024aee64*****6dc8ca20dbc328
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 6F42D500-1956-4B*****30-C09E755F4F4B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitAIImageAuditJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIImageAuditJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitAIImageAuditJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *SubmitAIImageAuditJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitAIImageAuditJobResponseBody) SetJobId(v string) *SubmitAIImageAuditJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitAIImageAuditJobResponseBody) SetRequestId(v string) *SubmitAIImageAuditJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitAIImageAuditJobResponseBody) Validate() error {
	return dara.Validate(s)
}

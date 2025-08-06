// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAIMediaAuditJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *SubmitAIMediaAuditJobResponseBody
	GetJobId() *string
	SetMediaId(v string) *SubmitAIMediaAuditJobResponseBody
	GetMediaId() *string
	SetRequestId(v string) *SubmitAIMediaAuditJobResponseBody
	GetRequestId() *string
}

type SubmitAIMediaAuditJobResponseBody struct {
	// The ID of the job.
	//
	// example:
	//
	// bdbc266af6893943a70176d92e99****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the media file.
	//
	// example:
	//
	// fe028d09441afffb138cd7ee****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// F4631053-8D9F-42B2-4A67281DB88E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitAIMediaAuditJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIMediaAuditJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitAIMediaAuditJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *SubmitAIMediaAuditJobResponseBody) GetMediaId() *string {
	return s.MediaId
}

func (s *SubmitAIMediaAuditJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitAIMediaAuditJobResponseBody) SetJobId(v string) *SubmitAIMediaAuditJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitAIMediaAuditJobResponseBody) SetMediaId(v string) *SubmitAIMediaAuditJobResponseBody {
	s.MediaId = &v
	return s
}

func (s *SubmitAIMediaAuditJobResponseBody) SetRequestId(v string) *SubmitAIMediaAuditJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitAIMediaAuditJobResponseBody) Validate() error {
	return dara.Validate(s)
}

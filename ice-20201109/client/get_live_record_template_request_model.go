// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLiveRecordTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *GetLiveRecordTemplateRequest
	GetJobId() *string
	SetTemplateId(v string) *GetLiveRecordTemplateRequest
	GetTemplateId() *string
}

type GetLiveRecordTemplateRequest struct {
	// The ID of the recording job. You can specify the JobId parameter to retrieve the snapshot of the template used by the job.
	//
	// example:
	//
	// ab0e3e76-1e9d-11ed-ba64-0c42a1b73d66
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 69e1f9fe-1e97-11ed-ba64-0c42a1b73d66
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s GetLiveRecordTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLiveRecordTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetLiveRecordTemplateRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetLiveRecordTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetLiveRecordTemplateRequest) SetJobId(v string) *GetLiveRecordTemplateRequest {
	s.JobId = &v
	return s
}

func (s *GetLiveRecordTemplateRequest) SetTemplateId(v string) *GetLiveRecordTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *GetLiveRecordTemplateRequest) Validate() error {
	return dara.Validate(s)
}

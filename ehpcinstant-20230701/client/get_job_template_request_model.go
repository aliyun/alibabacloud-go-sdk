// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobTemplateId(v string) *GetJobTemplateRequest
	GetJobTemplateId() *string
}

type GetJobTemplateRequest struct {
	// The job template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// jt-xxxx
	JobTemplateId *string `json:"JobTemplateId,omitempty" xml:"JobTemplateId,omitempty"`
}

func (s GetJobTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetJobTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetJobTemplateRequest) GetJobTemplateId() *string {
	return s.JobTemplateId
}

func (s *GetJobTemplateRequest) SetJobTemplateId(v string) *GetJobTemplateRequest {
	s.JobTemplateId = &v
	return s
}

func (s *GetJobTemplateRequest) Validate() error {
	return dara.Validate(s)
}

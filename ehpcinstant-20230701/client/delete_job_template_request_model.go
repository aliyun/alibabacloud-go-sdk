// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteJobTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobTemplateId(v string) *DeleteJobTemplateRequest
	GetJobTemplateId() *string
}

type DeleteJobTemplateRequest struct {
	// The ID of the job template.
	//
	// This parameter is required.
	//
	// example:
	//
	// jt-xxxx
	JobTemplateId *string `json:"JobTemplateId,omitempty" xml:"JobTemplateId,omitempty"`
}

func (s DeleteJobTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteJobTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteJobTemplateRequest) GetJobTemplateId() *string {
	return s.JobTemplateId
}

func (s *DeleteJobTemplateRequest) SetJobTemplateId(v string) *DeleteJobTemplateRequest {
	s.JobTemplateId = &v
	return s
}

func (s *DeleteJobTemplateRequest) Validate() error {
	return dara.Validate(s)
}

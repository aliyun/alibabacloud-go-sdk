// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateJobTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobTemplateId(v string) *CreateJobTemplateResponseBody
	GetJobTemplateId() *string
	SetRequestId(v string) *CreateJobTemplateResponseBody
	GetRequestId() *string
}

type CreateJobTemplateResponseBody struct {
	// The job template ID.
	//
	// example:
	//
	// jt-xxxx
	JobTemplateId *string `json:"JobTemplateId,omitempty" xml:"JobTemplateId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 896D338C-E4F4-41EC-A154-D605E5DE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateJobTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateJobTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateJobTemplateResponseBody) GetJobTemplateId() *string {
	return s.JobTemplateId
}

func (s *CreateJobTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateJobTemplateResponseBody) SetJobTemplateId(v string) *CreateJobTemplateResponseBody {
	s.JobTemplateId = &v
	return s
}

func (s *CreateJobTemplateResponseBody) SetRequestId(v string) *CreateJobTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateJobTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataQualityTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwner(v string) *CreateDataQualityTemplateRequest
	GetOwner() *string
	SetProjectId(v int64) *CreateDataQualityTemplateRequest
	GetProjectId() *int64
	SetSpec(v string) *CreateDataQualityTemplateRequest
	GetSpec() *string
}

type CreateDataQualityTemplateRequest struct {
	Owner     *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	ProjectId *int64  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Spec      *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s CreateDataQualityTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateDataQualityTemplateRequest) GetOwner() *string {
	return s.Owner
}

func (s *CreateDataQualityTemplateRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateDataQualityTemplateRequest) GetSpec() *string {
	return s.Spec
}

func (s *CreateDataQualityTemplateRequest) SetOwner(v string) *CreateDataQualityTemplateRequest {
	s.Owner = &v
	return s
}

func (s *CreateDataQualityTemplateRequest) SetProjectId(v int64) *CreateDataQualityTemplateRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateDataQualityTemplateRequest) SetSpec(v string) *CreateDataQualityTemplateRequest {
	s.Spec = &v
	return s
}

func (s *CreateDataQualityTemplateRequest) Validate() error {
	return dara.Validate(s)
}

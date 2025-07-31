// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataQualityTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *UpdateDataQualityTemplateRequest
	GetId() *string
	SetOwner(v string) *UpdateDataQualityTemplateRequest
	GetOwner() *string
	SetProjectId(v int64) *UpdateDataQualityTemplateRequest
	GetProjectId() *int64
	SetSpec(v string) *UpdateDataQualityTemplateRequest
	GetSpec() *string
}

type UpdateDataQualityTemplateRequest struct {
	// example:
	//
	// USER_DEFINED:2001
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 95279527****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 100001
	ProjectId *int64  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Spec      *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s UpdateDataQualityTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataQualityTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataQualityTemplateRequest) GetId() *string {
	return s.Id
}

func (s *UpdateDataQualityTemplateRequest) GetOwner() *string {
	return s.Owner
}

func (s *UpdateDataQualityTemplateRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateDataQualityTemplateRequest) GetSpec() *string {
	return s.Spec
}

func (s *UpdateDataQualityTemplateRequest) SetId(v string) *UpdateDataQualityTemplateRequest {
	s.Id = &v
	return s
}

func (s *UpdateDataQualityTemplateRequest) SetOwner(v string) *UpdateDataQualityTemplateRequest {
	s.Owner = &v
	return s
}

func (s *UpdateDataQualityTemplateRequest) SetProjectId(v int64) *UpdateDataQualityTemplateRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateDataQualityTemplateRequest) SetSpec(v string) *UpdateDataQualityTemplateRequest {
	s.Spec = &v
	return s
}

func (s *UpdateDataQualityTemplateRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTemplateRevision interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedTime(v string) *TemplateRevision
	GetCreatedTime() *string
	SetDescription(v string) *TemplateRevision
	GetDescription() *string
	SetKind(v string) *TemplateRevision
	GetKind() *string
	SetLabels(v map[string]*string) *TemplateRevision
	GetLabels() map[string]*string
	SetName(v string) *TemplateRevision
	GetName() *string
	SetSpec(v *TemplateSpec) *TemplateRevision
	GetSpec() *TemplateSpec
	SetStatus(v *TemplateRevisionStatus) *TemplateRevision
	GetStatus() *TemplateRevisionStatus
	SetTemplateName(v string) *TemplateRevision
	GetTemplateName() *string
	SetUid(v string) *TemplateRevision
	GetUid() *string
}

type TemplateRevision struct {
	// example:
	//
	// 2021-11-19T09:34:38Z
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// example:
	//
	// It is a template revision
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// TemplateRevision
	Kind   *string            `json:"kind,omitempty" xml:"kind,omitempty"`
	Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
	// example:
	//
	// demo-template
	Name   *string                 `json:"name,omitempty" xml:"name,omitempty"`
	Spec   *TemplateSpec           `json:"spec,omitempty" xml:"spec,omitempty"`
	Status *TemplateRevisionStatus `json:"status,omitempty" xml:"status,omitempty" type:"Struct"`
	// example:
	//
	// demo-template
	TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
	// example:
	//
	// 1455541096***548
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s TemplateRevision) String() string {
	return dara.Prettify(s)
}

func (s TemplateRevision) GoString() string {
	return s.String()
}

func (s *TemplateRevision) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *TemplateRevision) GetDescription() *string {
	return s.Description
}

func (s *TemplateRevision) GetKind() *string {
	return s.Kind
}

func (s *TemplateRevision) GetLabels() map[string]*string {
	return s.Labels
}

func (s *TemplateRevision) GetName() *string {
	return s.Name
}

func (s *TemplateRevision) GetSpec() *TemplateSpec {
	return s.Spec
}

func (s *TemplateRevision) GetStatus() *TemplateRevisionStatus {
	return s.Status
}

func (s *TemplateRevision) GetTemplateName() *string {
	return s.TemplateName
}

func (s *TemplateRevision) GetUid() *string {
	return s.Uid
}

func (s *TemplateRevision) SetCreatedTime(v string) *TemplateRevision {
	s.CreatedTime = &v
	return s
}

func (s *TemplateRevision) SetDescription(v string) *TemplateRevision {
	s.Description = &v
	return s
}

func (s *TemplateRevision) SetKind(v string) *TemplateRevision {
	s.Kind = &v
	return s
}

func (s *TemplateRevision) SetLabels(v map[string]*string) *TemplateRevision {
	s.Labels = v
	return s
}

func (s *TemplateRevision) SetName(v string) *TemplateRevision {
	s.Name = &v
	return s
}

func (s *TemplateRevision) SetSpec(v *TemplateSpec) *TemplateRevision {
	s.Spec = v
	return s
}

func (s *TemplateRevision) SetStatus(v *TemplateRevisionStatus) *TemplateRevision {
	s.Status = v
	return s
}

func (s *TemplateRevision) SetTemplateName(v string) *TemplateRevision {
	s.TemplateName = &v
	return s
}

func (s *TemplateRevision) SetUid(v string) *TemplateRevision {
	s.Uid = &v
	return s
}

func (s *TemplateRevision) Validate() error {
	return dara.Validate(s)
}

type TemplateRevisionStatus struct {
	// example:
	//
	// https://registry.serverless-devs.com/details.html?name=template-test&package_type=v3
	PackageUrl *string `json:"packageUrl,omitempty" xml:"packageUrl,omitempty"`
	// example:
	//
	// Published
	Phase *string `json:"phase,omitempty" xml:"phase,omitempty"`
	// example:
	//
	// p-default
	PipelineName *string `json:"pipelineName,omitempty" xml:"pipelineName,omitempty"`
	// example:
	//
	// https://cap.console.aliyun.com/template-detail?template=adasdasdaewe-adadqwe
	TemplateUrl *string `json:"templateUrl,omitempty" xml:"templateUrl,omitempty"`
}

func (s TemplateRevisionStatus) String() string {
	return dara.Prettify(s)
}

func (s TemplateRevisionStatus) GoString() string {
	return s.String()
}

func (s *TemplateRevisionStatus) GetPackageUrl() *string {
	return s.PackageUrl
}

func (s *TemplateRevisionStatus) GetPhase() *string {
	return s.Phase
}

func (s *TemplateRevisionStatus) GetPipelineName() *string {
	return s.PipelineName
}

func (s *TemplateRevisionStatus) GetTemplateUrl() *string {
	return s.TemplateUrl
}

func (s *TemplateRevisionStatus) SetPackageUrl(v string) *TemplateRevisionStatus {
	s.PackageUrl = &v
	return s
}

func (s *TemplateRevisionStatus) SetPhase(v string) *TemplateRevisionStatus {
	s.Phase = &v
	return s
}

func (s *TemplateRevisionStatus) SetPipelineName(v string) *TemplateRevisionStatus {
	s.PipelineName = &v
	return s
}

func (s *TemplateRevisionStatus) SetTemplateUrl(v string) *TemplateRevisionStatus {
	s.TemplateUrl = &v
	return s
}

func (s *TemplateRevisionStatus) Validate() error {
	return dara.Validate(s)
}

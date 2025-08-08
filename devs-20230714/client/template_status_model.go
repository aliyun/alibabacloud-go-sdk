// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTemplateStatus interface {
	dara.Model
	String() string
	GoString() string
	SetLatestDeployment(v *TemplateStatusLatestDeployment) *TemplateStatus
	GetLatestDeployment() *TemplateStatusLatestDeployment
	SetLatestVersion(v string) *TemplateStatus
	GetLatestVersion() *string
	SetPackageUrl(v string) *TemplateStatus
	GetPackageUrl() *string
	SetPhase(v string) *TemplateStatus
	GetPhase() *string
	SetTemplateUrl(v string) *TemplateStatus
	GetTemplateUrl() *string
}

type TemplateStatus struct {
	LatestDeployment *TemplateStatusLatestDeployment `json:"latestDeployment,omitempty" xml:"latestDeployment,omitempty" type:"Struct"`
	// example:
	//
	// 1.0.0
	LatestVersion *string `json:"latestVersion,omitempty" xml:"latestVersion,omitempty"`
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
	// https://cap.console.aliyun.com/template-detail?template=adasdasdaewe-adadqwe
	TemplateUrl *string `json:"templateUrl,omitempty" xml:"templateUrl,omitempty"`
}

func (s TemplateStatus) String() string {
	return dara.Prettify(s)
}

func (s TemplateStatus) GoString() string {
	return s.String()
}

func (s *TemplateStatus) GetLatestDeployment() *TemplateStatusLatestDeployment {
	return s.LatestDeployment
}

func (s *TemplateStatus) GetLatestVersion() *string {
	return s.LatestVersion
}

func (s *TemplateStatus) GetPackageUrl() *string {
	return s.PackageUrl
}

func (s *TemplateStatus) GetPhase() *string {
	return s.Phase
}

func (s *TemplateStatus) GetTemplateUrl() *string {
	return s.TemplateUrl
}

func (s *TemplateStatus) SetLatestDeployment(v *TemplateStatusLatestDeployment) *TemplateStatus {
	s.LatestDeployment = v
	return s
}

func (s *TemplateStatus) SetLatestVersion(v string) *TemplateStatus {
	s.LatestVersion = &v
	return s
}

func (s *TemplateStatus) SetPackageUrl(v string) *TemplateStatus {
	s.PackageUrl = &v
	return s
}

func (s *TemplateStatus) SetPhase(v string) *TemplateStatus {
	s.Phase = &v
	return s
}

func (s *TemplateStatus) SetTemplateUrl(v string) *TemplateStatus {
	s.TemplateUrl = &v
	return s
}

func (s *TemplateStatus) Validate() error {
	return dara.Validate(s)
}

type TemplateStatusLatestDeployment struct {
	FinishedTime *string `json:"finishedTime,omitempty" xml:"finishedTime,omitempty"`
	// example:
	//
	// BuildFinished
	Phase        *string `json:"phase,omitempty" xml:"phase,omitempty"`
	PipelineName *string `json:"pipelineName,omitempty" xml:"pipelineName,omitempty"`
	StartTime    *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s TemplateStatusLatestDeployment) String() string {
	return dara.Prettify(s)
}

func (s TemplateStatusLatestDeployment) GoString() string {
	return s.String()
}

func (s *TemplateStatusLatestDeployment) GetFinishedTime() *string {
	return s.FinishedTime
}

func (s *TemplateStatusLatestDeployment) GetPhase() *string {
	return s.Phase
}

func (s *TemplateStatusLatestDeployment) GetPipelineName() *string {
	return s.PipelineName
}

func (s *TemplateStatusLatestDeployment) GetStartTime() *string {
	return s.StartTime
}

func (s *TemplateStatusLatestDeployment) SetFinishedTime(v string) *TemplateStatusLatestDeployment {
	s.FinishedTime = &v
	return s
}

func (s *TemplateStatusLatestDeployment) SetPhase(v string) *TemplateStatusLatestDeployment {
	s.Phase = &v
	return s
}

func (s *TemplateStatusLatestDeployment) SetPipelineName(v string) *TemplateStatusLatestDeployment {
	s.PipelineName = &v
	return s
}

func (s *TemplateStatusLatestDeployment) SetStartTime(v string) *TemplateStatusLatestDeployment {
	s.StartTime = &v
	return s
}

func (s *TemplateStatusLatestDeployment) Validate() error {
	return dara.Validate(s)
}

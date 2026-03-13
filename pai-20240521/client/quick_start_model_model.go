// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuickStartModel interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *QuickStartModel
	GetAccessibility() *string
	SetDeploymentCount(v int64) *QuickStartModel
	GetDeploymentCount() *int64
	SetDomain(v string) *QuickStartModel
	GetDomain() *string
	SetExtraInfo(v map[string]interface{}) *QuickStartModel
	GetExtraInfo() map[string]interface{}
	SetGmtCreateTime(v string) *QuickStartModel
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *QuickStartModel
	GetGmtModifiedTime() *string
	SetLabels(v []*Label) *QuickStartModel
	GetLabels() []*Label
	SetLatestVersion(v *ModelVersion) *QuickStartModel
	GetLatestVersion() *ModelVersion
	SetModelDescription(v string) *QuickStartModel
	GetModelDescription() *string
	SetModelDoc(v string) *QuickStartModel
	GetModelDoc() *string
	SetModelId(v string) *QuickStartModel
	GetModelId() *string
	SetModelName(v string) *QuickStartModel
	GetModelName() *string
	SetOrderNumber(v int64) *QuickStartModel
	GetOrderNumber() *int64
	SetOrigin(v string) *QuickStartModel
	GetOrigin() *string
	SetOwnerId(v string) *QuickStartModel
	GetOwnerId() *string
	SetProvider(v string) *QuickStartModel
	GetProvider() *string
	SetQuickStartModelDescription(v map[string]interface{}) *QuickStartModel
	GetQuickStartModelDescription() map[string]interface{}
	SetQuickStartModelName(v map[string]interface{}) *QuickStartModel
	GetQuickStartModelName() map[string]interface{}
	SetRelatedModelId(v string) *QuickStartModel
	GetRelatedModelId() *string
	SetTask(v string) *QuickStartModel
	GetTask() *string
	SetTrainingCount(v int64) *QuickStartModel
	GetTrainingCount() *int64
	SetUserId(v string) *QuickStartModel
	GetUserId() *string
	SetViewCount(v int64) *QuickStartModel
	GetViewCount() *int64
	SetViewCount7Days(v int64) *QuickStartModel
	GetViewCount7Days() *int64
	SetWorkspaceId(v string) *QuickStartModel
	GetWorkspaceId() *string
}

type QuickStartModel struct {
	// example:
	//
	// PUBLIC
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// example:
	//
	// 100
	DeploymentCount *int64 `json:"DeploymentCount,omitempty" xml:"DeploymentCount,omitempty"`
	// example:
	//
	// cv
	Domain    *string                `json:"Domain,omitempty" xml:"Domain,omitempty"`
	ExtraInfo map[string]interface{} `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	// example:
	//
	// 2021-01-21T17:12:35Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2021-01-21T17:12:35Z
	GmtModifiedTime  *string       `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Labels           []*Label      `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	LatestVersion    *ModelVersion `json:"LatestVersion,omitempty" xml:"LatestVersion,omitempty"`
	ModelDescription *string       `json:"ModelDescription,omitempty" xml:"ModelDescription,omitempty"`
	// example:
	//
	// https://***.md
	ModelDoc *string `json:"ModelDoc,omitempty" xml:"ModelDoc,omitempty"`
	// example:
	//
	// model-1123*****
	ModelId   *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// example:
	//
	// 101
	OrderNumber *int64 `json:"OrderNumber,omitempty" xml:"OrderNumber,omitempty"`
	// example:
	//
	// ModelScope
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	// example:
	//
	// 1557702098******
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// pai
	Provider *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	// example:
	//
	// {"en-US":"Large Language Model"  "zh-CN":"大语言模型"}
	QuickStartModelDescription map[string]interface{} `json:"QuickStartModelDescription,omitempty" xml:"QuickStartModelDescription,omitempty"`
	// example:
	//
	// {"en-US":"qwen"，  "zh-CN":"千问"}
	QuickStartModelName map[string]interface{} `json:"QuickStartModelName,omitempty" xml:"QuickStartModelName,omitempty"`
	// example:
	//
	// model-12346
	RelatedModelId *string `json:"RelatedModelId,omitempty" xml:"RelatedModelId,omitempty"`
	// example:
	//
	// text-classifiaction
	Task *string `json:"Task,omitempty" xml:"Task,omitempty"`
	// example:
	//
	// 300
	TrainingCount *int64 `json:"TrainingCount,omitempty" xml:"TrainingCount,omitempty"`
	// example:
	//
	// 1557702098******
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// 500
	ViewCount *int64 `json:"ViewCount,omitempty" xml:"ViewCount,omitempty"`
	// example:
	//
	// 50
	ViewCount7Days *int64 `json:"ViewCount7Days,omitempty" xml:"ViewCount7Days,omitempty"`
	// example:
	//
	// 234**
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QuickStartModel) String() string {
	return dara.Prettify(s)
}

func (s QuickStartModel) GoString() string {
	return s.String()
}

func (s *QuickStartModel) GetAccessibility() *string {
	return s.Accessibility
}

func (s *QuickStartModel) GetDeploymentCount() *int64 {
	return s.DeploymentCount
}

func (s *QuickStartModel) GetDomain() *string {
	return s.Domain
}

func (s *QuickStartModel) GetExtraInfo() map[string]interface{} {
	return s.ExtraInfo
}

func (s *QuickStartModel) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *QuickStartModel) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *QuickStartModel) GetLabels() []*Label {
	return s.Labels
}

func (s *QuickStartModel) GetLatestVersion() *ModelVersion {
	return s.LatestVersion
}

func (s *QuickStartModel) GetModelDescription() *string {
	return s.ModelDescription
}

func (s *QuickStartModel) GetModelDoc() *string {
	return s.ModelDoc
}

func (s *QuickStartModel) GetModelId() *string {
	return s.ModelId
}

func (s *QuickStartModel) GetModelName() *string {
	return s.ModelName
}

func (s *QuickStartModel) GetOrderNumber() *int64 {
	return s.OrderNumber
}

func (s *QuickStartModel) GetOrigin() *string {
	return s.Origin
}

func (s *QuickStartModel) GetOwnerId() *string {
	return s.OwnerId
}

func (s *QuickStartModel) GetProvider() *string {
	return s.Provider
}

func (s *QuickStartModel) GetQuickStartModelDescription() map[string]interface{} {
	return s.QuickStartModelDescription
}

func (s *QuickStartModel) GetQuickStartModelName() map[string]interface{} {
	return s.QuickStartModelName
}

func (s *QuickStartModel) GetRelatedModelId() *string {
	return s.RelatedModelId
}

func (s *QuickStartModel) GetTask() *string {
	return s.Task
}

func (s *QuickStartModel) GetTrainingCount() *int64 {
	return s.TrainingCount
}

func (s *QuickStartModel) GetUserId() *string {
	return s.UserId
}

func (s *QuickStartModel) GetViewCount() *int64 {
	return s.ViewCount
}

func (s *QuickStartModel) GetViewCount7Days() *int64 {
	return s.ViewCount7Days
}

func (s *QuickStartModel) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QuickStartModel) SetAccessibility(v string) *QuickStartModel {
	s.Accessibility = &v
	return s
}

func (s *QuickStartModel) SetDeploymentCount(v int64) *QuickStartModel {
	s.DeploymentCount = &v
	return s
}

func (s *QuickStartModel) SetDomain(v string) *QuickStartModel {
	s.Domain = &v
	return s
}

func (s *QuickStartModel) SetExtraInfo(v map[string]interface{}) *QuickStartModel {
	s.ExtraInfo = v
	return s
}

func (s *QuickStartModel) SetGmtCreateTime(v string) *QuickStartModel {
	s.GmtCreateTime = &v
	return s
}

func (s *QuickStartModel) SetGmtModifiedTime(v string) *QuickStartModel {
	s.GmtModifiedTime = &v
	return s
}

func (s *QuickStartModel) SetLabels(v []*Label) *QuickStartModel {
	s.Labels = v
	return s
}

func (s *QuickStartModel) SetLatestVersion(v *ModelVersion) *QuickStartModel {
	s.LatestVersion = v
	return s
}

func (s *QuickStartModel) SetModelDescription(v string) *QuickStartModel {
	s.ModelDescription = &v
	return s
}

func (s *QuickStartModel) SetModelDoc(v string) *QuickStartModel {
	s.ModelDoc = &v
	return s
}

func (s *QuickStartModel) SetModelId(v string) *QuickStartModel {
	s.ModelId = &v
	return s
}

func (s *QuickStartModel) SetModelName(v string) *QuickStartModel {
	s.ModelName = &v
	return s
}

func (s *QuickStartModel) SetOrderNumber(v int64) *QuickStartModel {
	s.OrderNumber = &v
	return s
}

func (s *QuickStartModel) SetOrigin(v string) *QuickStartModel {
	s.Origin = &v
	return s
}

func (s *QuickStartModel) SetOwnerId(v string) *QuickStartModel {
	s.OwnerId = &v
	return s
}

func (s *QuickStartModel) SetProvider(v string) *QuickStartModel {
	s.Provider = &v
	return s
}

func (s *QuickStartModel) SetQuickStartModelDescription(v map[string]interface{}) *QuickStartModel {
	s.QuickStartModelDescription = v
	return s
}

func (s *QuickStartModel) SetQuickStartModelName(v map[string]interface{}) *QuickStartModel {
	s.QuickStartModelName = v
	return s
}

func (s *QuickStartModel) SetRelatedModelId(v string) *QuickStartModel {
	s.RelatedModelId = &v
	return s
}

func (s *QuickStartModel) SetTask(v string) *QuickStartModel {
	s.Task = &v
	return s
}

func (s *QuickStartModel) SetTrainingCount(v int64) *QuickStartModel {
	s.TrainingCount = &v
	return s
}

func (s *QuickStartModel) SetUserId(v string) *QuickStartModel {
	s.UserId = &v
	return s
}

func (s *QuickStartModel) SetViewCount(v int64) *QuickStartModel {
	s.ViewCount = &v
	return s
}

func (s *QuickStartModel) SetViewCount7Days(v int64) *QuickStartModel {
	s.ViewCount7Days = &v
	return s
}

func (s *QuickStartModel) SetWorkspaceId(v string) *QuickStartModel {
	s.WorkspaceId = &v
	return s
}

func (s *QuickStartModel) Validate() error {
	if s.Labels != nil {
		for _, item := range s.Labels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.LatestVersion != nil {
		if err := s.LatestVersion.Validate(); err != nil {
			return err
		}
	}
	return nil
}

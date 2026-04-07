// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModel interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *Model
	GetAccessibility() *string
	SetDomain(v string) *Model
	GetDomain() *string
	SetExtraInfo(v map[string]interface{}) *Model
	GetExtraInfo() map[string]interface{}
	SetGmtCreateTime(v string) *Model
	GetGmtCreateTime() *string
	SetGmtLatestVersionModifiedTime(v string) *Model
	GetGmtLatestVersionModifiedTime() *string
	SetGmtModifiedTime(v string) *Model
	GetGmtModifiedTime() *string
	SetLabels(v []*Label) *Model
	GetLabels() []*Label
	SetLatestVersion(v *ModelVersion) *Model
	GetLatestVersion() *ModelVersion
	SetModelDescription(v string) *Model
	GetModelDescription() *string
	SetModelDoc(v string) *Model
	GetModelDoc() *string
	SetModelId(v string) *Model
	GetModelId() *string
	SetModelName(v string) *Model
	GetModelName() *string
	SetModelType(v string) *Model
	GetModelType() *string
	SetOrderNumber(v int64) *Model
	GetOrderNumber() *int64
	SetOrigin(v string) *Model
	GetOrigin() *string
	SetOwnerId(v string) *Model
	GetOwnerId() *string
	SetParameterSize(v int64) *Model
	GetParameterSize() *int64
	SetProvider(v string) *Model
	GetProvider() *string
	SetTags(v []*Label) *Model
	GetTags() []*Label
	SetTask(v string) *Model
	GetTask() *string
	SetUserId(v string) *Model
	GetUserId() *string
	SetWorkspaceId(v string) *Model
	GetWorkspaceId() *string
}

type Model struct {
	// The workspace accessibility. Valid values:
	//
	// 	- PRIVATE (default): The model is accessible only to you and the administrator of the workspace.
	//
	// 	- PUBLIC: The model is accessible to all members of the workspace.
	//
	// example:
	//
	// PUBLIC
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// The domain where the model is applied, such as nlp (Natural Language Processing) and cv (Computer Vision).
	//
	// example:
	//
	// nlp
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The additional information.
	//
	// example:
	//
	// {
	//
	// 	"RatingCount": 2866,
	//
	// 	"Rating": 4.94,
	//
	// 	"FavoriteCount": 34992,
	//
	// 	"CommentCount": 754,
	//
	// 	"CoverUris": ["https://e***u.oss-cn-hangzhou.aliyuncs.com/drea***w.png"],
	//
	// 	"TippedAmountCount": 32,
	//
	// 	"DownloadCount": 606056
	//
	// }
	ExtraInfo map[string]interface{} `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	// The time when the model was created, in UTC. The time follows the ISO 8601 standard.
	//
	// example:
	//
	// 2021-01-21T17:12:35Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2021-01-21T17:12:35Z
	GmtLatestVersionModifiedTime *string `json:"GmtLatestVersionModifiedTime,omitempty" xml:"GmtLatestVersionModifiedTime,omitempty"`
	// The time when the model was last modified, in UTC. The time follows the ISO 8601 standard.
	//
	// example:
	//
	// 2021-01-21T17:12:35Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// The labels.
	Labels []*Label `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The latest version of the model.
	LatestVersion *ModelVersion `json:"LatestVersion,omitempty" xml:"LatestVersion,omitempty"`
	// The model description.
	ModelDescription *string `json:"ModelDescription,omitempty" xml:"ModelDescription,omitempty"`
	// The model document.
	//
	// example:
	//
	// https://***.md
	ModelDoc *string `json:"ModelDoc,omitempty" xml:"ModelDoc,omitempty"`
	// The model ID.
	//
	// example:
	//
	// model-1123*****
	ModelId *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	// The model name.
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// The model type, such as checkpoint and LoRA.
	//
	// example:
	//
	// Checkpoint
	ModelType *string `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	// The sequence number of the model.
	//
	// example:
	//
	// 101
	OrderNumber *int64 `json:"OrderNumber,omitempty" xml:"OrderNumber,omitempty"`
	// The community or organization to which the source model belongs, such as ModelScope or Hugging Face.
	//
	// example:
	//
	// ModelScope
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 1557702098******
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1024
	ParameterSize *int64 `json:"ParameterSize,omitempty" xml:"ParameterSize,omitempty"`
	// The model provider.
	//
	// example:
	//
	// pai
	Provider *string  `json:"Provider,omitempty" xml:"Provider,omitempty"`
	Tags     []*Label `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The task. The specific issue that the model resolves, such as text-classification.
	//
	// example:
	//
	// text-classifiaction
	Task *string `json:"Task,omitempty" xml:"Task,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 1557702098******
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 234**
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s Model) String() string {
	return dara.Prettify(s)
}

func (s Model) GoString() string {
	return s.String()
}

func (s *Model) GetAccessibility() *string {
	return s.Accessibility
}

func (s *Model) GetDomain() *string {
	return s.Domain
}

func (s *Model) GetExtraInfo() map[string]interface{} {
	return s.ExtraInfo
}

func (s *Model) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *Model) GetGmtLatestVersionModifiedTime() *string {
	return s.GmtLatestVersionModifiedTime
}

func (s *Model) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *Model) GetLabels() []*Label {
	return s.Labels
}

func (s *Model) GetLatestVersion() *ModelVersion {
	return s.LatestVersion
}

func (s *Model) GetModelDescription() *string {
	return s.ModelDescription
}

func (s *Model) GetModelDoc() *string {
	return s.ModelDoc
}

func (s *Model) GetModelId() *string {
	return s.ModelId
}

func (s *Model) GetModelName() *string {
	return s.ModelName
}

func (s *Model) GetModelType() *string {
	return s.ModelType
}

func (s *Model) GetOrderNumber() *int64 {
	return s.OrderNumber
}

func (s *Model) GetOrigin() *string {
	return s.Origin
}

func (s *Model) GetOwnerId() *string {
	return s.OwnerId
}

func (s *Model) GetParameterSize() *int64 {
	return s.ParameterSize
}

func (s *Model) GetProvider() *string {
	return s.Provider
}

func (s *Model) GetTags() []*Label {
	return s.Tags
}

func (s *Model) GetTask() *string {
	return s.Task
}

func (s *Model) GetUserId() *string {
	return s.UserId
}

func (s *Model) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *Model) SetAccessibility(v string) *Model {
	s.Accessibility = &v
	return s
}

func (s *Model) SetDomain(v string) *Model {
	s.Domain = &v
	return s
}

func (s *Model) SetExtraInfo(v map[string]interface{}) *Model {
	s.ExtraInfo = v
	return s
}

func (s *Model) SetGmtCreateTime(v string) *Model {
	s.GmtCreateTime = &v
	return s
}

func (s *Model) SetGmtLatestVersionModifiedTime(v string) *Model {
	s.GmtLatestVersionModifiedTime = &v
	return s
}

func (s *Model) SetGmtModifiedTime(v string) *Model {
	s.GmtModifiedTime = &v
	return s
}

func (s *Model) SetLabels(v []*Label) *Model {
	s.Labels = v
	return s
}

func (s *Model) SetLatestVersion(v *ModelVersion) *Model {
	s.LatestVersion = v
	return s
}

func (s *Model) SetModelDescription(v string) *Model {
	s.ModelDescription = &v
	return s
}

func (s *Model) SetModelDoc(v string) *Model {
	s.ModelDoc = &v
	return s
}

func (s *Model) SetModelId(v string) *Model {
	s.ModelId = &v
	return s
}

func (s *Model) SetModelName(v string) *Model {
	s.ModelName = &v
	return s
}

func (s *Model) SetModelType(v string) *Model {
	s.ModelType = &v
	return s
}

func (s *Model) SetOrderNumber(v int64) *Model {
	s.OrderNumber = &v
	return s
}

func (s *Model) SetOrigin(v string) *Model {
	s.Origin = &v
	return s
}

func (s *Model) SetOwnerId(v string) *Model {
	s.OwnerId = &v
	return s
}

func (s *Model) SetParameterSize(v int64) *Model {
	s.ParameterSize = &v
	return s
}

func (s *Model) SetProvider(v string) *Model {
	s.Provider = &v
	return s
}

func (s *Model) SetTags(v []*Label) *Model {
	s.Tags = v
	return s
}

func (s *Model) SetTask(v string) *Model {
	s.Task = &v
	return s
}

func (s *Model) SetUserId(v string) *Model {
	s.UserId = &v
	return s
}

func (s *Model) SetWorkspaceId(v string) *Model {
	s.WorkspaceId = &v
	return s
}

func (s *Model) Validate() error {
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
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

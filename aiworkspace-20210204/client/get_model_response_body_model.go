// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *GetModelResponseBody
	GetAccessibility() *string
	SetDomain(v string) *GetModelResponseBody
	GetDomain() *string
	SetExtraInfo(v map[string]interface{}) *GetModelResponseBody
	GetExtraInfo() map[string]interface{}
	SetGmtCreateTime(v string) *GetModelResponseBody
	GetGmtCreateTime() *string
	SetGmtLatestVersionModifiedTime(v string) *GetModelResponseBody
	GetGmtLatestVersionModifiedTime() *string
	SetGmtModifiedTime(v string) *GetModelResponseBody
	GetGmtModifiedTime() *string
	SetLabels(v []*Label) *GetModelResponseBody
	GetLabels() []*Label
	SetLatestVersion(v *ModelVersion) *GetModelResponseBody
	GetLatestVersion() *ModelVersion
	SetModelDescription(v string) *GetModelResponseBody
	GetModelDescription() *string
	SetModelDoc(v string) *GetModelResponseBody
	GetModelDoc() *string
	SetModelId(v string) *GetModelResponseBody
	GetModelId() *string
	SetModelName(v string) *GetModelResponseBody
	GetModelName() *string
	SetModelType(v string) *GetModelResponseBody
	GetModelType() *string
	SetOrderNumber(v int64) *GetModelResponseBody
	GetOrderNumber() *int64
	SetOrigin(v string) *GetModelResponseBody
	GetOrigin() *string
	SetOwnerId(v string) *GetModelResponseBody
	GetOwnerId() *string
	SetParameterSize(v int64) *GetModelResponseBody
	GetParameterSize() *int64
	SetProvider(v string) *GetModelResponseBody
	GetProvider() *string
	SetRequestId(v string) *GetModelResponseBody
	GetRequestId() *string
	SetTask(v string) *GetModelResponseBody
	GetTask() *string
	SetUserId(v string) *GetModelResponseBody
	GetUserId() *string
	SetWorkspaceId(v string) *GetModelResponseBody
	GetWorkspaceId() *string
}

type GetModelResponseBody struct {
	// The visibility of the workspace.
	//
	// 	- PRIVATE: The workspace is visible only to you and the administrator of the workspace.
	//
	// 	- PUBLIC: The workspace is visible to all users.
	//
	// example:
	//
	// PUBLIC
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// The domain. This parameter specifies the domain for which the model is developed. Valid values: nlp and cv. nlp indicates natural language processing and cv indicates computer vision.
	//
	// example:
	//
	// cv
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// Other information about the model.
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
	// The time when the model is created, in UTC. The time follows the ISO 8601 standard.
	//
	// example:
	//
	// 2021-01-30T12:51:33.028Z
	GmtCreateTime                *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtLatestVersionModifiedTime *string `json:"GmtLatestVersionModifiedTime,omitempty" xml:"GmtLatestVersionModifiedTime,omitempty"`
	// The time when the model is last modified, in UTC. The time follows the ISO 8601 standard.
	//
	// example:
	//
	// 2021-01-30T12:51:33.028Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// The model tags.
	Labels []*Label `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The latest version of the model.
	LatestVersion *ModelVersion `json:"LatestVersion,omitempty" xml:"LatestVersion,omitempty"`
	// The model description.
	ModelDescription *string `json:"ModelDescription,omitempty" xml:"ModelDescription,omitempty"`
	// The documentation of the model.
	//
	// example:
	//
	// https://***.md
	ModelDoc *string `json:"ModelDoc,omitempty" xml:"ModelDoc,omitempty"`
	// The model ID.
	//
	// example:
	//
	// model-rbvg5wzljz****ks92
	ModelId *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	// The model name.
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// The model type.
	//
	// example:
	//
	// Checkpoint
	ModelType *string `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	// The sequence number of the model.
	//
	// example:
	//
	// 1
	OrderNumber *int64 `json:"OrderNumber,omitempty" xml:"OrderNumber,omitempty"`
	// The source of the model. The community or organization to which the model belongs, such as ModelScope or HuggingFace.
	//
	// example:
	//
	// ModelScope
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 1234567890******
	OwnerId       *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ParameterSize *int64  `json:"ParameterSize,omitempty" xml:"ParameterSize,omitempty"`
	// The provider.
	//
	// example:
	//
	// pai
	Provider *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task of the model. This parameter describes specific issues that the model solves, such as text-classification.
	//
	// example:
	//
	// text-classification
	Task *string `json:"Task,omitempty" xml:"Task,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 1234567890******
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 234**
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetModelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetModelResponseBody) GoString() string {
	return s.String()
}

func (s *GetModelResponseBody) GetAccessibility() *string {
	return s.Accessibility
}

func (s *GetModelResponseBody) GetDomain() *string {
	return s.Domain
}

func (s *GetModelResponseBody) GetExtraInfo() map[string]interface{} {
	return s.ExtraInfo
}

func (s *GetModelResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetModelResponseBody) GetGmtLatestVersionModifiedTime() *string {
	return s.GmtLatestVersionModifiedTime
}

func (s *GetModelResponseBody) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *GetModelResponseBody) GetLabels() []*Label {
	return s.Labels
}

func (s *GetModelResponseBody) GetLatestVersion() *ModelVersion {
	return s.LatestVersion
}

func (s *GetModelResponseBody) GetModelDescription() *string {
	return s.ModelDescription
}

func (s *GetModelResponseBody) GetModelDoc() *string {
	return s.ModelDoc
}

func (s *GetModelResponseBody) GetModelId() *string {
	return s.ModelId
}

func (s *GetModelResponseBody) GetModelName() *string {
	return s.ModelName
}

func (s *GetModelResponseBody) GetModelType() *string {
	return s.ModelType
}

func (s *GetModelResponseBody) GetOrderNumber() *int64 {
	return s.OrderNumber
}

func (s *GetModelResponseBody) GetOrigin() *string {
	return s.Origin
}

func (s *GetModelResponseBody) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GetModelResponseBody) GetParameterSize() *int64 {
	return s.ParameterSize
}

func (s *GetModelResponseBody) GetProvider() *string {
	return s.Provider
}

func (s *GetModelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetModelResponseBody) GetTask() *string {
	return s.Task
}

func (s *GetModelResponseBody) GetUserId() *string {
	return s.UserId
}

func (s *GetModelResponseBody) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetModelResponseBody) SetAccessibility(v string) *GetModelResponseBody {
	s.Accessibility = &v
	return s
}

func (s *GetModelResponseBody) SetDomain(v string) *GetModelResponseBody {
	s.Domain = &v
	return s
}

func (s *GetModelResponseBody) SetExtraInfo(v map[string]interface{}) *GetModelResponseBody {
	s.ExtraInfo = v
	return s
}

func (s *GetModelResponseBody) SetGmtCreateTime(v string) *GetModelResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetModelResponseBody) SetGmtLatestVersionModifiedTime(v string) *GetModelResponseBody {
	s.GmtLatestVersionModifiedTime = &v
	return s
}

func (s *GetModelResponseBody) SetGmtModifiedTime(v string) *GetModelResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetModelResponseBody) SetLabels(v []*Label) *GetModelResponseBody {
	s.Labels = v
	return s
}

func (s *GetModelResponseBody) SetLatestVersion(v *ModelVersion) *GetModelResponseBody {
	s.LatestVersion = v
	return s
}

func (s *GetModelResponseBody) SetModelDescription(v string) *GetModelResponseBody {
	s.ModelDescription = &v
	return s
}

func (s *GetModelResponseBody) SetModelDoc(v string) *GetModelResponseBody {
	s.ModelDoc = &v
	return s
}

func (s *GetModelResponseBody) SetModelId(v string) *GetModelResponseBody {
	s.ModelId = &v
	return s
}

func (s *GetModelResponseBody) SetModelName(v string) *GetModelResponseBody {
	s.ModelName = &v
	return s
}

func (s *GetModelResponseBody) SetModelType(v string) *GetModelResponseBody {
	s.ModelType = &v
	return s
}

func (s *GetModelResponseBody) SetOrderNumber(v int64) *GetModelResponseBody {
	s.OrderNumber = &v
	return s
}

func (s *GetModelResponseBody) SetOrigin(v string) *GetModelResponseBody {
	s.Origin = &v
	return s
}

func (s *GetModelResponseBody) SetOwnerId(v string) *GetModelResponseBody {
	s.OwnerId = &v
	return s
}

func (s *GetModelResponseBody) SetParameterSize(v int64) *GetModelResponseBody {
	s.ParameterSize = &v
	return s
}

func (s *GetModelResponseBody) SetProvider(v string) *GetModelResponseBody {
	s.Provider = &v
	return s
}

func (s *GetModelResponseBody) SetRequestId(v string) *GetModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetModelResponseBody) SetTask(v string) *GetModelResponseBody {
	s.Task = &v
	return s
}

func (s *GetModelResponseBody) SetUserId(v string) *GetModelResponseBody {
	s.UserId = &v
	return s
}

func (s *GetModelResponseBody) SetWorkspaceId(v string) *GetModelResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *GetModelResponseBody) Validate() error {
	return dara.Validate(s)
}

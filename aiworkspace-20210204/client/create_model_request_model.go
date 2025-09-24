// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *CreateModelRequest
	GetAccessibility() *string
	SetDomain(v string) *CreateModelRequest
	GetDomain() *string
	SetExtraInfo(v map[string]interface{}) *CreateModelRequest
	GetExtraInfo() map[string]interface{}
	SetLabels(v []*Label) *CreateModelRequest
	GetLabels() []*Label
	SetModelDescription(v string) *CreateModelRequest
	GetModelDescription() *string
	SetModelDoc(v string) *CreateModelRequest
	GetModelDoc() *string
	SetModelName(v string) *CreateModelRequest
	GetModelName() *string
	SetModelType(v string) *CreateModelRequest
	GetModelType() *string
	SetOrderNumber(v int64) *CreateModelRequest
	GetOrderNumber() *int64
	SetOrigin(v string) *CreateModelRequest
	GetOrigin() *string
	SetParameterSize(v int64) *CreateModelRequest
	GetParameterSize() *int64
	SetTag(v []*Label) *CreateModelRequest
	GetTag() []*Label
	SetTask(v string) *CreateModelRequest
	GetTask() *string
	SetWorkspaceId(v string) *CreateModelRequest
	GetWorkspaceId() *string
}

type CreateModelRequest struct {
	// The visibility of the model in the workspace. Valid values:
	//
	// 	- PRIVATE (default): Visible only to you and the administrator of the workspace.
	//
	// 	- PUBLIC: Vvisible to all users in the workspace.
	//
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// The domain of the model. Describes the domain in which the model is for. Example: nlp (natural language processing), cv (computer vision), and others.
	//
	// example:
	//
	// nlp
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// Other information about the model.
	//
	// if can be null:
	// true
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
	// The tags. This parameter will be deprecated and replaced by Tag.
	Labels []*Label `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The model description, used to distinguish different models.
	ModelDescription *string `json:"ModelDescription,omitempty" xml:"ModelDescription,omitempty"`
	// The documentation of the model.
	//
	// example:
	//
	// https://*.md
	ModelDoc *string `json:"ModelDoc,omitempty" xml:"ModelDoc,omitempty"`
	// The name of the model. The name must be 1 to 127 characters in length.
	//
	// This parameter is required.
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// The model type. Example: Checkpoint or LoRA.
	//
	// example:
	//
	// Checkpoint
	ModelType *string `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	// The sequence number of the model. Can be used for custom sorting.
	//
	// example:
	//
	// 0
	OrderNumber *int64 `json:"OrderNumber,omitempty" xml:"OrderNumber,omitempty"`
	// The source of the model. The community or organization to which the source model belongs, such as ModelScope or HuggingFace.
	//
	// example:
	//
	// ModelScope
	Origin        *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	ParameterSize *int64  `json:"ParameterSize,omitempty" xml:"ParameterSize,omitempty"`
	// The tags.
	Tag []*Label `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The task of the model. Describes the specific problem that the model solves. Example: text-classification.
	//
	// example:
	//
	// text-classification
	Task *string `json:"Task,omitempty" xml:"Task,omitempty"`
	// The workspace ID. Call [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html) to obtain the workspace ID.
	//
	// example:
	//
	// 796**
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateModelRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateModelRequest) GoString() string {
	return s.String()
}

func (s *CreateModelRequest) GetAccessibility() *string {
	return s.Accessibility
}

func (s *CreateModelRequest) GetDomain() *string {
	return s.Domain
}

func (s *CreateModelRequest) GetExtraInfo() map[string]interface{} {
	return s.ExtraInfo
}

func (s *CreateModelRequest) GetLabels() []*Label {
	return s.Labels
}

func (s *CreateModelRequest) GetModelDescription() *string {
	return s.ModelDescription
}

func (s *CreateModelRequest) GetModelDoc() *string {
	return s.ModelDoc
}

func (s *CreateModelRequest) GetModelName() *string {
	return s.ModelName
}

func (s *CreateModelRequest) GetModelType() *string {
	return s.ModelType
}

func (s *CreateModelRequest) GetOrderNumber() *int64 {
	return s.OrderNumber
}

func (s *CreateModelRequest) GetOrigin() *string {
	return s.Origin
}

func (s *CreateModelRequest) GetParameterSize() *int64 {
	return s.ParameterSize
}

func (s *CreateModelRequest) GetTag() []*Label {
	return s.Tag
}

func (s *CreateModelRequest) GetTask() *string {
	return s.Task
}

func (s *CreateModelRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateModelRequest) SetAccessibility(v string) *CreateModelRequest {
	s.Accessibility = &v
	return s
}

func (s *CreateModelRequest) SetDomain(v string) *CreateModelRequest {
	s.Domain = &v
	return s
}

func (s *CreateModelRequest) SetExtraInfo(v map[string]interface{}) *CreateModelRequest {
	s.ExtraInfo = v
	return s
}

func (s *CreateModelRequest) SetLabels(v []*Label) *CreateModelRequest {
	s.Labels = v
	return s
}

func (s *CreateModelRequest) SetModelDescription(v string) *CreateModelRequest {
	s.ModelDescription = &v
	return s
}

func (s *CreateModelRequest) SetModelDoc(v string) *CreateModelRequest {
	s.ModelDoc = &v
	return s
}

func (s *CreateModelRequest) SetModelName(v string) *CreateModelRequest {
	s.ModelName = &v
	return s
}

func (s *CreateModelRequest) SetModelType(v string) *CreateModelRequest {
	s.ModelType = &v
	return s
}

func (s *CreateModelRequest) SetOrderNumber(v int64) *CreateModelRequest {
	s.OrderNumber = &v
	return s
}

func (s *CreateModelRequest) SetOrigin(v string) *CreateModelRequest {
	s.Origin = &v
	return s
}

func (s *CreateModelRequest) SetParameterSize(v int64) *CreateModelRequest {
	s.ParameterSize = &v
	return s
}

func (s *CreateModelRequest) SetTag(v []*Label) *CreateModelRequest {
	s.Tag = v
	return s
}

func (s *CreateModelRequest) SetTask(v string) *CreateModelRequest {
	s.Task = &v
	return s
}

func (s *CreateModelRequest) SetWorkspaceId(v string) *CreateModelRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateModelRequest) Validate() error {
	return dara.Validate(s)
}

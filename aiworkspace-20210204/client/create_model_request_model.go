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
	// - PRIVATE (default): The model is visible only to you and administrators in the workspace.
	//
	// - PUBLIC: The model is visible to everyone in the workspace.
	//
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// The domain. This describes the field that the model is designed for, such as nlp (Natural Language Processing) or cv (computer vision).
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
	// A list of labels. This parameter is deprecated and is replaced by the Tag parameter.
	Labels []*Label `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The description of the model. Use this to distinguish different models.
	//
	// example:
	//
	// News classification.
	ModelDescription *string `json:"ModelDescription,omitempty" xml:"ModelDescription,omitempty"`
	// The model documentation.
	//
	// example:
	//
	// https://*.md
	ModelDoc *string `json:"ModelDoc,omitempty" xml:"ModelDoc,omitempty"`
	// The name of the model. The name must be 1 to 127 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// News classification
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// The model type, such as Checkpoint or LoRA.
	//
	// example:
	//
	// Checkpoint
	ModelType *string `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	// The ordinal number of the model. You can use this for custom sorting.
	//
	// example:
	//
	// 0
	OrderNumber *int64 `json:"OrderNumber,omitempty" xml:"OrderNumber,omitempty"`
	// The source of the model. This specifies the community or organization that the source model belongs to, such as ModelScope and HuggingFace.
	//
	// example:
	//
	// ModelScope
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	// The number of parameters, in millions.
	//
	// example:
	//
	// 3000
	ParameterSize *int64 `json:"ParameterSize,omitempty" xml:"ParameterSize,omitempty"`
	// A list of tags.
	Tag []*Label `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The task. This describes the specific problem that the model solves, such as text-classification (text classification).
	//
	// example:
	//
	// text-classification
	Task *string `json:"Task,omitempty" xml:"Task,omitempty"`
	// The ID of the workspace. For more information about how to obtain a workspace ID, see [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html).
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
	if s.Labels != nil {
		for _, item := range s.Labels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *UpdateModelRequest
	GetAccessibility() *string
	SetDomain(v string) *UpdateModelRequest
	GetDomain() *string
	SetExtraInfo(v map[string]interface{}) *UpdateModelRequest
	GetExtraInfo() map[string]interface{}
	SetModelDescription(v string) *UpdateModelRequest
	GetModelDescription() *string
	SetModelDoc(v string) *UpdateModelRequest
	GetModelDoc() *string
	SetModelName(v string) *UpdateModelRequest
	GetModelName() *string
	SetModelType(v string) *UpdateModelRequest
	GetModelType() *string
	SetOrderNumber(v int64) *UpdateModelRequest
	GetOrderNumber() *int64
	SetOrigin(v string) *UpdateModelRequest
	GetOrigin() *string
	SetParameterSize(v int64) *UpdateModelRequest
	GetParameterSize() *int64
	SetTask(v string) *UpdateModelRequest
	GetTask() *string
}

type UpdateModelRequest struct {
	// The visibility of the model in the workspace. Valid values:
	//
	// 	- PRIVATE: The model is visible only to you and the administrator of the workspace.
	//
	// 	- PUBLIC: The model is visible to all users in the workspace.
	//
	// example:
	//
	// PUBLIC
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// The domain. This parameter describes the domain in which the model is applied. Valid values: nlp (natural language processing) and cv (computer vision).
	//
	// example:
	//
	// nlp
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
	// The model description.
	ModelDescription *string `json:"ModelDescription,omitempty" xml:"ModelDescription,omitempty"`
	// The documentation of the model.
	//
	// example:
	//
	// https://*.md
	ModelDoc *string `json:"ModelDoc,omitempty" xml:"ModelDoc,omitempty"`
	// The model name, which must be 1 to 127 characters in length.
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// The model type. Valid values: Checkpoint and LoRA.
	//
	// example:
	//
	// Checkpoint
	ModelType *string `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	// The sequence number of the model. This parameter can be used for custom sorting.
	//
	// example:
	//
	// 0
	OrderNumber *int64 `json:"OrderNumber,omitempty" xml:"OrderNumber,omitempty"`
	// The source of the model. This parameter describes the community or organization to which the source model belongs. Valid values: ModelScope and HuggingFace.
	//
	// example:
	//
	// ModelScope
	Origin        *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	ParameterSize *int64  `json:"ParameterSize,omitempty" xml:"ParameterSize,omitempty"`
	// The task. This parameter specifies the specific issue that the model resolves. Example: text-classification.
	//
	// example:
	//
	// text-classification
	Task *string `json:"Task,omitempty" xml:"Task,omitempty"`
}

func (s UpdateModelRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelRequest) GoString() string {
	return s.String()
}

func (s *UpdateModelRequest) GetAccessibility() *string {
	return s.Accessibility
}

func (s *UpdateModelRequest) GetDomain() *string {
	return s.Domain
}

func (s *UpdateModelRequest) GetExtraInfo() map[string]interface{} {
	return s.ExtraInfo
}

func (s *UpdateModelRequest) GetModelDescription() *string {
	return s.ModelDescription
}

func (s *UpdateModelRequest) GetModelDoc() *string {
	return s.ModelDoc
}

func (s *UpdateModelRequest) GetModelName() *string {
	return s.ModelName
}

func (s *UpdateModelRequest) GetModelType() *string {
	return s.ModelType
}

func (s *UpdateModelRequest) GetOrderNumber() *int64 {
	return s.OrderNumber
}

func (s *UpdateModelRequest) GetOrigin() *string {
	return s.Origin
}

func (s *UpdateModelRequest) GetParameterSize() *int64 {
	return s.ParameterSize
}

func (s *UpdateModelRequest) GetTask() *string {
	return s.Task
}

func (s *UpdateModelRequest) SetAccessibility(v string) *UpdateModelRequest {
	s.Accessibility = &v
	return s
}

func (s *UpdateModelRequest) SetDomain(v string) *UpdateModelRequest {
	s.Domain = &v
	return s
}

func (s *UpdateModelRequest) SetExtraInfo(v map[string]interface{}) *UpdateModelRequest {
	s.ExtraInfo = v
	return s
}

func (s *UpdateModelRequest) SetModelDescription(v string) *UpdateModelRequest {
	s.ModelDescription = &v
	return s
}

func (s *UpdateModelRequest) SetModelDoc(v string) *UpdateModelRequest {
	s.ModelDoc = &v
	return s
}

func (s *UpdateModelRequest) SetModelName(v string) *UpdateModelRequest {
	s.ModelName = &v
	return s
}

func (s *UpdateModelRequest) SetModelType(v string) *UpdateModelRequest {
	s.ModelType = &v
	return s
}

func (s *UpdateModelRequest) SetOrderNumber(v int64) *UpdateModelRequest {
	s.OrderNumber = &v
	return s
}

func (s *UpdateModelRequest) SetOrigin(v string) *UpdateModelRequest {
	s.Origin = &v
	return s
}

func (s *UpdateModelRequest) SetParameterSize(v int64) *UpdateModelRequest {
	s.ParameterSize = &v
	return s
}

func (s *UpdateModelRequest) SetTask(v string) *UpdateModelRequest {
	s.Task = &v
	return s
}

func (s *UpdateModelRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateModelInstanceRequest
	GetDescription() *string
	SetEasServiceId(v string) *UpdateModelInstanceRequest
	GetEasServiceId() *string
	SetEasServiceName(v string) *UpdateModelInstanceRequest
	GetEasServiceName() *string
	SetModelCallName(v string) *UpdateModelInstanceRequest
	GetModelCallName() *string
	SetModelCategoryId(v int64) *UpdateModelInstanceRequest
	GetModelCategoryId() *int64
	SetModelInstanceId(v int64) *UpdateModelInstanceRequest
	GetModelInstanceId() *int64
	SetModelToken(v string) *UpdateModelInstanceRequest
	GetModelToken() *string
	SetModelUrl(v string) *UpdateModelInstanceRequest
	GetModelUrl() *string
	SetModelVpcUrl(v string) *UpdateModelInstanceRequest
	GetModelVpcUrl() *string
	SetWorkspaceId(v int64) *UpdateModelInstanceRequest
	GetWorkspaceId() *int64
}

type UpdateModelInstanceRequest struct {
	// example:
	//
	// vllm==0.9.0
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// eas-m-12345678
	EasServiceId *string `json:"EasServiceId,omitempty" xml:"EasServiceId,omitempty"`
	// example:
	//
	// rai_content_detection_model
	EasServiceName *string `json:"EasServiceName,omitempty" xml:"EasServiceName,omitempty"`
	// example:
	//
	// demo
	ModelCallName *string `json:"ModelCallName,omitempty" xml:"ModelCallName,omitempty"`
	// example:
	//
	// 1
	ModelCategoryId *int64 `json:"ModelCategoryId,omitempty" xml:"ModelCategoryId,omitempty"`
	// example:
	//
	// 123
	ModelInstanceId *int64 `json:"ModelInstanceId,omitempty" xml:"ModelInstanceId,omitempty"`
	// example:
	//
	// MzJiMDI5MDliODc0MTlkYmI0ZDhlYmExYjczYTIyZTE3Zm********
	ModelToken *string `json:"ModelToken,omitempty" xml:"ModelToken,omitempty"`
	// example:
	//
	// http://12345*****.cn-shanghai.aliyuncs.com/api/predict/echo
	ModelUrl *string `json:"ModelUrl,omitempty" xml:"ModelUrl,omitempty"`
	// example:
	//
	// http://12345*****.vpc.cn-shanghai.aliyuncs.com/api/predict/demo
	ModelVpcUrl *string `json:"ModelVpcUrl,omitempty" xml:"ModelVpcUrl,omitempty"`
	// example:
	//
	// 608226
	WorkspaceId *int64 `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateModelInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateModelInstanceRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateModelInstanceRequest) GetEasServiceId() *string {
	return s.EasServiceId
}

func (s *UpdateModelInstanceRequest) GetEasServiceName() *string {
	return s.EasServiceName
}

func (s *UpdateModelInstanceRequest) GetModelCallName() *string {
	return s.ModelCallName
}

func (s *UpdateModelInstanceRequest) GetModelCategoryId() *int64 {
	return s.ModelCategoryId
}

func (s *UpdateModelInstanceRequest) GetModelInstanceId() *int64 {
	return s.ModelInstanceId
}

func (s *UpdateModelInstanceRequest) GetModelToken() *string {
	return s.ModelToken
}

func (s *UpdateModelInstanceRequest) GetModelUrl() *string {
	return s.ModelUrl
}

func (s *UpdateModelInstanceRequest) GetModelVpcUrl() *string {
	return s.ModelVpcUrl
}

func (s *UpdateModelInstanceRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *UpdateModelInstanceRequest) SetDescription(v string) *UpdateModelInstanceRequest {
	s.Description = &v
	return s
}

func (s *UpdateModelInstanceRequest) SetEasServiceId(v string) *UpdateModelInstanceRequest {
	s.EasServiceId = &v
	return s
}

func (s *UpdateModelInstanceRequest) SetEasServiceName(v string) *UpdateModelInstanceRequest {
	s.EasServiceName = &v
	return s
}

func (s *UpdateModelInstanceRequest) SetModelCallName(v string) *UpdateModelInstanceRequest {
	s.ModelCallName = &v
	return s
}

func (s *UpdateModelInstanceRequest) SetModelCategoryId(v int64) *UpdateModelInstanceRequest {
	s.ModelCategoryId = &v
	return s
}

func (s *UpdateModelInstanceRequest) SetModelInstanceId(v int64) *UpdateModelInstanceRequest {
	s.ModelInstanceId = &v
	return s
}

func (s *UpdateModelInstanceRequest) SetModelToken(v string) *UpdateModelInstanceRequest {
	s.ModelToken = &v
	return s
}

func (s *UpdateModelInstanceRequest) SetModelUrl(v string) *UpdateModelInstanceRequest {
	s.ModelUrl = &v
	return s
}

func (s *UpdateModelInstanceRequest) SetModelVpcUrl(v string) *UpdateModelInstanceRequest {
	s.ModelVpcUrl = &v
	return s
}

func (s *UpdateModelInstanceRequest) SetWorkspaceId(v int64) *UpdateModelInstanceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateModelInstanceRequest) Validate() error {
	return dara.Validate(s)
}

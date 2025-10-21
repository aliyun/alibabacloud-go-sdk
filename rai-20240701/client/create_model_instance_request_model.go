// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEasServiceId(v string) *CreateModelInstanceRequest
	GetEasServiceId() *string
	SetEasServiceName(v string) *CreateModelInstanceRequest
	GetEasServiceName() *string
	SetModelCallName(v string) *CreateModelInstanceRequest
	GetModelCallName() *string
	SetModelCategoryId(v int64) *CreateModelInstanceRequest
	GetModelCategoryId() *int64
	SetModelToken(v string) *CreateModelInstanceRequest
	GetModelToken() *string
	SetModelUrl(v string) *CreateModelInstanceRequest
	GetModelUrl() *string
	SetModelVpcUrl(v string) *CreateModelInstanceRequest
	GetModelVpcUrl() *string
	SetWorkspaceId(v int64) *CreateModelInstanceRequest
	GetWorkspaceId() *int64
}

type CreateModelInstanceRequest struct {
	// example:
	//
	// eas-m-12345678
	EasServiceId *string `json:"EasServiceId,omitempty" xml:"EasServiceId,omitempty"`
	// example:
	//
	// demo
	EasServiceName *string `json:"EasServiceName,omitempty" xml:"EasServiceName,omitempty"`
	// example:
	//
	// demo
	ModelCallName *string `json:"ModelCallName,omitempty" xml:"ModelCallName,omitempty"`
	// example:
	//
	// 1
	ModelCategoryId *int64 `json:"ModelCategoryId,omitempty" xml:"ModelCategoryId,omitempty"`
	// EAS Token
	//
	// example:
	//
	// MzJiMDI5MDliODc0MTlkYmI0ZDhlYmExYjczYTIyZTE3Zm********
	ModelToken *string `json:"ModelToken,omitempty" xml:"ModelToken,omitempty"`
	// example:
	//
	// http://12345*****.cn-shanghai.aliyuncs.com/api/predict/demo
	ModelUrl *string `json:"ModelUrl,omitempty" xml:"ModelUrl,omitempty"`
	// example:
	//
	// http://12345*****.vpc.cn-shanghai.aliyuncs.com/api/predict/demo
	ModelVpcUrl *string `json:"ModelVpcUrl,omitempty" xml:"ModelVpcUrl,omitempty"`
	// example:
	//
	// 643168
	WorkspaceId *int64 `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateModelInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateModelInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateModelInstanceRequest) GetEasServiceId() *string {
	return s.EasServiceId
}

func (s *CreateModelInstanceRequest) GetEasServiceName() *string {
	return s.EasServiceName
}

func (s *CreateModelInstanceRequest) GetModelCallName() *string {
	return s.ModelCallName
}

func (s *CreateModelInstanceRequest) GetModelCategoryId() *int64 {
	return s.ModelCategoryId
}

func (s *CreateModelInstanceRequest) GetModelToken() *string {
	return s.ModelToken
}

func (s *CreateModelInstanceRequest) GetModelUrl() *string {
	return s.ModelUrl
}

func (s *CreateModelInstanceRequest) GetModelVpcUrl() *string {
	return s.ModelVpcUrl
}

func (s *CreateModelInstanceRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *CreateModelInstanceRequest) SetEasServiceId(v string) *CreateModelInstanceRequest {
	s.EasServiceId = &v
	return s
}

func (s *CreateModelInstanceRequest) SetEasServiceName(v string) *CreateModelInstanceRequest {
	s.EasServiceName = &v
	return s
}

func (s *CreateModelInstanceRequest) SetModelCallName(v string) *CreateModelInstanceRequest {
	s.ModelCallName = &v
	return s
}

func (s *CreateModelInstanceRequest) SetModelCategoryId(v int64) *CreateModelInstanceRequest {
	s.ModelCategoryId = &v
	return s
}

func (s *CreateModelInstanceRequest) SetModelToken(v string) *CreateModelInstanceRequest {
	s.ModelToken = &v
	return s
}

func (s *CreateModelInstanceRequest) SetModelUrl(v string) *CreateModelInstanceRequest {
	s.ModelUrl = &v
	return s
}

func (s *CreateModelInstanceRequest) SetModelVpcUrl(v string) *CreateModelInstanceRequest {
	s.ModelVpcUrl = &v
	return s
}

func (s *CreateModelInstanceRequest) SetWorkspaceId(v int64) *CreateModelInstanceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateModelInstanceRequest) Validate() error {
	return dara.Validate(s)
}

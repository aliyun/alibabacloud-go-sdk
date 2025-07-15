// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBackendModelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackendId(v string) *CreateBackendModelRequest
	GetBackendId() *string
	SetBackendModelData(v string) *CreateBackendModelRequest
	GetBackendModelData() *string
	SetBackendType(v string) *CreateBackendModelRequest
	GetBackendType() *string
	SetDescription(v string) *CreateBackendModelRequest
	GetDescription() *string
	SetSecurityToken(v string) *CreateBackendModelRequest
	GetSecurityToken() *string
	SetStageName(v string) *CreateBackendModelRequest
	GetStageName() *string
}

type CreateBackendModelRequest struct {
	// example:
	//
	// 34e94fcd3e2e47a49824a89b8f92cb5e
	BackendId *string `json:"BackendId,omitempty" xml:"BackendId,omitempty"`
	// example:
	//
	// {\\"ServiceAddress\\":\\"http://apigateway-echo-redux.alicloudapi.com:8080\\"}
	BackendModelData *string `json:"BackendModelData,omitempty" xml:"BackendModelData,omitempty"`
	// example:
	//
	// HTTP
	BackendType *string `json:"BackendType,omitempty" xml:"BackendType,omitempty"`
	// example:
	//
	// model description
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// example:
	//
	// TEST
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s CreateBackendModelRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBackendModelRequest) GoString() string {
	return s.String()
}

func (s *CreateBackendModelRequest) GetBackendId() *string {
	return s.BackendId
}

func (s *CreateBackendModelRequest) GetBackendModelData() *string {
	return s.BackendModelData
}

func (s *CreateBackendModelRequest) GetBackendType() *string {
	return s.BackendType
}

func (s *CreateBackendModelRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateBackendModelRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *CreateBackendModelRequest) GetStageName() *string {
	return s.StageName
}

func (s *CreateBackendModelRequest) SetBackendId(v string) *CreateBackendModelRequest {
	s.BackendId = &v
	return s
}

func (s *CreateBackendModelRequest) SetBackendModelData(v string) *CreateBackendModelRequest {
	s.BackendModelData = &v
	return s
}

func (s *CreateBackendModelRequest) SetBackendType(v string) *CreateBackendModelRequest {
	s.BackendType = &v
	return s
}

func (s *CreateBackendModelRequest) SetDescription(v string) *CreateBackendModelRequest {
	s.Description = &v
	return s
}

func (s *CreateBackendModelRequest) SetSecurityToken(v string) *CreateBackendModelRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateBackendModelRequest) SetStageName(v string) *CreateBackendModelRequest {
	s.StageName = &v
	return s
}

func (s *CreateBackendModelRequest) Validate() error {
	return dara.Validate(s)
}

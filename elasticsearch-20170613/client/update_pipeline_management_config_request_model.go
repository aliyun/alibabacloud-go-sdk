// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePipelineManagementConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndpoints(v []*string) *UpdatePipelineManagementConfigRequest
	GetEndpoints() []*string
	SetEsInstanceId(v string) *UpdatePipelineManagementConfigRequest
	GetEsInstanceId() *string
	SetPassword(v string) *UpdatePipelineManagementConfigRequest
	GetPassword() *string
	SetPipelineIds(v []*string) *UpdatePipelineManagementConfigRequest
	GetPipelineIds() []*string
	SetPipelineManagementType(v string) *UpdatePipelineManagementConfigRequest
	GetPipelineManagementType() *string
	SetUserName(v string) *UpdatePipelineManagementConfigRequest
	GetUserName() *string
	SetClientToken(v string) *UpdatePipelineManagementConfigRequest
	GetClientToken() *string
}

type UpdatePipelineManagementConfigRequest struct {
	Endpoints    []*string `json:"endpoints,omitempty" xml:"endpoints,omitempty" type:"Repeated"`
	EsInstanceId *string   `json:"esInstanceId,omitempty" xml:"esInstanceId,omitempty"`
	// example:
	//
	// ******
	Password    *string   `json:"password,omitempty" xml:"password,omitempty"`
	PipelineIds []*string `json:"pipelineIds,omitempty" xml:"pipelineIds,omitempty" type:"Repeated"`
	// example:
	//
	// ES
	PipelineManagementType *string `json:"pipelineManagementType,omitempty" xml:"pipelineManagementType,omitempty"`
	// example:
	//
	// elastic
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdatePipelineManagementConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePipelineManagementConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdatePipelineManagementConfigRequest) GetEndpoints() []*string {
	return s.Endpoints
}

func (s *UpdatePipelineManagementConfigRequest) GetEsInstanceId() *string {
	return s.EsInstanceId
}

func (s *UpdatePipelineManagementConfigRequest) GetPassword() *string {
	return s.Password
}

func (s *UpdatePipelineManagementConfigRequest) GetPipelineIds() []*string {
	return s.PipelineIds
}

func (s *UpdatePipelineManagementConfigRequest) GetPipelineManagementType() *string {
	return s.PipelineManagementType
}

func (s *UpdatePipelineManagementConfigRequest) GetUserName() *string {
	return s.UserName
}

func (s *UpdatePipelineManagementConfigRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdatePipelineManagementConfigRequest) SetEndpoints(v []*string) *UpdatePipelineManagementConfigRequest {
	s.Endpoints = v
	return s
}

func (s *UpdatePipelineManagementConfigRequest) SetEsInstanceId(v string) *UpdatePipelineManagementConfigRequest {
	s.EsInstanceId = &v
	return s
}

func (s *UpdatePipelineManagementConfigRequest) SetPassword(v string) *UpdatePipelineManagementConfigRequest {
	s.Password = &v
	return s
}

func (s *UpdatePipelineManagementConfigRequest) SetPipelineIds(v []*string) *UpdatePipelineManagementConfigRequest {
	s.PipelineIds = v
	return s
}

func (s *UpdatePipelineManagementConfigRequest) SetPipelineManagementType(v string) *UpdatePipelineManagementConfigRequest {
	s.PipelineManagementType = &v
	return s
}

func (s *UpdatePipelineManagementConfigRequest) SetUserName(v string) *UpdatePipelineManagementConfigRequest {
	s.UserName = &v
	return s
}

func (s *UpdatePipelineManagementConfigRequest) SetClientToken(v string) *UpdatePipelineManagementConfigRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdatePipelineManagementConfigRequest) Validate() error {
	return dara.Validate(s)
}

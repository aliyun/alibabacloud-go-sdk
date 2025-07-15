// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationName(v string) *UpdateApplicationGroupRequest
	GetApplicationName() *string
	SetDeployedRevisionId(v string) *UpdateApplicationGroupRequest
	GetDeployedRevisionId() *string
	SetName(v string) *UpdateApplicationGroupRequest
	GetName() *string
	SetNewName(v string) *UpdateApplicationGroupRequest
	GetNewName() *string
	SetOperationName(v string) *UpdateApplicationGroupRequest
	GetOperationName() *string
	SetParameters(v map[string]interface{}) *UpdateApplicationGroupRequest
	GetParameters() map[string]interface{}
	SetRegionId(v string) *UpdateApplicationGroupRequest
	GetRegionId() *string
}

type UpdateApplicationGroupRequest struct {
	// The application name.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyApplication
	ApplicationName    *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	DeployedRevisionId *string `json:"DeployedRevisionId,omitempty" xml:"DeployedRevisionId,omitempty"`
	// The name of the application group.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyApplicationGroup
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The new name of the application group.
	//
	// example:
	//
	// UpdateMyApplicationGroup
	NewName *string `json:"NewName,omitempty" xml:"NewName,omitempty"`
	// The name of the configuration update operation.
	//
	// example:
	//
	// /business/v1/product/spus/{spu_id}
	OperationName *string `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	// The JSON string that consists of a set of parameters. Default value: {}.
	//
	// example:
	//
	// {"username": "xx"}
	Parameters map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID. Set the value to cn-hangzhou.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateApplicationGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateApplicationGroupRequest) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *UpdateApplicationGroupRequest) GetDeployedRevisionId() *string {
	return s.DeployedRevisionId
}

func (s *UpdateApplicationGroupRequest) GetName() *string {
	return s.Name
}

func (s *UpdateApplicationGroupRequest) GetNewName() *string {
	return s.NewName
}

func (s *UpdateApplicationGroupRequest) GetOperationName() *string {
	return s.OperationName
}

func (s *UpdateApplicationGroupRequest) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *UpdateApplicationGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateApplicationGroupRequest) SetApplicationName(v string) *UpdateApplicationGroupRequest {
	s.ApplicationName = &v
	return s
}

func (s *UpdateApplicationGroupRequest) SetDeployedRevisionId(v string) *UpdateApplicationGroupRequest {
	s.DeployedRevisionId = &v
	return s
}

func (s *UpdateApplicationGroupRequest) SetName(v string) *UpdateApplicationGroupRequest {
	s.Name = &v
	return s
}

func (s *UpdateApplicationGroupRequest) SetNewName(v string) *UpdateApplicationGroupRequest {
	s.NewName = &v
	return s
}

func (s *UpdateApplicationGroupRequest) SetOperationName(v string) *UpdateApplicationGroupRequest {
	s.OperationName = &v
	return s
}

func (s *UpdateApplicationGroupRequest) SetParameters(v map[string]interface{}) *UpdateApplicationGroupRequest {
	s.Parameters = v
	return s
}

func (s *UpdateApplicationGroupRequest) SetRegionId(v string) *UpdateApplicationGroupRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateApplicationGroupRequest) Validate() error {
	return dara.Validate(s)
}

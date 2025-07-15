// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationGroupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationName(v string) *UpdateApplicationGroupShrinkRequest
	GetApplicationName() *string
	SetDeployedRevisionId(v string) *UpdateApplicationGroupShrinkRequest
	GetDeployedRevisionId() *string
	SetName(v string) *UpdateApplicationGroupShrinkRequest
	GetName() *string
	SetNewName(v string) *UpdateApplicationGroupShrinkRequest
	GetNewName() *string
	SetOperationName(v string) *UpdateApplicationGroupShrinkRequest
	GetOperationName() *string
	SetParametersShrink(v string) *UpdateApplicationGroupShrinkRequest
	GetParametersShrink() *string
	SetRegionId(v string) *UpdateApplicationGroupShrinkRequest
	GetRegionId() *string
}

type UpdateApplicationGroupShrinkRequest struct {
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
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID. Set the value to cn-hangzhou.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateApplicationGroupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateApplicationGroupShrinkRequest) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *UpdateApplicationGroupShrinkRequest) GetDeployedRevisionId() *string {
	return s.DeployedRevisionId
}

func (s *UpdateApplicationGroupShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateApplicationGroupShrinkRequest) GetNewName() *string {
	return s.NewName
}

func (s *UpdateApplicationGroupShrinkRequest) GetOperationName() *string {
	return s.OperationName
}

func (s *UpdateApplicationGroupShrinkRequest) GetParametersShrink() *string {
	return s.ParametersShrink
}

func (s *UpdateApplicationGroupShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateApplicationGroupShrinkRequest) SetApplicationName(v string) *UpdateApplicationGroupShrinkRequest {
	s.ApplicationName = &v
	return s
}

func (s *UpdateApplicationGroupShrinkRequest) SetDeployedRevisionId(v string) *UpdateApplicationGroupShrinkRequest {
	s.DeployedRevisionId = &v
	return s
}

func (s *UpdateApplicationGroupShrinkRequest) SetName(v string) *UpdateApplicationGroupShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateApplicationGroupShrinkRequest) SetNewName(v string) *UpdateApplicationGroupShrinkRequest {
	s.NewName = &v
	return s
}

func (s *UpdateApplicationGroupShrinkRequest) SetOperationName(v string) *UpdateApplicationGroupShrinkRequest {
	s.OperationName = &v
	return s
}

func (s *UpdateApplicationGroupShrinkRequest) SetParametersShrink(v string) *UpdateApplicationGroupShrinkRequest {
	s.ParametersShrink = &v
	return s
}

func (s *UpdateApplicationGroupShrinkRequest) SetRegionId(v string) *UpdateApplicationGroupShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateApplicationGroupShrinkRequest) Validate() error {
	return dara.Validate(s)
}

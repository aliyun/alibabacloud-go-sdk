// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataServiceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiGatewayGroupId(v string) *CreateDataServiceGroupRequest
	GetApiGatewayGroupId() *string
	SetDescription(v string) *CreateDataServiceGroupRequest
	GetDescription() *string
	SetGroupName(v string) *CreateDataServiceGroupRequest
	GetGroupName() *string
	SetProjectId(v int64) *CreateDataServiceGroupRequest
	GetProjectId() *int64
	SetTenantId(v int64) *CreateDataServiceGroupRequest
	GetTenantId() *int64
}

type CreateDataServiceGroupRequest struct {
	// The ID of the API group that is associated with the business process in the API Gateway console. You can log on to the API Gateway console and go to the Group Details page to view the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000abcd
	ApiGatewayGroupId *string `json:"ApiGatewayGroupId,omitempty" xml:"ApiGatewayGroupId,omitempty"`
	// The description of the business process.
	//
	// example:
	//
	// Test business process
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the business process.
	//
	// This parameter is required.
	//
	// example:
	//
	// Business process name
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The ID of the workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10001
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The tenant ID. This parameter is deprecated.
	//
	// example:
	//
	// 10002
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s CreateDataServiceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataServiceGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateDataServiceGroupRequest) GetApiGatewayGroupId() *string {
	return s.ApiGatewayGroupId
}

func (s *CreateDataServiceGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDataServiceGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateDataServiceGroupRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateDataServiceGroupRequest) GetTenantId() *int64 {
	return s.TenantId
}

func (s *CreateDataServiceGroupRequest) SetApiGatewayGroupId(v string) *CreateDataServiceGroupRequest {
	s.ApiGatewayGroupId = &v
	return s
}

func (s *CreateDataServiceGroupRequest) SetDescription(v string) *CreateDataServiceGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateDataServiceGroupRequest) SetGroupName(v string) *CreateDataServiceGroupRequest {
	s.GroupName = &v
	return s
}

func (s *CreateDataServiceGroupRequest) SetProjectId(v int64) *CreateDataServiceGroupRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateDataServiceGroupRequest) SetTenantId(v int64) *CreateDataServiceGroupRequest {
	s.TenantId = &v
	return s
}

func (s *CreateDataServiceGroupRequest) Validate() error {
	return dara.Validate(s)
}

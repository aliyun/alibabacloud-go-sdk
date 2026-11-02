// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGatewayApiGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateGatewayApiGroupRequest
	GetAppId() *string
	SetCenters(v map[string]interface{}) *CreateGatewayApiGroupRequest
	GetCenters() map[string]interface{}
	SetDescription(v string) *CreateGatewayApiGroupRequest
	GetDescription() *string
	SetDomainName(v string) *CreateGatewayApiGroupRequest
	GetDomainName() *string
	SetEnableMTLS(v bool) *CreateGatewayApiGroupRequest
	GetEnableMTLS() *bool
	SetGroupName(v string) *CreateGatewayApiGroupRequest
	GetGroupName() *string
	SetHost(v string) *CreateGatewayApiGroupRequest
	GetHost() *string
	SetHttpsModel(v bool) *CreateGatewayApiGroupRequest
	GetHttpsModel() *bool
	SetInstanceId(v string) *CreateGatewayApiGroupRequest
	GetInstanceId() *string
	SetInstancePort(v int32) *CreateGatewayApiGroupRequest
	GetInstancePort() *int32
	SetRegionMode(v bool) *CreateGatewayApiGroupRequest
	GetRegionMode() *bool
	SetTenantId(v string) *CreateGatewayApiGroupRequest
	GetTenantId() *string
	SetTimeout(v int32) *CreateGatewayApiGroupRequest
	GetTimeout() *int32
	SetUpstreamType(v string) *CreateGatewayApiGroupRequest
	GetUpstreamType() *string
	SetVpcId(v string) *CreateGatewayApiGroupRequest
	GetVpcId() *string
	SetVpcModel(v bool) *CreateGatewayApiGroupRequest
	GetVpcModel() *bool
	SetWorkspaceId(v string) *CreateGatewayApiGroupRequest
	GetWorkspaceId() *string
}

type CreateGatewayApiGroupRequest struct {
	// This parameter is required.
	AppId       *string                `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Centers     map[string]interface{} `json:"Centers,omitempty" xml:"Centers,omitempty"`
	Description *string                `json:"Description,omitempty" xml:"Description,omitempty"`
	DomainName  *string                `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EnableMTLS  *bool                  `json:"EnableMTLS,omitempty" xml:"EnableMTLS,omitempty"`
	// This parameter is required.
	GroupName    *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	Host         *string `json:"Host,omitempty" xml:"Host,omitempty"`
	HttpsModel   *bool   `json:"HttpsModel,omitempty" xml:"HttpsModel,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstancePort *int32  `json:"InstancePort,omitempty" xml:"InstancePort,omitempty"`
	RegionMode   *bool   `json:"RegionMode,omitempty" xml:"RegionMode,omitempty"`
	// This parameter is required.
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// This parameter is required.
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// This parameter is required.
	UpstreamType *string `json:"UpstreamType,omitempty" xml:"UpstreamType,omitempty"`
	VpcId        *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpcModel     *bool   `json:"VpcModel,omitempty" xml:"VpcModel,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateGatewayApiGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayApiGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateGatewayApiGroupRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateGatewayApiGroupRequest) GetCenters() map[string]interface{} {
	return s.Centers
}

func (s *CreateGatewayApiGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateGatewayApiGroupRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *CreateGatewayApiGroupRequest) GetEnableMTLS() *bool {
	return s.EnableMTLS
}

func (s *CreateGatewayApiGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateGatewayApiGroupRequest) GetHost() *string {
	return s.Host
}

func (s *CreateGatewayApiGroupRequest) GetHttpsModel() *bool {
	return s.HttpsModel
}

func (s *CreateGatewayApiGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateGatewayApiGroupRequest) GetInstancePort() *int32 {
	return s.InstancePort
}

func (s *CreateGatewayApiGroupRequest) GetRegionMode() *bool {
	return s.RegionMode
}

func (s *CreateGatewayApiGroupRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateGatewayApiGroupRequest) GetTimeout() *int32 {
	return s.Timeout
}

func (s *CreateGatewayApiGroupRequest) GetUpstreamType() *string {
	return s.UpstreamType
}

func (s *CreateGatewayApiGroupRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateGatewayApiGroupRequest) GetVpcModel() *bool {
	return s.VpcModel
}

func (s *CreateGatewayApiGroupRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateGatewayApiGroupRequest) SetAppId(v string) *CreateGatewayApiGroupRequest {
	s.AppId = &v
	return s
}

func (s *CreateGatewayApiGroupRequest) SetCenters(v map[string]interface{}) *CreateGatewayApiGroupRequest {
	s.Centers = v
	return s
}

func (s *CreateGatewayApiGroupRequest) SetDescription(v string) *CreateGatewayApiGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateGatewayApiGroupRequest) SetDomainName(v string) *CreateGatewayApiGroupRequest {
	s.DomainName = &v
	return s
}

func (s *CreateGatewayApiGroupRequest) SetEnableMTLS(v bool) *CreateGatewayApiGroupRequest {
	s.EnableMTLS = &v
	return s
}

func (s *CreateGatewayApiGroupRequest) SetGroupName(v string) *CreateGatewayApiGroupRequest {
	s.GroupName = &v
	return s
}

func (s *CreateGatewayApiGroupRequest) SetHost(v string) *CreateGatewayApiGroupRequest {
	s.Host = &v
	return s
}

func (s *CreateGatewayApiGroupRequest) SetHttpsModel(v bool) *CreateGatewayApiGroupRequest {
	s.HttpsModel = &v
	return s
}

func (s *CreateGatewayApiGroupRequest) SetInstanceId(v string) *CreateGatewayApiGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateGatewayApiGroupRequest) SetInstancePort(v int32) *CreateGatewayApiGroupRequest {
	s.InstancePort = &v
	return s
}

func (s *CreateGatewayApiGroupRequest) SetRegionMode(v bool) *CreateGatewayApiGroupRequest {
	s.RegionMode = &v
	return s
}

func (s *CreateGatewayApiGroupRequest) SetTenantId(v string) *CreateGatewayApiGroupRequest {
	s.TenantId = &v
	return s
}

func (s *CreateGatewayApiGroupRequest) SetTimeout(v int32) *CreateGatewayApiGroupRequest {
	s.Timeout = &v
	return s
}

func (s *CreateGatewayApiGroupRequest) SetUpstreamType(v string) *CreateGatewayApiGroupRequest {
	s.UpstreamType = &v
	return s
}

func (s *CreateGatewayApiGroupRequest) SetVpcId(v string) *CreateGatewayApiGroupRequest {
	s.VpcId = &v
	return s
}

func (s *CreateGatewayApiGroupRequest) SetVpcModel(v bool) *CreateGatewayApiGroupRequest {
	s.VpcModel = &v
	return s
}

func (s *CreateGatewayApiGroupRequest) SetWorkspaceId(v string) *CreateGatewayApiGroupRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateGatewayApiGroupRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayApiGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateGatewayApiGroupRequest
	GetAppId() *string
	SetCenters(v map[string]interface{}) *UpdateGatewayApiGroupRequest
	GetCenters() map[string]interface{}
	SetDescription(v string) *UpdateGatewayApiGroupRequest
	GetDescription() *string
	SetDomainName(v string) *UpdateGatewayApiGroupRequest
	GetDomainName() *string
	SetEnableMTLS(v bool) *UpdateGatewayApiGroupRequest
	GetEnableMTLS() *bool
	SetGroupId(v int64) *UpdateGatewayApiGroupRequest
	GetGroupId() *int64
	SetGroupName(v string) *UpdateGatewayApiGroupRequest
	GetGroupName() *string
	SetHost(v string) *UpdateGatewayApiGroupRequest
	GetHost() *string
	SetHttpsModel(v bool) *UpdateGatewayApiGroupRequest
	GetHttpsModel() *bool
	SetInstanceId(v string) *UpdateGatewayApiGroupRequest
	GetInstanceId() *string
	SetInstancePort(v int32) *UpdateGatewayApiGroupRequest
	GetInstancePort() *int32
	SetRegionMode(v bool) *UpdateGatewayApiGroupRequest
	GetRegionMode() *bool
	SetTenantId(v string) *UpdateGatewayApiGroupRequest
	GetTenantId() *string
	SetTimeout(v int32) *UpdateGatewayApiGroupRequest
	GetTimeout() *int32
	SetUpstreamType(v string) *UpdateGatewayApiGroupRequest
	GetUpstreamType() *string
	SetVpcId(v string) *UpdateGatewayApiGroupRequest
	GetVpcId() *string
	SetVpcModel(v bool) *UpdateGatewayApiGroupRequest
	GetVpcModel() *bool
	SetWorkspaceId(v string) *UpdateGatewayApiGroupRequest
	GetWorkspaceId() *string
}

type UpdateGatewayApiGroupRequest struct {
	// This parameter is required.
	AppId       *string                `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Centers     map[string]interface{} `json:"Centers,omitempty" xml:"Centers,omitempty"`
	Description *string                `json:"Description,omitempty" xml:"Description,omitempty"`
	DomainName  *string                `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EnableMTLS  *bool                  `json:"EnableMTLS,omitempty" xml:"EnableMTLS,omitempty"`
	// This parameter is required.
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
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

func (s UpdateGatewayApiGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayApiGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayApiGroupRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateGatewayApiGroupRequest) GetCenters() map[string]interface{} {
	return s.Centers
}

func (s *UpdateGatewayApiGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateGatewayApiGroupRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *UpdateGatewayApiGroupRequest) GetEnableMTLS() *bool {
	return s.EnableMTLS
}

func (s *UpdateGatewayApiGroupRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *UpdateGatewayApiGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *UpdateGatewayApiGroupRequest) GetHost() *string {
	return s.Host
}

func (s *UpdateGatewayApiGroupRequest) GetHttpsModel() *bool {
	return s.HttpsModel
}

func (s *UpdateGatewayApiGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateGatewayApiGroupRequest) GetInstancePort() *int32 {
	return s.InstancePort
}

func (s *UpdateGatewayApiGroupRequest) GetRegionMode() *bool {
	return s.RegionMode
}

func (s *UpdateGatewayApiGroupRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateGatewayApiGroupRequest) GetTimeout() *int32 {
	return s.Timeout
}

func (s *UpdateGatewayApiGroupRequest) GetUpstreamType() *string {
	return s.UpstreamType
}

func (s *UpdateGatewayApiGroupRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *UpdateGatewayApiGroupRequest) GetVpcModel() *bool {
	return s.VpcModel
}

func (s *UpdateGatewayApiGroupRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateGatewayApiGroupRequest) SetAppId(v string) *UpdateGatewayApiGroupRequest {
	s.AppId = &v
	return s
}

func (s *UpdateGatewayApiGroupRequest) SetCenters(v map[string]interface{}) *UpdateGatewayApiGroupRequest {
	s.Centers = v
	return s
}

func (s *UpdateGatewayApiGroupRequest) SetDescription(v string) *UpdateGatewayApiGroupRequest {
	s.Description = &v
	return s
}

func (s *UpdateGatewayApiGroupRequest) SetDomainName(v string) *UpdateGatewayApiGroupRequest {
	s.DomainName = &v
	return s
}

func (s *UpdateGatewayApiGroupRequest) SetEnableMTLS(v bool) *UpdateGatewayApiGroupRequest {
	s.EnableMTLS = &v
	return s
}

func (s *UpdateGatewayApiGroupRequest) SetGroupId(v int64) *UpdateGatewayApiGroupRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateGatewayApiGroupRequest) SetGroupName(v string) *UpdateGatewayApiGroupRequest {
	s.GroupName = &v
	return s
}

func (s *UpdateGatewayApiGroupRequest) SetHost(v string) *UpdateGatewayApiGroupRequest {
	s.Host = &v
	return s
}

func (s *UpdateGatewayApiGroupRequest) SetHttpsModel(v bool) *UpdateGatewayApiGroupRequest {
	s.HttpsModel = &v
	return s
}

func (s *UpdateGatewayApiGroupRequest) SetInstanceId(v string) *UpdateGatewayApiGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateGatewayApiGroupRequest) SetInstancePort(v int32) *UpdateGatewayApiGroupRequest {
	s.InstancePort = &v
	return s
}

func (s *UpdateGatewayApiGroupRequest) SetRegionMode(v bool) *UpdateGatewayApiGroupRequest {
	s.RegionMode = &v
	return s
}

func (s *UpdateGatewayApiGroupRequest) SetTenantId(v string) *UpdateGatewayApiGroupRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateGatewayApiGroupRequest) SetTimeout(v int32) *UpdateGatewayApiGroupRequest {
	s.Timeout = &v
	return s
}

func (s *UpdateGatewayApiGroupRequest) SetUpstreamType(v string) *UpdateGatewayApiGroupRequest {
	s.UpstreamType = &v
	return s
}

func (s *UpdateGatewayApiGroupRequest) SetVpcId(v string) *UpdateGatewayApiGroupRequest {
	s.VpcId = &v
	return s
}

func (s *UpdateGatewayApiGroupRequest) SetVpcModel(v bool) *UpdateGatewayApiGroupRequest {
	s.VpcModel = &v
	return s
}

func (s *UpdateGatewayApiGroupRequest) SetWorkspaceId(v string) *UpdateGatewayApiGroupRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateGatewayApiGroupRequest) Validate() error {
	return dara.Validate(s)
}

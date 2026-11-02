// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayApiGroupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateGatewayApiGroupShrinkRequest
	GetAppId() *string
	SetCentersShrink(v string) *UpdateGatewayApiGroupShrinkRequest
	GetCentersShrink() *string
	SetDescription(v string) *UpdateGatewayApiGroupShrinkRequest
	GetDescription() *string
	SetDomainName(v string) *UpdateGatewayApiGroupShrinkRequest
	GetDomainName() *string
	SetEnableMTLS(v bool) *UpdateGatewayApiGroupShrinkRequest
	GetEnableMTLS() *bool
	SetGroupId(v int64) *UpdateGatewayApiGroupShrinkRequest
	GetGroupId() *int64
	SetGroupName(v string) *UpdateGatewayApiGroupShrinkRequest
	GetGroupName() *string
	SetHost(v string) *UpdateGatewayApiGroupShrinkRequest
	GetHost() *string
	SetHttpsModel(v bool) *UpdateGatewayApiGroupShrinkRequest
	GetHttpsModel() *bool
	SetInstanceId(v string) *UpdateGatewayApiGroupShrinkRequest
	GetInstanceId() *string
	SetInstancePort(v int32) *UpdateGatewayApiGroupShrinkRequest
	GetInstancePort() *int32
	SetRegionMode(v bool) *UpdateGatewayApiGroupShrinkRequest
	GetRegionMode() *bool
	SetTenantId(v string) *UpdateGatewayApiGroupShrinkRequest
	GetTenantId() *string
	SetTimeout(v int32) *UpdateGatewayApiGroupShrinkRequest
	GetTimeout() *int32
	SetUpstreamType(v string) *UpdateGatewayApiGroupShrinkRequest
	GetUpstreamType() *string
	SetVpcId(v string) *UpdateGatewayApiGroupShrinkRequest
	GetVpcId() *string
	SetVpcModel(v bool) *UpdateGatewayApiGroupShrinkRequest
	GetVpcModel() *bool
	SetWorkspaceId(v string) *UpdateGatewayApiGroupShrinkRequest
	GetWorkspaceId() *string
}

type UpdateGatewayApiGroupShrinkRequest struct {
	// This parameter is required.
	AppId         *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CentersShrink *string `json:"Centers,omitempty" xml:"Centers,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EnableMTLS    *bool   `json:"EnableMTLS,omitempty" xml:"EnableMTLS,omitempty"`
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

func (s UpdateGatewayApiGroupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayApiGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayApiGroupShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateGatewayApiGroupShrinkRequest) GetCentersShrink() *string {
	return s.CentersShrink
}

func (s *UpdateGatewayApiGroupShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateGatewayApiGroupShrinkRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *UpdateGatewayApiGroupShrinkRequest) GetEnableMTLS() *bool {
	return s.EnableMTLS
}

func (s *UpdateGatewayApiGroupShrinkRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *UpdateGatewayApiGroupShrinkRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *UpdateGatewayApiGroupShrinkRequest) GetHost() *string {
	return s.Host
}

func (s *UpdateGatewayApiGroupShrinkRequest) GetHttpsModel() *bool {
	return s.HttpsModel
}

func (s *UpdateGatewayApiGroupShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateGatewayApiGroupShrinkRequest) GetInstancePort() *int32 {
	return s.InstancePort
}

func (s *UpdateGatewayApiGroupShrinkRequest) GetRegionMode() *bool {
	return s.RegionMode
}

func (s *UpdateGatewayApiGroupShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateGatewayApiGroupShrinkRequest) GetTimeout() *int32 {
	return s.Timeout
}

func (s *UpdateGatewayApiGroupShrinkRequest) GetUpstreamType() *string {
	return s.UpstreamType
}

func (s *UpdateGatewayApiGroupShrinkRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *UpdateGatewayApiGroupShrinkRequest) GetVpcModel() *bool {
	return s.VpcModel
}

func (s *UpdateGatewayApiGroupShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateGatewayApiGroupShrinkRequest) SetAppId(v string) *UpdateGatewayApiGroupShrinkRequest {
	s.AppId = &v
	return s
}

func (s *UpdateGatewayApiGroupShrinkRequest) SetCentersShrink(v string) *UpdateGatewayApiGroupShrinkRequest {
	s.CentersShrink = &v
	return s
}

func (s *UpdateGatewayApiGroupShrinkRequest) SetDescription(v string) *UpdateGatewayApiGroupShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateGatewayApiGroupShrinkRequest) SetDomainName(v string) *UpdateGatewayApiGroupShrinkRequest {
	s.DomainName = &v
	return s
}

func (s *UpdateGatewayApiGroupShrinkRequest) SetEnableMTLS(v bool) *UpdateGatewayApiGroupShrinkRequest {
	s.EnableMTLS = &v
	return s
}

func (s *UpdateGatewayApiGroupShrinkRequest) SetGroupId(v int64) *UpdateGatewayApiGroupShrinkRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateGatewayApiGroupShrinkRequest) SetGroupName(v string) *UpdateGatewayApiGroupShrinkRequest {
	s.GroupName = &v
	return s
}

func (s *UpdateGatewayApiGroupShrinkRequest) SetHost(v string) *UpdateGatewayApiGroupShrinkRequest {
	s.Host = &v
	return s
}

func (s *UpdateGatewayApiGroupShrinkRequest) SetHttpsModel(v bool) *UpdateGatewayApiGroupShrinkRequest {
	s.HttpsModel = &v
	return s
}

func (s *UpdateGatewayApiGroupShrinkRequest) SetInstanceId(v string) *UpdateGatewayApiGroupShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateGatewayApiGroupShrinkRequest) SetInstancePort(v int32) *UpdateGatewayApiGroupShrinkRequest {
	s.InstancePort = &v
	return s
}

func (s *UpdateGatewayApiGroupShrinkRequest) SetRegionMode(v bool) *UpdateGatewayApiGroupShrinkRequest {
	s.RegionMode = &v
	return s
}

func (s *UpdateGatewayApiGroupShrinkRequest) SetTenantId(v string) *UpdateGatewayApiGroupShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateGatewayApiGroupShrinkRequest) SetTimeout(v int32) *UpdateGatewayApiGroupShrinkRequest {
	s.Timeout = &v
	return s
}

func (s *UpdateGatewayApiGroupShrinkRequest) SetUpstreamType(v string) *UpdateGatewayApiGroupShrinkRequest {
	s.UpstreamType = &v
	return s
}

func (s *UpdateGatewayApiGroupShrinkRequest) SetVpcId(v string) *UpdateGatewayApiGroupShrinkRequest {
	s.VpcId = &v
	return s
}

func (s *UpdateGatewayApiGroupShrinkRequest) SetVpcModel(v bool) *UpdateGatewayApiGroupShrinkRequest {
	s.VpcModel = &v
	return s
}

func (s *UpdateGatewayApiGroupShrinkRequest) SetWorkspaceId(v string) *UpdateGatewayApiGroupShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateGatewayApiGroupShrinkRequest) Validate() error {
	return dara.Validate(s)
}

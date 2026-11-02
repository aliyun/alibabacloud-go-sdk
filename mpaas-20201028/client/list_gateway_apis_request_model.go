// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayApisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiStatus(v string) *ListGatewayApisRequest
	GetApiStatus() *string
	SetApiType(v string) *ListGatewayApisRequest
	GetApiType() *string
	SetAppId(v string) *ListGatewayApisRequest
	GetAppId() *string
	SetGroupId(v int64) *ListGatewayApisRequest
	GetGroupId() *int64
	SetKeyword(v string) *ListGatewayApisRequest
	GetKeyword() *string
	SetPageNumber(v int32) *ListGatewayApisRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListGatewayApisRequest
	GetPageSize() *int32
	SetSysName(v string) *ListGatewayApisRequest
	GetSysName() *string
	SetTenantId(v string) *ListGatewayApisRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *ListGatewayApisRequest
	GetWorkspaceId() *string
}

type ListGatewayApisRequest struct {
	ApiStatus *string `json:"ApiStatus,omitempty" xml:"ApiStatus,omitempty"`
	ApiType   *string `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	// This parameter is required.
	AppId      *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	GroupId    *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Keyword    *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SysName    *string `json:"SysName,omitempty" xml:"SysName,omitempty"`
	// This parameter is required.
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListGatewayApisRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayApisRequest) GoString() string {
	return s.String()
}

func (s *ListGatewayApisRequest) GetApiStatus() *string {
	return s.ApiStatus
}

func (s *ListGatewayApisRequest) GetApiType() *string {
	return s.ApiType
}

func (s *ListGatewayApisRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListGatewayApisRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *ListGatewayApisRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListGatewayApisRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListGatewayApisRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListGatewayApisRequest) GetSysName() *string {
	return s.SysName
}

func (s *ListGatewayApisRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListGatewayApisRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListGatewayApisRequest) SetApiStatus(v string) *ListGatewayApisRequest {
	s.ApiStatus = &v
	return s
}

func (s *ListGatewayApisRequest) SetApiType(v string) *ListGatewayApisRequest {
	s.ApiType = &v
	return s
}

func (s *ListGatewayApisRequest) SetAppId(v string) *ListGatewayApisRequest {
	s.AppId = &v
	return s
}

func (s *ListGatewayApisRequest) SetGroupId(v int64) *ListGatewayApisRequest {
	s.GroupId = &v
	return s
}

func (s *ListGatewayApisRequest) SetKeyword(v string) *ListGatewayApisRequest {
	s.Keyword = &v
	return s
}

func (s *ListGatewayApisRequest) SetPageNumber(v int32) *ListGatewayApisRequest {
	s.PageNumber = &v
	return s
}

func (s *ListGatewayApisRequest) SetPageSize(v int32) *ListGatewayApisRequest {
	s.PageSize = &v
	return s
}

func (s *ListGatewayApisRequest) SetSysName(v string) *ListGatewayApisRequest {
	s.SysName = &v
	return s
}

func (s *ListGatewayApisRequest) SetTenantId(v string) *ListGatewayApisRequest {
	s.TenantId = &v
	return s
}

func (s *ListGatewayApisRequest) SetWorkspaceId(v string) *ListGatewayApisRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListGatewayApisRequest) Validate() error {
	return dara.Validate(s)
}

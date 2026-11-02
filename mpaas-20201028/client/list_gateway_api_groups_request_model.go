// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayApiGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListGatewayApiGroupsRequest
	GetAppId() *string
	SetKeyword(v string) *ListGatewayApiGroupsRequest
	GetKeyword() *string
	SetPageNumber(v int32) *ListGatewayApiGroupsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListGatewayApiGroupsRequest
	GetPageSize() *int32
	SetTenantId(v string) *ListGatewayApiGroupsRequest
	GetTenantId() *string
	SetUpstreamType(v string) *ListGatewayApiGroupsRequest
	GetUpstreamType() *string
	SetWorkspaceId(v string) *ListGatewayApiGroupsRequest
	GetWorkspaceId() *string
}

type ListGatewayApiGroupsRequest struct {
	// This parameter is required.
	AppId      *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Keyword    *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	TenantId     *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	UpstreamType *string `json:"UpstreamType,omitempty" xml:"UpstreamType,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListGatewayApiGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayApiGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListGatewayApiGroupsRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListGatewayApiGroupsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListGatewayApiGroupsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListGatewayApiGroupsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListGatewayApiGroupsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListGatewayApiGroupsRequest) GetUpstreamType() *string {
	return s.UpstreamType
}

func (s *ListGatewayApiGroupsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListGatewayApiGroupsRequest) SetAppId(v string) *ListGatewayApiGroupsRequest {
	s.AppId = &v
	return s
}

func (s *ListGatewayApiGroupsRequest) SetKeyword(v string) *ListGatewayApiGroupsRequest {
	s.Keyword = &v
	return s
}

func (s *ListGatewayApiGroupsRequest) SetPageNumber(v int32) *ListGatewayApiGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListGatewayApiGroupsRequest) SetPageSize(v int32) *ListGatewayApiGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListGatewayApiGroupsRequest) SetTenantId(v string) *ListGatewayApiGroupsRequest {
	s.TenantId = &v
	return s
}

func (s *ListGatewayApiGroupsRequest) SetUpstreamType(v string) *ListGatewayApiGroupsRequest {
	s.UpstreamType = &v
	return s
}

func (s *ListGatewayApiGroupsRequest) SetWorkspaceId(v string) *ListGatewayApiGroupsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListGatewayApiGroupsRequest) Validate() error {
	return dara.Validate(s)
}

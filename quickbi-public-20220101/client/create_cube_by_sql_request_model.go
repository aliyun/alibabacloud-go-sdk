// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCubeBySqlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCaption(v string) *CreateCubeBySqlRequest
	GetCaption() *string
	SetCustomSql(v string) *CreateCubeBySqlRequest
	GetCustomSql() *string
	SetDsId(v string) *CreateCubeBySqlRequest
	GetDsId() *string
	SetUserId(v string) *CreateCubeBySqlRequest
	GetUserId() *string
	SetWorkspaceId(v string) *CreateCubeBySqlRequest
	GetWorkspaceId() *string
}

type CreateCubeBySqlRequest struct {
	// The name of the custom dataset.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Caption *string `json:"Caption,omitempty" xml:"Caption,omitempty"`
	// The custom SQL statement.
	//
	// This parameter is required.
	//
	// example:
	//
	// select 	- from qqq
	CustomSql *string `json:"CustomSql,omitempty" xml:"CustomSql,omitempty"`
	// The ID of the data source.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7AAB95D-*****-****-*4FC0C976
	DsId *string `json:"DsId,omitempty" xml:"DsId,omitempty"`
	// The ID of a Quick BI user with permissions to create datasets. This is not your Alibaba Cloud account ID. Call the [QueryUserInfoByAccount](https://next.api.aliyun.com/api/quickbi-public/2022-01-01/QueryUserInfoByAccount?spm=api-workbench.api_explorer.0.0.672f50daGq9ooV\\&params=%7B%7D\\&tab=DOC\\&sdkStyle=old\\&RegionId=cn-hangzhou) operation to obtain the user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// asdaf-asda*****asd
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The ID of the workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// 95296e95-ca89-4c7d-8af9-dedf0ad0****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateCubeBySqlRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCubeBySqlRequest) GoString() string {
	return s.String()
}

func (s *CreateCubeBySqlRequest) GetCaption() *string {
	return s.Caption
}

func (s *CreateCubeBySqlRequest) GetCustomSql() *string {
	return s.CustomSql
}

func (s *CreateCubeBySqlRequest) GetDsId() *string {
	return s.DsId
}

func (s *CreateCubeBySqlRequest) GetUserId() *string {
	return s.UserId
}

func (s *CreateCubeBySqlRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateCubeBySqlRequest) SetCaption(v string) *CreateCubeBySqlRequest {
	s.Caption = &v
	return s
}

func (s *CreateCubeBySqlRequest) SetCustomSql(v string) *CreateCubeBySqlRequest {
	s.CustomSql = &v
	return s
}

func (s *CreateCubeBySqlRequest) SetDsId(v string) *CreateCubeBySqlRequest {
	s.DsId = &v
	return s
}

func (s *CreateCubeBySqlRequest) SetUserId(v string) *CreateCubeBySqlRequest {
	s.UserId = &v
	return s
}

func (s *CreateCubeBySqlRequest) SetWorkspaceId(v string) *CreateCubeBySqlRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateCubeBySqlRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCubeBySqlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCubeId(v string) *UpdateCubeBySqlRequest
	GetCubeId() *string
	SetCustomSql(v string) *UpdateCubeBySqlRequest
	GetCustomSql() *string
	SetDsId(v string) *UpdateCubeBySqlRequest
	GetDsId() *string
	SetPlaceholders(v string) *UpdateCubeBySqlRequest
	GetPlaceholders() *string
	SetUserId(v string) *UpdateCubeBySqlRequest
	GetUserId() *string
	SetWorkspaceId(v string) *UpdateCubeBySqlRequest
	GetWorkspaceId() *string
}

type UpdateCubeBySqlRequest struct {
	// The dataset ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7c7223ae-****-3c744528014b
	CubeId *string `json:"CubeId,omitempty" xml:"CubeId,omitempty"`
	// The custom SQL statement.
	//
	// This parameter is required.
	//
	// example:
	//
	// select 	- from sdsd
	CustomSql *string `json:"CustomSql,omitempty" xml:"CustomSql,omitempty"`
	// The data source ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7AAB95D-*****-****-*4FC0C976
	DsId *string `json:"DsId,omitempty" xml:"DsId,omitempty"`
	// The placeholder parameters. For more information, see the supplementary description below.
	//
	// example:
	//
	// [
	//
	//     {
	//
	//         "name": "test",
	//
	//         "style": "placeholder",
	//
	//         "type": "string"
	//
	//     }
	//
	// ]
	Placeholders *string `json:"Placeholders,omitempty" xml:"Placeholders,omitempty"`
	// The Quick BI UserId of a user who has permissions to create datasets. This is not your Alibaba Cloud account ID. Call the [QueryUserInfoByAccount](https://next.api.aliyun.com/api/quickbi-public/2022-01-01/QueryUserInfoByAccount?spm=api-workbench.api_explorer.0.0.672f50daGq9ooV&params=%7B%7D&tab=DOC&sdkStyle=old&RegionId=cn-hangzhou) operation to obtain the UserId.
	//
	// This parameter is required.
	//
	// example:
	//
	// 95296e95-ca89-4c7d-8af9-dedf0ad0****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 95296e95-ca89-4c7d-8af9-dedf0ad0****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateCubeBySqlRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCubeBySqlRequest) GoString() string {
	return s.String()
}

func (s *UpdateCubeBySqlRequest) GetCubeId() *string {
	return s.CubeId
}

func (s *UpdateCubeBySqlRequest) GetCustomSql() *string {
	return s.CustomSql
}

func (s *UpdateCubeBySqlRequest) GetDsId() *string {
	return s.DsId
}

func (s *UpdateCubeBySqlRequest) GetPlaceholders() *string {
	return s.Placeholders
}

func (s *UpdateCubeBySqlRequest) GetUserId() *string {
	return s.UserId
}

func (s *UpdateCubeBySqlRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateCubeBySqlRequest) SetCubeId(v string) *UpdateCubeBySqlRequest {
	s.CubeId = &v
	return s
}

func (s *UpdateCubeBySqlRequest) SetCustomSql(v string) *UpdateCubeBySqlRequest {
	s.CustomSql = &v
	return s
}

func (s *UpdateCubeBySqlRequest) SetDsId(v string) *UpdateCubeBySqlRequest {
	s.DsId = &v
	return s
}

func (s *UpdateCubeBySqlRequest) SetPlaceholders(v string) *UpdateCubeBySqlRequest {
	s.Placeholders = &v
	return s
}

func (s *UpdateCubeBySqlRequest) SetUserId(v string) *UpdateCubeBySqlRequest {
	s.UserId = &v
	return s
}

func (s *UpdateCubeBySqlRequest) SetWorkspaceId(v string) *UpdateCubeBySqlRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateCubeBySqlRequest) Validate() error {
	return dara.Validate(s)
}

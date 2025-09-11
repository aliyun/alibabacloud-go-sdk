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
	// This parameter is required.
	Caption *string `json:"Caption,omitempty" xml:"Caption,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// select 	- from qqq
	CustomSql *string `json:"CustomSql,omitempty" xml:"CustomSql,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 7AAB95D-*****-****-*4FC0C976
	DsId *string `json:"DsId,omitempty" xml:"DsId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// asdaf-asda*****asd
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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

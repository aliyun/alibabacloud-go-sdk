// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshDataSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataSourceId(v string) *RefreshDataSourceRequest
	GetDataSourceId() *string
	SetLang(v string) *RefreshDataSourceRequest
	GetLang() *string
	SetRegionId(v string) *RefreshDataSourceRequest
	GetRegionId() *string
	SetRoleFor(v int64) *RefreshDataSourceRequest
	GetRoleFor() *int64
}

type RefreshDataSourceRequest struct {
	// The ID of the data source.
	//
	// example:
	//
	// ds-jl67vixpe1scwysgyu3x
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// The language of the messages. Valid values:
	//
	// - **zh*	- (default): Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The region where the Data Management Center for threat analysis is located. Select the region of the Data Management Center based on the region of your assets. Valid values:
	//
	// - cn-hangzhou: Your assets are in the Chinese mainland or China (Hong Kong).
	//
	// - ap-southeast-1: Your assets are in a region outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of a member. An administrator can use this parameter to switch to the perspective of the specified member.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s RefreshDataSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s RefreshDataSourceRequest) GoString() string {
	return s.String()
}

func (s *RefreshDataSourceRequest) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *RefreshDataSourceRequest) GetLang() *string {
	return s.Lang
}

func (s *RefreshDataSourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RefreshDataSourceRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *RefreshDataSourceRequest) SetDataSourceId(v string) *RefreshDataSourceRequest {
	s.DataSourceId = &v
	return s
}

func (s *RefreshDataSourceRequest) SetLang(v string) *RefreshDataSourceRequest {
	s.Lang = &v
	return s
}

func (s *RefreshDataSourceRequest) SetRegionId(v string) *RefreshDataSourceRequest {
	s.RegionId = &v
	return s
}

func (s *RefreshDataSourceRequest) SetRoleFor(v int64) *RefreshDataSourceRequest {
	s.RoleFor = &v
	return s
}

func (s *RefreshDataSourceRequest) Validate() error {
	return dara.Validate(s)
}

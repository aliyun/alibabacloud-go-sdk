// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataSourceId(v string) *DeleteDataSourceRequest
	GetDataSourceId() *string
	SetLang(v string) *DeleteDataSourceRequest
	GetLang() *string
	SetRegionId(v string) *DeleteDataSourceRequest
	GetRegionId() *string
	SetRoleFor(v int64) *DeleteDataSourceRequest
	GetRoleFor() *int64
}

type DeleteDataSourceRequest struct {
	// The ID of the data source.
	//
	// example:
	//
	// ds-txejfbrh94k5cx58a4qh
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// The language of the response. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The region where the Data Management center for threat analysis is located. Select a region based on the location of your assets. Valid values:
	//
	// - cn-hangzhou: Your assets are in the Chinese mainland.
	//
	// - ap-southeast-1: Your assets are in a region outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of the member whose permissions you want to use. This parameter is available only to administrators.
	//
	// example:
	//
	// 173326*******
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s DeleteDataSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataSourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataSourceRequest) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *DeleteDataSourceRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteDataSourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteDataSourceRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *DeleteDataSourceRequest) SetDataSourceId(v string) *DeleteDataSourceRequest {
	s.DataSourceId = &v
	return s
}

func (s *DeleteDataSourceRequest) SetLang(v string) *DeleteDataSourceRequest {
	s.Lang = &v
	return s
}

func (s *DeleteDataSourceRequest) SetRegionId(v string) *DeleteDataSourceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDataSourceRequest) SetRoleFor(v int64) *DeleteDataSourceRequest {
	s.RoleFor = &v
	return s
}

func (s *DeleteDataSourceRequest) Validate() error {
	return dara.Validate(s)
}

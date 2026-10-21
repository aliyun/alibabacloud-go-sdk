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
	// The data source ID.
	//
	// Note: This parameter identifies the target of the delete operation and is required. Only data sources with a DataSourceType of custom can be deleted. Data sources of the preset type cannot be deleted.
	//
	// example:
	//
	// ds-txejfbrh94k5cx58a4qh
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// The language of the response messages. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// Note: This parameter is required. If it is not provided, the API returns a parameter validation error: "The specified Lang can not be empty."
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The region where the data management center for Cloud Threat Detection and Response (CTDR) is located. Select the management center region based on the region where your assets reside. Valid values:
	//
	// - cn-hangzhou: Assets are located in the Chinese mainland.
	//
	// - ap-southeast-1: Assets are located outside China.
	//
	// - cn-shanghai: The actual data storage region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID that the administrator uses to switch to the view of another member.
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

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataIngestionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataIngestionId(v string) *DeleteDataIngestionRequest
	GetDataIngestionId() *string
	SetLang(v string) *DeleteDataIngestionRequest
	GetLang() *string
	SetRegionId(v string) *DeleteDataIngestionRequest
	GetRegionId() *string
	SetRoleFor(v int64) *DeleteDataIngestionRequest
	GetRoleFor() *int64
}

type DeleteDataIngestionRequest struct {
	// The data ingestion ID.
	//
	// example:
	//
	// alibaba_cloud_sas_netstat_ingestion_173326*******
	DataIngestionId *string `json:"DataIngestionId,omitempty" xml:"DataIngestionId,omitempty"`
	// The language of the response message. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The region of the Management Hub for threat analysis. Select the region of the Management Hub based on the region where your assets are located. Valid values:
	//
	// - cn-hangzhou: Your assets are in the Chinese mainland.
	//
	// - ap-southeast-1: Your assets are in a region outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of a member. This parameter allows an administrator to operate on behalf of the member.
	//
	// example:
	//
	// 173326*******
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s DeleteDataIngestionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataIngestionRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataIngestionRequest) GetDataIngestionId() *string {
	return s.DataIngestionId
}

func (s *DeleteDataIngestionRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteDataIngestionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteDataIngestionRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *DeleteDataIngestionRequest) SetDataIngestionId(v string) *DeleteDataIngestionRequest {
	s.DataIngestionId = &v
	return s
}

func (s *DeleteDataIngestionRequest) SetLang(v string) *DeleteDataIngestionRequest {
	s.Lang = &v
	return s
}

func (s *DeleteDataIngestionRequest) SetRegionId(v string) *DeleteDataIngestionRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDataIngestionRequest) SetRoleFor(v int64) *DeleteDataIngestionRequest {
	s.RoleFor = &v
	return s
}

func (s *DeleteDataIngestionRequest) Validate() error {
	return dara.Validate(s)
}

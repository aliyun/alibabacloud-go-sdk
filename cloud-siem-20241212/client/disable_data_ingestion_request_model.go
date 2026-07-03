// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableDataIngestionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataIngestionId(v string) *DisableDataIngestionRequest
	GetDataIngestionId() *string
	SetLang(v string) *DisableDataIngestionRequest
	GetLang() *string
	SetRegionId(v string) *DisableDataIngestionRequest
	GetRegionId() *string
	SetRoleFor(v int64) *DisableDataIngestionRequest
	GetRoleFor() *int64
}

type DisableDataIngestionRequest struct {
	// The data ingestion ID.
	//
	// example:
	//
	// alibaba_cloud_sas_netstat_ingestion_173326*******
	DataIngestionId *string `json:"DataIngestionId,omitempty" xml:"DataIngestionId,omitempty"`
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
	// The region of the Data Management center for threat analysis. Select the region of the management center based on the region where your asset is located. Valid values:
	//
	// - cn-hangzhou: Your asset is in the Chinese mainland.
	//
	// - ap-southeast-1: Your asset is outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of the member account that the administrator wants to switch to.
	//
	// example:
	//
	// 173326*******
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s DisableDataIngestionRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableDataIngestionRequest) GoString() string {
	return s.String()
}

func (s *DisableDataIngestionRequest) GetDataIngestionId() *string {
	return s.DataIngestionId
}

func (s *DisableDataIngestionRequest) GetLang() *string {
	return s.Lang
}

func (s *DisableDataIngestionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DisableDataIngestionRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *DisableDataIngestionRequest) SetDataIngestionId(v string) *DisableDataIngestionRequest {
	s.DataIngestionId = &v
	return s
}

func (s *DisableDataIngestionRequest) SetLang(v string) *DisableDataIngestionRequest {
	s.Lang = &v
	return s
}

func (s *DisableDataIngestionRequest) SetRegionId(v string) *DisableDataIngestionRequest {
	s.RegionId = &v
	return s
}

func (s *DisableDataIngestionRequest) SetRoleFor(v int64) *DisableDataIngestionRequest {
	s.RoleFor = &v
	return s
}

func (s *DisableDataIngestionRequest) Validate() error {
	return dara.Validate(s)
}

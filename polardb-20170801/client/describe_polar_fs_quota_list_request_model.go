// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarFsQuotaListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribePolarFsQuotaListRequest
	GetDBClusterId() *string
	SetPageNumber(v int32) *DescribePolarFsQuotaListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribePolarFsQuotaListRequest
	GetPageSize() *int32
	SetPolarFsInstanceId(v string) *DescribePolarFsQuotaListRequest
	GetPolarFsInstanceId() *string
	SetQuotaMode(v string) *DescribePolarFsQuotaListRequest
	GetQuotaMode() *string
	SetRegionId(v string) *DescribePolarFsQuotaListRequest
	GetRegionId() *string
}

type DescribePolarFsQuotaListRequest struct {
	// example:
	//
	// pc-************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	PageNumber  *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pfs-2ze0i74ka607*****
	PolarFsInstanceId *string `json:"PolarFsInstanceId,omitempty" xml:"PolarFsInstanceId,omitempty"`
	// example:
	//
	// quotaPolicy
	QuotaMode *string `json:"QuotaMode,omitempty" xml:"QuotaMode,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribePolarFsQuotaListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarFsQuotaListRequest) GoString() string {
	return s.String()
}

func (s *DescribePolarFsQuotaListRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribePolarFsQuotaListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribePolarFsQuotaListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePolarFsQuotaListRequest) GetPolarFsInstanceId() *string {
	return s.PolarFsInstanceId
}

func (s *DescribePolarFsQuotaListRequest) GetQuotaMode() *string {
	return s.QuotaMode
}

func (s *DescribePolarFsQuotaListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePolarFsQuotaListRequest) SetDBClusterId(v string) *DescribePolarFsQuotaListRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribePolarFsQuotaListRequest) SetPageNumber(v int32) *DescribePolarFsQuotaListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribePolarFsQuotaListRequest) SetPageSize(v int32) *DescribePolarFsQuotaListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePolarFsQuotaListRequest) SetPolarFsInstanceId(v string) *DescribePolarFsQuotaListRequest {
	s.PolarFsInstanceId = &v
	return s
}

func (s *DescribePolarFsQuotaListRequest) SetQuotaMode(v string) *DescribePolarFsQuotaListRequest {
	s.QuotaMode = &v
	return s
}

func (s *DescribePolarFsQuotaListRequest) SetRegionId(v string) *DescribePolarFsQuotaListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePolarFsQuotaListRequest) Validate() error {
	return dara.Validate(s)
}

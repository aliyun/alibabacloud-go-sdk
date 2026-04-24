// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarFsQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribePolarFsQuotaRequest
	GetDBClusterId() *string
	SetPath(v string) *DescribePolarFsQuotaRequest
	GetPath() *string
	SetPolarFsInstanceId(v string) *DescribePolarFsQuotaRequest
	GetPolarFsInstanceId() *string
	SetQuotaType(v string) *DescribePolarFsQuotaRequest
	GetQuotaType() *string
	SetRegionId(v string) *DescribePolarFsQuotaRequest
	GetRegionId() *string
}

type DescribePolarFsQuotaRequest struct {
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// /data
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pfs-2ze0i74ka607*****
	PolarFsInstanceId *string `json:"PolarFsInstanceId,omitempty" xml:"PolarFsInstanceId,omitempty"`
	// example:
	//
	// quotaPolicy
	QuotaType *string `json:"QuotaType,omitempty" xml:"QuotaType,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribePolarFsQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarFsQuotaRequest) GoString() string {
	return s.String()
}

func (s *DescribePolarFsQuotaRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribePolarFsQuotaRequest) GetPath() *string {
	return s.Path
}

func (s *DescribePolarFsQuotaRequest) GetPolarFsInstanceId() *string {
	return s.PolarFsInstanceId
}

func (s *DescribePolarFsQuotaRequest) GetQuotaType() *string {
	return s.QuotaType
}

func (s *DescribePolarFsQuotaRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePolarFsQuotaRequest) SetDBClusterId(v string) *DescribePolarFsQuotaRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribePolarFsQuotaRequest) SetPath(v string) *DescribePolarFsQuotaRequest {
	s.Path = &v
	return s
}

func (s *DescribePolarFsQuotaRequest) SetPolarFsInstanceId(v string) *DescribePolarFsQuotaRequest {
	s.PolarFsInstanceId = &v
	return s
}

func (s *DescribePolarFsQuotaRequest) SetQuotaType(v string) *DescribePolarFsQuotaRequest {
	s.QuotaType = &v
	return s
}

func (s *DescribePolarFsQuotaRequest) SetRegionId(v string) *DescribePolarFsQuotaRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePolarFsQuotaRequest) Validate() error {
	return dara.Validate(s)
}

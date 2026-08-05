// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarFsMappingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribePolarFsMappingRequest
	GetDBClusterId() *string
	SetPageNumber(v int32) *DescribePolarFsMappingRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribePolarFsMappingRequest
	GetPageSize() *int32
	SetPolarFsInstanceId(v string) *DescribePolarFsMappingRequest
	GetPolarFsInstanceId() *string
}

type DescribePolarFsMappingRequest struct {
	// The cluster ID.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to query the details of all clusters under your account, including the cluster ID.
	//
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// 	- **30**
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// Default value: **30**.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The PolarFS instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pfs-2ze0i7*****
	PolarFsInstanceId *string `json:"PolarFsInstanceId,omitempty" xml:"PolarFsInstanceId,omitempty"`
}

func (s DescribePolarFsMappingRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarFsMappingRequest) GoString() string {
	return s.String()
}

func (s *DescribePolarFsMappingRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribePolarFsMappingRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribePolarFsMappingRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePolarFsMappingRequest) GetPolarFsInstanceId() *string {
	return s.PolarFsInstanceId
}

func (s *DescribePolarFsMappingRequest) SetDBClusterId(v string) *DescribePolarFsMappingRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribePolarFsMappingRequest) SetPageNumber(v int32) *DescribePolarFsMappingRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribePolarFsMappingRequest) SetPageSize(v int32) *DescribePolarFsMappingRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePolarFsMappingRequest) SetPolarFsInstanceId(v string) *DescribePolarFsMappingRequest {
	s.PolarFsInstanceId = &v
	return s
}

func (s *DescribePolarFsMappingRequest) Validate() error {
	return dara.Validate(s)
}

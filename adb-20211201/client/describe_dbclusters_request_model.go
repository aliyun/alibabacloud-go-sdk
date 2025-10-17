// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClustersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterDescription(v string) *DescribeDBClustersRequest
	GetDBClusterDescription() *string
	SetDBClusterIds(v string) *DescribeDBClustersRequest
	GetDBClusterIds() *string
	SetDBClusterStatus(v string) *DescribeDBClustersRequest
	GetDBClusterStatus() *string
	SetDBClusterVersion(v string) *DescribeDBClustersRequest
	GetDBClusterVersion() *string
	SetPageNumber(v int32) *DescribeDBClustersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDBClustersRequest
	GetPageSize() *int32
	SetProductVersion(v string) *DescribeDBClustersRequest
	GetProductVersion() *string
	SetRegionId(v string) *DescribeDBClustersRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeDBClustersRequest
	GetResourceGroupId() *string
	SetTag(v []*DescribeDBClustersRequestTag) *DescribeDBClustersRequest
	GetTag() []*DescribeDBClustersRequestTag
}

type DescribeDBClustersRequest struct {
	// The description of the cluster.
	//
	// 	- The description cannot start with `http://` or `https://`.
	//
	// 	- The description must be 2 to 256 characters in length
	//
	// example:
	//
	// test
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// If you do not specify this parameter, the information about all clusters that reside in the region is returned.
	//
	// example:
	//
	// amv-bp1r053byu48p****
	DBClusterIds *string `json:"DBClusterIds,omitempty" xml:"DBClusterIds,omitempty"`
	// The status of the cluster. Valid values:
	//
	// 	- **Preparing**
	//
	// 	- **Creating**
	//
	// 	- **Running**
	//
	// 	- **Deleting**
	//
	// 	- **Restoring**
	//
	// 	- **ClassChanging**
	//
	// 	- **NetAddressCreating**
	//
	// 	- **NetAddressDeleting**
	//
	// 	- **NetAddressModifying**
	//
	// example:
	//
	// Running
	DBClusterStatus *string `json:"DBClusterStatus,omitempty" xml:"DBClusterStatus,omitempty"`
	// The version number corresponding to the edition of the cluster. Valid values:
	//
	// 	- **3.0**: Data Warehouse Edition.
	//
	// 	- **5.0*	- (default): includes Data Lakehouse Edition, Enterprise Edition, and Basic Edition.
	//
	// 	- **All**: all editions, including Data Warehouse Edition, Data Lakehouse Edition, Enterprise Edition, and Basic Edition.
	//
	// example:
	//
	// 5.0
	DBClusterVersion *string `json:"DBClusterVersion,omitempty" xml:"DBClusterVersion,omitempty"`
	// The page number. Pages start from page 1. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// 	- **30*	- (default)
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The edition of the cluster. Valid values:
	//
	// 	- **EnterpriseVersion**: Enterprise Edition.
	//
	// 	- **BasicVersion**: Basic Edition.
	//
	// >  If you leave this parameter empty, the information about clusters of all editions is returned.
	//
	// example:
	//
	// BasicVersion
	ProductVersion *string `json:"ProductVersion,omitempty" xml:"ProductVersion,omitempty"`
	// The region ID of the cluster.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/454314.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID. If you do not specify this parameter, the information about all resource groups in the cluster is returned.
	//
	// example:
	//
	// rg-4690g37929****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags that are added to the cluster.
	Tag []*DescribeDBClustersRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDBClustersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClustersRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersRequest) GetDBClusterDescription() *string {
	return s.DBClusterDescription
}

func (s *DescribeDBClustersRequest) GetDBClusterIds() *string {
	return s.DBClusterIds
}

func (s *DescribeDBClustersRequest) GetDBClusterStatus() *string {
	return s.DBClusterStatus
}

func (s *DescribeDBClustersRequest) GetDBClusterVersion() *string {
	return s.DBClusterVersion
}

func (s *DescribeDBClustersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDBClustersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDBClustersRequest) GetProductVersion() *string {
	return s.ProductVersion
}

func (s *DescribeDBClustersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBClustersRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDBClustersRequest) GetTag() []*DescribeDBClustersRequestTag {
	return s.Tag
}

func (s *DescribeDBClustersRequest) SetDBClusterDescription(v string) *DescribeDBClustersRequest {
	s.DBClusterDescription = &v
	return s
}

func (s *DescribeDBClustersRequest) SetDBClusterIds(v string) *DescribeDBClustersRequest {
	s.DBClusterIds = &v
	return s
}

func (s *DescribeDBClustersRequest) SetDBClusterStatus(v string) *DescribeDBClustersRequest {
	s.DBClusterStatus = &v
	return s
}

func (s *DescribeDBClustersRequest) SetDBClusterVersion(v string) *DescribeDBClustersRequest {
	s.DBClusterVersion = &v
	return s
}

func (s *DescribeDBClustersRequest) SetPageNumber(v int32) *DescribeDBClustersRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBClustersRequest) SetPageSize(v int32) *DescribeDBClustersRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDBClustersRequest) SetProductVersion(v string) *DescribeDBClustersRequest {
	s.ProductVersion = &v
	return s
}

func (s *DescribeDBClustersRequest) SetRegionId(v string) *DescribeDBClustersRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClustersRequest) SetResourceGroupId(v string) *DescribeDBClustersRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBClustersRequest) SetTag(v []*DescribeDBClustersRequestTag) *DescribeDBClustersRequest {
	s.Tag = v
	return s
}

func (s *DescribeDBClustersRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBClustersRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// tag1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// test1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBClustersRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClustersRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeDBClustersRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeDBClustersRequestTag) SetKey(v string) *DescribeDBClustersRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeDBClustersRequestTag) SetValue(v string) *DescribeDBClustersRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeDBClustersRequestTag) Validate() error {
	return dara.Validate(s)
}

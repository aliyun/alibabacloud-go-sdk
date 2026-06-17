// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGlobalDatabaseNetworksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeGlobalDatabaseNetworksResponseBodyItems) *DescribeGlobalDatabaseNetworksResponseBody
	GetItems() []*DescribeGlobalDatabaseNetworksResponseBodyItems
	SetPageNumber(v int32) *DescribeGlobalDatabaseNetworksResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeGlobalDatabaseNetworksResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeGlobalDatabaseNetworksResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *DescribeGlobalDatabaseNetworksResponseBody
	GetTotalRecordCount() *int32
}

type DescribeGlobalDatabaseNetworksResponseBody struct {
	// A list of GDNs.
	Items []*DescribeGlobalDatabaseNetworksResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records on the current page.
	//
	// example:
	//
	// 30
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 69A85BAF-1089-4CDF-A82F-0A140F******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 1
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeGlobalDatabaseNetworksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalDatabaseNetworksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDatabaseNetworksResponseBody) GetItems() []*DescribeGlobalDatabaseNetworksResponseBodyItems {
	return s.Items
}

func (s *DescribeGlobalDatabaseNetworksResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeGlobalDatabaseNetworksResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeGlobalDatabaseNetworksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGlobalDatabaseNetworksResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeGlobalDatabaseNetworksResponseBody) SetItems(v []*DescribeGlobalDatabaseNetworksResponseBodyItems) *DescribeGlobalDatabaseNetworksResponseBody {
	s.Items = v
	return s
}

func (s *DescribeGlobalDatabaseNetworksResponseBody) SetPageNumber(v int32) *DescribeGlobalDatabaseNetworksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworksResponseBody) SetPageRecordCount(v int32) *DescribeGlobalDatabaseNetworksResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworksResponseBody) SetRequestId(v string) *DescribeGlobalDatabaseNetworksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworksResponseBody) SetTotalRecordCount(v int32) *DescribeGlobalDatabaseNetworksResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworksResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeGlobalDatabaseNetworksResponseBodyItems struct {
	// The time at which the GDN was created, in UTC. The format is `YYYY-MM-DDTHH:mm:ssZ`.
	//
	// example:
	//
	// 2020-03-23T05:46:54Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// A list of clusters in the GDN.
	DBClusters []*DescribeGlobalDatabaseNetworksResponseBodyItemsDBClusters `json:"DBClusters,omitempty" xml:"DBClusters,omitempty" type:"Repeated"`
	// The database engine type. Only **MySQL*	- is supported.
	//
	// example:
	//
	// MySQL
	DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	// The database engine version. Only version **8.0*	- is supported.
	//
	// example:
	//
	// 8.0
	DBVersion *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	// The GDN description. Requirements:
	//
	// - Cannot start with http\\:// or https\\://.
	//
	// - Must start with a letter or a Chinese character.
	//
	// - Can contain letters, Chinese characters, digits, underscores (_), or hyphens (-).
	//
	// - Must be 2 to 126 characters long.
	//
	// example:
	//
	// test
	GDNDescription *string `json:"GDNDescription,omitempty" xml:"GDNDescription,omitempty"`
	// The GDN ID.
	//
	// example:
	//
	// gdn-****************
	GDNId *string `json:"GDNId,omitempty" xml:"GDNId,omitempty"`
	// The status of the GDN. Valid values:
	//
	// - **creating**: The GDN is being created.
	//
	// - **active**: The GDN is running.
	//
	// - **deleting**: The GDN is being deleted.
	//
	// - **locked**: The GDN is locked. This status prevents any operations on clusters in the GDN.
	//
	// - **removing_member**: A standby cluster is being removed from the GDN.
	//
	// example:
	//
	// active
	GDNStatus *string `json:"GDNStatus,omitempty" xml:"GDNStatus,omitempty"`
	// The tags applied to the GDN.
	Labels *DescribeGlobalDatabaseNetworksResponseBodyItemsLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Struct"`
}

func (s DescribeGlobalDatabaseNetworksResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalDatabaseNetworksResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDatabaseNetworksResponseBodyItems) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeGlobalDatabaseNetworksResponseBodyItems) GetDBClusters() []*DescribeGlobalDatabaseNetworksResponseBodyItemsDBClusters {
	return s.DBClusters
}

func (s *DescribeGlobalDatabaseNetworksResponseBodyItems) GetDBType() *string {
	return s.DBType
}

func (s *DescribeGlobalDatabaseNetworksResponseBodyItems) GetDBVersion() *string {
	return s.DBVersion
}

func (s *DescribeGlobalDatabaseNetworksResponseBodyItems) GetGDNDescription() *string {
	return s.GDNDescription
}

func (s *DescribeGlobalDatabaseNetworksResponseBodyItems) GetGDNId() *string {
	return s.GDNId
}

func (s *DescribeGlobalDatabaseNetworksResponseBodyItems) GetGDNStatus() *string {
	return s.GDNStatus
}

func (s *DescribeGlobalDatabaseNetworksResponseBodyItems) GetLabels() *DescribeGlobalDatabaseNetworksResponseBodyItemsLabels {
	return s.Labels
}

func (s *DescribeGlobalDatabaseNetworksResponseBodyItems) SetCreateTime(v string) *DescribeGlobalDatabaseNetworksResponseBodyItems {
	s.CreateTime = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworksResponseBodyItems) SetDBClusters(v []*DescribeGlobalDatabaseNetworksResponseBodyItemsDBClusters) *DescribeGlobalDatabaseNetworksResponseBodyItems {
	s.DBClusters = v
	return s
}

func (s *DescribeGlobalDatabaseNetworksResponseBodyItems) SetDBType(v string) *DescribeGlobalDatabaseNetworksResponseBodyItems {
	s.DBType = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworksResponseBodyItems) SetDBVersion(v string) *DescribeGlobalDatabaseNetworksResponseBodyItems {
	s.DBVersion = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworksResponseBodyItems) SetGDNDescription(v string) *DescribeGlobalDatabaseNetworksResponseBodyItems {
	s.GDNDescription = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworksResponseBodyItems) SetGDNId(v string) *DescribeGlobalDatabaseNetworksResponseBodyItems {
	s.GDNId = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworksResponseBodyItems) SetGDNStatus(v string) *DescribeGlobalDatabaseNetworksResponseBodyItems {
	s.GDNStatus = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworksResponseBodyItems) SetLabels(v *DescribeGlobalDatabaseNetworksResponseBodyItemsLabels) *DescribeGlobalDatabaseNetworksResponseBodyItems {
	s.Labels = v
	return s
}

func (s *DescribeGlobalDatabaseNetworksResponseBodyItems) Validate() error {
	if s.DBClusters != nil {
		for _, item := range s.DBClusters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Labels != nil {
		if err := s.Labels.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeGlobalDatabaseNetworksResponseBodyItemsDBClusters struct {
	// The cluster ID.
	//
	// example:
	//
	// pc-****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The role of the cluster. Valid values:
	//
	// - **primary**: The primary cluster.
	//
	// - **standby**: The standby cluster.
	//
	// > A GDN consists of one primary cluster and up to four standby clusters. For more information, see [Global Database Network](https://help.aliyun.com/document_detail/160381.html).
	//
	// example:
	//
	// primary
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s DescribeGlobalDatabaseNetworksResponseBodyItemsDBClusters) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalDatabaseNetworksResponseBodyItemsDBClusters) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDatabaseNetworksResponseBodyItemsDBClusters) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeGlobalDatabaseNetworksResponseBodyItemsDBClusters) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeGlobalDatabaseNetworksResponseBodyItemsDBClusters) GetRole() *string {
	return s.Role
}

func (s *DescribeGlobalDatabaseNetworksResponseBodyItemsDBClusters) SetDBClusterId(v string) *DescribeGlobalDatabaseNetworksResponseBodyItemsDBClusters {
	s.DBClusterId = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworksResponseBodyItemsDBClusters) SetRegionId(v string) *DescribeGlobalDatabaseNetworksResponseBodyItemsDBClusters {
	s.RegionId = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworksResponseBodyItemsDBClusters) SetRole(v string) *DescribeGlobalDatabaseNetworksResponseBodyItemsDBClusters {
	s.Role = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworksResponseBodyItemsDBClusters) Validate() error {
	return dara.Validate(s)
}

type DescribeGlobalDatabaseNetworksResponseBodyItemsLabels struct {
	// The GDN version.
	//
	// example:
	//
	// 2.0
	GDNVersion *string `json:"GDNVersion,omitempty" xml:"GDNVersion,omitempty"`
}

func (s DescribeGlobalDatabaseNetworksResponseBodyItemsLabels) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalDatabaseNetworksResponseBodyItemsLabels) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDatabaseNetworksResponseBodyItemsLabels) GetGDNVersion() *string {
	return s.GDNVersion
}

func (s *DescribeGlobalDatabaseNetworksResponseBodyItemsLabels) SetGDNVersion(v string) *DescribeGlobalDatabaseNetworksResponseBodyItemsLabels {
	s.GDNVersion = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworksResponseBodyItemsLabels) Validate() error {
	return dara.Validate(s)
}

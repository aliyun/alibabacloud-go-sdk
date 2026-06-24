// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataLimitSetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataLimitSet(v *DescribeDataLimitSetResponseBodyDataLimitSet) *DescribeDataLimitSetResponseBody
	GetDataLimitSet() *DescribeDataLimitSetResponseBodyDataLimitSet
	SetRequestId(v string) *DescribeDataLimitSetResponseBody
	GetRequestId() *string
}

type DescribeDataLimitSetResponseBody struct {
	// Information about the authorized data assets.
	DataLimitSet *DescribeDataLimitSetResponseBodyDataLimitSet `json:"DataLimitSet,omitempty" xml:"DataLimitSet,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 769FB3C1-F4C9-42DF-9B72-7077A8989C13
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDataLimitSetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataLimitSetResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataLimitSetResponseBody) GetDataLimitSet() *DescribeDataLimitSetResponseBodyDataLimitSet {
	return s.DataLimitSet
}

func (s *DescribeDataLimitSetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDataLimitSetResponseBody) SetDataLimitSet(v *DescribeDataLimitSetResponseBodyDataLimitSet) *DescribeDataLimitSetResponseBody {
	s.DataLimitSet = v
	return s
}

func (s *DescribeDataLimitSetResponseBody) SetRequestId(v string) *DescribeDataLimitSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataLimitSetResponseBody) Validate() error {
	if s.DataLimitSet != nil {
		if err := s.DataLimitSet.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDataLimitSetResponseBodyDataLimitSet struct {
	// A list of authorized data assets.
	DataLimitList []*DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList `json:"DataLimitList,omitempty" xml:"DataLimitList,omitempty" type:"Repeated"`
	// A list of authorized OSS buckets.
	OssBucketList []*DescribeDataLimitSetResponseBodyDataLimitSetOssBucketList `json:"OssBucketList,omitempty" xml:"OssBucketList,omitempty" type:"Repeated"`
	// A list of regions that support scanning.
	RegionList []*DescribeDataLimitSetResponseBodyDataLimitSetRegionList `json:"RegionList,omitempty" xml:"RegionList,omitempty" type:"Repeated"`
	// The type of the data asset. Valid values:
	//
	// - **1**: MaxCompute.
	//
	// - **2**: OSS.
	//
	// - **3**: ADS.
	//
	// - **4**: OTS.
	//
	// - **5**: RDS.
	//
	// - **6**: SELF_DB.
	//
	// example:
	//
	// 2
	ResourceType *int64 `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The code for the data asset type. Valid values:
	//
	// - **MaxCompute**
	//
	// - **OSS**
	//
	// - **ADS**
	//
	// - **OTS**
	//
	// - **RDS**
	//
	// - **SELF_DB**
	//
	// example:
	//
	// OSS
	ResourceTypeCode *string `json:"ResourceTypeCode,omitempty" xml:"ResourceTypeCode,omitempty"`
	// The total number of assets found.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDataLimitSetResponseBodyDataLimitSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataLimitSetResponseBodyDataLimitSet) GoString() string {
	return s.String()
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSet) GetDataLimitList() []*DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList {
	return s.DataLimitList
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSet) GetOssBucketList() []*DescribeDataLimitSetResponseBodyDataLimitSetOssBucketList {
	return s.OssBucketList
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSet) GetRegionList() []*DescribeDataLimitSetResponseBodyDataLimitSetRegionList {
	return s.RegionList
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSet) GetResourceType() *int64 {
	return s.ResourceType
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSet) GetResourceTypeCode() *string {
	return s.ResourceTypeCode
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSet) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSet) SetDataLimitList(v []*DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList) *DescribeDataLimitSetResponseBodyDataLimitSet {
	s.DataLimitList = v
	return s
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSet) SetOssBucketList(v []*DescribeDataLimitSetResponseBodyDataLimitSetOssBucketList) *DescribeDataLimitSetResponseBodyDataLimitSet {
	s.OssBucketList = v
	return s
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSet) SetRegionList(v []*DescribeDataLimitSetResponseBodyDataLimitSetRegionList) *DescribeDataLimitSetResponseBodyDataLimitSet {
	s.RegionList = v
	return s
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSet) SetResourceType(v int64) *DescribeDataLimitSetResponseBodyDataLimitSet {
	s.ResourceType = &v
	return s
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSet) SetResourceTypeCode(v string) *DescribeDataLimitSetResponseBodyDataLimitSet {
	s.ResourceTypeCode = &v
	return s
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSet) SetTotalCount(v int32) *DescribeDataLimitSetResponseBodyDataLimitSet {
	s.TotalCount = &v
	return s
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSet) Validate() error {
	if s.DataLimitList != nil {
		for _, item := range s.DataLimitList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.OssBucketList != nil {
		for _, item := range s.OssBucketList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RegionList != nil {
		for _, item := range s.RegionList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList struct {
	// The status of the connectivity test between Security Center and the authorized data asset.
	//
	// - **2**: The connectivity test is in progress.
	//
	// - **3**: The connectivity test is passed.
	//
	// - **4**: The connectivity test has failed.
	//
	// example:
	//
	// 3
	CheckStatus *int32 `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	// The name of the connectivity test status.
	//
	// example:
	//
	// Connected
	CheckStatusName *string `json:"CheckStatusName,omitempty" xml:"CheckStatusName,omitempty"`
	// The connection string for the data asset.
	//
	// example:
	//
	// jdbc:mysql://10.*.*.94:3306/test_demo
	Connector *string `json:"Connector,omitempty" xml:"Connector,omitempty"`
	// The time when the data asset was created. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1625587423000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The unique ID of the data asset.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the region where the data asset is located.
	//
	// example:
	//
	// cn-hangzhou
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The ID of the parent asset.
	//
	// example:
	//
	// db
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The ID of the region where the data asset is located.
	//
	// example:
	//
	// cn-****
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the data asset. Valid values:
	//
	// - **1**: MaxCompute.
	//
	// - **2**: OSS.
	//
	// - **3**: ADS.
	//
	// - **4**: OTS.
	//
	// - **5**: RDS.
	//
	// - **6**: SELF_DB.
	//
	// example:
	//
	// 2
	ResourceType *int64 `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The code for the data asset type. Valid values:
	//
	// - **MaxCompute**
	//
	// - **OSS**
	//
	// - **ADS**
	//
	// - **OTS**
	//
	// - **RDS**
	//
	// - **SELF_DB**
	//
	// example:
	//
	// OSS
	ResourceTypeCode *string `json:"ResourceTypeCode,omitempty" xml:"ResourceTypeCode,omitempty"`
	// The username of the data owner.
	//
	// example:
	//
	// tsts
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList) GoString() string {
	return s.String()
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList) GetCheckStatus() *int32 {
	return s.CheckStatus
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList) GetCheckStatusName() *string {
	return s.CheckStatusName
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList) GetConnector() *string {
	return s.Connector
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList) GetId() *int64 {
	return s.Id
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList) GetLocalName() *string {
	return s.LocalName
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList) GetParentId() *string {
	return s.ParentId
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList) GetResourceType() *int64 {
	return s.ResourceType
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList) GetResourceTypeCode() *string {
	return s.ResourceTypeCode
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList) GetUserName() *string {
	return s.UserName
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList) SetCheckStatus(v int32) *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList {
	s.CheckStatus = &v
	return s
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList) SetCheckStatusName(v string) *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList {
	s.CheckStatusName = &v
	return s
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList) SetConnector(v string) *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList {
	s.Connector = &v
	return s
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList) SetGmtCreate(v int64) *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList {
	s.GmtCreate = &v
	return s
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList) SetId(v int64) *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList {
	s.Id = &v
	return s
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList) SetLocalName(v string) *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList {
	s.LocalName = &v
	return s
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList) SetParentId(v string) *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList {
	s.ParentId = &v
	return s
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList) SetRegionId(v string) *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList {
	s.RegionId = &v
	return s
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList) SetResourceType(v int64) *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList {
	s.ResourceType = &v
	return s
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList) SetResourceTypeCode(v string) *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList {
	s.ResourceTypeCode = &v
	return s
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList) SetUserName(v string) *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList {
	s.UserName = &v
	return s
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetDataLimitList) Validate() error {
	return dara.Validate(s)
}

type DescribeDataLimitSetResponseBodyDataLimitSetOssBucketList struct {
	// The name of the OSS bucket.
	//
	// example:
	//
	// oss-bucket
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The ID of the region where the OSS bucket is located.
	//
	// example:
	//
	// cn-****
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDataLimitSetResponseBodyDataLimitSetOssBucketList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataLimitSetResponseBodyDataLimitSetOssBucketList) GoString() string {
	return s.String()
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetOssBucketList) GetBucketName() *string {
	return s.BucketName
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetOssBucketList) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetOssBucketList) SetBucketName(v string) *DescribeDataLimitSetResponseBodyDataLimitSetOssBucketList {
	s.BucketName = &v
	return s
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetOssBucketList) SetRegionId(v string) *DescribeDataLimitSetResponseBodyDataLimitSetOssBucketList {
	s.RegionId = &v
	return s
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetOssBucketList) Validate() error {
	return dara.Validate(s)
}

type DescribeDataLimitSetResponseBodyDataLimitSetRegionList struct {
	// The region name.
	//
	// example:
	//
	// cn-hangzhou
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-****
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDataLimitSetResponseBodyDataLimitSetRegionList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataLimitSetResponseBodyDataLimitSetRegionList) GoString() string {
	return s.String()
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetRegionList) GetLocalName() *string {
	return s.LocalName
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetRegionList) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetRegionList) SetLocalName(v string) *DescribeDataLimitSetResponseBodyDataLimitSetRegionList {
	s.LocalName = &v
	return s
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetRegionList) SetRegionId(v string) *DescribeDataLimitSetResponseBodyDataLimitSetRegionList {
	s.RegionId = &v
	return s
}

func (s *DescribeDataLimitSetResponseBodyDataLimitSetRegionList) Validate() error {
	return dara.Validate(s)
}

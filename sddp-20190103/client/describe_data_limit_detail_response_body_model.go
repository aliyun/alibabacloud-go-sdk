// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataLimitDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataLimit(v *DescribeDataLimitDetailResponseBodyDataLimit) *DescribeDataLimitDetailResponseBody
	GetDataLimit() *DescribeDataLimitDetailResponseBodyDataLimit
	SetRequestId(v string) *DescribeDataLimitDetailResponseBody
	GetRequestId() *string
}

type DescribeDataLimitDetailResponseBody struct {
	// The details of the data asset.
	DataLimit *DescribeDataLimitDetailResponseBodyDataLimit `json:"DataLimit,omitempty" xml:"DataLimit,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 769FB3C1-F4C9-42DF-9B72-7077A8989C13
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDataLimitDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataLimitDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataLimitDetailResponseBody) GetDataLimit() *DescribeDataLimitDetailResponseBodyDataLimit {
	return s.DataLimit
}

func (s *DescribeDataLimitDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDataLimitDetailResponseBody) SetDataLimit(v *DescribeDataLimitDetailResponseBodyDataLimit) *DescribeDataLimitDetailResponseBody {
	s.DataLimit = v
	return s
}

func (s *DescribeDataLimitDetailResponseBody) SetRequestId(v string) *DescribeDataLimitDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataLimitDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDataLimitDetailResponseBodyDataLimit struct {
	// The status of the connectivity test between the data asset and DSC. Valid values:
	//
	// 	- **2**: indicates that the data asset was being connected.
	//
	// 	- **3**: indicates that the data asset was connected to DSC.
	//
	// 	- **4**: indicates that the data asset failed to be connected.
	//
	// example:
	//
	// 3
	CheckStatus *int32 `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	// The result that indicates the status of the connectivity test between the data asset and DSC. Valid values:
	//
	// 	- **Passed**
	//
	// 	- **Failed**
	//
	// 	- **Testing**
	//
	// example:
	//
	// Passed
	CheckStatusName *string `json:"CheckStatusName,omitempty" xml:"CheckStatusName,omitempty"`
	// The time when the data asset was connected to DSC. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 145600000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The ID of the data asset.
	//
	// example:
	//
	// 12300
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The region in which the data asset resides.
	//
	// example:
	//
	// China (Qingdao)
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The ID and name of the data asset in the service to which the data asset belongs.
	//
	// example:
	//
	// rm-m5eup49p6o274****.RDS_example
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The port number that is used to connect to the database.
	//
	// example:
	//
	// 3306
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the region in which the data asset resides.
	//
	// example:
	//
	// cn-qingdao
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the service to which the data asset belongs. Valid values:
	//
	// 	- **1**: MaxCompute
	//
	// 	- **2**: OSS
	//
	// 	- **3**: AnalyticDB for MySQL
	//
	// 	- **4**: Tablestore
	//
	// 	- **5**: ApsaraDB RDS
	//
	// example:
	//
	// 1
	ResourceType *int64 `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The service to which the data asset belongs. Valid values:
	//
	// 	- **MaxCompute**
	//
	// 	- **OSS**
	//
	// 	- **ADS**
	//
	// 	- **OTS**
	//
	// 	- **RDS**
	//
	// example:
	//
	// RDS
	ResourceTypeCode *string `json:"ResourceTypeCode,omitempty" xml:"ResourceTypeCode,omitempty"`
	// The account of the user who manages the data asset.
	//
	// example:
	//
	// User01
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeDataLimitDetailResponseBodyDataLimit) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataLimitDetailResponseBodyDataLimit) GoString() string {
	return s.String()
}

func (s *DescribeDataLimitDetailResponseBodyDataLimit) GetCheckStatus() *int32 {
	return s.CheckStatus
}

func (s *DescribeDataLimitDetailResponseBodyDataLimit) GetCheckStatusName() *string {
	return s.CheckStatusName
}

func (s *DescribeDataLimitDetailResponseBodyDataLimit) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeDataLimitDetailResponseBodyDataLimit) GetId() *int64 {
	return s.Id
}

func (s *DescribeDataLimitDetailResponseBodyDataLimit) GetLocalName() *string {
	return s.LocalName
}

func (s *DescribeDataLimitDetailResponseBodyDataLimit) GetParentId() *string {
	return s.ParentId
}

func (s *DescribeDataLimitDetailResponseBodyDataLimit) GetPort() *int32 {
	return s.Port
}

func (s *DescribeDataLimitDetailResponseBodyDataLimit) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDataLimitDetailResponseBodyDataLimit) GetResourceType() *int64 {
	return s.ResourceType
}

func (s *DescribeDataLimitDetailResponseBodyDataLimit) GetResourceTypeCode() *string {
	return s.ResourceTypeCode
}

func (s *DescribeDataLimitDetailResponseBodyDataLimit) GetUserName() *string {
	return s.UserName
}

func (s *DescribeDataLimitDetailResponseBodyDataLimit) SetCheckStatus(v int32) *DescribeDataLimitDetailResponseBodyDataLimit {
	s.CheckStatus = &v
	return s
}

func (s *DescribeDataLimitDetailResponseBodyDataLimit) SetCheckStatusName(v string) *DescribeDataLimitDetailResponseBodyDataLimit {
	s.CheckStatusName = &v
	return s
}

func (s *DescribeDataLimitDetailResponseBodyDataLimit) SetGmtCreate(v int64) *DescribeDataLimitDetailResponseBodyDataLimit {
	s.GmtCreate = &v
	return s
}

func (s *DescribeDataLimitDetailResponseBodyDataLimit) SetId(v int64) *DescribeDataLimitDetailResponseBodyDataLimit {
	s.Id = &v
	return s
}

func (s *DescribeDataLimitDetailResponseBodyDataLimit) SetLocalName(v string) *DescribeDataLimitDetailResponseBodyDataLimit {
	s.LocalName = &v
	return s
}

func (s *DescribeDataLimitDetailResponseBodyDataLimit) SetParentId(v string) *DescribeDataLimitDetailResponseBodyDataLimit {
	s.ParentId = &v
	return s
}

func (s *DescribeDataLimitDetailResponseBodyDataLimit) SetPort(v int32) *DescribeDataLimitDetailResponseBodyDataLimit {
	s.Port = &v
	return s
}

func (s *DescribeDataLimitDetailResponseBodyDataLimit) SetRegionId(v string) *DescribeDataLimitDetailResponseBodyDataLimit {
	s.RegionId = &v
	return s
}

func (s *DescribeDataLimitDetailResponseBodyDataLimit) SetResourceType(v int64) *DescribeDataLimitDetailResponseBodyDataLimit {
	s.ResourceType = &v
	return s
}

func (s *DescribeDataLimitDetailResponseBodyDataLimit) SetResourceTypeCode(v string) *DescribeDataLimitDetailResponseBodyDataLimit {
	s.ResourceTypeCode = &v
	return s
}

func (s *DescribeDataLimitDetailResponseBodyDataLimit) SetUserName(v string) *DescribeDataLimitDetailResponseBodyDataLimit {
	s.UserName = &v
	return s
}

func (s *DescribeDataLimitDetailResponseBodyDataLimit) Validate() error {
	return dara.Validate(s)
}

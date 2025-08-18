// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstances(v []*DescribeAppInstancesResponseBodyInstances) *DescribeAppInstancesResponseBody
	GetInstances() []*DescribeAppInstancesResponseBodyInstances
	SetMaxResults(v int64) *DescribeAppInstancesResponseBody
	GetMaxResults() *int64
	SetPageNumber(v int64) *DescribeAppInstancesResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeAppInstancesResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeAppInstancesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeAppInstancesResponseBody
	GetTotalCount() *int64
}

type DescribeAppInstancesResponseBody struct {
	Instances []*DescribeAppInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// example:
	//
	// None
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 32DEFB4A-861F-5D1D-ADD5-918E4FD7AB8C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 20
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAppInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppInstancesResponseBody) GetInstances() []*DescribeAppInstancesResponseBodyInstances {
	return s.Instances
}

func (s *DescribeAppInstancesResponseBody) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *DescribeAppInstancesResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeAppInstancesResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeAppInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAppInstancesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeAppInstancesResponseBody) SetInstances(v []*DescribeAppInstancesResponseBodyInstances) *DescribeAppInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeAppInstancesResponseBody) SetMaxResults(v int64) *DescribeAppInstancesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeAppInstancesResponseBody) SetPageNumber(v int64) *DescribeAppInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAppInstancesResponseBody) SetPageSize(v int64) *DescribeAppInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAppInstancesResponseBody) SetRequestId(v string) *DescribeAppInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppInstancesResponseBody) SetTotalCount(v int64) *DescribeAppInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAppInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeAppInstancesResponseBodyInstances struct {
	// example:
	//
	// test-supabase
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// supabase
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// example:
	//
	// pgm-2ze49qv594vi****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// rdsai.supabase.basic
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// example:
	//
	// 20241231
	InstanceMinorVersion *string `json:"InstanceMinorVersion,omitempty" xml:"InstanceMinorVersion,omitempty"`
	// example:
	//
	// ra-supabase-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// 8.152. XXX.XXX:8000
	PublicConnectionString *string `json:"PublicConnectionString,omitempty" xml:"PublicConnectionString,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// vsw-2zeaepb8k4ku05ov2****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// 172.16.XXX.XXX:8000
	VpcConnectionString *string `json:"VpcConnectionString,omitempty" xml:"VpcConnectionString,omitempty"`
}

func (s DescribeAppInstancesResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeAppInstancesResponseBodyInstances) GetAppName() *string {
	return s.AppName
}

func (s *DescribeAppInstancesResponseBodyInstances) GetAppType() *string {
	return s.AppType
}

func (s *DescribeAppInstancesResponseBodyInstances) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeAppInstancesResponseBodyInstances) GetInstanceClass() *string {
	return s.InstanceClass
}

func (s *DescribeAppInstancesResponseBodyInstances) GetInstanceMinorVersion() *string {
	return s.InstanceMinorVersion
}

func (s *DescribeAppInstancesResponseBodyInstances) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeAppInstancesResponseBodyInstances) GetPublicConnectionString() *string {
	return s.PublicConnectionString
}

func (s *DescribeAppInstancesResponseBodyInstances) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAppInstancesResponseBodyInstances) GetStatus() *string {
	return s.Status
}

func (s *DescribeAppInstancesResponseBodyInstances) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeAppInstancesResponseBodyInstances) GetVpcConnectionString() *string {
	return s.VpcConnectionString
}

func (s *DescribeAppInstancesResponseBodyInstances) SetAppName(v string) *DescribeAppInstancesResponseBodyInstances {
	s.AppName = &v
	return s
}

func (s *DescribeAppInstancesResponseBodyInstances) SetAppType(v string) *DescribeAppInstancesResponseBodyInstances {
	s.AppType = &v
	return s
}

func (s *DescribeAppInstancesResponseBodyInstances) SetDBInstanceName(v string) *DescribeAppInstancesResponseBodyInstances {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeAppInstancesResponseBodyInstances) SetInstanceClass(v string) *DescribeAppInstancesResponseBodyInstances {
	s.InstanceClass = &v
	return s
}

func (s *DescribeAppInstancesResponseBodyInstances) SetInstanceMinorVersion(v string) *DescribeAppInstancesResponseBodyInstances {
	s.InstanceMinorVersion = &v
	return s
}

func (s *DescribeAppInstancesResponseBodyInstances) SetInstanceName(v string) *DescribeAppInstancesResponseBodyInstances {
	s.InstanceName = &v
	return s
}

func (s *DescribeAppInstancesResponseBodyInstances) SetPublicConnectionString(v string) *DescribeAppInstancesResponseBodyInstances {
	s.PublicConnectionString = &v
	return s
}

func (s *DescribeAppInstancesResponseBodyInstances) SetRegionId(v string) *DescribeAppInstancesResponseBodyInstances {
	s.RegionId = &v
	return s
}

func (s *DescribeAppInstancesResponseBodyInstances) SetStatus(v string) *DescribeAppInstancesResponseBodyInstances {
	s.Status = &v
	return s
}

func (s *DescribeAppInstancesResponseBodyInstances) SetVSwitchId(v string) *DescribeAppInstancesResponseBodyInstances {
	s.VSwitchId = &v
	return s
}

func (s *DescribeAppInstancesResponseBodyInstances) SetVpcConnectionString(v string) *DescribeAppInstancesResponseBodyInstances {
	s.VpcConnectionString = &v
	return s
}

func (s *DescribeAppInstancesResponseBodyInstances) Validate() error {
	return dara.Validate(s)
}

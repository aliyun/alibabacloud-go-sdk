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
	// The information about the RDS Supabase instances.
	Instances []*DescribeAppInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// A reserved parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// None
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 32DEFB4A-861F-5D1D-ADD5-918E4FD7AB8C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
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
	if s.Instances != nil {
		for _, item := range s.Instances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAppInstancesResponseBodyInstances struct {
	// The name of the AI application.
	//
	// example:
	//
	// test-supabase
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The application type. Only **supabase*	- is supported. For more information, see [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html).
	//
	// example:
	//
	// supabase
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// The ID of the RDS for PostgreSQL instance with which the RDS Supabase instances are associated.
	//
	// example:
	//
	// pgm-2ze49qv594vi****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The instance type.
	//
	// example:
	//
	// rdsai.supabase.basic
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// The minor version number of RDS Supabase instance.
	//
	// example:
	//
	// 20241231
	InstanceMinorVersion *string `json:"InstanceMinorVersion,omitempty" xml:"InstanceMinorVersion,omitempty"`
	// The ID of the RDS Supabase instance.
	//
	// example:
	//
	// ra-supabase-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The public endpoint of the AI application.
	//
	// example:
	//
	// 8.152. XXX.XXX:8000
	PublicConnectionString *string `json:"PublicConnectionString,omitempty" xml:"PublicConnectionString,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The instance status. For more information, see [Instance state table](https://help.aliyun.com/document_detail/2623972.html).
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-2zeaepb8k4ku05ov2****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The internal endpoint of the AI application.
	//
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

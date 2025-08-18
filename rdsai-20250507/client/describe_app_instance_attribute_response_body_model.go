// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppInstanceAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DescribeAppInstanceAttributeResponseBody
	GetAppName() *string
	SetAppType(v string) *DescribeAppInstanceAttributeResponseBody
	GetAppType() *string
	SetDBInstanceName(v string) *DescribeAppInstanceAttributeResponseBody
	GetDBInstanceName() *string
	SetInstanceClass(v string) *DescribeAppInstanceAttributeResponseBody
	GetInstanceClass() *string
	SetInstanceMinorVersion(v string) *DescribeAppInstanceAttributeResponseBody
	GetInstanceMinorVersion() *string
	SetInstanceName(v string) *DescribeAppInstanceAttributeResponseBody
	GetInstanceName() *string
	SetPublicConnectionString(v string) *DescribeAppInstanceAttributeResponseBody
	GetPublicConnectionString() *string
	SetRegionId(v string) *DescribeAppInstanceAttributeResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeAppInstanceAttributeResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeAppInstanceAttributeResponseBody
	GetStatus() *string
	SetVSwitchId(v string) *DescribeAppInstanceAttributeResponseBody
	GetVSwitchId() *string
	SetVpcConnectionString(v string) *DescribeAppInstanceAttributeResponseBody
	GetVpcConnectionString() *string
	SetZoneId(v string) *DescribeAppInstanceAttributeResponseBody
	GetZoneId() *string
}

type DescribeAppInstanceAttributeResponseBody struct {
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
	// FE9C65D7-930F-57A5-A207-8C396329241C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// cn-beijing-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeAppInstanceAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppInstanceAttributeResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *DescribeAppInstanceAttributeResponseBody) GetAppType() *string {
	return s.AppType
}

func (s *DescribeAppInstanceAttributeResponseBody) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeAppInstanceAttributeResponseBody) GetInstanceClass() *string {
	return s.InstanceClass
}

func (s *DescribeAppInstanceAttributeResponseBody) GetInstanceMinorVersion() *string {
	return s.InstanceMinorVersion
}

func (s *DescribeAppInstanceAttributeResponseBody) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeAppInstanceAttributeResponseBody) GetPublicConnectionString() *string {
	return s.PublicConnectionString
}

func (s *DescribeAppInstanceAttributeResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAppInstanceAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAppInstanceAttributeResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeAppInstanceAttributeResponseBody) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeAppInstanceAttributeResponseBody) GetVpcConnectionString() *string {
	return s.VpcConnectionString
}

func (s *DescribeAppInstanceAttributeResponseBody) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeAppInstanceAttributeResponseBody) SetAppName(v string) *DescribeAppInstanceAttributeResponseBody {
	s.AppName = &v
	return s
}

func (s *DescribeAppInstanceAttributeResponseBody) SetAppType(v string) *DescribeAppInstanceAttributeResponseBody {
	s.AppType = &v
	return s
}

func (s *DescribeAppInstanceAttributeResponseBody) SetDBInstanceName(v string) *DescribeAppInstanceAttributeResponseBody {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeAppInstanceAttributeResponseBody) SetInstanceClass(v string) *DescribeAppInstanceAttributeResponseBody {
	s.InstanceClass = &v
	return s
}

func (s *DescribeAppInstanceAttributeResponseBody) SetInstanceMinorVersion(v string) *DescribeAppInstanceAttributeResponseBody {
	s.InstanceMinorVersion = &v
	return s
}

func (s *DescribeAppInstanceAttributeResponseBody) SetInstanceName(v string) *DescribeAppInstanceAttributeResponseBody {
	s.InstanceName = &v
	return s
}

func (s *DescribeAppInstanceAttributeResponseBody) SetPublicConnectionString(v string) *DescribeAppInstanceAttributeResponseBody {
	s.PublicConnectionString = &v
	return s
}

func (s *DescribeAppInstanceAttributeResponseBody) SetRegionId(v string) *DescribeAppInstanceAttributeResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeAppInstanceAttributeResponseBody) SetRequestId(v string) *DescribeAppInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppInstanceAttributeResponseBody) SetStatus(v string) *DescribeAppInstanceAttributeResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeAppInstanceAttributeResponseBody) SetVSwitchId(v string) *DescribeAppInstanceAttributeResponseBody {
	s.VSwitchId = &v
	return s
}

func (s *DescribeAppInstanceAttributeResponseBody) SetVpcConnectionString(v string) *DescribeAppInstanceAttributeResponseBody {
	s.VpcConnectionString = &v
	return s
}

func (s *DescribeAppInstanceAttributeResponseBody) SetZoneId(v string) *DescribeAppInstanceAttributeResponseBody {
	s.ZoneId = &v
	return s
}

func (s *DescribeAppInstanceAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}

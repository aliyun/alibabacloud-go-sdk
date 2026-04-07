// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMem0InfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeMem0InfoResponseBodyData) *DescribeMem0InfoResponseBody
	GetData() *DescribeMem0InfoResponseBodyData
	SetRequestId(v string) *DescribeMem0InfoResponseBody
	GetRequestId() *string
}

type DescribeMem0InfoResponseBody struct {
	Data *DescribeMem0InfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// C457B28E-9CAB-4B77-B5C6-5D71B7870B6E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMem0InfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMem0InfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMem0InfoResponseBody) GetData() *DescribeMem0InfoResponseBodyData {
	return s.Data
}

func (s *DescribeMem0InfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMem0InfoResponseBody) SetData(v *DescribeMem0InfoResponseBodyData) *DescribeMem0InfoResponseBody {
	s.Data = v
	return s
}

func (s *DescribeMem0InfoResponseBody) SetRequestId(v string) *DescribeMem0InfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMem0InfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMem0InfoResponseBodyData struct {
	Instance *DescribeMem0InfoResponseBodyDataInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Struct"`
}

func (s DescribeMem0InfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeMem0InfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeMem0InfoResponseBodyData) GetInstance() *DescribeMem0InfoResponseBodyDataInstance {
	return s.Instance
}

func (s *DescribeMem0InfoResponseBodyData) SetInstance(v *DescribeMem0InfoResponseBodyDataInstance) *DescribeMem0InfoResponseBodyData {
	s.Instance = v
	return s
}

func (s *DescribeMem0InfoResponseBodyData) Validate() error {
	if s.Instance != nil {
		if err := s.Instance.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMem0InfoResponseBodyDataInstance struct {
	// example:
	//
	// mysql.x2.large.2c
	ClassCode *string                                              `json:"ClassCode,omitempty" xml:"ClassCode,omitempty"`
	ConnAddrs []*DescribeMem0InfoResponseBodyDataInstanceConnAddrs `json:"ConnAddrs,omitempty" xml:"ConnAddrs,omitempty" type:"Repeated"`
	// example:
	//
	// 2026-02-17T02:00:20Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// pxc-spsil01pww4hfz-mem
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1
	NodeCount *int32 `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	// example:
	//
	// ap-southeast-1
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// local_ssd
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// VPC ID
	//
	// example:
	//
	// vpc-t4n4hf9xey7ea3lp4bwwx
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// example:
	//
	// vsw-t4ny14pr37spmjsbv5dc2
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// ap-southeast-1a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeMem0InfoResponseBodyDataInstance) String() string {
	return dara.Prettify(s)
}

func (s DescribeMem0InfoResponseBodyDataInstance) GoString() string {
	return s.String()
}

func (s *DescribeMem0InfoResponseBodyDataInstance) GetClassCode() *string {
	return s.ClassCode
}

func (s *DescribeMem0InfoResponseBodyDataInstance) GetConnAddrs() []*DescribeMem0InfoResponseBodyDataInstanceConnAddrs {
	return s.ConnAddrs
}

func (s *DescribeMem0InfoResponseBodyDataInstance) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeMem0InfoResponseBodyDataInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeMem0InfoResponseBodyDataInstance) GetNodeCount() *int32 {
	return s.NodeCount
}

func (s *DescribeMem0InfoResponseBodyDataInstance) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeMem0InfoResponseBodyDataInstance) GetStatus() *string {
	return s.Status
}

func (s *DescribeMem0InfoResponseBodyDataInstance) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeMem0InfoResponseBodyDataInstance) GetVPCId() *string {
	return s.VPCId
}

func (s *DescribeMem0InfoResponseBodyDataInstance) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeMem0InfoResponseBodyDataInstance) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeMem0InfoResponseBodyDataInstance) SetClassCode(v string) *DescribeMem0InfoResponseBodyDataInstance {
	s.ClassCode = &v
	return s
}

func (s *DescribeMem0InfoResponseBodyDataInstance) SetConnAddrs(v []*DescribeMem0InfoResponseBodyDataInstanceConnAddrs) *DescribeMem0InfoResponseBodyDataInstance {
	s.ConnAddrs = v
	return s
}

func (s *DescribeMem0InfoResponseBodyDataInstance) SetCreateTime(v string) *DescribeMem0InfoResponseBodyDataInstance {
	s.CreateTime = &v
	return s
}

func (s *DescribeMem0InfoResponseBodyDataInstance) SetInstanceId(v string) *DescribeMem0InfoResponseBodyDataInstance {
	s.InstanceId = &v
	return s
}

func (s *DescribeMem0InfoResponseBodyDataInstance) SetNodeCount(v int32) *DescribeMem0InfoResponseBodyDataInstance {
	s.NodeCount = &v
	return s
}

func (s *DescribeMem0InfoResponseBodyDataInstance) SetRegionId(v string) *DescribeMem0InfoResponseBodyDataInstance {
	s.RegionId = &v
	return s
}

func (s *DescribeMem0InfoResponseBodyDataInstance) SetStatus(v string) *DescribeMem0InfoResponseBodyDataInstance {
	s.Status = &v
	return s
}

func (s *DescribeMem0InfoResponseBodyDataInstance) SetStorageType(v string) *DescribeMem0InfoResponseBodyDataInstance {
	s.StorageType = &v
	return s
}

func (s *DescribeMem0InfoResponseBodyDataInstance) SetVPCId(v string) *DescribeMem0InfoResponseBodyDataInstance {
	s.VPCId = &v
	return s
}

func (s *DescribeMem0InfoResponseBodyDataInstance) SetVSwitchId(v string) *DescribeMem0InfoResponseBodyDataInstance {
	s.VSwitchId = &v
	return s
}

func (s *DescribeMem0InfoResponseBodyDataInstance) SetZoneId(v string) *DescribeMem0InfoResponseBodyDataInstance {
	s.ZoneId = &v
	return s
}

func (s *DescribeMem0InfoResponseBodyDataInstance) Validate() error {
	if s.ConnAddrs != nil {
		for _, item := range s.ConnAddrs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMem0InfoResponseBodyDataInstanceConnAddrs struct {
	// example:
	//
	// pxc-spsil01pww4hfz.polarx.singapore.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// example:
	//
	// 3306
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// example:
	//
	// VPC
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// VPC ID
	//
	// example:
	//
	// vpc-t4n4hf9xey7ea3lp4bwwx
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// example:
	//
	// vsw-t4ny14pr37spmjsbv5dc2
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// pxc-spsil01pww4hfzjayd-cn-20251013180429
	VpcInstanceId *string `json:"VpcInstanceId,omitempty" xml:"VpcInstanceId,omitempty"`
}

func (s DescribeMem0InfoResponseBodyDataInstanceConnAddrs) String() string {
	return dara.Prettify(s)
}

func (s DescribeMem0InfoResponseBodyDataInstanceConnAddrs) GoString() string {
	return s.String()
}

func (s *DescribeMem0InfoResponseBodyDataInstanceConnAddrs) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *DescribeMem0InfoResponseBodyDataInstanceConnAddrs) GetPort() *int32 {
	return s.Port
}

func (s *DescribeMem0InfoResponseBodyDataInstanceConnAddrs) GetType() *string {
	return s.Type
}

func (s *DescribeMem0InfoResponseBodyDataInstanceConnAddrs) GetVPCId() *string {
	return s.VPCId
}

func (s *DescribeMem0InfoResponseBodyDataInstanceConnAddrs) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeMem0InfoResponseBodyDataInstanceConnAddrs) GetVpcInstanceId() *string {
	return s.VpcInstanceId
}

func (s *DescribeMem0InfoResponseBodyDataInstanceConnAddrs) SetConnectionString(v string) *DescribeMem0InfoResponseBodyDataInstanceConnAddrs {
	s.ConnectionString = &v
	return s
}

func (s *DescribeMem0InfoResponseBodyDataInstanceConnAddrs) SetPort(v int32) *DescribeMem0InfoResponseBodyDataInstanceConnAddrs {
	s.Port = &v
	return s
}

func (s *DescribeMem0InfoResponseBodyDataInstanceConnAddrs) SetType(v string) *DescribeMem0InfoResponseBodyDataInstanceConnAddrs {
	s.Type = &v
	return s
}

func (s *DescribeMem0InfoResponseBodyDataInstanceConnAddrs) SetVPCId(v string) *DescribeMem0InfoResponseBodyDataInstanceConnAddrs {
	s.VPCId = &v
	return s
}

func (s *DescribeMem0InfoResponseBodyDataInstanceConnAddrs) SetVSwitchId(v string) *DescribeMem0InfoResponseBodyDataInstanceConnAddrs {
	s.VSwitchId = &v
	return s
}

func (s *DescribeMem0InfoResponseBodyDataInstanceConnAddrs) SetVpcInstanceId(v string) *DescribeMem0InfoResponseBodyDataInstanceConnAddrs {
	s.VpcInstanceId = &v
	return s
}

func (s *DescribeMem0InfoResponseBodyDataInstanceConnAddrs) Validate() error {
	return dara.Validate(s)
}

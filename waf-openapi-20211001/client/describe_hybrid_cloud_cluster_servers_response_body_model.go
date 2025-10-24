// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridCloudClusterServersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterServers(v []*DescribeHybridCloudClusterServersResponseBodyClusterServers) *DescribeHybridCloudClusterServersResponseBody
	GetClusterServers() []*DescribeHybridCloudClusterServersResponseBodyClusterServers
	SetRequestId(v string) *DescribeHybridCloudClusterServersResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeHybridCloudClusterServersResponseBody
	GetTotalCount() *int32
}

type DescribeHybridCloudClusterServersResponseBody struct {
	ClusterServers []*DescribeHybridCloudClusterServersResponseBodyClusterServers `json:"ClusterServers,omitempty" xml:"ClusterServers,omitempty" type:"Repeated"`
	// example:
	//
	// C3B0FDD2-11CE-59A5-BEB5-*****1A969
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 3
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeHybridCloudClusterServersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudClusterServersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudClusterServersResponseBody) GetClusterServers() []*DescribeHybridCloudClusterServersResponseBodyClusterServers {
	return s.ClusterServers
}

func (s *DescribeHybridCloudClusterServersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHybridCloudClusterServersResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeHybridCloudClusterServersResponseBody) SetClusterServers(v []*DescribeHybridCloudClusterServersResponseBodyClusterServers) *DescribeHybridCloudClusterServersResponseBody {
	s.ClusterServers = v
	return s
}

func (s *DescribeHybridCloudClusterServersResponseBody) SetRequestId(v string) *DescribeHybridCloudClusterServersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHybridCloudClusterServersResponseBody) SetTotalCount(v int32) *DescribeHybridCloudClusterServersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeHybridCloudClusterServersResponseBody) Validate() error {
	if s.ClusterServers != nil {
		for _, item := range s.ClusterServers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeHybridCloudClusterServersResponseBodyClusterServers struct {
	// example:
	//
	// asiapacific
	Continents *string `json:"Continents,omitempty" xml:"Continents,omitempty"`
	// example:
	//
	// 410
	ContinentsValue *int32 `json:"ContinentsValue,omitempty" xml:"ContinentsValue,omitempty"`
	// example:
	//
	// 4
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// example:
	//
	// 1665987186000
	CreateTimestamp *int64  `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	CustomName      *string `json:"CustomName,omitempty" xml:"CustomName,omitempty"`
	// example:
	//
	// 1
	GroupId   *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// example:
	//
	// protect
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// example:
	//
	// qsh5-sec-waf-*****-6
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// example:
	//
	// 1.*.*1。
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// example:
	//
	// running
	JobStatus *string `json:"JobStatus,omitempty" xml:"JobStatus,omitempty"`
	// example:
	//
	// **:**:00:02:eb:c7
	Mac *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	// example:
	//
	// 33547386880
	Memory *int64 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// example:
	//
	// 28ab688cd403563e8e173*****79600
	Mid *string `json:"Mid,omitempty" xml:"Mid,omitempty"`
	// example:
	//
	// aliyun。
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// example:
	//
	// 0
	OperatorValue *int32 `json:"OperatorValue,omitempty" xml:"OperatorValue,omitempty"`
	// example:
	//
	// beijing
	RegionCode *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
	// example:
	//
	// 0
	RegionCodeValue *int32 `json:"RegionCodeValue,omitempty" xml:"RegionCodeValue,omitempty"`
	// example:
	//
	// ok
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1711953897000
	UpdateTimestamp *int64 `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
}

func (s DescribeHybridCloudClusterServersResponseBodyClusterServers) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudClusterServersResponseBodyClusterServers) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudClusterServersResponseBodyClusterServers) GetContinents() *string {
	return s.Continents
}

func (s *DescribeHybridCloudClusterServersResponseBodyClusterServers) GetContinentsValue() *int32 {
	return s.ContinentsValue
}

func (s *DescribeHybridCloudClusterServersResponseBodyClusterServers) GetCpu() *int32 {
	return s.Cpu
}

func (s *DescribeHybridCloudClusterServersResponseBodyClusterServers) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DescribeHybridCloudClusterServersResponseBodyClusterServers) GetCustomName() *string {
	return s.CustomName
}

func (s *DescribeHybridCloudClusterServersResponseBodyClusterServers) GetGroupId() *int64 {
	return s.GroupId
}

func (s *DescribeHybridCloudClusterServersResponseBodyClusterServers) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeHybridCloudClusterServersResponseBodyClusterServers) GetGroupType() *string {
	return s.GroupType
}

func (s *DescribeHybridCloudClusterServersResponseBodyClusterServers) GetHostName() *string {
	return s.HostName
}

func (s *DescribeHybridCloudClusterServersResponseBodyClusterServers) GetIp() *string {
	return s.Ip
}

func (s *DescribeHybridCloudClusterServersResponseBodyClusterServers) GetJobStatus() *string {
	return s.JobStatus
}

func (s *DescribeHybridCloudClusterServersResponseBodyClusterServers) GetMac() *string {
	return s.Mac
}

func (s *DescribeHybridCloudClusterServersResponseBodyClusterServers) GetMemory() *int64 {
	return s.Memory
}

func (s *DescribeHybridCloudClusterServersResponseBodyClusterServers) GetMid() *string {
	return s.Mid
}

func (s *DescribeHybridCloudClusterServersResponseBodyClusterServers) GetOperator() *string {
	return s.Operator
}

func (s *DescribeHybridCloudClusterServersResponseBodyClusterServers) GetOperatorValue() *int32 {
	return s.OperatorValue
}

func (s *DescribeHybridCloudClusterServersResponseBodyClusterServers) GetRegionCode() *string {
	return s.RegionCode
}

func (s *DescribeHybridCloudClusterServersResponseBodyClusterServers) GetRegionCodeValue() *int32 {
	return s.RegionCodeValue
}

func (s *DescribeHybridCloudClusterServersResponseBodyClusterServers) GetStatus() *string {
	return s.Status
}

func (s *DescribeHybridCloudClusterServersResponseBodyClusterServers) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *DescribeHybridCloudClusterServersResponseBodyClusterServers) SetContinents(v string) *DescribeHybridCloudClusterServersResponseBodyClusterServers {
	s.Continents = &v
	return s
}

func (s *DescribeHybridCloudClusterServersResponseBodyClusterServers) SetContinentsValue(v int32) *DescribeHybridCloudClusterServersResponseBodyClusterServers {
	s.ContinentsValue = &v
	return s
}

func (s *DescribeHybridCloudClusterServersResponseBodyClusterServers) SetCpu(v int32) *DescribeHybridCloudClusterServersResponseBodyClusterServers {
	s.Cpu = &v
	return s
}

func (s *DescribeHybridCloudClusterServersResponseBodyClusterServers) SetCreateTimestamp(v int64) *DescribeHybridCloudClusterServersResponseBodyClusterServers {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeHybridCloudClusterServersResponseBodyClusterServers) SetCustomName(v string) *DescribeHybridCloudClusterServersResponseBodyClusterServers {
	s.CustomName = &v
	return s
}

func (s *DescribeHybridCloudClusterServersResponseBodyClusterServers) SetGroupId(v int64) *DescribeHybridCloudClusterServersResponseBodyClusterServers {
	s.GroupId = &v
	return s
}

func (s *DescribeHybridCloudClusterServersResponseBodyClusterServers) SetGroupName(v string) *DescribeHybridCloudClusterServersResponseBodyClusterServers {
	s.GroupName = &v
	return s
}

func (s *DescribeHybridCloudClusterServersResponseBodyClusterServers) SetGroupType(v string) *DescribeHybridCloudClusterServersResponseBodyClusterServers {
	s.GroupType = &v
	return s
}

func (s *DescribeHybridCloudClusterServersResponseBodyClusterServers) SetHostName(v string) *DescribeHybridCloudClusterServersResponseBodyClusterServers {
	s.HostName = &v
	return s
}

func (s *DescribeHybridCloudClusterServersResponseBodyClusterServers) SetIp(v string) *DescribeHybridCloudClusterServersResponseBodyClusterServers {
	s.Ip = &v
	return s
}

func (s *DescribeHybridCloudClusterServersResponseBodyClusterServers) SetJobStatus(v string) *DescribeHybridCloudClusterServersResponseBodyClusterServers {
	s.JobStatus = &v
	return s
}

func (s *DescribeHybridCloudClusterServersResponseBodyClusterServers) SetMac(v string) *DescribeHybridCloudClusterServersResponseBodyClusterServers {
	s.Mac = &v
	return s
}

func (s *DescribeHybridCloudClusterServersResponseBodyClusterServers) SetMemory(v int64) *DescribeHybridCloudClusterServersResponseBodyClusterServers {
	s.Memory = &v
	return s
}

func (s *DescribeHybridCloudClusterServersResponseBodyClusterServers) SetMid(v string) *DescribeHybridCloudClusterServersResponseBodyClusterServers {
	s.Mid = &v
	return s
}

func (s *DescribeHybridCloudClusterServersResponseBodyClusterServers) SetOperator(v string) *DescribeHybridCloudClusterServersResponseBodyClusterServers {
	s.Operator = &v
	return s
}

func (s *DescribeHybridCloudClusterServersResponseBodyClusterServers) SetOperatorValue(v int32) *DescribeHybridCloudClusterServersResponseBodyClusterServers {
	s.OperatorValue = &v
	return s
}

func (s *DescribeHybridCloudClusterServersResponseBodyClusterServers) SetRegionCode(v string) *DescribeHybridCloudClusterServersResponseBodyClusterServers {
	s.RegionCode = &v
	return s
}

func (s *DescribeHybridCloudClusterServersResponseBodyClusterServers) SetRegionCodeValue(v int32) *DescribeHybridCloudClusterServersResponseBodyClusterServers {
	s.RegionCodeValue = &v
	return s
}

func (s *DescribeHybridCloudClusterServersResponseBodyClusterServers) SetStatus(v string) *DescribeHybridCloudClusterServersResponseBodyClusterServers {
	s.Status = &v
	return s
}

func (s *DescribeHybridCloudClusterServersResponseBodyClusterServers) SetUpdateTimestamp(v int64) *DescribeHybridCloudClusterServersResponseBodyClusterServers {
	s.UpdateTimestamp = &v
	return s
}

func (s *DescribeHybridCloudClusterServersResponseBodyClusterServers) Validate() error {
	return dara.Validate(s)
}

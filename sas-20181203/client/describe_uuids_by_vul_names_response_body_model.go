// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUuidsByVulNamesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMachineInfoStatistics(v []*DescribeUuidsByVulNamesResponseBodyMachineInfoStatistics) *DescribeUuidsByVulNamesResponseBody
	GetMachineInfoStatistics() []*DescribeUuidsByVulNamesResponseBodyMachineInfoStatistics
	SetRequestId(v string) *DescribeUuidsByVulNamesResponseBody
	GetRequestId() *string
	SetVulCount(v int32) *DescribeUuidsByVulNamesResponseBody
	GetVulCount() *int32
}

type DescribeUuidsByVulNamesResponseBody struct {
	// The statistics about the servers.
	MachineInfoStatistics []*DescribeUuidsByVulNamesResponseBodyMachineInfoStatistics `json:"MachineInfoStatistics,omitempty" xml:"MachineInfoStatistics,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 97286A-4A6B-4A4-95FA-EC7E3E2451
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of vulnerabilities on the server.
	//
	// example:
	//
	// 2
	VulCount *int32 `json:"VulCount,omitempty" xml:"VulCount,omitempty"`
}

func (s DescribeUuidsByVulNamesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUuidsByVulNamesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUuidsByVulNamesResponseBody) GetMachineInfoStatistics() []*DescribeUuidsByVulNamesResponseBodyMachineInfoStatistics {
	return s.MachineInfoStatistics
}

func (s *DescribeUuidsByVulNamesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUuidsByVulNamesResponseBody) GetVulCount() *int32 {
	return s.VulCount
}

func (s *DescribeUuidsByVulNamesResponseBody) SetMachineInfoStatistics(v []*DescribeUuidsByVulNamesResponseBodyMachineInfoStatistics) *DescribeUuidsByVulNamesResponseBody {
	s.MachineInfoStatistics = v
	return s
}

func (s *DescribeUuidsByVulNamesResponseBody) SetRequestId(v string) *DescribeUuidsByVulNamesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUuidsByVulNamesResponseBody) SetVulCount(v int32) *DescribeUuidsByVulNamesResponseBody {
	s.VulCount = &v
	return s
}

func (s *DescribeUuidsByVulNamesResponseBody) Validate() error {
	if s.MachineInfoStatistics != nil {
		for _, item := range s.MachineInfoStatistics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeUuidsByVulNamesResponseBodyMachineInfoStatistics struct {
	// The public IP address of the server on which the exception was detected.
	//
	// example:
	//
	// 47.98.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address of the server on which the exception was detected.
	//
	// example:
	//
	// 172.18.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The instance ID of the server.
	//
	// example:
	//
	// i-wz9gd1os5talju****
	MachineInstanceId *string `json:"MachineInstanceId,omitempty" xml:"MachineInstanceId,omitempty"`
	// The IP address of the server.
	//
	// example:
	//
	// 192.168.XX.XX
	MachineIp *string `json:"MachineIp,omitempty" xml:"MachineIp,omitempty"`
	// The name of the server.
	//
	// example:
	//
	// TestMachine
	MachineName *string `json:"MachineName,omitempty" xml:"MachineName,omitempty"`
	// The operating system that the server runs.
	//
	// example:
	//
	// windows
	Os *string `json:"Os,omitempty" xml:"Os,omitempty"`
	// The region ID of the server.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The UUID of the server.
	//
	// example:
	//
	// 18375c64-eaa2-4702-92b0-4ee7******
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeUuidsByVulNamesResponseBodyMachineInfoStatistics) String() string {
	return dara.Prettify(s)
}

func (s DescribeUuidsByVulNamesResponseBodyMachineInfoStatistics) GoString() string {
	return s.String()
}

func (s *DescribeUuidsByVulNamesResponseBodyMachineInfoStatistics) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribeUuidsByVulNamesResponseBodyMachineInfoStatistics) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *DescribeUuidsByVulNamesResponseBodyMachineInfoStatistics) GetMachineInstanceId() *string {
	return s.MachineInstanceId
}

func (s *DescribeUuidsByVulNamesResponseBodyMachineInfoStatistics) GetMachineIp() *string {
	return s.MachineIp
}

func (s *DescribeUuidsByVulNamesResponseBodyMachineInfoStatistics) GetMachineName() *string {
	return s.MachineName
}

func (s *DescribeUuidsByVulNamesResponseBodyMachineInfoStatistics) GetOs() *string {
	return s.Os
}

func (s *DescribeUuidsByVulNamesResponseBodyMachineInfoStatistics) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeUuidsByVulNamesResponseBodyMachineInfoStatistics) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeUuidsByVulNamesResponseBodyMachineInfoStatistics) SetInternetIp(v string) *DescribeUuidsByVulNamesResponseBodyMachineInfoStatistics {
	s.InternetIp = &v
	return s
}

func (s *DescribeUuidsByVulNamesResponseBodyMachineInfoStatistics) SetIntranetIp(v string) *DescribeUuidsByVulNamesResponseBodyMachineInfoStatistics {
	s.IntranetIp = &v
	return s
}

func (s *DescribeUuidsByVulNamesResponseBodyMachineInfoStatistics) SetMachineInstanceId(v string) *DescribeUuidsByVulNamesResponseBodyMachineInfoStatistics {
	s.MachineInstanceId = &v
	return s
}

func (s *DescribeUuidsByVulNamesResponseBodyMachineInfoStatistics) SetMachineIp(v string) *DescribeUuidsByVulNamesResponseBodyMachineInfoStatistics {
	s.MachineIp = &v
	return s
}

func (s *DescribeUuidsByVulNamesResponseBodyMachineInfoStatistics) SetMachineName(v string) *DescribeUuidsByVulNamesResponseBodyMachineInfoStatistics {
	s.MachineName = &v
	return s
}

func (s *DescribeUuidsByVulNamesResponseBodyMachineInfoStatistics) SetOs(v string) *DescribeUuidsByVulNamesResponseBodyMachineInfoStatistics {
	s.Os = &v
	return s
}

func (s *DescribeUuidsByVulNamesResponseBodyMachineInfoStatistics) SetRegionId(v string) *DescribeUuidsByVulNamesResponseBodyMachineInfoStatistics {
	s.RegionId = &v
	return s
}

func (s *DescribeUuidsByVulNamesResponseBodyMachineInfoStatistics) SetUuid(v string) *DescribeUuidsByVulNamesResponseBodyMachineInfoStatistics {
	s.Uuid = &v
	return s
}

func (s *DescribeUuidsByVulNamesResponseBodyMachineInfoStatistics) Validate() error {
	return dara.Validate(s)
}

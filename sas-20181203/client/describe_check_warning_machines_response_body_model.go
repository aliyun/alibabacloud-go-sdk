// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCheckWarningMachinesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *DescribeCheckWarningMachinesResponseBody
	GetCount() *int32
	SetCurrentPage(v int32) *DescribeCheckWarningMachinesResponseBody
	GetCurrentPage() *int32
	SetMachines(v []*DescribeCheckWarningMachinesResponseBodyMachines) *DescribeCheckWarningMachinesResponseBody
	GetMachines() []*DescribeCheckWarningMachinesResponseBodyMachines
	SetPageSize(v int32) *DescribeCheckWarningMachinesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeCheckWarningMachinesResponseBody
	GetRequestId() *string
}

type DescribeCheckWarningMachinesResponseBody struct {
	// The number of the servers on which the same risk item is detected.
	//
	// example:
	//
	// 10
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The servers on which the same risk item is detected.
	Machines []*DescribeCheckWarningMachinesResponseBodyMachines `json:"Machines,omitempty" xml:"Machines,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// B3D989EB-2E59-584C-8438-5DFAA34A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCheckWarningMachinesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCheckWarningMachinesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCheckWarningMachinesResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *DescribeCheckWarningMachinesResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeCheckWarningMachinesResponseBody) GetMachines() []*DescribeCheckWarningMachinesResponseBodyMachines {
	return s.Machines
}

func (s *DescribeCheckWarningMachinesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCheckWarningMachinesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCheckWarningMachinesResponseBody) SetCount(v int32) *DescribeCheckWarningMachinesResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeCheckWarningMachinesResponseBody) SetCurrentPage(v int32) *DescribeCheckWarningMachinesResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCheckWarningMachinesResponseBody) SetMachines(v []*DescribeCheckWarningMachinesResponseBodyMachines) *DescribeCheckWarningMachinesResponseBody {
	s.Machines = v
	return s
}

func (s *DescribeCheckWarningMachinesResponseBody) SetPageSize(v int32) *DescribeCheckWarningMachinesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCheckWarningMachinesResponseBody) SetRequestId(v string) *DescribeCheckWarningMachinesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCheckWarningMachinesResponseBody) Validate() error {
	if s.Machines != nil {
		for _, item := range s.Machines {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCheckWarningMachinesResponseBodyMachines struct {
	// Indicates whether Security Center is authorized to protect the asset. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Bind *bool `json:"Bind,omitempty" xml:"Bind,omitempty"`
	// The instance ID of the server.
	//
	// example:
	//
	// i-bp1dp42hgn1xcq8l****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name of the server.
	//
	// example:
	//
	// win-oracl****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The public IP address of the server.
	//
	// example:
	//
	// 8.210.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address of the server.
	//
	// example:
	//
	// 172.25.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The ID of the region in which the server resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The UUID of the server.
	//
	// example:
	//
	// D0D6E6E4-CB8C-4897-B852-46AEFDA0****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeCheckWarningMachinesResponseBodyMachines) String() string {
	return dara.Prettify(s)
}

func (s DescribeCheckWarningMachinesResponseBodyMachines) GoString() string {
	return s.String()
}

func (s *DescribeCheckWarningMachinesResponseBodyMachines) GetBind() *bool {
	return s.Bind
}

func (s *DescribeCheckWarningMachinesResponseBodyMachines) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeCheckWarningMachinesResponseBodyMachines) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeCheckWarningMachinesResponseBodyMachines) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribeCheckWarningMachinesResponseBodyMachines) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *DescribeCheckWarningMachinesResponseBodyMachines) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCheckWarningMachinesResponseBodyMachines) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeCheckWarningMachinesResponseBodyMachines) SetBind(v bool) *DescribeCheckWarningMachinesResponseBodyMachines {
	s.Bind = &v
	return s
}

func (s *DescribeCheckWarningMachinesResponseBodyMachines) SetInstanceId(v string) *DescribeCheckWarningMachinesResponseBodyMachines {
	s.InstanceId = &v
	return s
}

func (s *DescribeCheckWarningMachinesResponseBodyMachines) SetInstanceName(v string) *DescribeCheckWarningMachinesResponseBodyMachines {
	s.InstanceName = &v
	return s
}

func (s *DescribeCheckWarningMachinesResponseBodyMachines) SetInternetIp(v string) *DescribeCheckWarningMachinesResponseBodyMachines {
	s.InternetIp = &v
	return s
}

func (s *DescribeCheckWarningMachinesResponseBodyMachines) SetIntranetIp(v string) *DescribeCheckWarningMachinesResponseBodyMachines {
	s.IntranetIp = &v
	return s
}

func (s *DescribeCheckWarningMachinesResponseBodyMachines) SetRegionId(v string) *DescribeCheckWarningMachinesResponseBodyMachines {
	s.RegionId = &v
	return s
}

func (s *DescribeCheckWarningMachinesResponseBodyMachines) SetUuid(v string) *DescribeCheckWarningMachinesResponseBodyMachines {
	s.Uuid = &v
	return s
}

func (s *DescribeCheckWarningMachinesResponseBodyMachines) Validate() error {
	return dara.Validate(s)
}

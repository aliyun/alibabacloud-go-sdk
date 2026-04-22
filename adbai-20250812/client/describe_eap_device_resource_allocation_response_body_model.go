// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEapDeviceResourceAllocationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeEapDeviceResourceAllocationResponseBodyItems) *DescribeEapDeviceResourceAllocationResponseBody
	GetItems() []*DescribeEapDeviceResourceAllocationResponseBodyItems
	SetRequestId(v string) *DescribeEapDeviceResourceAllocationResponseBody
	GetRequestId() *string
}

type DescribeEapDeviceResourceAllocationResponseBody struct {
	Items []*DescribeEapDeviceResourceAllocationResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 19E994DC-A816-56DB-9F90-5F8E403DD19D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeEapDeviceResourceAllocationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEapDeviceResourceAllocationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEapDeviceResourceAllocationResponseBody) GetItems() []*DescribeEapDeviceResourceAllocationResponseBodyItems {
	return s.Items
}

func (s *DescribeEapDeviceResourceAllocationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEapDeviceResourceAllocationResponseBody) SetItems(v []*DescribeEapDeviceResourceAllocationResponseBodyItems) *DescribeEapDeviceResourceAllocationResponseBody {
	s.Items = v
	return s
}

func (s *DescribeEapDeviceResourceAllocationResponseBody) SetRequestId(v string) *DescribeEapDeviceResourceAllocationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEapDeviceResourceAllocationResponseBody) Validate() error {
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

type DescribeEapDeviceResourceAllocationResponseBodyItems struct {
	// example:
	//
	// 3
	DeviceCount *int32 `json:"DeviceCount,omitempty" xml:"DeviceCount,omitempty"`
	// example:
	//
	// 2.0
	HeadAcu *float64 `json:"HeadAcu,omitempty" xml:"HeadAcu,omitempty"`
	// example:
	//
	// 2
	HeadCpu *int32 `json:"HeadCpu,omitempty" xml:"HeadCpu,omitempty"`
	// example:
	//
	// "medium"
	HeadSpecName *string `json:"HeadSpecName,omitempty" xml:"HeadSpecName,omitempty"`
	// example:
	//
	// 26
	TotalAcu *float64 `json:"TotalAcu,omitempty" xml:"TotalAcu,omitempty"`
	// example:
	//
	// 26
	TotalDeployedCpu *int32 `json:"TotalDeployedCpu,omitempty" xml:"TotalDeployedCpu,omitempty"`
	// example:
	//
	// 24
	TotalTargetCpu *int32 `json:"TotalTargetCpu,omitempty" xml:"TotalTargetCpu,omitempty"`
	// example:
	//
	// 8.0
	WebserverAcu *float64 `json:"WebserverAcu,omitempty" xml:"WebserverAcu,omitempty"`
	// example:
	//
	// 8
	WebserverCpu *int32 `json:"WebserverCpu,omitempty" xml:"WebserverCpu,omitempty"`
	// example:
	//
	// "xlarge"
	WebserverSpecName *string `json:"WebserverSpecName,omitempty" xml:"WebserverSpecName,omitempty"`
	// example:
	//
	// 16.0
	WorkerAcu *float64 `json:"WorkerAcu,omitempty" xml:"WorkerAcu,omitempty"`
	// example:
	//
	// 2
	WorkerCount *int32 `json:"WorkerCount,omitempty" xml:"WorkerCount,omitempty"`
	// example:
	//
	// 8
	WorkerCpu *int32 `json:"WorkerCpu,omitempty" xml:"WorkerCpu,omitempty"`
	// example:
	//
	// "xlarge"
	WorkerSpecName *string `json:"WorkerSpecName,omitempty" xml:"WorkerSpecName,omitempty"`
}

func (s DescribeEapDeviceResourceAllocationResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeEapDeviceResourceAllocationResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeEapDeviceResourceAllocationResponseBodyItems) GetDeviceCount() *int32 {
	return s.DeviceCount
}

func (s *DescribeEapDeviceResourceAllocationResponseBodyItems) GetHeadAcu() *float64 {
	return s.HeadAcu
}

func (s *DescribeEapDeviceResourceAllocationResponseBodyItems) GetHeadCpu() *int32 {
	return s.HeadCpu
}

func (s *DescribeEapDeviceResourceAllocationResponseBodyItems) GetHeadSpecName() *string {
	return s.HeadSpecName
}

func (s *DescribeEapDeviceResourceAllocationResponseBodyItems) GetTotalAcu() *float64 {
	return s.TotalAcu
}

func (s *DescribeEapDeviceResourceAllocationResponseBodyItems) GetTotalDeployedCpu() *int32 {
	return s.TotalDeployedCpu
}

func (s *DescribeEapDeviceResourceAllocationResponseBodyItems) GetTotalTargetCpu() *int32 {
	return s.TotalTargetCpu
}

func (s *DescribeEapDeviceResourceAllocationResponseBodyItems) GetWebserverAcu() *float64 {
	return s.WebserverAcu
}

func (s *DescribeEapDeviceResourceAllocationResponseBodyItems) GetWebserverCpu() *int32 {
	return s.WebserverCpu
}

func (s *DescribeEapDeviceResourceAllocationResponseBodyItems) GetWebserverSpecName() *string {
	return s.WebserverSpecName
}

func (s *DescribeEapDeviceResourceAllocationResponseBodyItems) GetWorkerAcu() *float64 {
	return s.WorkerAcu
}

func (s *DescribeEapDeviceResourceAllocationResponseBodyItems) GetWorkerCount() *int32 {
	return s.WorkerCount
}

func (s *DescribeEapDeviceResourceAllocationResponseBodyItems) GetWorkerCpu() *int32 {
	return s.WorkerCpu
}

func (s *DescribeEapDeviceResourceAllocationResponseBodyItems) GetWorkerSpecName() *string {
	return s.WorkerSpecName
}

func (s *DescribeEapDeviceResourceAllocationResponseBodyItems) SetDeviceCount(v int32) *DescribeEapDeviceResourceAllocationResponseBodyItems {
	s.DeviceCount = &v
	return s
}

func (s *DescribeEapDeviceResourceAllocationResponseBodyItems) SetHeadAcu(v float64) *DescribeEapDeviceResourceAllocationResponseBodyItems {
	s.HeadAcu = &v
	return s
}

func (s *DescribeEapDeviceResourceAllocationResponseBodyItems) SetHeadCpu(v int32) *DescribeEapDeviceResourceAllocationResponseBodyItems {
	s.HeadCpu = &v
	return s
}

func (s *DescribeEapDeviceResourceAllocationResponseBodyItems) SetHeadSpecName(v string) *DescribeEapDeviceResourceAllocationResponseBodyItems {
	s.HeadSpecName = &v
	return s
}

func (s *DescribeEapDeviceResourceAllocationResponseBodyItems) SetTotalAcu(v float64) *DescribeEapDeviceResourceAllocationResponseBodyItems {
	s.TotalAcu = &v
	return s
}

func (s *DescribeEapDeviceResourceAllocationResponseBodyItems) SetTotalDeployedCpu(v int32) *DescribeEapDeviceResourceAllocationResponseBodyItems {
	s.TotalDeployedCpu = &v
	return s
}

func (s *DescribeEapDeviceResourceAllocationResponseBodyItems) SetTotalTargetCpu(v int32) *DescribeEapDeviceResourceAllocationResponseBodyItems {
	s.TotalTargetCpu = &v
	return s
}

func (s *DescribeEapDeviceResourceAllocationResponseBodyItems) SetWebserverAcu(v float64) *DescribeEapDeviceResourceAllocationResponseBodyItems {
	s.WebserverAcu = &v
	return s
}

func (s *DescribeEapDeviceResourceAllocationResponseBodyItems) SetWebserverCpu(v int32) *DescribeEapDeviceResourceAllocationResponseBodyItems {
	s.WebserverCpu = &v
	return s
}

func (s *DescribeEapDeviceResourceAllocationResponseBodyItems) SetWebserverSpecName(v string) *DescribeEapDeviceResourceAllocationResponseBodyItems {
	s.WebserverSpecName = &v
	return s
}

func (s *DescribeEapDeviceResourceAllocationResponseBodyItems) SetWorkerAcu(v float64) *DescribeEapDeviceResourceAllocationResponseBodyItems {
	s.WorkerAcu = &v
	return s
}

func (s *DescribeEapDeviceResourceAllocationResponseBodyItems) SetWorkerCount(v int32) *DescribeEapDeviceResourceAllocationResponseBodyItems {
	s.WorkerCount = &v
	return s
}

func (s *DescribeEapDeviceResourceAllocationResponseBodyItems) SetWorkerCpu(v int32) *DescribeEapDeviceResourceAllocationResponseBodyItems {
	s.WorkerCpu = &v
	return s
}

func (s *DescribeEapDeviceResourceAllocationResponseBodyItems) SetWorkerSpecName(v string) *DescribeEapDeviceResourceAllocationResponseBodyItems {
	s.WorkerSpecName = &v
	return s
}

func (s *DescribeEapDeviceResourceAllocationResponseBodyItems) Validate() error {
	return dara.Validate(s)
}

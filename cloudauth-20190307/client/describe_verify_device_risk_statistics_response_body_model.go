// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVerifyDeviceRiskStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeVerifyDeviceRiskStatisticsResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeVerifyDeviceRiskStatisticsResponseBodyResultObject) *DescribeVerifyDeviceRiskStatisticsResponseBody
	GetResultObject() *DescribeVerifyDeviceRiskStatisticsResponseBodyResultObject
}

type DescribeVerifyDeviceRiskStatisticsResponseBody struct {
	// ID of this request.
	//
	// example:
	//
	// B3193814-AE54-50C5-9070-68B69C07287D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Authentication result.
	ResultObject *DescribeVerifyDeviceRiskStatisticsResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s DescribeVerifyDeviceRiskStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifyDeviceRiskStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVerifyDeviceRiskStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVerifyDeviceRiskStatisticsResponseBody) GetResultObject() *DescribeVerifyDeviceRiskStatisticsResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeVerifyDeviceRiskStatisticsResponseBody) SetRequestId(v string) *DescribeVerifyDeviceRiskStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVerifyDeviceRiskStatisticsResponseBody) SetResultObject(v *DescribeVerifyDeviceRiskStatisticsResponseBodyResultObject) *DescribeVerifyDeviceRiskStatisticsResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeVerifyDeviceRiskStatisticsResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVerifyDeviceRiskStatisticsResponseBodyResultObject struct {
	// Suspected fake face percentage: total number of suspected fake faces / total number of risks.
	//
	// example:
	//
	// 0
	FaceAttackRate *string `json:"FaceAttackRate,omitempty" xml:"FaceAttackRate,omitempty"`
	// Total number of suspected fake identities.
	//
	// example:
	//
	// 0
	IdFakeRate *string `json:"IdFakeRate,omitempty" xml:"IdFakeRate,omitempty"`
	// Data items in the response.
	Items []*DescribeVerifyDeviceRiskStatisticsResponseBodyResultObjectItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// Number of risks.
	//
	// example:
	//
	// 2
	RiskCount *int64 `json:"RiskCount,omitempty" xml:"RiskCount,omitempty"`
	// Root percentage: total number of root / total number of risks.
	//
	// example:
	//
	// 0
	RootRate *string `json:"RootRate,omitempty" xml:"RootRate,omitempty"`
	// Simulator percentage: total number of simulators / total number of risks.
	//
	// example:
	//
	// 0
	SimulatorRate *string `json:"SimulatorRate,omitempty" xml:"SimulatorRate,omitempty"`
	// Virtual video percentage: total number of virtual videos / total number of risks.
	//
	// example:
	//
	// 23.2
	VirtualVideoRate *string `json:"VirtualVideoRate,omitempty" xml:"VirtualVideoRate,omitempty"`
}

func (s DescribeVerifyDeviceRiskStatisticsResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifyDeviceRiskStatisticsResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeVerifyDeviceRiskStatisticsResponseBodyResultObject) GetFaceAttackRate() *string {
	return s.FaceAttackRate
}

func (s *DescribeVerifyDeviceRiskStatisticsResponseBodyResultObject) GetIdFakeRate() *string {
	return s.IdFakeRate
}

func (s *DescribeVerifyDeviceRiskStatisticsResponseBodyResultObject) GetItems() []*DescribeVerifyDeviceRiskStatisticsResponseBodyResultObjectItems {
	return s.Items
}

func (s *DescribeVerifyDeviceRiskStatisticsResponseBodyResultObject) GetRiskCount() *int64 {
	return s.RiskCount
}

func (s *DescribeVerifyDeviceRiskStatisticsResponseBodyResultObject) GetRootRate() *string {
	return s.RootRate
}

func (s *DescribeVerifyDeviceRiskStatisticsResponseBodyResultObject) GetSimulatorRate() *string {
	return s.SimulatorRate
}

func (s *DescribeVerifyDeviceRiskStatisticsResponseBodyResultObject) GetVirtualVideoRate() *string {
	return s.VirtualVideoRate
}

func (s *DescribeVerifyDeviceRiskStatisticsResponseBodyResultObject) SetFaceAttackRate(v string) *DescribeVerifyDeviceRiskStatisticsResponseBodyResultObject {
	s.FaceAttackRate = &v
	return s
}

func (s *DescribeVerifyDeviceRiskStatisticsResponseBodyResultObject) SetIdFakeRate(v string) *DescribeVerifyDeviceRiskStatisticsResponseBodyResultObject {
	s.IdFakeRate = &v
	return s
}

func (s *DescribeVerifyDeviceRiskStatisticsResponseBodyResultObject) SetItems(v []*DescribeVerifyDeviceRiskStatisticsResponseBodyResultObjectItems) *DescribeVerifyDeviceRiskStatisticsResponseBodyResultObject {
	s.Items = v
	return s
}

func (s *DescribeVerifyDeviceRiskStatisticsResponseBodyResultObject) SetRiskCount(v int64) *DescribeVerifyDeviceRiskStatisticsResponseBodyResultObject {
	s.RiskCount = &v
	return s
}

func (s *DescribeVerifyDeviceRiskStatisticsResponseBodyResultObject) SetRootRate(v string) *DescribeVerifyDeviceRiskStatisticsResponseBodyResultObject {
	s.RootRate = &v
	return s
}

func (s *DescribeVerifyDeviceRiskStatisticsResponseBodyResultObject) SetSimulatorRate(v string) *DescribeVerifyDeviceRiskStatisticsResponseBodyResultObject {
	s.SimulatorRate = &v
	return s
}

func (s *DescribeVerifyDeviceRiskStatisticsResponseBodyResultObject) SetVirtualVideoRate(v string) *DescribeVerifyDeviceRiskStatisticsResponseBodyResultObject {
	s.VirtualVideoRate = &v
	return s
}

func (s *DescribeVerifyDeviceRiskStatisticsResponseBodyResultObject) Validate() error {
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

type DescribeVerifyDeviceRiskStatisticsResponseBodyResultObjectItems struct {
	// Daily call count.
	//
	// example:
	//
	// 11
	DailyCallCount *int64 `json:"DailyCallCount,omitempty" xml:"DailyCallCount,omitempty"`
	// Date.
	//
	// example:
	//
	// 2025-10-10
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// Abnormal device risk ratio.
	//
	// example:
	//
	// 0
	DeviceRiskRate *string `json:"DeviceRiskRate,omitempty" xml:"DeviceRiskRate,omitempty"`
	// Abnormal identity risk ratio.
	//
	// example:
	//
	// 0
	IdentityRiskRate *string `json:"IdentityRiskRate,omitempty" xml:"IdentityRiskRate,omitempty"`
}

func (s DescribeVerifyDeviceRiskStatisticsResponseBodyResultObjectItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifyDeviceRiskStatisticsResponseBodyResultObjectItems) GoString() string {
	return s.String()
}

func (s *DescribeVerifyDeviceRiskStatisticsResponseBodyResultObjectItems) GetDailyCallCount() *int64 {
	return s.DailyCallCount
}

func (s *DescribeVerifyDeviceRiskStatisticsResponseBodyResultObjectItems) GetDate() *string {
	return s.Date
}

func (s *DescribeVerifyDeviceRiskStatisticsResponseBodyResultObjectItems) GetDeviceRiskRate() *string {
	return s.DeviceRiskRate
}

func (s *DescribeVerifyDeviceRiskStatisticsResponseBodyResultObjectItems) GetIdentityRiskRate() *string {
	return s.IdentityRiskRate
}

func (s *DescribeVerifyDeviceRiskStatisticsResponseBodyResultObjectItems) SetDailyCallCount(v int64) *DescribeVerifyDeviceRiskStatisticsResponseBodyResultObjectItems {
	s.DailyCallCount = &v
	return s
}

func (s *DescribeVerifyDeviceRiskStatisticsResponseBodyResultObjectItems) SetDate(v string) *DescribeVerifyDeviceRiskStatisticsResponseBodyResultObjectItems {
	s.Date = &v
	return s
}

func (s *DescribeVerifyDeviceRiskStatisticsResponseBodyResultObjectItems) SetDeviceRiskRate(v string) *DescribeVerifyDeviceRiskStatisticsResponseBodyResultObjectItems {
	s.DeviceRiskRate = &v
	return s
}

func (s *DescribeVerifyDeviceRiskStatisticsResponseBodyResultObjectItems) SetIdentityRiskRate(v string) *DescribeVerifyDeviceRiskStatisticsResponseBodyResultObjectItems {
	s.IdentityRiskRate = &v
	return s
}

func (s *DescribeVerifyDeviceRiskStatisticsResponseBodyResultObjectItems) Validate() error {
	return dara.Validate(s)
}

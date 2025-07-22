// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFaultDiagnosisFactorDistributionStatResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeFaultDiagnosisFactorDistributionStatResponseBody
	GetRequestId() *string
	SetStatList(v []*DescribeFaultDiagnosisFactorDistributionStatResponseBodyStatList) *DescribeFaultDiagnosisFactorDistributionStatResponseBody
	GetStatList() []*DescribeFaultDiagnosisFactorDistributionStatResponseBodyStatList
}

type DescribeFaultDiagnosisFactorDistributionStatResponseBody struct {
	// example:
	//
	// 231470C1-ACFB-4C9F-844F-4CFE1E3804C5
	RequestId *string                                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StatList  []*DescribeFaultDiagnosisFactorDistributionStatResponseBodyStatList `json:"StatList,omitempty" xml:"StatList,omitempty" type:"Repeated"`
}

func (s DescribeFaultDiagnosisFactorDistributionStatResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFaultDiagnosisFactorDistributionStatResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisFactorDistributionStatResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFaultDiagnosisFactorDistributionStatResponseBody) GetStatList() []*DescribeFaultDiagnosisFactorDistributionStatResponseBodyStatList {
	return s.StatList
}

func (s *DescribeFaultDiagnosisFactorDistributionStatResponseBody) SetRequestId(v string) *DescribeFaultDiagnosisFactorDistributionStatResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFaultDiagnosisFactorDistributionStatResponseBody) SetStatList(v []*DescribeFaultDiagnosisFactorDistributionStatResponseBodyStatList) *DescribeFaultDiagnosisFactorDistributionStatResponseBody {
	s.StatList = v
	return s
}

func (s *DescribeFaultDiagnosisFactorDistributionStatResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeFaultDiagnosisFactorDistributionStatResponseBodyStatList struct {
	// example:
	//
	// 1
	FactorId *string `json:"FactorId,omitempty" xml:"FactorId,omitempty"`
	// example:
	//
	// 100
	UserCount *int32 `json:"UserCount,omitempty" xml:"UserCount,omitempty"`
	// example:
	//
	// 0.9239
	UserRatio *float32 `json:"UserRatio,omitempty" xml:"UserRatio,omitempty"`
}

func (s DescribeFaultDiagnosisFactorDistributionStatResponseBodyStatList) String() string {
	return dara.Prettify(s)
}

func (s DescribeFaultDiagnosisFactorDistributionStatResponseBodyStatList) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisFactorDistributionStatResponseBodyStatList) GetFactorId() *string {
	return s.FactorId
}

func (s *DescribeFaultDiagnosisFactorDistributionStatResponseBodyStatList) GetUserCount() *int32 {
	return s.UserCount
}

func (s *DescribeFaultDiagnosisFactorDistributionStatResponseBodyStatList) GetUserRatio() *float32 {
	return s.UserRatio
}

func (s *DescribeFaultDiagnosisFactorDistributionStatResponseBodyStatList) SetFactorId(v string) *DescribeFaultDiagnosisFactorDistributionStatResponseBodyStatList {
	s.FactorId = &v
	return s
}

func (s *DescribeFaultDiagnosisFactorDistributionStatResponseBodyStatList) SetUserCount(v int32) *DescribeFaultDiagnosisFactorDistributionStatResponseBodyStatList {
	s.UserCount = &v
	return s
}

func (s *DescribeFaultDiagnosisFactorDistributionStatResponseBodyStatList) SetUserRatio(v float32) *DescribeFaultDiagnosisFactorDistributionStatResponseBodyStatList {
	s.UserRatio = &v
	return s
}

func (s *DescribeFaultDiagnosisFactorDistributionStatResponseBodyStatList) Validate() error {
	return dara.Validate(s)
}

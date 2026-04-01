// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEpnInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEPNInstances(v *DescribeEpnInstancesResponseBodyEPNInstances) *DescribeEpnInstancesResponseBody
	GetEPNInstances() *DescribeEpnInstancesResponseBodyEPNInstances
	SetPageNumber(v int32) *DescribeEpnInstancesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeEpnInstancesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeEpnInstancesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeEpnInstancesResponseBody
	GetTotalCount() *int32
}

type DescribeEpnInstancesResponseBody struct {
	EPNInstances *DescribeEpnInstancesResponseBodyEPNInstances `json:"EPNInstances,omitempty" xml:"EPNInstances,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 40
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A1707FC0-430C-423A-B624-284046B20399
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of pages returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeEpnInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEpnInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEpnInstancesResponseBody) GetEPNInstances() *DescribeEpnInstancesResponseBodyEPNInstances {
	return s.EPNInstances
}

func (s *DescribeEpnInstancesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeEpnInstancesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeEpnInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEpnInstancesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeEpnInstancesResponseBody) SetEPNInstances(v *DescribeEpnInstancesResponseBodyEPNInstances) *DescribeEpnInstancesResponseBody {
	s.EPNInstances = v
	return s
}

func (s *DescribeEpnInstancesResponseBody) SetPageNumber(v int32) *DescribeEpnInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeEpnInstancesResponseBody) SetPageSize(v int32) *DescribeEpnInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeEpnInstancesResponseBody) SetRequestId(v string) *DescribeEpnInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEpnInstancesResponseBody) SetTotalCount(v int32) *DescribeEpnInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeEpnInstancesResponseBody) Validate() error {
	if s.EPNInstances != nil {
		if err := s.EPNInstances.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeEpnInstancesResponseBodyEPNInstances struct {
	EPNInstance []*DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance `json:"EPNInstance,omitempty" xml:"EPNInstance,omitempty" type:"Repeated"`
}

func (s DescribeEpnInstancesResponseBodyEPNInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeEpnInstancesResponseBodyEPNInstances) GoString() string {
	return s.String()
}

func (s *DescribeEpnInstancesResponseBodyEPNInstances) GetEPNInstance() []*DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance {
	return s.EPNInstance
}

func (s *DescribeEpnInstancesResponseBodyEPNInstances) SetEPNInstance(v []*DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance) *DescribeEpnInstancesResponseBodyEPNInstances {
	s.EPNInstance = v
	return s
}

func (s *DescribeEpnInstancesResponseBodyEPNInstances) Validate() error {
	if s.EPNInstance != nil {
		for _, item := range s.EPNInstance {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance struct {
	CreationTime            *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	EPNInstanceId           *string `json:"EPNInstanceId,omitempty" xml:"EPNInstanceId,omitempty"`
	EPNInstanceName         *string `json:"EPNInstanceName,omitempty" xml:"EPNInstanceName,omitempty"`
	EPNInstanceType         *string `json:"EPNInstanceType,omitempty" xml:"EPNInstanceType,omitempty"`
	EndTime                 *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InternetMaxBandwidthOut *int32  `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	ModifyTime              *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	NetworkingModel         *string `json:"NetworkingModel,omitempty" xml:"NetworkingModel,omitempty"`
	StartTime               *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status                  *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance) String() string {
	return dara.Prettify(s)
}

func (s DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance) GoString() string {
	return s.String()
}

func (s *DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance) GetEPNInstanceId() *string {
	return s.EPNInstanceId
}

func (s *DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance) GetEPNInstanceName() *string {
	return s.EPNInstanceName
}

func (s *DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance) GetEPNInstanceType() *string {
	return s.EPNInstanceType
}

func (s *DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance) GetInternetMaxBandwidthOut() *int32 {
	return s.InternetMaxBandwidthOut
}

func (s *DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance) GetNetworkingModel() *string {
	return s.NetworkingModel
}

func (s *DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance) GetStatus() *string {
	return s.Status
}

func (s *DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance) SetCreationTime(v string) *DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance {
	s.CreationTime = &v
	return s
}

func (s *DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance) SetEPNInstanceId(v string) *DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance {
	s.EPNInstanceId = &v
	return s
}

func (s *DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance) SetEPNInstanceName(v string) *DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance {
	s.EPNInstanceName = &v
	return s
}

func (s *DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance) SetEPNInstanceType(v string) *DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance {
	s.EPNInstanceType = &v
	return s
}

func (s *DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance) SetEndTime(v string) *DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance {
	s.EndTime = &v
	return s
}

func (s *DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance) SetInternetMaxBandwidthOut(v int32) *DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance) SetModifyTime(v string) *DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance {
	s.ModifyTime = &v
	return s
}

func (s *DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance) SetNetworkingModel(v string) *DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance {
	s.NetworkingModel = &v
	return s
}

func (s *DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance) SetStartTime(v string) *DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance {
	s.StartTime = &v
	return s
}

func (s *DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance) SetStatus(v string) *DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance {
	s.Status = &v
	return s
}

func (s *DescribeEpnInstancesResponseBodyEPNInstancesEPNInstance) Validate() error {
	return dara.Validate(s)
}

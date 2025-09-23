// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBatchSlsDispatchStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeBatchSlsDispatchStatusResponseBody
	GetRequestId() *string
	SetSlsConfigStatusList(v []*DescribeBatchSlsDispatchStatusResponseBodySlsConfigStatusList) *DescribeBatchSlsDispatchStatusResponseBody
	GetSlsConfigStatusList() []*DescribeBatchSlsDispatchStatusResponseBodySlsConfigStatusList
	SetTotalCount(v int32) *DescribeBatchSlsDispatchStatusResponseBody
	GetTotalCount() *int32
}

type DescribeBatchSlsDispatchStatusResponseBody struct {
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId           *string                                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SlsConfigStatusList []*DescribeBatchSlsDispatchStatusResponseBodySlsConfigStatusList `json:"SlsConfigStatusList,omitempty" xml:"SlsConfigStatusList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeBatchSlsDispatchStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBatchSlsDispatchStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBatchSlsDispatchStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBatchSlsDispatchStatusResponseBody) GetSlsConfigStatusList() []*DescribeBatchSlsDispatchStatusResponseBodySlsConfigStatusList {
	return s.SlsConfigStatusList
}

func (s *DescribeBatchSlsDispatchStatusResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeBatchSlsDispatchStatusResponseBody) SetRequestId(v string) *DescribeBatchSlsDispatchStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBatchSlsDispatchStatusResponseBody) SetSlsConfigStatusList(v []*DescribeBatchSlsDispatchStatusResponseBodySlsConfigStatusList) *DescribeBatchSlsDispatchStatusResponseBody {
	s.SlsConfigStatusList = v
	return s
}

func (s *DescribeBatchSlsDispatchStatusResponseBody) SetTotalCount(v int32) *DescribeBatchSlsDispatchStatusResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeBatchSlsDispatchStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeBatchSlsDispatchStatusResponseBodySlsConfigStatusList struct {
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s DescribeBatchSlsDispatchStatusResponseBodySlsConfigStatusList) String() string {
	return dara.Prettify(s)
}

func (s DescribeBatchSlsDispatchStatusResponseBodySlsConfigStatusList) GoString() string {
	return s.String()
}

func (s *DescribeBatchSlsDispatchStatusResponseBodySlsConfigStatusList) GetDomain() *string {
	return s.Domain
}

func (s *DescribeBatchSlsDispatchStatusResponseBodySlsConfigStatusList) GetEnable() *bool {
	return s.Enable
}

func (s *DescribeBatchSlsDispatchStatusResponseBodySlsConfigStatusList) SetDomain(v string) *DescribeBatchSlsDispatchStatusResponseBodySlsConfigStatusList {
	s.Domain = &v
	return s
}

func (s *DescribeBatchSlsDispatchStatusResponseBodySlsConfigStatusList) SetEnable(v bool) *DescribeBatchSlsDispatchStatusResponseBodySlsConfigStatusList {
	s.Enable = &v
	return s
}

func (s *DescribeBatchSlsDispatchStatusResponseBodySlsConfigStatusList) Validate() error {
	return dara.Validate(s)
}

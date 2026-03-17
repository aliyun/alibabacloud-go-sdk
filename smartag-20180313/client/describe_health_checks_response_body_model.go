// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHealthChecksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHealthChecks(v *DescribeHealthChecksResponseBodyHealthChecks) *DescribeHealthChecksResponseBody
	GetHealthChecks() *DescribeHealthChecksResponseBodyHealthChecks
	SetPageNumber(v int32) *DescribeHealthChecksResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeHealthChecksResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeHealthChecksResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeHealthChecksResponseBody
	GetTotalCount() *int32
}

type DescribeHealthChecksResponseBody struct {
	HealthChecks *DescribeHealthChecksResponseBodyHealthChecks `json:"HealthChecks,omitempty" xml:"HealthChecks,omitempty" type:"Struct"`
	// The page number of the returned page. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page. Default value: **10**. Maximum value: **50**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0B275FE6-FC86-4921-BC70-1B3DF68B078F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeHealthChecksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHealthChecksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHealthChecksResponseBody) GetHealthChecks() *DescribeHealthChecksResponseBodyHealthChecks {
	return s.HealthChecks
}

func (s *DescribeHealthChecksResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeHealthChecksResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeHealthChecksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHealthChecksResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeHealthChecksResponseBody) SetHealthChecks(v *DescribeHealthChecksResponseBodyHealthChecks) *DescribeHealthChecksResponseBody {
	s.HealthChecks = v
	return s
}

func (s *DescribeHealthChecksResponseBody) SetPageNumber(v int32) *DescribeHealthChecksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeHealthChecksResponseBody) SetPageSize(v int32) *DescribeHealthChecksResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeHealthChecksResponseBody) SetRequestId(v string) *DescribeHealthChecksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHealthChecksResponseBody) SetTotalCount(v int32) *DescribeHealthChecksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeHealthChecksResponseBody) Validate() error {
	if s.HealthChecks != nil {
		if err := s.HealthChecks.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeHealthChecksResponseBodyHealthChecks struct {
	HealthCheck []*DescribeHealthChecksResponseBodyHealthChecksHealthCheck `json:"HealthCheck,omitempty" xml:"HealthCheck,omitempty" type:"Repeated"`
}

func (s DescribeHealthChecksResponseBodyHealthChecks) String() string {
	return dara.Prettify(s)
}

func (s DescribeHealthChecksResponseBodyHealthChecks) GoString() string {
	return s.String()
}

func (s *DescribeHealthChecksResponseBodyHealthChecks) GetHealthCheck() []*DescribeHealthChecksResponseBodyHealthChecksHealthCheck {
	return s.HealthCheck
}

func (s *DescribeHealthChecksResponseBodyHealthChecks) SetHealthCheck(v []*DescribeHealthChecksResponseBodyHealthChecksHealthCheck) *DescribeHealthChecksResponseBodyHealthChecks {
	s.HealthCheck = v
	return s
}

func (s *DescribeHealthChecksResponseBodyHealthChecks) Validate() error {
	if s.HealthCheck != nil {
		for _, item := range s.HealthCheck {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeHealthChecksResponseBodyHealthChecksHealthCheck struct {
	Description        *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DstIpAddr          *string `json:"DstIpAddr,omitempty" xml:"DstIpAddr,omitempty"`
	DstPort            *int32  `json:"DstPort,omitempty" xml:"DstPort,omitempty"`
	FailCountThreshold *int32  `json:"FailCountThreshold,omitempty" xml:"FailCountThreshold,omitempty"`
	HcInstanceId       *string `json:"HcInstanceId,omitempty" xml:"HcInstanceId,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ProbeCount         *int32  `json:"ProbeCount,omitempty" xml:"ProbeCount,omitempty"`
	ProbeInterval      *int32  `json:"ProbeInterval,omitempty" xml:"ProbeInterval,omitempty"`
	ProbeTimeout       *int32  `json:"ProbeTimeout,omitempty" xml:"ProbeTimeout,omitempty"`
	RelationCount      *int32  `json:"RelationCount,omitempty" xml:"RelationCount,omitempty"`
	RttFailThreshold   *int32  `json:"RttFailThreshold,omitempty" xml:"RttFailThreshold,omitempty"`
	RttThreshold       *int32  `json:"RttThreshold,omitempty" xml:"RttThreshold,omitempty"`
	SmartAGId          *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
	SrcIpAddr          *string `json:"SrcIpAddr,omitempty" xml:"SrcIpAddr,omitempty"`
	SrcPort            *int32  `json:"SrcPort,omitempty" xml:"SrcPort,omitempty"`
	Status             *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type               *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeHealthChecksResponseBodyHealthChecksHealthCheck) String() string {
	return dara.Prettify(s)
}

func (s DescribeHealthChecksResponseBodyHealthChecksHealthCheck) GoString() string {
	return s.String()
}

func (s *DescribeHealthChecksResponseBodyHealthChecksHealthCheck) GetDescription() *string {
	return s.Description
}

func (s *DescribeHealthChecksResponseBodyHealthChecksHealthCheck) GetDstIpAddr() *string {
	return s.DstIpAddr
}

func (s *DescribeHealthChecksResponseBodyHealthChecksHealthCheck) GetDstPort() *int32 {
	return s.DstPort
}

func (s *DescribeHealthChecksResponseBodyHealthChecksHealthCheck) GetFailCountThreshold() *int32 {
	return s.FailCountThreshold
}

func (s *DescribeHealthChecksResponseBodyHealthChecksHealthCheck) GetHcInstanceId() *string {
	return s.HcInstanceId
}

func (s *DescribeHealthChecksResponseBodyHealthChecksHealthCheck) GetName() *string {
	return s.Name
}

func (s *DescribeHealthChecksResponseBodyHealthChecksHealthCheck) GetProbeCount() *int32 {
	return s.ProbeCount
}

func (s *DescribeHealthChecksResponseBodyHealthChecksHealthCheck) GetProbeInterval() *int32 {
	return s.ProbeInterval
}

func (s *DescribeHealthChecksResponseBodyHealthChecksHealthCheck) GetProbeTimeout() *int32 {
	return s.ProbeTimeout
}

func (s *DescribeHealthChecksResponseBodyHealthChecksHealthCheck) GetRelationCount() *int32 {
	return s.RelationCount
}

func (s *DescribeHealthChecksResponseBodyHealthChecksHealthCheck) GetRttFailThreshold() *int32 {
	return s.RttFailThreshold
}

func (s *DescribeHealthChecksResponseBodyHealthChecksHealthCheck) GetRttThreshold() *int32 {
	return s.RttThreshold
}

func (s *DescribeHealthChecksResponseBodyHealthChecksHealthCheck) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DescribeHealthChecksResponseBodyHealthChecksHealthCheck) GetSrcIpAddr() *string {
	return s.SrcIpAddr
}

func (s *DescribeHealthChecksResponseBodyHealthChecksHealthCheck) GetSrcPort() *int32 {
	return s.SrcPort
}

func (s *DescribeHealthChecksResponseBodyHealthChecksHealthCheck) GetStatus() *string {
	return s.Status
}

func (s *DescribeHealthChecksResponseBodyHealthChecksHealthCheck) GetType() *string {
	return s.Type
}

func (s *DescribeHealthChecksResponseBodyHealthChecksHealthCheck) SetDescription(v string) *DescribeHealthChecksResponseBodyHealthChecksHealthCheck {
	s.Description = &v
	return s
}

func (s *DescribeHealthChecksResponseBodyHealthChecksHealthCheck) SetDstIpAddr(v string) *DescribeHealthChecksResponseBodyHealthChecksHealthCheck {
	s.DstIpAddr = &v
	return s
}

func (s *DescribeHealthChecksResponseBodyHealthChecksHealthCheck) SetDstPort(v int32) *DescribeHealthChecksResponseBodyHealthChecksHealthCheck {
	s.DstPort = &v
	return s
}

func (s *DescribeHealthChecksResponseBodyHealthChecksHealthCheck) SetFailCountThreshold(v int32) *DescribeHealthChecksResponseBodyHealthChecksHealthCheck {
	s.FailCountThreshold = &v
	return s
}

func (s *DescribeHealthChecksResponseBodyHealthChecksHealthCheck) SetHcInstanceId(v string) *DescribeHealthChecksResponseBodyHealthChecksHealthCheck {
	s.HcInstanceId = &v
	return s
}

func (s *DescribeHealthChecksResponseBodyHealthChecksHealthCheck) SetName(v string) *DescribeHealthChecksResponseBodyHealthChecksHealthCheck {
	s.Name = &v
	return s
}

func (s *DescribeHealthChecksResponseBodyHealthChecksHealthCheck) SetProbeCount(v int32) *DescribeHealthChecksResponseBodyHealthChecksHealthCheck {
	s.ProbeCount = &v
	return s
}

func (s *DescribeHealthChecksResponseBodyHealthChecksHealthCheck) SetProbeInterval(v int32) *DescribeHealthChecksResponseBodyHealthChecksHealthCheck {
	s.ProbeInterval = &v
	return s
}

func (s *DescribeHealthChecksResponseBodyHealthChecksHealthCheck) SetProbeTimeout(v int32) *DescribeHealthChecksResponseBodyHealthChecksHealthCheck {
	s.ProbeTimeout = &v
	return s
}

func (s *DescribeHealthChecksResponseBodyHealthChecksHealthCheck) SetRelationCount(v int32) *DescribeHealthChecksResponseBodyHealthChecksHealthCheck {
	s.RelationCount = &v
	return s
}

func (s *DescribeHealthChecksResponseBodyHealthChecksHealthCheck) SetRttFailThreshold(v int32) *DescribeHealthChecksResponseBodyHealthChecksHealthCheck {
	s.RttFailThreshold = &v
	return s
}

func (s *DescribeHealthChecksResponseBodyHealthChecksHealthCheck) SetRttThreshold(v int32) *DescribeHealthChecksResponseBodyHealthChecksHealthCheck {
	s.RttThreshold = &v
	return s
}

func (s *DescribeHealthChecksResponseBodyHealthChecksHealthCheck) SetSmartAGId(v string) *DescribeHealthChecksResponseBodyHealthChecksHealthCheck {
	s.SmartAGId = &v
	return s
}

func (s *DescribeHealthChecksResponseBodyHealthChecksHealthCheck) SetSrcIpAddr(v string) *DescribeHealthChecksResponseBodyHealthChecksHealthCheck {
	s.SrcIpAddr = &v
	return s
}

func (s *DescribeHealthChecksResponseBodyHealthChecksHealthCheck) SetSrcPort(v int32) *DescribeHealthChecksResponseBodyHealthChecksHealthCheck {
	s.SrcPort = &v
	return s
}

func (s *DescribeHealthChecksResponseBodyHealthChecksHealthCheck) SetStatus(v string) *DescribeHealthChecksResponseBodyHealthChecksHealthCheck {
	s.Status = &v
	return s
}

func (s *DescribeHealthChecksResponseBodyHealthChecksHealthCheck) SetType(v string) *DescribeHealthChecksResponseBodyHealthChecksHealthCheck {
	s.Type = &v
	return s
}

func (s *DescribeHealthChecksResponseBodyHealthChecksHealthCheck) Validate() error {
	return dara.Validate(s)
}

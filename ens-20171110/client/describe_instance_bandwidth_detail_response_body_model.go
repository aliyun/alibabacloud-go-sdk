// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceBandwidthDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidths(v []*DescribeInstanceBandwidthDetailResponseBodyBandwidths) *DescribeInstanceBandwidthDetailResponseBody
	GetBandwidths() []*DescribeInstanceBandwidthDetailResponseBodyBandwidths
	SetPageNumber(v int32) *DescribeInstanceBandwidthDetailResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeInstanceBandwidthDetailResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeInstanceBandwidthDetailResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeInstanceBandwidthDetailResponseBody
	GetTotalCount() *int32
}

type DescribeInstanceBandwidthDetailResponseBody struct {
	Bandwidths []*DescribeInstanceBandwidthDetailResponseBodyBandwidths `json:"Bandwidths,omitempty" xml:"Bandwidths,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// C0003E8B-B930-4F59-ADC0-0E209A9012A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeInstanceBandwidthDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceBandwidthDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceBandwidthDetailResponseBody) GetBandwidths() []*DescribeInstanceBandwidthDetailResponseBodyBandwidths {
	return s.Bandwidths
}

func (s *DescribeInstanceBandwidthDetailResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeInstanceBandwidthDetailResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInstanceBandwidthDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceBandwidthDetailResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeInstanceBandwidthDetailResponseBody) SetBandwidths(v []*DescribeInstanceBandwidthDetailResponseBodyBandwidths) *DescribeInstanceBandwidthDetailResponseBody {
	s.Bandwidths = v
	return s
}

func (s *DescribeInstanceBandwidthDetailResponseBody) SetPageNumber(v int32) *DescribeInstanceBandwidthDetailResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstanceBandwidthDetailResponseBody) SetPageSize(v int32) *DescribeInstanceBandwidthDetailResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeInstanceBandwidthDetailResponseBody) SetRequestId(v string) *DescribeInstanceBandwidthDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceBandwidthDetailResponseBody) SetTotalCount(v int32) *DescribeInstanceBandwidthDetailResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeInstanceBandwidthDetailResponseBody) Validate() error {
	if s.Bandwidths != nil {
		for _, item := range s.Bandwidths {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstanceBandwidthDetailResponseBodyBandwidths struct {
	// example:
	//
	// 1972653484384661
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// example:
	//
	// 2024-01-11 00:00:00
	BizTime *string `json:"BizTime,omitempty" xml:"BizTime,omitempty"`
	// example:
	//
	// cn-yichang-2
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// example:
	//
	// 1
	FlowType *int32 `json:"FlowType,omitempty" xml:"FlowType,omitempty"`
	// example:
	//
	// i-6ecpqvkicnchxccozrpxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// vm
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// example:
	//
	// 203.107.***
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// example:
	//
	// cmcc
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// example:
	//
	// 43795230
	RxBw *int64 `json:"RxBw,omitempty" xml:"RxBw,omitempty"`
	// example:
	//
	// vm
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// example:
	//
	// 25415638
	TxBw *int64 `json:"TxBw,omitempty" xml:"TxBw,omitempty"`
}

func (s DescribeInstanceBandwidthDetailResponseBodyBandwidths) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceBandwidthDetailResponseBodyBandwidths) GoString() string {
	return s.String()
}

func (s *DescribeInstanceBandwidthDetailResponseBodyBandwidths) GetAliUid() *int64 {
	return s.AliUid
}

func (s *DescribeInstanceBandwidthDetailResponseBodyBandwidths) GetBizTime() *string {
	return s.BizTime
}

func (s *DescribeInstanceBandwidthDetailResponseBodyBandwidths) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeInstanceBandwidthDetailResponseBodyBandwidths) GetFlowType() *int32 {
	return s.FlowType
}

func (s *DescribeInstanceBandwidthDetailResponseBodyBandwidths) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceBandwidthDetailResponseBodyBandwidths) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeInstanceBandwidthDetailResponseBodyBandwidths) GetIp() *string {
	return s.Ip
}

func (s *DescribeInstanceBandwidthDetailResponseBodyBandwidths) GetIsp() *string {
	return s.Isp
}

func (s *DescribeInstanceBandwidthDetailResponseBodyBandwidths) GetRxBw() *int64 {
	return s.RxBw
}

func (s *DescribeInstanceBandwidthDetailResponseBodyBandwidths) GetServiceType() *string {
	return s.ServiceType
}

func (s *DescribeInstanceBandwidthDetailResponseBodyBandwidths) GetTxBw() *int64 {
	return s.TxBw
}

func (s *DescribeInstanceBandwidthDetailResponseBodyBandwidths) SetAliUid(v int64) *DescribeInstanceBandwidthDetailResponseBodyBandwidths {
	s.AliUid = &v
	return s
}

func (s *DescribeInstanceBandwidthDetailResponseBodyBandwidths) SetBizTime(v string) *DescribeInstanceBandwidthDetailResponseBodyBandwidths {
	s.BizTime = &v
	return s
}

func (s *DescribeInstanceBandwidthDetailResponseBodyBandwidths) SetEnsRegionId(v string) *DescribeInstanceBandwidthDetailResponseBodyBandwidths {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeInstanceBandwidthDetailResponseBodyBandwidths) SetFlowType(v int32) *DescribeInstanceBandwidthDetailResponseBodyBandwidths {
	s.FlowType = &v
	return s
}

func (s *DescribeInstanceBandwidthDetailResponseBodyBandwidths) SetInstanceId(v string) *DescribeInstanceBandwidthDetailResponseBodyBandwidths {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceBandwidthDetailResponseBodyBandwidths) SetInstanceType(v string) *DescribeInstanceBandwidthDetailResponseBodyBandwidths {
	s.InstanceType = &v
	return s
}

func (s *DescribeInstanceBandwidthDetailResponseBodyBandwidths) SetIp(v string) *DescribeInstanceBandwidthDetailResponseBodyBandwidths {
	s.Ip = &v
	return s
}

func (s *DescribeInstanceBandwidthDetailResponseBodyBandwidths) SetIsp(v string) *DescribeInstanceBandwidthDetailResponseBodyBandwidths {
	s.Isp = &v
	return s
}

func (s *DescribeInstanceBandwidthDetailResponseBodyBandwidths) SetRxBw(v int64) *DescribeInstanceBandwidthDetailResponseBodyBandwidths {
	s.RxBw = &v
	return s
}

func (s *DescribeInstanceBandwidthDetailResponseBodyBandwidths) SetServiceType(v string) *DescribeInstanceBandwidthDetailResponseBodyBandwidths {
	s.ServiceType = &v
	return s
}

func (s *DescribeInstanceBandwidthDetailResponseBodyBandwidths) SetTxBw(v int64) *DescribeInstanceBandwidthDetailResponseBodyBandwidths {
	s.TxBw = &v
	return s
}

func (s *DescribeInstanceBandwidthDetailResponseBodyBandwidths) Validate() error {
	return dara.Validate(s)
}

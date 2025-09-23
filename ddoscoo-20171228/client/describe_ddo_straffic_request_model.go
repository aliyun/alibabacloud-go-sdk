// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDDoSTrafficRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEip(v string) *DescribeDDoSTrafficRequest
	GetEip() *string
	SetEndTime(v int64) *DescribeDDoSTrafficRequest
	GetEndTime() *int64
	SetInterval(v int32) *DescribeDDoSTrafficRequest
	GetInterval() *int32
	SetResourceGroupId(v string) *DescribeDDoSTrafficRequest
	GetResourceGroupId() *string
	SetSourceIp(v string) *DescribeDDoSTrafficRequest
	GetSourceIp() *string
	SetStartTime(v int64) *DescribeDDoSTrafficRequest
	GetStartTime() *int64
}

type DescribeDDoSTrafficRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1.1.1.1
	Eip *string `json:"Eip,omitempty" xml:"Eip,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3289457398
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 60
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// example:
	//
	// xx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3289457324
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDDoSTrafficRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDoSTrafficRequest) GoString() string {
	return s.String()
}

func (s *DescribeDDoSTrafficRequest) GetEip() *string {
	return s.Eip
}

func (s *DescribeDDoSTrafficRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeDDoSTrafficRequest) GetInterval() *int32 {
	return s.Interval
}

func (s *DescribeDDoSTrafficRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDDoSTrafficRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeDDoSTrafficRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDDoSTrafficRequest) SetEip(v string) *DescribeDDoSTrafficRequest {
	s.Eip = &v
	return s
}

func (s *DescribeDDoSTrafficRequest) SetEndTime(v int64) *DescribeDDoSTrafficRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDDoSTrafficRequest) SetInterval(v int32) *DescribeDDoSTrafficRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDDoSTrafficRequest) SetResourceGroupId(v string) *DescribeDDoSTrafficRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDDoSTrafficRequest) SetSourceIp(v string) *DescribeDDoSTrafficRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDDoSTrafficRequest) SetStartTime(v int64) *DescribeDDoSTrafficRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDDoSTrafficRequest) Validate() error {
	return dara.Validate(s)
}

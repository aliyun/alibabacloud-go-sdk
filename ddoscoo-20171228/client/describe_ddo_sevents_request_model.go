// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDDoSEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEip(v string) *DescribeDDoSEventsRequest
	GetEip() *string
	SetEndTime(v int64) *DescribeDDoSEventsRequest
	GetEndTime() *int64
	SetOffset(v int32) *DescribeDDoSEventsRequest
	GetOffset() *int32
	SetPageSize(v string) *DescribeDDoSEventsRequest
	GetPageSize() *string
	SetResourceGroupId(v string) *DescribeDDoSEventsRequest
	GetResourceGroupId() *string
	SetSourceIp(v string) *DescribeDDoSEventsRequest
	GetSourceIp() *string
	SetStartTime(v int64) *DescribeDDoSEventsRequest
	GetStartTime() *int64
}

type DescribeDDoSEventsRequest struct {
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
	// 3289457324
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Offset *int32 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
	// 3289457398
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDDoSEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDoSEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDDoSEventsRequest) GetEip() *string {
	return s.Eip
}

func (s *DescribeDDoSEventsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeDDoSEventsRequest) GetOffset() *int32 {
	return s.Offset
}

func (s *DescribeDDoSEventsRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeDDoSEventsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDDoSEventsRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeDDoSEventsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDDoSEventsRequest) SetEip(v string) *DescribeDDoSEventsRequest {
	s.Eip = &v
	return s
}

func (s *DescribeDDoSEventsRequest) SetEndTime(v int64) *DescribeDDoSEventsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDDoSEventsRequest) SetOffset(v int32) *DescribeDDoSEventsRequest {
	s.Offset = &v
	return s
}

func (s *DescribeDDoSEventsRequest) SetPageSize(v string) *DescribeDDoSEventsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDDoSEventsRequest) SetResourceGroupId(v string) *DescribeDDoSEventsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDDoSEventsRequest) SetSourceIp(v string) *DescribeDDoSEventsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDDoSEventsRequest) SetStartTime(v int64) *DescribeDDoSEventsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDDoSEventsRequest) Validate() error {
	return dara.Validate(s)
}

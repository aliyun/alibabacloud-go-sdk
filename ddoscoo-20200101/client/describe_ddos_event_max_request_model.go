// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDDosEventMaxRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *DescribeDDosEventMaxRequest
	GetEndTime() *int64
	SetIp(v string) *DescribeDDosEventMaxRequest
	GetIp() *string
	SetStartTime(v int64) *DescribeDDosEventMaxRequest
	GetStartTime() *int64
}

type DescribeDDosEventMaxRequest struct {
	// The end of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1604073600
	EndTime *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Ip      *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The beginning of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1598889600
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDDosEventMaxRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDosEventMaxRequest) GoString() string {
	return s.String()
}

func (s *DescribeDDosEventMaxRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeDDosEventMaxRequest) GetIp() *string {
	return s.Ip
}

func (s *DescribeDDosEventMaxRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDDosEventMaxRequest) SetEndTime(v int64) *DescribeDDosEventMaxRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDDosEventMaxRequest) SetIp(v string) *DescribeDDosEventMaxRequest {
	s.Ip = &v
	return s
}

func (s *DescribeDDosEventMaxRequest) SetStartTime(v int64) *DescribeDDosEventMaxRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDDosEventMaxRequest) Validate() error {
	return dara.Validate(s)
}

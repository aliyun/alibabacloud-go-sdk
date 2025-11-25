// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeElasticQpsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetElasticQps(v []*DescribeElasticQpsResponseBodyElasticQps) *DescribeElasticQpsResponseBody
	GetElasticQps() []*DescribeElasticQpsResponseBodyElasticQps
	SetRequestId(v string) *DescribeElasticQpsResponseBody
	GetRequestId() *string
}

type DescribeElasticQpsResponseBody struct {
	// The information about the burstable QPS.
	ElasticQps []*DescribeElasticQpsResponseBodyElasticQps `json:"ElasticQps,omitempty" xml:"ElasticQps,omitempty" type:"Repeated"`
	// The request ID, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 2E7F7F7B-39A8-5D92-BAB4-D89D9DCE7D4F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeElasticQpsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticQpsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeElasticQpsResponseBody) GetElasticQps() []*DescribeElasticQpsResponseBodyElasticQps {
	return s.ElasticQps
}

func (s *DescribeElasticQpsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeElasticQpsResponseBody) SetElasticQps(v []*DescribeElasticQpsResponseBodyElasticQps) *DescribeElasticQpsResponseBody {
	s.ElasticQps = v
	return s
}

func (s *DescribeElasticQpsResponseBody) SetRequestId(v string) *DescribeElasticQpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeElasticQpsResponseBody) Validate() error {
	if s.ElasticQps != nil {
		for _, item := range s.ElasticQps {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeElasticQpsResponseBodyElasticQps struct {
	// The index number of the returned data.
	//
	// example:
	//
	// 1
	Index *int64 `json:"Index,omitempty" xml:"Index,omitempty"`
	// The peak QPS of the normal service.
	//
	// example:
	//
	// 23
	MaxNormalQps *int64 `json:"MaxNormalQps,omitempty" xml:"MaxNormalQps,omitempty"`
	// The peak inbound QPS.
	//
	// example:
	//
	// 100
	MaxQps *int64 `json:"MaxQps,omitempty" xml:"MaxQps,omitempty"`
	// The total number of requests during the step size period.
	//
	// example:
	//
	// 15104
	Pv *int64 `json:"Pv,omitempty" xml:"Pv,omitempty"`
	// The total number of HTTP 2xx status codes during the step size period.
	//
	// example:
	//
	// 455
	Status2 *int64 `json:"Status2,omitempty" xml:"Status2,omitempty"`
	// The total number of HTTP 3xx status codes during the step size period.
	//
	// example:
	//
	// 100
	Status3 *int64 `json:"Status3,omitempty" xml:"Status3,omitempty"`
	// The total number of HTTP 4xx status codes during the step size period.
	//
	// example:
	//
	// 34
	Status4 *int64 `json:"Status4,omitempty" xml:"Status4,omitempty"`
	// The total number of HTTP 5xx status codes during the step size period.
	//
	// example:
	//
	// 0
	Status5 *int64 `json:"Status5,omitempty" xml:"Status5,omitempty"`
	// The total number of origin requests during the step size period.
	//
	// example:
	//
	// 1223
	Ups *int64 `json:"Ups,omitempty" xml:"Ups,omitempty"`
}

func (s DescribeElasticQpsResponseBodyElasticQps) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticQpsResponseBodyElasticQps) GoString() string {
	return s.String()
}

func (s *DescribeElasticQpsResponseBodyElasticQps) GetIndex() *int64 {
	return s.Index
}

func (s *DescribeElasticQpsResponseBodyElasticQps) GetMaxNormalQps() *int64 {
	return s.MaxNormalQps
}

func (s *DescribeElasticQpsResponseBodyElasticQps) GetMaxQps() *int64 {
	return s.MaxQps
}

func (s *DescribeElasticQpsResponseBodyElasticQps) GetPv() *int64 {
	return s.Pv
}

func (s *DescribeElasticQpsResponseBodyElasticQps) GetStatus2() *int64 {
	return s.Status2
}

func (s *DescribeElasticQpsResponseBodyElasticQps) GetStatus3() *int64 {
	return s.Status3
}

func (s *DescribeElasticQpsResponseBodyElasticQps) GetStatus4() *int64 {
	return s.Status4
}

func (s *DescribeElasticQpsResponseBodyElasticQps) GetStatus5() *int64 {
	return s.Status5
}

func (s *DescribeElasticQpsResponseBodyElasticQps) GetUps() *int64 {
	return s.Ups
}

func (s *DescribeElasticQpsResponseBodyElasticQps) SetIndex(v int64) *DescribeElasticQpsResponseBodyElasticQps {
	s.Index = &v
	return s
}

func (s *DescribeElasticQpsResponseBodyElasticQps) SetMaxNormalQps(v int64) *DescribeElasticQpsResponseBodyElasticQps {
	s.MaxNormalQps = &v
	return s
}

func (s *DescribeElasticQpsResponseBodyElasticQps) SetMaxQps(v int64) *DescribeElasticQpsResponseBodyElasticQps {
	s.MaxQps = &v
	return s
}

func (s *DescribeElasticQpsResponseBodyElasticQps) SetPv(v int64) *DescribeElasticQpsResponseBodyElasticQps {
	s.Pv = &v
	return s
}

func (s *DescribeElasticQpsResponseBodyElasticQps) SetStatus2(v int64) *DescribeElasticQpsResponseBodyElasticQps {
	s.Status2 = &v
	return s
}

func (s *DescribeElasticQpsResponseBodyElasticQps) SetStatus3(v int64) *DescribeElasticQpsResponseBodyElasticQps {
	s.Status3 = &v
	return s
}

func (s *DescribeElasticQpsResponseBodyElasticQps) SetStatus4(v int64) *DescribeElasticQpsResponseBodyElasticQps {
	s.Status4 = &v
	return s
}

func (s *DescribeElasticQpsResponseBodyElasticQps) SetStatus5(v int64) *DescribeElasticQpsResponseBodyElasticQps {
	s.Status5 = &v
	return s
}

func (s *DescribeElasticQpsResponseBodyElasticQps) SetUps(v int64) *DescribeElasticQpsResponseBodyElasticQps {
	s.Ups = &v
	return s
}

func (s *DescribeElasticQpsResponseBodyElasticQps) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStreamVodListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRecords(v []*DescribeStreamVodListResponseBodyRecords) *DescribeStreamVodListResponseBody
	GetRecords() []*DescribeStreamVodListResponseBodyRecords
	SetRequestId(v string) *DescribeStreamVodListResponseBody
	GetRequestId() *string
}

type DescribeStreamVodListResponseBody struct {
	Records []*DescribeStreamVodListResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeStreamVodListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeStreamVodListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeStreamVodListResponseBody) GetRecords() []*DescribeStreamVodListResponseBodyRecords {
	return s.Records
}

func (s *DescribeStreamVodListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeStreamVodListResponseBody) SetRecords(v []*DescribeStreamVodListResponseBodyRecords) *DescribeStreamVodListResponseBody {
	s.Records = v
	return s
}

func (s *DescribeStreamVodListResponseBody) SetRequestId(v string) *DescribeStreamVodListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeStreamVodListResponseBody) Validate() error {
	if s.Records != nil {
		for _, item := range s.Records {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeStreamVodListResponseBodyRecords struct {
	// example:
	//
	// 1634873413
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1639077653
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeStreamVodListResponseBodyRecords) String() string {
	return dara.Prettify(s)
}

func (s DescribeStreamVodListResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *DescribeStreamVodListResponseBodyRecords) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeStreamVodListResponseBodyRecords) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeStreamVodListResponseBodyRecords) SetEndTime(v int64) *DescribeStreamVodListResponseBodyRecords {
	s.EndTime = &v
	return s
}

func (s *DescribeStreamVodListResponseBodyRecords) SetStartTime(v int64) *DescribeStreamVodListResponseBodyRecords {
	s.StartTime = &v
	return s
}

func (s *DescribeStreamVodListResponseBodyRecords) Validate() error {
	return dara.Validate(s)
}

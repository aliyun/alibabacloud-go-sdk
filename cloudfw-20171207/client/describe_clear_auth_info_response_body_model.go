// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClearAuthInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int32) *DescribeClearAuthInfoResponseBody
	GetEndTime() *int32
	SetLeftTimes(v int32) *DescribeClearAuthInfoResponseBody
	GetLeftTimes() *int32
	SetRequestId(v string) *DescribeClearAuthInfoResponseBody
	GetRequestId() *string
}

type DescribeClearAuthInfoResponseBody struct {
	// example:
	//
	// 1755964800
	EndTime *int32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 20
	LeftTimes *int32 `json:"LeftTimes,omitempty" xml:"LeftTimes,omitempty"`
	// example:
	//
	// 8DDEE254-5639-5548-82D1-AAAC7347****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeClearAuthInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClearAuthInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClearAuthInfoResponseBody) GetEndTime() *int32 {
	return s.EndTime
}

func (s *DescribeClearAuthInfoResponseBody) GetLeftTimes() *int32 {
	return s.LeftTimes
}

func (s *DescribeClearAuthInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeClearAuthInfoResponseBody) SetEndTime(v int32) *DescribeClearAuthInfoResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeClearAuthInfoResponseBody) SetLeftTimes(v int32) *DescribeClearAuthInfoResponseBody {
	s.LeftTimes = &v
	return s
}

func (s *DescribeClearAuthInfoResponseBody) SetRequestId(v string) *DescribeClearAuthInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClearAuthInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

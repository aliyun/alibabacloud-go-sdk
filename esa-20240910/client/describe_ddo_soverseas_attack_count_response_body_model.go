// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDDoSOverseasAttackCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDDoSOverseasAttackCountResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeDDoSOverseasAttackCountResponseBody
	GetTotalCount() *int32
	SetUsedCount(v int32) *DescribeDDoSOverseasAttackCountResponseBody
	GetUsedCount() *int32
}

type DescribeDDoSOverseasAttackCountResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// EBCECE18-DA96-5307-8506-01E5A0E21530
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of attacks supported by the plan.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The number of attacks that have been used.
	//
	// example:
	//
	// 0
	UsedCount *int32 `json:"UsedCount,omitempty" xml:"UsedCount,omitempty"`
}

func (s DescribeDDoSOverseasAttackCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDoSOverseasAttackCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDDoSOverseasAttackCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDDoSOverseasAttackCountResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDDoSOverseasAttackCountResponseBody) GetUsedCount() *int32 {
	return s.UsedCount
}

func (s *DescribeDDoSOverseasAttackCountResponseBody) SetRequestId(v string) *DescribeDDoSOverseasAttackCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDDoSOverseasAttackCountResponseBody) SetTotalCount(v int32) *DescribeDDoSOverseasAttackCountResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDDoSOverseasAttackCountResponseBody) SetUsedCount(v int32) *DescribeDDoSOverseasAttackCountResponseBody {
	s.UsedCount = &v
	return s
}

func (s *DescribeDDoSOverseasAttackCountResponseBody) Validate() error {
	return dara.Validate(s)
}

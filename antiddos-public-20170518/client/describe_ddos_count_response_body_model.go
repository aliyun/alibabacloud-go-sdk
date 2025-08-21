// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDdosCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDdosCount(v *DescribeDdosCountResponseBodyDdosCount) *DescribeDdosCountResponseBody
	GetDdosCount() *DescribeDdosCountResponseBodyDdosCount
	SetRequestId(v string) *DescribeDdosCountResponseBody
	GetRequestId() *string
}

type DescribeDdosCountResponseBody struct {
	// The number of assets that are under DDoS attacks.
	DdosCount *DescribeDdosCountResponseBodyDdosCount `json:"DdosCount,omitempty" xml:"DdosCount,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 7D66C762-324E-51CE-B461-2007996087B2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDdosCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDdosCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDdosCountResponseBody) GetDdosCount() *DescribeDdosCountResponseBodyDdosCount {
	return s.DdosCount
}

func (s *DescribeDdosCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDdosCountResponseBody) SetDdosCount(v *DescribeDdosCountResponseBodyDdosCount) *DescribeDdosCountResponseBody {
	s.DdosCount = v
	return s
}

func (s *DescribeDdosCountResponseBody) SetRequestId(v string) *DescribeDdosCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDdosCountResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDdosCountResponseBodyDdosCount struct {
	// The number of assets for which blackhole filtering is triggered.
	//
	// example:
	//
	// 0
	BlackholeCount *int32 `json:"BlackholeCount,omitempty" xml:"BlackholeCount,omitempty"`
	// The number of assets for which traffic scrubbing is triggered.
	//
	// example:
	//
	// 4
	DefenseCount *int32 `json:"DefenseCount,omitempty" xml:"DefenseCount,omitempty"`
	// The total number of assets.
	//
	// example:
	//
	// 0
	InstacenCount *int32 `json:"InstacenCount,omitempty" xml:"InstacenCount,omitempty"`
}

func (s DescribeDdosCountResponseBodyDdosCount) String() string {
	return dara.Prettify(s)
}

func (s DescribeDdosCountResponseBodyDdosCount) GoString() string {
	return s.String()
}

func (s *DescribeDdosCountResponseBodyDdosCount) GetBlackholeCount() *int32 {
	return s.BlackholeCount
}

func (s *DescribeDdosCountResponseBodyDdosCount) GetDefenseCount() *int32 {
	return s.DefenseCount
}

func (s *DescribeDdosCountResponseBodyDdosCount) GetInstacenCount() *int32 {
	return s.InstacenCount
}

func (s *DescribeDdosCountResponseBodyDdosCount) SetBlackholeCount(v int32) *DescribeDdosCountResponseBodyDdosCount {
	s.BlackholeCount = &v
	return s
}

func (s *DescribeDdosCountResponseBodyDdosCount) SetDefenseCount(v int32) *DescribeDdosCountResponseBodyDdosCount {
	s.DefenseCount = &v
	return s
}

func (s *DescribeDdosCountResponseBodyDdosCount) SetInstacenCount(v int32) *DescribeDdosCountResponseBodyDdosCount {
	s.InstacenCount = &v
	return s
}

func (s *DescribeDdosCountResponseBodyDdosCount) Validate() error {
	return dara.Validate(s)
}

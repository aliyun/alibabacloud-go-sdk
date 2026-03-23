// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCInstanceDdosCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDdosCount(v *DescribeRCInstanceDdosCountResponseBodyDdosCount) *DescribeRCInstanceDdosCountResponseBody
	GetDdosCount() *DescribeRCInstanceDdosCountResponseBodyDdosCount
	SetRequestId(v string) *DescribeRCInstanceDdosCountResponseBody
	GetRequestId() *string
}

type DescribeRCInstanceDdosCountResponseBody struct {
	DdosCount *DescribeRCInstanceDdosCountResponseBodyDdosCount `json:"DdosCount,omitempty" xml:"DdosCount,omitempty" type:"Struct"`
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRCInstanceDdosCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstanceDdosCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRCInstanceDdosCountResponseBody) GetDdosCount() *DescribeRCInstanceDdosCountResponseBodyDdosCount {
	return s.DdosCount
}

func (s *DescribeRCInstanceDdosCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRCInstanceDdosCountResponseBody) SetDdosCount(v *DescribeRCInstanceDdosCountResponseBodyDdosCount) *DescribeRCInstanceDdosCountResponseBody {
	s.DdosCount = v
	return s
}

func (s *DescribeRCInstanceDdosCountResponseBody) SetRequestId(v string) *DescribeRCInstanceDdosCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRCInstanceDdosCountResponseBody) Validate() error {
	if s.DdosCount != nil {
		if err := s.DdosCount.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRCInstanceDdosCountResponseBodyDdosCount struct {
	BlackholeCount *string `json:"BlackholeCount,omitempty" xml:"BlackholeCount,omitempty"`
	DefenseCount   *string `json:"DefenseCount,omitempty" xml:"DefenseCount,omitempty"`
	InstacenCount  *string `json:"InstacenCount,omitempty" xml:"InstacenCount,omitempty"`
}

func (s DescribeRCInstanceDdosCountResponseBodyDdosCount) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstanceDdosCountResponseBodyDdosCount) GoString() string {
	return s.String()
}

func (s *DescribeRCInstanceDdosCountResponseBodyDdosCount) GetBlackholeCount() *string {
	return s.BlackholeCount
}

func (s *DescribeRCInstanceDdosCountResponseBodyDdosCount) GetDefenseCount() *string {
	return s.DefenseCount
}

func (s *DescribeRCInstanceDdosCountResponseBodyDdosCount) GetInstacenCount() *string {
	return s.InstacenCount
}

func (s *DescribeRCInstanceDdosCountResponseBodyDdosCount) SetBlackholeCount(v string) *DescribeRCInstanceDdosCountResponseBodyDdosCount {
	s.BlackholeCount = &v
	return s
}

func (s *DescribeRCInstanceDdosCountResponseBodyDdosCount) SetDefenseCount(v string) *DescribeRCInstanceDdosCountResponseBodyDdosCount {
	s.DefenseCount = &v
	return s
}

func (s *DescribeRCInstanceDdosCountResponseBodyDdosCount) SetInstacenCount(v string) *DescribeRCInstanceDdosCountResponseBodyDdosCount {
	s.InstacenCount = &v
	return s
}

func (s *DescribeRCInstanceDdosCountResponseBodyDdosCount) Validate() error {
	return dara.Validate(s)
}

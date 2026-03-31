// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserTrafficResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeUserTrafficResponseBody
	GetRequestId() *string
	SetUserTraffic(v []*DescribeUserTrafficResponseBodyUserTraffic) *DescribeUserTrafficResponseBody
	GetUserTraffic() []*DescribeUserTrafficResponseBodyUserTraffic
}

type DescribeUserTrafficResponseBody struct {
	// example:
	//
	// EC10C9EA-A367-52D5-***-***
	RequestId   *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UserTraffic []*DescribeUserTrafficResponseBodyUserTraffic `json:"UserTraffic,omitempty" xml:"UserTraffic,omitempty" type:"Repeated"`
}

func (s DescribeUserTrafficResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserTrafficResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserTrafficResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserTrafficResponseBody) GetUserTraffic() []*DescribeUserTrafficResponseBodyUserTraffic {
	return s.UserTraffic
}

func (s *DescribeUserTrafficResponseBody) SetRequestId(v string) *DescribeUserTrafficResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserTrafficResponseBody) SetUserTraffic(v []*DescribeUserTrafficResponseBodyUserTraffic) *DescribeUserTrafficResponseBody {
	s.UserTraffic = v
	return s
}

func (s *DescribeUserTrafficResponseBody) Validate() error {
	if s.UserTraffic != nil {
		for _, item := range s.UserTraffic {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeUserTrafficResponseBodyUserTraffic struct {
	// example:
	//
	// 10
	Index *int64 `json:"Index,omitempty" xml:"Index,omitempty"`
	// example:
	//
	// 1024
	Pv *int64 `json:"Pv,omitempty" xml:"Pv,omitempty"`
}

func (s DescribeUserTrafficResponseBodyUserTraffic) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserTrafficResponseBodyUserTraffic) GoString() string {
	return s.String()
}

func (s *DescribeUserTrafficResponseBodyUserTraffic) GetIndex() *int64 {
	return s.Index
}

func (s *DescribeUserTrafficResponseBodyUserTraffic) GetPv() *int64 {
	return s.Pv
}

func (s *DescribeUserTrafficResponseBodyUserTraffic) SetIndex(v int64) *DescribeUserTrafficResponseBodyUserTraffic {
	s.Index = &v
	return s
}

func (s *DescribeUserTrafficResponseBodyUserTraffic) SetPv(v int64) *DescribeUserTrafficResponseBodyUserTraffic {
	s.Pv = &v
	return s
}

func (s *DescribeUserTrafficResponseBodyUserTraffic) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVisitUasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeVisitUasResponseBody
	GetRequestId() *string
	SetUas(v []*DescribeVisitUasResponseBodyUas) *DescribeVisitUasResponseBody
	GetUas() []*DescribeVisitUasResponseBodyUas
}

type DescribeVisitUasResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 2847CE98-AFAE-5A64-B80E-60461717F9DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The array of the top 10 user agents that are used to initiate requests.
	Uas []*DescribeVisitUasResponseBodyUas `json:"Uas,omitempty" xml:"Uas,omitempty" type:"Repeated"`
}

func (s DescribeVisitUasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVisitUasResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVisitUasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVisitUasResponseBody) GetUas() []*DescribeVisitUasResponseBodyUas {
	return s.Uas
}

func (s *DescribeVisitUasResponseBody) SetRequestId(v string) *DescribeVisitUasResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVisitUasResponseBody) SetUas(v []*DescribeVisitUasResponseBodyUas) *DescribeVisitUasResponseBody {
	s.Uas = v
	return s
}

func (s *DescribeVisitUasResponseBody) Validate() error {
	if s.Uas != nil {
		for _, item := range s.Uas {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVisitUasResponseBodyUas struct {
	// The number of requests that use the user agent.
	//
	// example:
	//
	// 698455
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The user agent.
	//
	// example:
	//
	// chrome
	Ua *string `json:"Ua,omitempty" xml:"Ua,omitempty"`
}

func (s DescribeVisitUasResponseBodyUas) String() string {
	return dara.Prettify(s)
}

func (s DescribeVisitUasResponseBodyUas) GoString() string {
	return s.String()
}

func (s *DescribeVisitUasResponseBodyUas) GetCount() *int64 {
	return s.Count
}

func (s *DescribeVisitUasResponseBodyUas) GetUa() *string {
	return s.Ua
}

func (s *DescribeVisitUasResponseBodyUas) SetCount(v int64) *DescribeVisitUasResponseBodyUas {
	s.Count = &v
	return s
}

func (s *DescribeVisitUasResponseBodyUas) SetUa(v string) *DescribeVisitUasResponseBodyUas {
	s.Ua = &v
	return s
}

func (s *DescribeVisitUasResponseBodyUas) Validate() error {
	return dara.Validate(s)
}

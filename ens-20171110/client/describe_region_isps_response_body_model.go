// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionIspsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsps(v []*DescribeRegionIspsResponseBodyIsps) *DescribeRegionIspsResponseBody
	GetIsps() []*DescribeRegionIspsResponseBodyIsps
	SetRequestId(v string) *DescribeRegionIspsResponseBody
	GetRequestId() *string
}

type DescribeRegionIspsResponseBody struct {
	// The list of ISPs.
	Isps []*DescribeRegionIspsResponseBodyIsps `json:"Isps,omitempty" xml:"Isps,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// DC90CC7E-23B6-5A90-9097-A17CE4A161C4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRegionIspsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionIspsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionIspsResponseBody) GetIsps() []*DescribeRegionIspsResponseBodyIsps {
	return s.Isps
}

func (s *DescribeRegionIspsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRegionIspsResponseBody) SetIsps(v []*DescribeRegionIspsResponseBodyIsps) *DescribeRegionIspsResponseBody {
	s.Isps = v
	return s
}

func (s *DescribeRegionIspsResponseBody) SetRequestId(v string) *DescribeRegionIspsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionIspsResponseBody) Validate() error {
	if s.Isps != nil {
		for _, item := range s.Isps {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRegionIspsResponseBodyIsps struct {
	// The code of the ISP.
	//
	// example:
	//
	// cmcc
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The name of the ISP.
	//
	// example:
	//
	// move
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeRegionIspsResponseBodyIsps) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionIspsResponseBodyIsps) GoString() string {
	return s.String()
}

func (s *DescribeRegionIspsResponseBodyIsps) GetCode() *string {
	return s.Code
}

func (s *DescribeRegionIspsResponseBodyIsps) GetName() *string {
	return s.Name
}

func (s *DescribeRegionIspsResponseBodyIsps) SetCode(v string) *DescribeRegionIspsResponseBodyIsps {
	s.Code = &v
	return s
}

func (s *DescribeRegionIspsResponseBodyIsps) SetName(v string) *DescribeRegionIspsResponseBodyIsps {
	s.Name = &v
	return s
}

func (s *DescribeRegionIspsResponseBodyIsps) Validate() error {
	return dara.Validate(s)
}

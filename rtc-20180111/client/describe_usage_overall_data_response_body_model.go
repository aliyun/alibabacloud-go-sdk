// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUsageOverallDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeUsageOverallDataResponseBody
	GetRequestId() *string
	SetUsageOverallData(v []*DescribeUsageOverallDataResponseBodyUsageOverallData) *DescribeUsageOverallDataResponseBody
	GetUsageOverallData() []*DescribeUsageOverallDataResponseBodyUsageOverallData
}

type DescribeUsageOverallDataResponseBody struct {
	// example:
	//
	// 231470C1-ACFB-4C9F-844F-4CFE1E3804C5
	RequestId        *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UsageOverallData []*DescribeUsageOverallDataResponseBodyUsageOverallData `json:"UsageOverallData,omitempty" xml:"UsageOverallData,omitempty" type:"Repeated"`
}

func (s DescribeUsageOverallDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUsageOverallDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUsageOverallDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUsageOverallDataResponseBody) GetUsageOverallData() []*DescribeUsageOverallDataResponseBodyUsageOverallData {
	return s.UsageOverallData
}

func (s *DescribeUsageOverallDataResponseBody) SetRequestId(v string) *DescribeUsageOverallDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUsageOverallDataResponseBody) SetUsageOverallData(v []*DescribeUsageOverallDataResponseBodyUsageOverallData) *DescribeUsageOverallDataResponseBody {
	s.UsageOverallData = v
	return s
}

func (s *DescribeUsageOverallDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeUsageOverallDataResponseBodyUsageOverallData struct {
	Nodes []*DescribeUsageOverallDataResponseBodyUsageOverallDataNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// example:
	//
	// ONLINE_USER_PEAK
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeUsageOverallDataResponseBodyUsageOverallData) String() string {
	return dara.Prettify(s)
}

func (s DescribeUsageOverallDataResponseBodyUsageOverallData) GoString() string {
	return s.String()
}

func (s *DescribeUsageOverallDataResponseBodyUsageOverallData) GetNodes() []*DescribeUsageOverallDataResponseBodyUsageOverallDataNodes {
	return s.Nodes
}

func (s *DescribeUsageOverallDataResponseBodyUsageOverallData) GetType() *string {
	return s.Type
}

func (s *DescribeUsageOverallDataResponseBodyUsageOverallData) SetNodes(v []*DescribeUsageOverallDataResponseBodyUsageOverallDataNodes) *DescribeUsageOverallDataResponseBodyUsageOverallData {
	s.Nodes = v
	return s
}

func (s *DescribeUsageOverallDataResponseBodyUsageOverallData) SetType(v string) *DescribeUsageOverallDataResponseBodyUsageOverallData {
	s.Type = &v
	return s
}

func (s *DescribeUsageOverallDataResponseBodyUsageOverallData) Validate() error {
	return dara.Validate(s)
}

type DescribeUsageOverallDataResponseBodyUsageOverallDataNodes struct {
	// example:
	//
	// 1615824000
	X *string `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 1
	Y *string `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s DescribeUsageOverallDataResponseBodyUsageOverallDataNodes) String() string {
	return dara.Prettify(s)
}

func (s DescribeUsageOverallDataResponseBodyUsageOverallDataNodes) GoString() string {
	return s.String()
}

func (s *DescribeUsageOverallDataResponseBodyUsageOverallDataNodes) GetX() *string {
	return s.X
}

func (s *DescribeUsageOverallDataResponseBodyUsageOverallDataNodes) GetY() *string {
	return s.Y
}

func (s *DescribeUsageOverallDataResponseBodyUsageOverallDataNodes) SetX(v string) *DescribeUsageOverallDataResponseBodyUsageOverallDataNodes {
	s.X = &v
	return s
}

func (s *DescribeUsageOverallDataResponseBodyUsageOverallDataNodes) SetY(v string) *DescribeUsageOverallDataResponseBodyUsageOverallDataNodes {
	s.Y = &v
	return s
}

func (s *DescribeUsageOverallDataResponseBodyUsageOverallDataNodes) Validate() error {
	return dara.Validate(s)
}

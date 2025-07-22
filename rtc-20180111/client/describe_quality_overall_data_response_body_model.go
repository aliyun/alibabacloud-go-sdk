// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeQualityOverallDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetQualityOverallData(v []*DescribeQualityOverallDataResponseBodyQualityOverallData) *DescribeQualityOverallDataResponseBody
	GetQualityOverallData() []*DescribeQualityOverallDataResponseBodyQualityOverallData
	SetRequestId(v string) *DescribeQualityOverallDataResponseBody
	GetRequestId() *string
}

type DescribeQualityOverallDataResponseBody struct {
	QualityOverallData []*DescribeQualityOverallDataResponseBodyQualityOverallData `json:"QualityOverallData,omitempty" xml:"QualityOverallData,omitempty" type:"Repeated"`
	// example:
	//
	// 231470C1-ACFB-4C9F-844F-4CFE1E3804C5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeQualityOverallDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeQualityOverallDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeQualityOverallDataResponseBody) GetQualityOverallData() []*DescribeQualityOverallDataResponseBodyQualityOverallData {
	return s.QualityOverallData
}

func (s *DescribeQualityOverallDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeQualityOverallDataResponseBody) SetQualityOverallData(v []*DescribeQualityOverallDataResponseBodyQualityOverallData) *DescribeQualityOverallDataResponseBody {
	s.QualityOverallData = v
	return s
}

func (s *DescribeQualityOverallDataResponseBody) SetRequestId(v string) *DescribeQualityOverallDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeQualityOverallDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeQualityOverallDataResponseBodyQualityOverallData struct {
	// example:
	//
	// 0.9376
	Average *string                                                          `json:"Average,omitempty" xml:"Average,omitempty"`
	Nodes   []*DescribeQualityOverallDataResponseBodyQualityOverallDataNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// example:
	//
	// JOIN_CHANNEL_SUC_RATE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeQualityOverallDataResponseBodyQualityOverallData) String() string {
	return dara.Prettify(s)
}

func (s DescribeQualityOverallDataResponseBodyQualityOverallData) GoString() string {
	return s.String()
}

func (s *DescribeQualityOverallDataResponseBodyQualityOverallData) GetAverage() *string {
	return s.Average
}

func (s *DescribeQualityOverallDataResponseBodyQualityOverallData) GetNodes() []*DescribeQualityOverallDataResponseBodyQualityOverallDataNodes {
	return s.Nodes
}

func (s *DescribeQualityOverallDataResponseBodyQualityOverallData) GetType() *string {
	return s.Type
}

func (s *DescribeQualityOverallDataResponseBodyQualityOverallData) SetAverage(v string) *DescribeQualityOverallDataResponseBodyQualityOverallData {
	s.Average = &v
	return s
}

func (s *DescribeQualityOverallDataResponseBodyQualityOverallData) SetNodes(v []*DescribeQualityOverallDataResponseBodyQualityOverallDataNodes) *DescribeQualityOverallDataResponseBodyQualityOverallData {
	s.Nodes = v
	return s
}

func (s *DescribeQualityOverallDataResponseBodyQualityOverallData) SetType(v string) *DescribeQualityOverallDataResponseBodyQualityOverallData {
	s.Type = &v
	return s
}

func (s *DescribeQualityOverallDataResponseBodyQualityOverallData) Validate() error {
	return dara.Validate(s)
}

type DescribeQualityOverallDataResponseBodyQualityOverallDataNodes struct {
	// example:
	//
	// 1615831200
	X *string `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 1.0000
	Y *string `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s DescribeQualityOverallDataResponseBodyQualityOverallDataNodes) String() string {
	return dara.Prettify(s)
}

func (s DescribeQualityOverallDataResponseBodyQualityOverallDataNodes) GoString() string {
	return s.String()
}

func (s *DescribeQualityOverallDataResponseBodyQualityOverallDataNodes) GetX() *string {
	return s.X
}

func (s *DescribeQualityOverallDataResponseBodyQualityOverallDataNodes) GetY() *string {
	return s.Y
}

func (s *DescribeQualityOverallDataResponseBodyQualityOverallDataNodes) SetX(v string) *DescribeQualityOverallDataResponseBodyQualityOverallDataNodes {
	s.X = &v
	return s
}

func (s *DescribeQualityOverallDataResponseBodyQualityOverallDataNodes) SetY(v string) *DescribeQualityOverallDataResponseBodyQualityOverallDataNodes {
	s.Y = &v
	return s
}

func (s *DescribeQualityOverallDataResponseBodyQualityOverallDataNodes) Validate() error {
	return dara.Validate(s)
}

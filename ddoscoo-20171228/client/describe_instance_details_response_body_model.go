// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceDetails(v []*DescribeInstanceDetailsResponseBodyInstanceDetails) *DescribeInstanceDetailsResponseBody
	GetInstanceDetails() []*DescribeInstanceDetailsResponseBodyInstanceDetails
	SetRequestId(v string) *DescribeInstanceDetailsResponseBody
	GetRequestId() *string
}

type DescribeInstanceDetailsResponseBody struct {
	InstanceDetails []*DescribeInstanceDetailsResponseBodyInstanceDetails `json:"InstanceDetails,omitempty" xml:"InstanceDetails,omitempty" type:"Repeated"`
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDetailsResponseBody) GetInstanceDetails() []*DescribeInstanceDetailsResponseBodyInstanceDetails {
	return s.InstanceDetails
}

func (s *DescribeInstanceDetailsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceDetailsResponseBody) SetInstanceDetails(v []*DescribeInstanceDetailsResponseBodyInstanceDetails) *DescribeInstanceDetailsResponseBody {
	s.InstanceDetails = v
	return s
}

func (s *DescribeInstanceDetailsResponseBody) SetRequestId(v string) *DescribeInstanceDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceDetailsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceDetailsResponseBodyInstanceDetails struct {
	EipInfoList []*DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfoList `json:"EipInfoList,omitempty" xml:"EipInfoList,omitempty" type:"Repeated"`
	// example:
	//
	// ddoscoo-cn-XXXXX
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// coop-line-001
	Line *string `json:"Line,omitempty" xml:"Line,omitempty"`
}

func (s DescribeInstanceDetailsResponseBodyInstanceDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceDetailsResponseBodyInstanceDetails) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDetailsResponseBodyInstanceDetails) GetEipInfoList() []*DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfoList {
	return s.EipInfoList
}

func (s *DescribeInstanceDetailsResponseBodyInstanceDetails) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceDetailsResponseBodyInstanceDetails) GetLine() *string {
	return s.Line
}

func (s *DescribeInstanceDetailsResponseBodyInstanceDetails) SetEipInfoList(v []*DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfoList) *DescribeInstanceDetailsResponseBodyInstanceDetails {
	s.EipInfoList = v
	return s
}

func (s *DescribeInstanceDetailsResponseBodyInstanceDetails) SetInstanceId(v string) *DescribeInstanceDetailsResponseBodyInstanceDetails {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceDetailsResponseBodyInstanceDetails) SetLine(v string) *DescribeInstanceDetailsResponseBodyInstanceDetails {
	s.Line = &v
	return s
}

func (s *DescribeInstanceDetailsResponseBodyInstanceDetails) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfoList struct {
	// example:
	//
	// 1.1.1.1
	Eip *string `json:"Eip,omitempty" xml:"Eip,omitempty"`
	// example:
	//
	// normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfoList) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfoList) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfoList) GetEip() *string {
	return s.Eip
}

func (s *DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfoList) GetStatus() *string {
	return s.Status
}

func (s *DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfoList) SetEip(v string) *DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfoList {
	s.Eip = &v
	return s
}

func (s *DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfoList) SetStatus(v string) *DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfoList {
	s.Status = &v
	return s
}

func (s *DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfoList) Validate() error {
	return dara.Validate(s)
}

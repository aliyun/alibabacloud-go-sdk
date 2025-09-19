// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageRiskLevelStatisticResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImageRiskLevelList(v []*DescribeImageRiskLevelStatisticResponseBodyImageRiskLevelList) *DescribeImageRiskLevelStatisticResponseBody
	GetImageRiskLevelList() []*DescribeImageRiskLevelStatisticResponseBodyImageRiskLevelList
	SetRequestId(v string) *DescribeImageRiskLevelStatisticResponseBody
	GetRequestId() *string
}

type DescribeImageRiskLevelStatisticResponseBody struct {
	// The information about risks at the image level. The risks include vulnerabilities, baselines risks, and malicious file risks.
	ImageRiskLevelList []*DescribeImageRiskLevelStatisticResponseBodyImageRiskLevelList `json:"ImageRiskLevelList,omitempty" xml:"ImageRiskLevelList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// A60DA4EC-7CD8-577D-AD73-******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeImageRiskLevelStatisticResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageRiskLevelStatisticResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageRiskLevelStatisticResponseBody) GetImageRiskLevelList() []*DescribeImageRiskLevelStatisticResponseBodyImageRiskLevelList {
	return s.ImageRiskLevelList
}

func (s *DescribeImageRiskLevelStatisticResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeImageRiskLevelStatisticResponseBody) SetImageRiskLevelList(v []*DescribeImageRiskLevelStatisticResponseBodyImageRiskLevelList) *DescribeImageRiskLevelStatisticResponseBody {
	s.ImageRiskLevelList = v
	return s
}

func (s *DescribeImageRiskLevelStatisticResponseBody) SetRequestId(v string) *DescribeImageRiskLevelStatisticResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImageRiskLevelStatisticResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeImageRiskLevelStatisticResponseBodyImageRiskLevelList struct {
	// The number of images at the risk level.
	//
	// example:
	//
	// 12
	Cnt *int32 `json:"Cnt,omitempty" xml:"Cnt,omitempty"`
	// The risk level of the image. Valid values:
	//
	// 	- **3**: high risk.
	//
	// 	- **2**: medium risk.
	//
	// 	- **1**: low risk.
	//
	// 	- **0**: no risk.
	//
	// example:
	//
	// 0
	ImageRiskLevel *string `json:"ImageRiskLevel,omitempty" xml:"ImageRiskLevel,omitempty"`
}

func (s DescribeImageRiskLevelStatisticResponseBodyImageRiskLevelList) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageRiskLevelStatisticResponseBodyImageRiskLevelList) GoString() string {
	return s.String()
}

func (s *DescribeImageRiskLevelStatisticResponseBodyImageRiskLevelList) GetCnt() *int32 {
	return s.Cnt
}

func (s *DescribeImageRiskLevelStatisticResponseBodyImageRiskLevelList) GetImageRiskLevel() *string {
	return s.ImageRiskLevel
}

func (s *DescribeImageRiskLevelStatisticResponseBodyImageRiskLevelList) SetCnt(v int32) *DescribeImageRiskLevelStatisticResponseBodyImageRiskLevelList {
	s.Cnt = &v
	return s
}

func (s *DescribeImageRiskLevelStatisticResponseBodyImageRiskLevelList) SetImageRiskLevel(v string) *DescribeImageRiskLevelStatisticResponseBodyImageRiskLevelList {
	s.ImageRiskLevel = &v
	return s
}

func (s *DescribeImageRiskLevelStatisticResponseBodyImageRiskLevelList) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInterAuthStatisticsZoneOverviewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeInterAuthStatisticsZoneOverviewResponseBodyData) *DescribeInterAuthStatisticsZoneOverviewResponseBody
	GetData() *DescribeInterAuthStatisticsZoneOverviewResponseBodyData
	SetRequestId(v string) *DescribeInterAuthStatisticsZoneOverviewResponseBody
	GetRequestId() *string
}

type DescribeInterAuthStatisticsZoneOverviewResponseBody struct {
	Data *DescribeInterAuthStatisticsZoneOverviewResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInterAuthStatisticsZoneOverviewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInterAuthStatisticsZoneOverviewResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInterAuthStatisticsZoneOverviewResponseBody) GetData() *DescribeInterAuthStatisticsZoneOverviewResponseBodyData {
	return s.Data
}

func (s *DescribeInterAuthStatisticsZoneOverviewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInterAuthStatisticsZoneOverviewResponseBody) SetData(v *DescribeInterAuthStatisticsZoneOverviewResponseBodyData) *DescribeInterAuthStatisticsZoneOverviewResponseBody {
	s.Data = v
	return s
}

func (s *DescribeInterAuthStatisticsZoneOverviewResponseBody) SetRequestId(v string) *DescribeInterAuthStatisticsZoneOverviewResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInterAuthStatisticsZoneOverviewResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInterAuthStatisticsZoneOverviewResponseBodyData struct {
	// example:
	//
	// 6
	RefusedDomainCount *int64 `json:"RefusedDomainCount,omitempty" xml:"RefusedDomainCount,omitempty"`
	// example:
	//
	// 66
	SuddenDropDomainCount *int64 `json:"SuddenDropDomainCount,omitempty" xml:"SuddenDropDomainCount,omitempty"`
	// example:
	//
	// 56
	SuddenIncreaseDomainCount *int64 `json:"SuddenIncreaseDomainCount,omitempty" xml:"SuddenIncreaseDomainCount,omitempty"`
}

func (s DescribeInterAuthStatisticsZoneOverviewResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeInterAuthStatisticsZoneOverviewResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeInterAuthStatisticsZoneOverviewResponseBodyData) GetRefusedDomainCount() *int64 {
	return s.RefusedDomainCount
}

func (s *DescribeInterAuthStatisticsZoneOverviewResponseBodyData) GetSuddenDropDomainCount() *int64 {
	return s.SuddenDropDomainCount
}

func (s *DescribeInterAuthStatisticsZoneOverviewResponseBodyData) GetSuddenIncreaseDomainCount() *int64 {
	return s.SuddenIncreaseDomainCount
}

func (s *DescribeInterAuthStatisticsZoneOverviewResponseBodyData) SetRefusedDomainCount(v int64) *DescribeInterAuthStatisticsZoneOverviewResponseBodyData {
	s.RefusedDomainCount = &v
	return s
}

func (s *DescribeInterAuthStatisticsZoneOverviewResponseBodyData) SetSuddenDropDomainCount(v int64) *DescribeInterAuthStatisticsZoneOverviewResponseBodyData {
	s.SuddenDropDomainCount = &v
	return s
}

func (s *DescribeInterAuthStatisticsZoneOverviewResponseBodyData) SetSuddenIncreaseDomainCount(v int64) *DescribeInterAuthStatisticsZoneOverviewResponseBodyData {
	s.SuddenIncreaseDomainCount = &v
	return s
}

func (s *DescribeInterAuthStatisticsZoneOverviewResponseBodyData) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePvtzStatisticsZoneOverviewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribePvtzStatisticsZoneOverviewResponseBodyData) *DescribePvtzStatisticsZoneOverviewResponseBody
	GetData() *DescribePvtzStatisticsZoneOverviewResponseBodyData
	SetRequestId(v string) *DescribePvtzStatisticsZoneOverviewResponseBody
	GetRequestId() *string
}

type DescribePvtzStatisticsZoneOverviewResponseBody struct {
	Data *DescribePvtzStatisticsZoneOverviewResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// C6F1D541-E7A6-447A-A2B5-9F7A20B2A8FB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePvtzStatisticsZoneOverviewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePvtzStatisticsZoneOverviewResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePvtzStatisticsZoneOverviewResponseBody) GetData() *DescribePvtzStatisticsZoneOverviewResponseBodyData {
	return s.Data
}

func (s *DescribePvtzStatisticsZoneOverviewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePvtzStatisticsZoneOverviewResponseBody) SetData(v *DescribePvtzStatisticsZoneOverviewResponseBodyData) *DescribePvtzStatisticsZoneOverviewResponseBody {
	s.Data = v
	return s
}

func (s *DescribePvtzStatisticsZoneOverviewResponseBody) SetRequestId(v string) *DescribePvtzStatisticsZoneOverviewResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePvtzStatisticsZoneOverviewResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePvtzStatisticsZoneOverviewResponseBodyData struct {
	// example:
	//
	// 2
	RefusedDomainCount *int64 `json:"RefusedDomainCount,omitempty" xml:"RefusedDomainCount,omitempty"`
	// example:
	//
	// 3
	SuddenDropDomainCount *int64 `json:"SuddenDropDomainCount,omitempty" xml:"SuddenDropDomainCount,omitempty"`
	// example:
	//
	// 5
	SuddenIncreaseDomainCount *int64 `json:"SuddenIncreaseDomainCount,omitempty" xml:"SuddenIncreaseDomainCount,omitempty"`
}

func (s DescribePvtzStatisticsZoneOverviewResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribePvtzStatisticsZoneOverviewResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribePvtzStatisticsZoneOverviewResponseBodyData) GetRefusedDomainCount() *int64 {
	return s.RefusedDomainCount
}

func (s *DescribePvtzStatisticsZoneOverviewResponseBodyData) GetSuddenDropDomainCount() *int64 {
	return s.SuddenDropDomainCount
}

func (s *DescribePvtzStatisticsZoneOverviewResponseBodyData) GetSuddenIncreaseDomainCount() *int64 {
	return s.SuddenIncreaseDomainCount
}

func (s *DescribePvtzStatisticsZoneOverviewResponseBodyData) SetRefusedDomainCount(v int64) *DescribePvtzStatisticsZoneOverviewResponseBodyData {
	s.RefusedDomainCount = &v
	return s
}

func (s *DescribePvtzStatisticsZoneOverviewResponseBodyData) SetSuddenDropDomainCount(v int64) *DescribePvtzStatisticsZoneOverviewResponseBodyData {
	s.SuddenDropDomainCount = &v
	return s
}

func (s *DescribePvtzStatisticsZoneOverviewResponseBodyData) SetSuddenIncreaseDomainCount(v int64) *DescribePvtzStatisticsZoneOverviewResponseBodyData {
	s.SuddenIncreaseDomainCount = &v
	return s
}

func (s *DescribePvtzStatisticsZoneOverviewResponseBodyData) Validate() error {
	return dara.Validate(s)
}

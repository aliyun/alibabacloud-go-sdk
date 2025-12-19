// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVSwitchesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeVSwitchesResponseBodyItems) *DescribeVSwitchesResponseBody
	GetItems() []*DescribeVSwitchesResponseBodyItems
	SetMaxResults(v int32) *DescribeVSwitchesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeVSwitchesResponseBody
	GetNextToken() *string
	SetPageNumber(v string) *DescribeVSwitchesResponseBody
	GetPageNumber() *string
	SetPageSize(v string) *DescribeVSwitchesResponseBody
	GetPageSize() *string
	SetRequestId(v string) *DescribeVSwitchesResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v string) *DescribeVSwitchesResponseBody
	GetTotalRecordCount() *string
}

type DescribeVSwitchesResponseBody struct {
	Items []*DescribeVSwitchesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 459a0909c0315bfbe0b5e8707e68181f
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 4773E4EC-025D-509F-AEA9-D53123FDFB0F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 15
	TotalRecordCount *string `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeVSwitchesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVSwitchesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchesResponseBody) GetItems() []*DescribeVSwitchesResponseBodyItems {
	return s.Items
}

func (s *DescribeVSwitchesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeVSwitchesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeVSwitchesResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeVSwitchesResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeVSwitchesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVSwitchesResponseBody) GetTotalRecordCount() *string {
	return s.TotalRecordCount
}

func (s *DescribeVSwitchesResponseBody) SetItems(v []*DescribeVSwitchesResponseBodyItems) *DescribeVSwitchesResponseBody {
	s.Items = v
	return s
}

func (s *DescribeVSwitchesResponseBody) SetMaxResults(v int32) *DescribeVSwitchesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeVSwitchesResponseBody) SetNextToken(v string) *DescribeVSwitchesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeVSwitchesResponseBody) SetPageNumber(v string) *DescribeVSwitchesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeVSwitchesResponseBody) SetPageSize(v string) *DescribeVSwitchesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeVSwitchesResponseBody) SetRequestId(v string) *DescribeVSwitchesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVSwitchesResponseBody) SetTotalRecordCount(v string) *DescribeVSwitchesResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeVSwitchesResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVSwitchesResponseBodyItems struct {
	// example:
	//
	// vsw-bp1usf8eabhvibkkfde96
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// subnet
	VSwitchName *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeVSwitchesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeVSwitchesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchesResponseBodyItems) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeVSwitchesResponseBodyItems) GetVSwitchName() *string {
	return s.VSwitchName
}

func (s *DescribeVSwitchesResponseBodyItems) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeVSwitchesResponseBodyItems) SetVSwitchId(v string) *DescribeVSwitchesResponseBodyItems {
	s.VSwitchId = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyItems) SetVSwitchName(v string) *DescribeVSwitchesResponseBodyItems {
	s.VSwitchName = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyItems) SetZoneId(v string) *DescribeVSwitchesResponseBodyItems {
	s.ZoneId = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyItems) Validate() error {
	return dara.Validate(s)
}

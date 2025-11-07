// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVerifyFailStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeVerifyFailStatisticsResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeVerifyFailStatisticsResponseBodyResultObject) *DescribeVerifyFailStatisticsResponseBody
	GetResultObject() *DescribeVerifyFailStatisticsResponseBodyResultObject
}

type DescribeVerifyFailStatisticsResponseBody struct {
	// ID of this request.
	//
	// example:
	//
	// C2C596D1-B14B-5D79-9672-61D7686912B2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Authentication result.
	ResultObject *DescribeVerifyFailStatisticsResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s DescribeVerifyFailStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifyFailStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVerifyFailStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVerifyFailStatisticsResponseBody) GetResultObject() *DescribeVerifyFailStatisticsResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeVerifyFailStatisticsResponseBody) SetRequestId(v string) *DescribeVerifyFailStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVerifyFailStatisticsResponseBody) SetResultObject(v *DescribeVerifyFailStatisticsResponseBodyResultObject) *DescribeVerifyFailStatisticsResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeVerifyFailStatisticsResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVerifyFailStatisticsResponseBodyResultObject struct {
	// Line chart of failure reasons.
	Column *DescribeVerifyFailStatisticsResponseBodyResultObjectColumn `json:"Column,omitempty" xml:"Column,omitempty" type:"Struct"`
	// Bar chart of failure reasons.
	Line *DescribeVerifyFailStatisticsResponseBodyResultObjectLine `json:"Line,omitempty" xml:"Line,omitempty" type:"Struct"`
}

func (s DescribeVerifyFailStatisticsResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifyFailStatisticsResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeVerifyFailStatisticsResponseBodyResultObject) GetColumn() *DescribeVerifyFailStatisticsResponseBodyResultObjectColumn {
	return s.Column
}

func (s *DescribeVerifyFailStatisticsResponseBodyResultObject) GetLine() *DescribeVerifyFailStatisticsResponseBodyResultObjectLine {
	return s.Line
}

func (s *DescribeVerifyFailStatisticsResponseBodyResultObject) SetColumn(v *DescribeVerifyFailStatisticsResponseBodyResultObjectColumn) *DescribeVerifyFailStatisticsResponseBodyResultObject {
	s.Column = v
	return s
}

func (s *DescribeVerifyFailStatisticsResponseBodyResultObject) SetLine(v *DescribeVerifyFailStatisticsResponseBodyResultObjectLine) *DescribeVerifyFailStatisticsResponseBodyResultObject {
	s.Line = v
	return s
}

func (s *DescribeVerifyFailStatisticsResponseBodyResultObject) Validate() error {
	if s.Column != nil {
		if err := s.Column.Validate(); err != nil {
			return err
		}
	}
	if s.Line != nil {
		if err := s.Line.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVerifyFailStatisticsResponseBodyResultObjectColumn struct {
	// Column information.
	Items []*DescribeVerifyFailStatisticsResponseBodyResultObjectColumnItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// Total count.
	//
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeVerifyFailStatisticsResponseBodyResultObjectColumn) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifyFailStatisticsResponseBodyResultObjectColumn) GoString() string {
	return s.String()
}

func (s *DescribeVerifyFailStatisticsResponseBodyResultObjectColumn) GetItems() []*DescribeVerifyFailStatisticsResponseBodyResultObjectColumnItems {
	return s.Items
}

func (s *DescribeVerifyFailStatisticsResponseBodyResultObjectColumn) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeVerifyFailStatisticsResponseBodyResultObjectColumn) SetItems(v []*DescribeVerifyFailStatisticsResponseBodyResultObjectColumnItems) *DescribeVerifyFailStatisticsResponseBodyResultObjectColumn {
	s.Items = v
	return s
}

func (s *DescribeVerifyFailStatisticsResponseBodyResultObjectColumn) SetTotalCount(v int64) *DescribeVerifyFailStatisticsResponseBodyResultObjectColumn {
	s.TotalCount = &v
	return s
}

func (s *DescribeVerifyFailStatisticsResponseBodyResultObjectColumn) Validate() error {
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

type DescribeVerifyFailStatisticsResponseBodyResultObjectColumnItems struct {
	// Error code.
	//
	// example:
	//
	// 404
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Failure count.
	//
	// example:
	//
	// 1
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// Date: Date
	//
	// example:
	//
	// 2025-10-16
	Rate *string `json:"Rate,omitempty" xml:"Rate,omitempty"`
}

func (s DescribeVerifyFailStatisticsResponseBodyResultObjectColumnItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifyFailStatisticsResponseBodyResultObjectColumnItems) GoString() string {
	return s.String()
}

func (s *DescribeVerifyFailStatisticsResponseBodyResultObjectColumnItems) GetCode() *string {
	return s.Code
}

func (s *DescribeVerifyFailStatisticsResponseBodyResultObjectColumnItems) GetCount() *int64 {
	return s.Count
}

func (s *DescribeVerifyFailStatisticsResponseBodyResultObjectColumnItems) GetRate() *string {
	return s.Rate
}

func (s *DescribeVerifyFailStatisticsResponseBodyResultObjectColumnItems) SetCode(v string) *DescribeVerifyFailStatisticsResponseBodyResultObjectColumnItems {
	s.Code = &v
	return s
}

func (s *DescribeVerifyFailStatisticsResponseBodyResultObjectColumnItems) SetCount(v int64) *DescribeVerifyFailStatisticsResponseBodyResultObjectColumnItems {
	s.Count = &v
	return s
}

func (s *DescribeVerifyFailStatisticsResponseBodyResultObjectColumnItems) SetRate(v string) *DescribeVerifyFailStatisticsResponseBodyResultObjectColumnItems {
	s.Rate = &v
	return s
}

func (s *DescribeVerifyFailStatisticsResponseBodyResultObjectColumnItems) Validate() error {
	return dara.Validate(s)
}

type DescribeVerifyFailStatisticsResponseBodyResultObjectLine struct {
	// Column information.
	Items []*DescribeVerifyFailStatisticsResponseBodyResultObjectLineItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// Total count.
	//
	// example:
	//
	// 0
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeVerifyFailStatisticsResponseBodyResultObjectLine) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifyFailStatisticsResponseBodyResultObjectLine) GoString() string {
	return s.String()
}

func (s *DescribeVerifyFailStatisticsResponseBodyResultObjectLine) GetItems() []*DescribeVerifyFailStatisticsResponseBodyResultObjectLineItems {
	return s.Items
}

func (s *DescribeVerifyFailStatisticsResponseBodyResultObjectLine) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeVerifyFailStatisticsResponseBodyResultObjectLine) SetItems(v []*DescribeVerifyFailStatisticsResponseBodyResultObjectLineItems) *DescribeVerifyFailStatisticsResponseBodyResultObjectLine {
	s.Items = v
	return s
}

func (s *DescribeVerifyFailStatisticsResponseBodyResultObjectLine) SetTotalCount(v int64) *DescribeVerifyFailStatisticsResponseBodyResultObjectLine {
	s.TotalCount = &v
	return s
}

func (s *DescribeVerifyFailStatisticsResponseBodyResultObjectLine) Validate() error {
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

type DescribeVerifyFailStatisticsResponseBodyResultObjectLineItems struct {
	// Error code.
	//
	// example:
	//
	// 404
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Returned quantity data.
	Data []*DescribeVerifyFailStatisticsResponseBodyResultObjectLineItemsData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s DescribeVerifyFailStatisticsResponseBodyResultObjectLineItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifyFailStatisticsResponseBodyResultObjectLineItems) GoString() string {
	return s.String()
}

func (s *DescribeVerifyFailStatisticsResponseBodyResultObjectLineItems) GetCode() *string {
	return s.Code
}

func (s *DescribeVerifyFailStatisticsResponseBodyResultObjectLineItems) GetData() []*DescribeVerifyFailStatisticsResponseBodyResultObjectLineItemsData {
	return s.Data
}

func (s *DescribeVerifyFailStatisticsResponseBodyResultObjectLineItems) SetCode(v string) *DescribeVerifyFailStatisticsResponseBodyResultObjectLineItems {
	s.Code = &v
	return s
}

func (s *DescribeVerifyFailStatisticsResponseBodyResultObjectLineItems) SetData(v []*DescribeVerifyFailStatisticsResponseBodyResultObjectLineItemsData) *DescribeVerifyFailStatisticsResponseBodyResultObjectLineItems {
	s.Data = v
	return s
}

func (s *DescribeVerifyFailStatisticsResponseBodyResultObjectLineItems) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVerifyFailStatisticsResponseBodyResultObjectLineItemsData struct {
	// Error code.
	//
	// example:
	//
	// 404
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Count.
	//
	// example:
	//
	// 9
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// Date.
	//
	// example:
	//
	// 2025-10-16
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
}

func (s DescribeVerifyFailStatisticsResponseBodyResultObjectLineItemsData) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifyFailStatisticsResponseBodyResultObjectLineItemsData) GoString() string {
	return s.String()
}

func (s *DescribeVerifyFailStatisticsResponseBodyResultObjectLineItemsData) GetCode() *string {
	return s.Code
}

func (s *DescribeVerifyFailStatisticsResponseBodyResultObjectLineItemsData) GetCount() *int64 {
	return s.Count
}

func (s *DescribeVerifyFailStatisticsResponseBodyResultObjectLineItemsData) GetDate() *string {
	return s.Date
}

func (s *DescribeVerifyFailStatisticsResponseBodyResultObjectLineItemsData) SetCode(v string) *DescribeVerifyFailStatisticsResponseBodyResultObjectLineItemsData {
	s.Code = &v
	return s
}

func (s *DescribeVerifyFailStatisticsResponseBodyResultObjectLineItemsData) SetCount(v int64) *DescribeVerifyFailStatisticsResponseBodyResultObjectLineItemsData {
	s.Count = &v
	return s
}

func (s *DescribeVerifyFailStatisticsResponseBodyResultObjectLineItemsData) SetDate(v string) *DescribeVerifyFailStatisticsResponseBodyResultObjectLineItemsData {
	s.Date = &v
	return s
}

func (s *DescribeVerifyFailStatisticsResponseBodyResultObjectLineItemsData) Validate() error {
	return dara.Validate(s)
}

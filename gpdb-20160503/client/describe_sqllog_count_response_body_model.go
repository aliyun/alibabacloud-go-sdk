// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSQLLogCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeSQLLogCountResponseBody
	GetDBClusterId() *string
	SetEndTime(v string) *DescribeSQLLogCountResponseBody
	GetEndTime() *string
	SetItems(v []*DescribeSQLLogCountResponseBodyItems) *DescribeSQLLogCountResponseBody
	GetItems() []*DescribeSQLLogCountResponseBodyItems
	SetRequestId(v string) *DescribeSQLLogCountResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeSQLLogCountResponseBody
	GetStartTime() *string
}

type DescribeSQLLogCountResponseBody struct {
	// The instance ID.
	//
	// example:
	//
	// gp-xxxxxxxx
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end time of the query.
	//
	// example:
	//
	// 2020-12-14T11:22Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the instance.
	Items []*DescribeSQLLogCountResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 7565770E-7C45-462D-BA4A-8A**********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start time of the query.
	//
	// example:
	//
	// 2020-12-12T11:22Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSQLLogCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLLogCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogCountResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeSQLLogCountResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeSQLLogCountResponseBody) GetItems() []*DescribeSQLLogCountResponseBodyItems {
	return s.Items
}

func (s *DescribeSQLLogCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSQLLogCountResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeSQLLogCountResponseBody) SetDBClusterId(v string) *DescribeSQLLogCountResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSQLLogCountResponseBody) SetEndTime(v string) *DescribeSQLLogCountResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeSQLLogCountResponseBody) SetItems(v []*DescribeSQLLogCountResponseBodyItems) *DescribeSQLLogCountResponseBody {
	s.Items = v
	return s
}

func (s *DescribeSQLLogCountResponseBody) SetRequestId(v string) *DescribeSQLLogCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSQLLogCountResponseBody) SetStartTime(v string) *DescribeSQLLogCountResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeSQLLogCountResponseBody) Validate() error {
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

type DescribeSQLLogCountResponseBodyItems struct {
	// The name of the table.
	//
	// example:
	//
	// gp-xxxxxxxx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Details of the audit logs.
	Series []*DescribeSQLLogCountResponseBodyItemsSeries `json:"Series,omitempty" xml:"Series,omitempty" type:"Repeated"`
}

func (s DescribeSQLLogCountResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLLogCountResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogCountResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *DescribeSQLLogCountResponseBodyItems) GetSeries() []*DescribeSQLLogCountResponseBodyItemsSeries {
	return s.Series
}

func (s *DescribeSQLLogCountResponseBodyItems) SetName(v string) *DescribeSQLLogCountResponseBodyItems {
	s.Name = &v
	return s
}

func (s *DescribeSQLLogCountResponseBodyItems) SetSeries(v []*DescribeSQLLogCountResponseBodyItemsSeries) *DescribeSQLLogCountResponseBodyItems {
	s.Series = v
	return s
}

func (s *DescribeSQLLogCountResponseBodyItems) Validate() error {
	if s.Series != nil {
		for _, item := range s.Series {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSQLLogCountResponseBodyItemsSeries struct {
	// Details of the audit logs.
	Values []*DescribeSQLLogCountResponseBodyItemsSeriesValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeSQLLogCountResponseBodyItemsSeries) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLLogCountResponseBodyItemsSeries) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogCountResponseBodyItemsSeries) GetValues() []*DescribeSQLLogCountResponseBodyItemsSeriesValues {
	return s.Values
}

func (s *DescribeSQLLogCountResponseBodyItemsSeries) SetValues(v []*DescribeSQLLogCountResponseBodyItemsSeriesValues) *DescribeSQLLogCountResponseBodyItemsSeries {
	s.Values = v
	return s
}

func (s *DescribeSQLLogCountResponseBodyItemsSeries) Validate() error {
	if s.Values != nil {
		for _, item := range s.Values {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSQLLogCountResponseBodyItemsSeriesValues struct {
	// The time when the audit logs were generated and the number of the audit logs.
	Point []*string `json:"Point,omitempty" xml:"Point,omitempty" type:"Repeated"`
}

func (s DescribeSQLLogCountResponseBodyItemsSeriesValues) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLLogCountResponseBodyItemsSeriesValues) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogCountResponseBodyItemsSeriesValues) GetPoint() []*string {
	return s.Point
}

func (s *DescribeSQLLogCountResponseBodyItemsSeriesValues) SetPoint(v []*string) *DescribeSQLLogCountResponseBodyItemsSeriesValues {
	s.Point = v
	return s
}

func (s *DescribeSQLLogCountResponseBodyItemsSeriesValues) Validate() error {
	return dara.Validate(s)
}

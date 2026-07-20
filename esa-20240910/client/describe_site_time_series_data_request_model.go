// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSiteTimeSeriesDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeSiteTimeSeriesDataRequest
	GetEndTime() *string
	SetFields(v []*DescribeSiteTimeSeriesDataRequestFields) *DescribeSiteTimeSeriesDataRequest
	GetFields() []*DescribeSiteTimeSeriesDataRequestFields
	SetInterval(v string) *DescribeSiteTimeSeriesDataRequest
	GetInterval() *string
	SetSiteId(v string) *DescribeSiteTimeSeriesDataRequest
	GetSiteId() *string
	SetStartTime(v string) *DescribeSiteTimeSeriesDataRequest
	GetStartTime() *string
}

type DescribeSiteTimeSeriesDataRequest struct {
	// The end time for the data query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC+0.
	//
	// > The end time must be later than the start time.
	//
	// example:
	//
	// 2023-04-09T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The query metrics.
	//
	// This parameter is required.
	Fields []*DescribeSiteTimeSeriesDataRequestFields `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	// The time granularity for the query data, in seconds.
	//
	// Based on the maximum time span of a single query, this parameter supports the values 60 (1 minute), 300 (5 minutes), 3600 (1 hour), and 86400 (1 day). For more information, see the **Supported time granularity*	- section above.
	//
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The site ID. You can call the [ListSites](~~ListSites~~) operation to obtain the site ID.
	//
	// If this parameter is left empty, account-level data is queried.
	//
	// example:
	//
	// 1150376036*****
	SiteId *string `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The start time for the data query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC+0.
	//
	// example:
	//
	// 2023-04-08T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSiteTimeSeriesDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteTimeSeriesDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeSiteTimeSeriesDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeSiteTimeSeriesDataRequest) GetFields() []*DescribeSiteTimeSeriesDataRequestFields {
	return s.Fields
}

func (s *DescribeSiteTimeSeriesDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeSiteTimeSeriesDataRequest) GetSiteId() *string {
	return s.SiteId
}

func (s *DescribeSiteTimeSeriesDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeSiteTimeSeriesDataRequest) SetEndTime(v string) *DescribeSiteTimeSeriesDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSiteTimeSeriesDataRequest) SetFields(v []*DescribeSiteTimeSeriesDataRequestFields) *DescribeSiteTimeSeriesDataRequest {
	s.Fields = v
	return s
}

func (s *DescribeSiteTimeSeriesDataRequest) SetInterval(v string) *DescribeSiteTimeSeriesDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeSiteTimeSeriesDataRequest) SetSiteId(v string) *DescribeSiteTimeSeriesDataRequest {
	s.SiteId = &v
	return s
}

func (s *DescribeSiteTimeSeriesDataRequest) SetStartTime(v string) *DescribeSiteTimeSeriesDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSiteTimeSeriesDataRequest) Validate() error {
	if s.Fields != nil {
		for _, item := range s.Fields {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSiteTimeSeriesDataRequestFields struct {
	// The query dimension.
	Dimension []*string `json:"Dimension,omitempty" xml:"Dimension,omitempty" type:"Repeated"`
	// The metric name.
	//
	// > For specific dimensions, see [Data analytics field description](https://help.aliyun.com/document_detail/2878520.html).
	//
	// example:
	//
	// Traffic
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
}

func (s DescribeSiteTimeSeriesDataRequestFields) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteTimeSeriesDataRequestFields) GoString() string {
	return s.String()
}

func (s *DescribeSiteTimeSeriesDataRequestFields) GetDimension() []*string {
	return s.Dimension
}

func (s *DescribeSiteTimeSeriesDataRequestFields) GetFieldName() *string {
	return s.FieldName
}

func (s *DescribeSiteTimeSeriesDataRequestFields) SetDimension(v []*string) *DescribeSiteTimeSeriesDataRequestFields {
	s.Dimension = v
	return s
}

func (s *DescribeSiteTimeSeriesDataRequestFields) SetFieldName(v string) *DescribeSiteTimeSeriesDataRequestFields {
	s.FieldName = &v
	return s
}

func (s *DescribeSiteTimeSeriesDataRequestFields) Validate() error {
	return dara.Validate(s)
}

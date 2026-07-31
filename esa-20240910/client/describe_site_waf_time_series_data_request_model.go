// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSiteWafTimeSeriesDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeSiteWafTimeSeriesDataRequest
	GetEndTime() *string
	SetFields(v []*DescribeSiteWafTimeSeriesDataRequestFields) *DescribeSiteWafTimeSeriesDataRequest
	GetFields() []*DescribeSiteWafTimeSeriesDataRequestFields
	SetInterval(v string) *DescribeSiteWafTimeSeriesDataRequest
	GetInterval() *string
	SetSiteId(v string) *DescribeSiteWafTimeSeriesDataRequest
	GetSiteId() *string
	SetStartTime(v string) *DescribeSiteWafTimeSeriesDataRequest
	GetStartTime() *string
}

type DescribeSiteWafTimeSeriesDataRequest struct {
	// The end of the time range to query.
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
	Fields []*DescribeSiteWafTimeSeriesDataRequestFields `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	// The time granularity for querying data, in seconds.
	//
	// Based on the maximum time span of a single query, this parameter supports the values 60 (1 minute), 300 (5 minutes), 3600 (1 hour), and 86400 (1 day). For more information, see the **Supported time granularity*	- section above.
	//
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The site ID. You can call the [ListSites](~~ListSites~~) operation to obtain the site ID.
	//
	// If this parameter is left empty, user-level data is queried.
	//
	// example:
	//
	// 11089268296****
	SiteId *string `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The beginning of the time range to query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC+0.
	//
	// example:
	//
	// 2023-04-08T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSiteWafTimeSeriesDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteWafTimeSeriesDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeSiteWafTimeSeriesDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeSiteWafTimeSeriesDataRequest) GetFields() []*DescribeSiteWafTimeSeriesDataRequestFields {
	return s.Fields
}

func (s *DescribeSiteWafTimeSeriesDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeSiteWafTimeSeriesDataRequest) GetSiteId() *string {
	return s.SiteId
}

func (s *DescribeSiteWafTimeSeriesDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeSiteWafTimeSeriesDataRequest) SetEndTime(v string) *DescribeSiteWafTimeSeriesDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSiteWafTimeSeriesDataRequest) SetFields(v []*DescribeSiteWafTimeSeriesDataRequestFields) *DescribeSiteWafTimeSeriesDataRequest {
	s.Fields = v
	return s
}

func (s *DescribeSiteWafTimeSeriesDataRequest) SetInterval(v string) *DescribeSiteWafTimeSeriesDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeSiteWafTimeSeriesDataRequest) SetSiteId(v string) *DescribeSiteWafTimeSeriesDataRequest {
	s.SiteId = &v
	return s
}

func (s *DescribeSiteWafTimeSeriesDataRequest) SetStartTime(v string) *DescribeSiteWafTimeSeriesDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSiteWafTimeSeriesDataRequest) Validate() error {
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

type DescribeSiteWafTimeSeriesDataRequestFields struct {
	// The query dimensions.
	Dimension []*string `json:"Dimension,omitempty" xml:"Dimension,omitempty" type:"Repeated"`
	// The metric name.
	//
	// >For specific dimensions, see [Data analytics field description](https://help.aliyun.com/document_detail/2878520.html).
	//
	// example:
	//
	// Requests
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
}

func (s DescribeSiteWafTimeSeriesDataRequestFields) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteWafTimeSeriesDataRequestFields) GoString() string {
	return s.String()
}

func (s *DescribeSiteWafTimeSeriesDataRequestFields) GetDimension() []*string {
	return s.Dimension
}

func (s *DescribeSiteWafTimeSeriesDataRequestFields) GetFieldName() *string {
	return s.FieldName
}

func (s *DescribeSiteWafTimeSeriesDataRequestFields) SetDimension(v []*string) *DescribeSiteWafTimeSeriesDataRequestFields {
	s.Dimension = v
	return s
}

func (s *DescribeSiteWafTimeSeriesDataRequestFields) SetFieldName(v string) *DescribeSiteWafTimeSeriesDataRequestFields {
	s.FieldName = &v
	return s
}

func (s *DescribeSiteWafTimeSeriesDataRequestFields) Validate() error {
	return dara.Validate(s)
}

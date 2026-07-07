// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSiteWafTopDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeSiteWafTopDataRequest
	GetEndTime() *string
	SetFields(v []*DescribeSiteWafTopDataRequestFields) *DescribeSiteWafTopDataRequest
	GetFields() []*DescribeSiteWafTopDataRequestFields
	SetInterval(v string) *DescribeSiteWafTopDataRequest
	GetInterval() *string
	SetLimit(v string) *DescribeSiteWafTopDataRequest
	GetLimit() *string
	SetSiteId(v string) *DescribeSiteWafTopDataRequest
	GetSiteId() *string
	SetStartTime(v string) *DescribeSiteWafTopDataRequest
	GetStartTime() *string
}

type DescribeSiteWafTopDataRequest struct {
	// The end of the time range to query.
	//
	// Specify the time in ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC+0.
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
	Fields []*DescribeSiteWafTopDataRequestFields `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	// The time granularity for querying data. Unit: seconds.
	//
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The number of top data entries to query.
	//
	// example:
	//
	// 5
	Limit *string `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The site ID. You can call the [ListSites](~~ListSites~~) operation to obtain the site ID.
	//
	//
	// If this parameter is left empty, user-level data is queried.
	//
	// example:
	//
	// 1150376036*****
	SiteId *string `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The beginning of the time range to query.
	//
	// Specify the time in ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC+0.
	//
	// example:
	//
	// 2023-04-08T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSiteWafTopDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteWafTopDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeSiteWafTopDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeSiteWafTopDataRequest) GetFields() []*DescribeSiteWafTopDataRequestFields {
	return s.Fields
}

func (s *DescribeSiteWafTopDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeSiteWafTopDataRequest) GetLimit() *string {
	return s.Limit
}

func (s *DescribeSiteWafTopDataRequest) GetSiteId() *string {
	return s.SiteId
}

func (s *DescribeSiteWafTopDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeSiteWafTopDataRequest) SetEndTime(v string) *DescribeSiteWafTopDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSiteWafTopDataRequest) SetFields(v []*DescribeSiteWafTopDataRequestFields) *DescribeSiteWafTopDataRequest {
	s.Fields = v
	return s
}

func (s *DescribeSiteWafTopDataRequest) SetInterval(v string) *DescribeSiteWafTopDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeSiteWafTopDataRequest) SetLimit(v string) *DescribeSiteWafTopDataRequest {
	s.Limit = &v
	return s
}

func (s *DescribeSiteWafTopDataRequest) SetSiteId(v string) *DescribeSiteWafTopDataRequest {
	s.SiteId = &v
	return s
}

func (s *DescribeSiteWafTopDataRequest) SetStartTime(v string) *DescribeSiteWafTopDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSiteWafTopDataRequest) Validate() error {
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

type DescribeSiteWafTopDataRequestFields struct {
	// The query dimension.
	Dimension []*string `json:"Dimension,omitempty" xml:"Dimension,omitempty" type:"Repeated"`
	// The query metric value.
	//
	// > For specific dimensions, see [Data analytics field description](https://help.aliyun.com/document_detail/2878520.html).
	//
	// example:
	//
	// Requests
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
}

func (s DescribeSiteWafTopDataRequestFields) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteWafTopDataRequestFields) GoString() string {
	return s.String()
}

func (s *DescribeSiteWafTopDataRequestFields) GetDimension() []*string {
	return s.Dimension
}

func (s *DescribeSiteWafTopDataRequestFields) GetFieldName() *string {
	return s.FieldName
}

func (s *DescribeSiteWafTopDataRequestFields) SetDimension(v []*string) *DescribeSiteWafTopDataRequestFields {
	s.Dimension = v
	return s
}

func (s *DescribeSiteWafTopDataRequestFields) SetFieldName(v string) *DescribeSiteWafTopDataRequestFields {
	s.FieldName = &v
	return s
}

func (s *DescribeSiteWafTopDataRequestFields) Validate() error {
	return dara.Validate(s)
}

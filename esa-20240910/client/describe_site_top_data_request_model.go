// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSiteTopDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeSiteTopDataRequest
	GetEndTime() *string
	SetFields(v []*DescribeSiteTopDataRequestFields) *DescribeSiteTopDataRequest
	GetFields() []*DescribeSiteTopDataRequestFields
	SetInterval(v string) *DescribeSiteTopDataRequest
	GetInterval() *string
	SetLimit(v string) *DescribeSiteTopDataRequest
	GetLimit() *string
	SetSiteId(v string) *DescribeSiteTopDataRequest
	GetSiteId() *string
	SetStartTime(v string) *DescribeSiteTopDataRequest
	GetStartTime() *string
}

type DescribeSiteTopDataRequest struct {
	// example:
	//
	// 2023-04-09T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	Fields []*DescribeSiteTopDataRequestFields `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	Limit    *string `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// example:
	//
	// 1150376036*****
	SiteId *string `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// example:
	//
	// 2023-04-08T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSiteTopDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteTopDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeSiteTopDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeSiteTopDataRequest) GetFields() []*DescribeSiteTopDataRequestFields {
	return s.Fields
}

func (s *DescribeSiteTopDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeSiteTopDataRequest) GetLimit() *string {
	return s.Limit
}

func (s *DescribeSiteTopDataRequest) GetSiteId() *string {
	return s.SiteId
}

func (s *DescribeSiteTopDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeSiteTopDataRequest) SetEndTime(v string) *DescribeSiteTopDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSiteTopDataRequest) SetFields(v []*DescribeSiteTopDataRequestFields) *DescribeSiteTopDataRequest {
	s.Fields = v
	return s
}

func (s *DescribeSiteTopDataRequest) SetInterval(v string) *DescribeSiteTopDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeSiteTopDataRequest) SetLimit(v string) *DescribeSiteTopDataRequest {
	s.Limit = &v
	return s
}

func (s *DescribeSiteTopDataRequest) SetSiteId(v string) *DescribeSiteTopDataRequest {
	s.SiteId = &v
	return s
}

func (s *DescribeSiteTopDataRequest) SetStartTime(v string) *DescribeSiteTopDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSiteTopDataRequest) Validate() error {
	return dara.Validate(s)
}

type DescribeSiteTopDataRequestFields struct {
	Dimension []*string `json:"Dimension,omitempty" xml:"Dimension,omitempty" type:"Repeated"`
	// example:
	//
	// Traffic
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
}

func (s DescribeSiteTopDataRequestFields) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteTopDataRequestFields) GoString() string {
	return s.String()
}

func (s *DescribeSiteTopDataRequestFields) GetDimension() []*string {
	return s.Dimension
}

func (s *DescribeSiteTopDataRequestFields) GetFieldName() *string {
	return s.FieldName
}

func (s *DescribeSiteTopDataRequestFields) SetDimension(v []*string) *DescribeSiteTopDataRequestFields {
	s.Dimension = v
	return s
}

func (s *DescribeSiteTopDataRequestFields) SetFieldName(v string) *DescribeSiteTopDataRequestFields {
	s.FieldName = &v
	return s
}

func (s *DescribeSiteTopDataRequestFields) Validate() error {
	return dara.Validate(s)
}

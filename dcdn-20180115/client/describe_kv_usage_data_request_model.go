// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKvUsageDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessType(v string) *DescribeKvUsageDataRequest
	GetAccessType() *string
	SetEndTime(v string) *DescribeKvUsageDataRequest
	GetEndTime() *string
	SetField(v string) *DescribeKvUsageDataRequest
	GetField() *string
	SetNamespaceId(v string) *DescribeKvUsageDataRequest
	GetNamespaceId() *string
	SetResponseType(v string) *DescribeKvUsageDataRequest
	GetResponseType() *string
	SetSplitBy(v string) *DescribeKvUsageDataRequest
	GetSplitBy() *string
	SetStartTime(v string) *DescribeKvUsageDataRequest
	GetStartTime() *string
}

type DescribeKvUsageDataRequest struct {
	// The request method. If the parameter is empty, data about all methods is returned. Valid values:
	//
	// 	- **get**
	//
	// 	- **put**
	//
	// 	- **list**
	//
	// 	- **delete**
	//
	// example:
	//
	// get
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2022-08-10T23:59:59Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The type of the request data. Set the value to **acc**.
	//
	// This parameter is required.
	//
	// example:
	//
	// acc
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// The namespace ID. If the parameter is empty, data about all namespaces is returned.
	//
	// You can specify a maximum number of 30 namespace IDs and separate them with commas (,).
	//
	// example:
	//
	// 12423131231****
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The type of the response data. Valid values:
	//
	// 	- **detail**: detailed data
	//
	// 	- **total**: summary data
	//
	// Default value: **detail**.
	//
	// example:
	//
	// detail
	ResponseType *string `json:"ResponseType,omitempty" xml:"ResponseType,omitempty"`
	// The key that is used to group data. Valid values: **type*	- and **namespace**.
	//
	// 	- **type**: Data is grouped by time. The data in the last 5 minutes is returned.
	//
	// 	- **namespace**: Data is grouped by namespace and is not padded with zeros.
	//
	// 	- Default value: **type**.
	//
	// If **ResponseType*	- is set to **total**, data to return is not grouped by **namespace*	- but by **type**.
	//
	// example:
	//
	// type
	SplitBy *string `json:"SplitBy,omitempty" xml:"SplitBy,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// The minimum data granularity is 1 hour. If you do not specify this parameter, the data in the last seven days is returned.
	//
	// example:
	//
	// 2022-08-10T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeKvUsageDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeKvUsageDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeKvUsageDataRequest) GetAccessType() *string {
	return s.AccessType
}

func (s *DescribeKvUsageDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeKvUsageDataRequest) GetField() *string {
	return s.Field
}

func (s *DescribeKvUsageDataRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *DescribeKvUsageDataRequest) GetResponseType() *string {
	return s.ResponseType
}

func (s *DescribeKvUsageDataRequest) GetSplitBy() *string {
	return s.SplitBy
}

func (s *DescribeKvUsageDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeKvUsageDataRequest) SetAccessType(v string) *DescribeKvUsageDataRequest {
	s.AccessType = &v
	return s
}

func (s *DescribeKvUsageDataRequest) SetEndTime(v string) *DescribeKvUsageDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeKvUsageDataRequest) SetField(v string) *DescribeKvUsageDataRequest {
	s.Field = &v
	return s
}

func (s *DescribeKvUsageDataRequest) SetNamespaceId(v string) *DescribeKvUsageDataRequest {
	s.NamespaceId = &v
	return s
}

func (s *DescribeKvUsageDataRequest) SetResponseType(v string) *DescribeKvUsageDataRequest {
	s.ResponseType = &v
	return s
}

func (s *DescribeKvUsageDataRequest) SetSplitBy(v string) *DescribeKvUsageDataRequest {
	s.SplitBy = &v
	return s
}

func (s *DescribeKvUsageDataRequest) SetStartTime(v string) *DescribeKvUsageDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeKvUsageDataRequest) Validate() error {
	return dara.Validate(s)
}

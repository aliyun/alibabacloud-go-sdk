// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUserOmsDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataType(v string) *QueryUserOmsDataRequest
	GetDataType() *string
	SetEndTime(v string) *QueryUserOmsDataRequest
	GetEndTime() *string
	SetMarker(v string) *QueryUserOmsDataRequest
	GetMarker() *string
	SetOwnerId(v int64) *QueryUserOmsDataRequest
	GetOwnerId() *int64
	SetPageSize(v int32) *QueryUserOmsDataRequest
	GetPageSize() *int32
	SetStartTime(v string) *QueryUserOmsDataRequest
	GetStartTime() *string
	SetTable(v string) *QueryUserOmsDataRequest
	GetTable() *string
}

type QueryUserOmsDataRequest struct {
	// The time type of the usage data. Set the parameter based on the description in the documentation of the specified service. Valid values:
	//
	// 	- Raw
	//
	// 	- Hour
	//
	// 	- Day
	//
	// 	- Month
	//
	// This parameter is required.
	//
	// example:
	//
	// Hour
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The end of the time range to query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-02-21T12:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the record from which the usage data starts to return. The usage data records whose names are alphabetically after the value of the Marker parameter are returned. By default, the usage data starts to return from the earliest record.
	//
	// example:
	//
	// NextToken
	Marker  *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of entries to return on each page. Valid values: 1 to 200. Default value: 100.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The beginning of the time range to query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-02-20T12:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The service whose usage data you want to query and the details of the usage data. The parameter value is usually set to the code of a service. Various usage models are provided for different services.
	//
	// This parameter is required.
	//
	// example:
	//
	// rds
	Table *string `json:"Table,omitempty" xml:"Table,omitempty"`
}

func (s QueryUserOmsDataRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryUserOmsDataRequest) GoString() string {
	return s.String()
}

func (s *QueryUserOmsDataRequest) GetDataType() *string {
	return s.DataType
}

func (s *QueryUserOmsDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *QueryUserOmsDataRequest) GetMarker() *string {
	return s.Marker
}

func (s *QueryUserOmsDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryUserOmsDataRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryUserOmsDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *QueryUserOmsDataRequest) GetTable() *string {
	return s.Table
}

func (s *QueryUserOmsDataRequest) SetDataType(v string) *QueryUserOmsDataRequest {
	s.DataType = &v
	return s
}

func (s *QueryUserOmsDataRequest) SetEndTime(v string) *QueryUserOmsDataRequest {
	s.EndTime = &v
	return s
}

func (s *QueryUserOmsDataRequest) SetMarker(v string) *QueryUserOmsDataRequest {
	s.Marker = &v
	return s
}

func (s *QueryUserOmsDataRequest) SetOwnerId(v int64) *QueryUserOmsDataRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryUserOmsDataRequest) SetPageSize(v int32) *QueryUserOmsDataRequest {
	s.PageSize = &v
	return s
}

func (s *QueryUserOmsDataRequest) SetStartTime(v string) *QueryUserOmsDataRequest {
	s.StartTime = &v
	return s
}

func (s *QueryUserOmsDataRequest) SetTable(v string) *QueryUserOmsDataRequest {
	s.Table = &v
	return s
}

func (s *QueryUserOmsDataRequest) Validate() error {
	return dara.Validate(s)
}

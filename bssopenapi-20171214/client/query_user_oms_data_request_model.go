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
	// The data timestamp type of the metering data. Refer to the specific service documentation for details. Valid values:
	//
	// - Raw: hourly
	//
	// - Hour: hourly
	//
	// - Day: daily
	//
	// - Month: monthly
	//
	// This parameter is required.
	//
	// example:
	//
	// Hour
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The end time of the request.
	//
	// The date format must comply with the ISO 8601 standard and use UTC time. Format: yyyy-mm-ddThh:mm:ssZ.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-02-21T12:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The marker that specifies the starting point for returning results in alphabetical order after the NextToken. By default, results are returned from the beginning.
	//
	// example:
	//
	// NextToken
	Marker  *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of rows per page for paging in a paged query. Valid values: 1 to 200. Default value: 100.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The start time of the request.
	//
	// The date format must comply with the ISO 8601 standard and use UTC time. Format: yyyy-mm-ddThh:mm:ssZ.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-02-20T12:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The product and detail information for the usage data, which is typically the ProductCode. Different services support multiple metering models.
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

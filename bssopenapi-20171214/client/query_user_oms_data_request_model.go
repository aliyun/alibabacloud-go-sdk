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
	// This parameter is required.
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// This parameter is required.
	EndTime  *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Marker   *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// This parameter is required.
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

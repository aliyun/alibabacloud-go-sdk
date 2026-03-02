// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPdpLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFrom(v int64) *ListPdpLogsRequest
	GetFrom() *int64
	SetIp(v string) *ListPdpLogsRequest
	GetIp() *string
	SetPageNumber(v int32) *ListPdpLogsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPdpLogsRequest
	GetPageSize() *int32
	SetQuery(v string) *ListPdpLogsRequest
	GetQuery() *string
	SetServiceGroupId(v int64) *ListPdpLogsRequest
	GetServiceGroupId() *int64
	SetSourceType(v string) *ListPdpLogsRequest
	GetSourceType() *string
	SetTo(v int64) *ListPdpLogsRequest
	GetTo() *int64
}

type ListPdpLogsRequest struct {
	From           *int64  `json:"from,omitempty" xml:"from,omitempty"`
	Ip             *string `json:"ip,omitempty" xml:"ip,omitempty"`
	PageNumber     *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize       *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Query          *string `json:"query,omitempty" xml:"query,omitempty"`
	ServiceGroupId *int64  `json:"serviceGroupId,omitempty" xml:"serviceGroupId,omitempty"`
	SourceType     *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	To             *int64  `json:"to,omitempty" xml:"to,omitempty"`
}

func (s ListPdpLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPdpLogsRequest) GoString() string {
	return s.String()
}

func (s *ListPdpLogsRequest) GetFrom() *int64 {
	return s.From
}

func (s *ListPdpLogsRequest) GetIp() *string {
	return s.Ip
}

func (s *ListPdpLogsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPdpLogsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPdpLogsRequest) GetQuery() *string {
	return s.Query
}

func (s *ListPdpLogsRequest) GetServiceGroupId() *int64 {
	return s.ServiceGroupId
}

func (s *ListPdpLogsRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *ListPdpLogsRequest) GetTo() *int64 {
	return s.To
}

func (s *ListPdpLogsRequest) SetFrom(v int64) *ListPdpLogsRequest {
	s.From = &v
	return s
}

func (s *ListPdpLogsRequest) SetIp(v string) *ListPdpLogsRequest {
	s.Ip = &v
	return s
}

func (s *ListPdpLogsRequest) SetPageNumber(v int32) *ListPdpLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPdpLogsRequest) SetPageSize(v int32) *ListPdpLogsRequest {
	s.PageSize = &v
	return s
}

func (s *ListPdpLogsRequest) SetQuery(v string) *ListPdpLogsRequest {
	s.Query = &v
	return s
}

func (s *ListPdpLogsRequest) SetServiceGroupId(v int64) *ListPdpLogsRequest {
	s.ServiceGroupId = &v
	return s
}

func (s *ListPdpLogsRequest) SetSourceType(v string) *ListPdpLogsRequest {
	s.SourceType = &v
	return s
}

func (s *ListPdpLogsRequest) SetTo(v int64) *ListPdpLogsRequest {
	s.To = &v
	return s
}

func (s *ListPdpLogsRequest) Validate() error {
	return dara.Validate(s)
}

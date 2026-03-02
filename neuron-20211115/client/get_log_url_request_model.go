// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLogUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIp(v string) *GetLogUrlRequest
	GetIp() *string
	SetOrderBy(v string) *GetLogUrlRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *GetLogUrlRequest
	GetOrderDirection() *string
	SetPageNumber(v int32) *GetLogUrlRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *GetLogUrlRequest
	GetPageSize() *int32
	SetQuery(v string) *GetLogUrlRequest
	GetQuery() *string
	SetServiceGroupId(v int64) *GetLogUrlRequest
	GetServiceGroupId() *int64
	SetSourceType(v string) *GetLogUrlRequest
	GetSourceType() *string
	SetTo(v int32) *GetLogUrlRequest
	GetTo() *int32
}

type GetLogUrlRequest struct {
	// ip
	//
	// example:
	//
	// 172.35.1.139
	Ip             *string `json:"ip,omitempty" xml:"ip,omitempty"`
	OrderBy        *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	OrderDirection *string `json:"orderDirection,omitempty" xml:"orderDirection,omitempty"`
	PageNumber     *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize       *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// Exception
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// This parameter is required.
	ServiceGroupId *int64  `json:"serviceGroupId,omitempty" xml:"serviceGroupId,omitempty"`
	SourceType     *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// example:
	//
	// 1667985359
	To *int32 `json:"to,omitempty" xml:"to,omitempty"`
}

func (s GetLogUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLogUrlRequest) GoString() string {
	return s.String()
}

func (s *GetLogUrlRequest) GetIp() *string {
	return s.Ip
}

func (s *GetLogUrlRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *GetLogUrlRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *GetLogUrlRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetLogUrlRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetLogUrlRequest) GetQuery() *string {
	return s.Query
}

func (s *GetLogUrlRequest) GetServiceGroupId() *int64 {
	return s.ServiceGroupId
}

func (s *GetLogUrlRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *GetLogUrlRequest) GetTo() *int32 {
	return s.To
}

func (s *GetLogUrlRequest) SetIp(v string) *GetLogUrlRequest {
	s.Ip = &v
	return s
}

func (s *GetLogUrlRequest) SetOrderBy(v string) *GetLogUrlRequest {
	s.OrderBy = &v
	return s
}

func (s *GetLogUrlRequest) SetOrderDirection(v string) *GetLogUrlRequest {
	s.OrderDirection = &v
	return s
}

func (s *GetLogUrlRequest) SetPageNumber(v int32) *GetLogUrlRequest {
	s.PageNumber = &v
	return s
}

func (s *GetLogUrlRequest) SetPageSize(v int32) *GetLogUrlRequest {
	s.PageSize = &v
	return s
}

func (s *GetLogUrlRequest) SetQuery(v string) *GetLogUrlRequest {
	s.Query = &v
	return s
}

func (s *GetLogUrlRequest) SetServiceGroupId(v int64) *GetLogUrlRequest {
	s.ServiceGroupId = &v
	return s
}

func (s *GetLogUrlRequest) SetSourceType(v string) *GetLogUrlRequest {
	s.SourceType = &v
	return s
}

func (s *GetLogUrlRequest) SetTo(v int32) *GetLogUrlRequest {
	s.To = &v
	return s
}

func (s *GetLogUrlRequest) Validate() error {
	return dara.Validate(s)
}

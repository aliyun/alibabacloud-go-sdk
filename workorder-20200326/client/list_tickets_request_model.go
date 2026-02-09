// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTicketsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAfterTime(v int64) *ListTicketsRequest
	GetCreatedAfterTime() *int64
	SetCreatedBeforeTime(v int64) *ListTicketsRequest
	GetCreatedBeforeTime() *int64
	SetIds(v string) *ListTicketsRequest
	GetIds() *string
	SetLanguage(v string) *ListTicketsRequest
	GetLanguage() *string
	SetPageSize(v int32) *ListTicketsRequest
	GetPageSize() *int32
	SetPageStart(v int32) *ListTicketsRequest
	GetPageStart() *int32
	SetProductCode(v string) *ListTicketsRequest
	GetProductCode() *string
	SetSubUserId(v string) *ListTicketsRequest
	GetSubUserId() *string
	SetTicketStatus(v string) *ListTicketsRequest
	GetTicketStatus() *string
}

type ListTicketsRequest struct {
	// example:
	//
	// 1585790280
	CreatedAfterTime *int64 `json:"CreatedAfterTime,omitempty" xml:"CreatedAfterTime,omitempty"`
	// example:
	//
	// 1586049480
	CreatedBeforeTime *int64 `json:"CreatedBeforeTime,omitempty" xml:"CreatedBeforeTime,omitempty"`
	// example:
	//
	// ETARNPP,RTARNP
	Ids *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	// example:
	//
	// zh
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1
	PageStart *int32 `json:"PageStart,omitempty" xml:"PageStart,omitempty"`
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// example:
	//
	// 252448085032933023
	SubUserId *string `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
	// example:
	//
	// wait_confirm,dealing
	TicketStatus *string `json:"TicketStatus,omitempty" xml:"TicketStatus,omitempty"`
}

func (s ListTicketsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTicketsRequest) GoString() string {
	return s.String()
}

func (s *ListTicketsRequest) GetCreatedAfterTime() *int64 {
	return s.CreatedAfterTime
}

func (s *ListTicketsRequest) GetCreatedBeforeTime() *int64 {
	return s.CreatedBeforeTime
}

func (s *ListTicketsRequest) GetIds() *string {
	return s.Ids
}

func (s *ListTicketsRequest) GetLanguage() *string {
	return s.Language
}

func (s *ListTicketsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTicketsRequest) GetPageStart() *int32 {
	return s.PageStart
}

func (s *ListTicketsRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListTicketsRequest) GetSubUserId() *string {
	return s.SubUserId
}

func (s *ListTicketsRequest) GetTicketStatus() *string {
	return s.TicketStatus
}

func (s *ListTicketsRequest) SetCreatedAfterTime(v int64) *ListTicketsRequest {
	s.CreatedAfterTime = &v
	return s
}

func (s *ListTicketsRequest) SetCreatedBeforeTime(v int64) *ListTicketsRequest {
	s.CreatedBeforeTime = &v
	return s
}

func (s *ListTicketsRequest) SetIds(v string) *ListTicketsRequest {
	s.Ids = &v
	return s
}

func (s *ListTicketsRequest) SetLanguage(v string) *ListTicketsRequest {
	s.Language = &v
	return s
}

func (s *ListTicketsRequest) SetPageSize(v int32) *ListTicketsRequest {
	s.PageSize = &v
	return s
}

func (s *ListTicketsRequest) SetPageStart(v int32) *ListTicketsRequest {
	s.PageStart = &v
	return s
}

func (s *ListTicketsRequest) SetProductCode(v string) *ListTicketsRequest {
	s.ProductCode = &v
	return s
}

func (s *ListTicketsRequest) SetSubUserId(v string) *ListTicketsRequest {
	s.SubUserId = &v
	return s
}

func (s *ListTicketsRequest) SetTicketStatus(v string) *ListTicketsRequest {
	s.TicketStatus = &v
	return s
}

func (s *ListTicketsRequest) Validate() error {
	return dara.Validate(s)
}

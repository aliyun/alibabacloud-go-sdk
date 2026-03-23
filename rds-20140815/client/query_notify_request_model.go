// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryNotifyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFrom(v string) *QueryNotifyRequest
	GetFrom() *string
	SetPageNumber(v int32) *QueryNotifyRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *QueryNotifyRequest
	GetPageSize() *int32
	SetTo(v string) *QueryNotifyRequest
	GetTo() *string
	SetWithConfirmed(v bool) *QueryNotifyRequest
	GetWithConfirmed() *bool
}

type QueryNotifyRequest struct {
	// This parameter is required.
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	To *string `json:"To,omitempty" xml:"To,omitempty"`
	// This parameter is required.
	WithConfirmed *bool `json:"WithConfirmed,omitempty" xml:"WithConfirmed,omitempty"`
}

func (s QueryNotifyRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryNotifyRequest) GoString() string {
	return s.String()
}

func (s *QueryNotifyRequest) GetFrom() *string {
	return s.From
}

func (s *QueryNotifyRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *QueryNotifyRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryNotifyRequest) GetTo() *string {
	return s.To
}

func (s *QueryNotifyRequest) GetWithConfirmed() *bool {
	return s.WithConfirmed
}

func (s *QueryNotifyRequest) SetFrom(v string) *QueryNotifyRequest {
	s.From = &v
	return s
}

func (s *QueryNotifyRequest) SetPageNumber(v int32) *QueryNotifyRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryNotifyRequest) SetPageSize(v int32) *QueryNotifyRequest {
	s.PageSize = &v
	return s
}

func (s *QueryNotifyRequest) SetTo(v string) *QueryNotifyRequest {
	s.To = &v
	return s
}

func (s *QueryNotifyRequest) SetWithConfirmed(v bool) *QueryNotifyRequest {
	s.WithConfirmed = &v
	return s
}

func (s *QueryNotifyRequest) Validate() error {
	return dara.Validate(s)
}

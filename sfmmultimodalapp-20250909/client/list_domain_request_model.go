// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListDomainRequest
	GetAppId() *string
	SetPageNumber(v int32) *ListDomainRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDomainRequest
	GetPageSize() *int32
	SetType(v string) *ListDomainRequest
	GetType() *string
}

type ListDomainRequest struct {
	AppId      *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDomainRequest) GoString() string {
	return s.String()
}

func (s *ListDomainRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListDomainRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDomainRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDomainRequest) GetType() *string {
	return s.Type
}

func (s *ListDomainRequest) SetAppId(v string) *ListDomainRequest {
	s.AppId = &v
	return s
}

func (s *ListDomainRequest) SetPageNumber(v int32) *ListDomainRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDomainRequest) SetPageSize(v int32) *ListDomainRequest {
	s.PageSize = &v
	return s
}

func (s *ListDomainRequest) SetType(v string) *ListDomainRequest {
	s.Type = &v
	return s
}

func (s *ListDomainRequest) Validate() error {
	return dara.Validate(s)
}

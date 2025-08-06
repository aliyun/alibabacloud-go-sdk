// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVodEsTemplateInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLanguage(v string) *ListVodEsTemplateInfoRequest
	GetLanguage() *string
	SetOp(v string) *ListVodEsTemplateInfoRequest
	GetOp() *string
	SetOwnerId(v int64) *ListVodEsTemplateInfoRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *ListVodEsTemplateInfoRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListVodEsTemplateInfoRequest
	GetPageSize() *int32
}

type ListVodEsTemplateInfoRequest struct {
	Language   *string `json:"Language,omitempty" xml:"Language,omitempty"`
	Op         *string `json:"Op,omitempty" xml:"Op,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListVodEsTemplateInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVodEsTemplateInfoRequest) GoString() string {
	return s.String()
}

func (s *ListVodEsTemplateInfoRequest) GetLanguage() *string {
	return s.Language
}

func (s *ListVodEsTemplateInfoRequest) GetOp() *string {
	return s.Op
}

func (s *ListVodEsTemplateInfoRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListVodEsTemplateInfoRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListVodEsTemplateInfoRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListVodEsTemplateInfoRequest) SetLanguage(v string) *ListVodEsTemplateInfoRequest {
	s.Language = &v
	return s
}

func (s *ListVodEsTemplateInfoRequest) SetOp(v string) *ListVodEsTemplateInfoRequest {
	s.Op = &v
	return s
}

func (s *ListVodEsTemplateInfoRequest) SetOwnerId(v int64) *ListVodEsTemplateInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *ListVodEsTemplateInfoRequest) SetPageNumber(v int32) *ListVodEsTemplateInfoRequest {
	s.PageNumber = &v
	return s
}

func (s *ListVodEsTemplateInfoRequest) SetPageSize(v int32) *ListVodEsTemplateInfoRequest {
	s.PageSize = &v
	return s
}

func (s *ListVodEsTemplateInfoRequest) Validate() error {
	return dara.Validate(s)
}

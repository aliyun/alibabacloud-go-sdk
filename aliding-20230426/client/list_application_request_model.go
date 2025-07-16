// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppFilter(v string) *ListApplicationRequest
	GetAppFilter() *string
	SetAppNameSearchKeyword(v string) *ListApplicationRequest
	GetAppNameSearchKeyword() *string
	SetCorpId(v string) *ListApplicationRequest
	GetCorpId() *string
	SetPageNumber(v int32) *ListApplicationRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListApplicationRequest
	GetPageSize() *int32
	SetToken(v string) *ListApplicationRequest
	GetToken() *string
}

type ListApplicationRequest struct {
	// example:
	//
	// createdByMe
	AppFilter *string `json:"AppFilter,omitempty" xml:"AppFilter,omitempty"`
	// example:
	//
	// keyword
	AppNameSearchKeyword *string `json:"AppNameSearchKeyword,omitempty" xml:"AppNameSearchKeyword,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// corpid
	CorpId *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	// example:
	//
	// keyword
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// keyword
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// keyword
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s ListApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationRequest) GetAppFilter() *string {
	return s.AppFilter
}

func (s *ListApplicationRequest) GetAppNameSearchKeyword() *string {
	return s.AppNameSearchKeyword
}

func (s *ListApplicationRequest) GetCorpId() *string {
	return s.CorpId
}

func (s *ListApplicationRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListApplicationRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListApplicationRequest) GetToken() *string {
	return s.Token
}

func (s *ListApplicationRequest) SetAppFilter(v string) *ListApplicationRequest {
	s.AppFilter = &v
	return s
}

func (s *ListApplicationRequest) SetAppNameSearchKeyword(v string) *ListApplicationRequest {
	s.AppNameSearchKeyword = &v
	return s
}

func (s *ListApplicationRequest) SetCorpId(v string) *ListApplicationRequest {
	s.CorpId = &v
	return s
}

func (s *ListApplicationRequest) SetPageNumber(v int32) *ListApplicationRequest {
	s.PageNumber = &v
	return s
}

func (s *ListApplicationRequest) SetPageSize(v int32) *ListApplicationRequest {
	s.PageSize = &v
	return s
}

func (s *ListApplicationRequest) SetToken(v string) *ListApplicationRequest {
	s.Token = &v
	return s
}

func (s *ListApplicationRequest) Validate() error {
	return dara.Validate(s)
}

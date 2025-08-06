// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlayEventListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLanguage(v string) *DescribePlayEventListRequest
	GetLanguage() *string
	SetPageNo(v int64) *DescribePlayEventListRequest
	GetPageNo() *int64
	SetPageSize(v int64) *DescribePlayEventListRequest
	GetPageSize() *int64
	SetPlayTs(v string) *DescribePlayEventListRequest
	GetPlayTs() *string
	SetSessionId(v string) *DescribePlayEventListRequest
	GetSessionId() *string
}

type DescribePlayEventListRequest struct {
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// This parameter is required.
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// This parameter is required.
	PageSize *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PlayTs   *string `json:"PlayTs,omitempty" xml:"PlayTs,omitempty"`
	// This parameter is required.
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s DescribePlayEventListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePlayEventListRequest) GoString() string {
	return s.String()
}

func (s *DescribePlayEventListRequest) GetLanguage() *string {
	return s.Language
}

func (s *DescribePlayEventListRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *DescribePlayEventListRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribePlayEventListRequest) GetPlayTs() *string {
	return s.PlayTs
}

func (s *DescribePlayEventListRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *DescribePlayEventListRequest) SetLanguage(v string) *DescribePlayEventListRequest {
	s.Language = &v
	return s
}

func (s *DescribePlayEventListRequest) SetPageNo(v int64) *DescribePlayEventListRequest {
	s.PageNo = &v
	return s
}

func (s *DescribePlayEventListRequest) SetPageSize(v int64) *DescribePlayEventListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePlayEventListRequest) SetPlayTs(v string) *DescribePlayEventListRequest {
	s.PlayTs = &v
	return s
}

func (s *DescribePlayEventListRequest) SetSessionId(v string) *DescribePlayEventListRequest {
	s.SessionId = &v
	return s
}

func (s *DescribePlayEventListRequest) Validate() error {
	return dara.Validate(s)
}

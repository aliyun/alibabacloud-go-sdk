// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamsBlockListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeLiveStreamsBlockListRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DescribeLiveStreamsBlockListRequest
	GetOwnerId() *int64
	SetPageNum(v int32) *DescribeLiveStreamsBlockListRequest
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeLiveStreamsBlockListRequest
	GetPageSize() *int32
	SetSecurityToken(v string) *DescribeLiveStreamsBlockListRequest
	GetSecurityToken() *string
}

type DescribeLiveStreamsBlockListRequest struct {
	// The main streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 2
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries to return on each page. Valid values: 1 to 3000. Default value: 2000.
	//
	// example:
	//
	// 10
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeLiveStreamsBlockListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamsBlockListRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamsBlockListRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveStreamsBlockListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveStreamsBlockListRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeLiveStreamsBlockListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLiveStreamsBlockListRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeLiveStreamsBlockListRequest) SetDomainName(v string) *DescribeLiveStreamsBlockListRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveStreamsBlockListRequest) SetOwnerId(v int64) *DescribeLiveStreamsBlockListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveStreamsBlockListRequest) SetPageNum(v int32) *DescribeLiveStreamsBlockListRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeLiveStreamsBlockListRequest) SetPageSize(v int32) *DescribeLiveStreamsBlockListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveStreamsBlockListRequest) SetSecurityToken(v string) *DescribeLiveStreamsBlockListRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeLiveStreamsBlockListRequest) Validate() error {
	return dara.Validate(s)
}

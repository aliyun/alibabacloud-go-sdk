// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUnionTaskListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBrandUserId(v int64) *QueryUnionTaskListRequest
	GetBrandUserId() *int64
	SetBrandUserNick(v string) *QueryUnionTaskListRequest
	GetBrandUserNick() *string
	SetChannelId(v string) *QueryUnionTaskListRequest
	GetChannelId() *string
	SetPageIndex(v int32) *QueryUnionTaskListRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *QueryUnionTaskListRequest
	GetPageSize() *int32
	SetProxyUserId(v int64) *QueryUnionTaskListRequest
	GetProxyUserId() *int64
}

type QueryUnionTaskListRequest struct {
	BrandUserId   *int64  `json:"BrandUserId,omitempty" xml:"BrandUserId,omitempty"`
	BrandUserNick *string `json:"BrandUserNick,omitempty" xml:"BrandUserNick,omitempty"`
	// This parameter is required.
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// This parameter is required.
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// This parameter is required.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	ProxyUserId *int64 `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
}

func (s QueryUnionTaskListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryUnionTaskListRequest) GoString() string {
	return s.String()
}

func (s *QueryUnionTaskListRequest) GetBrandUserId() *int64 {
	return s.BrandUserId
}

func (s *QueryUnionTaskListRequest) GetBrandUserNick() *string {
	return s.BrandUserNick
}

func (s *QueryUnionTaskListRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *QueryUnionTaskListRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *QueryUnionTaskListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryUnionTaskListRequest) GetProxyUserId() *int64 {
	return s.ProxyUserId
}

func (s *QueryUnionTaskListRequest) SetBrandUserId(v int64) *QueryUnionTaskListRequest {
	s.BrandUserId = &v
	return s
}

func (s *QueryUnionTaskListRequest) SetBrandUserNick(v string) *QueryUnionTaskListRequest {
	s.BrandUserNick = &v
	return s
}

func (s *QueryUnionTaskListRequest) SetChannelId(v string) *QueryUnionTaskListRequest {
	s.ChannelId = &v
	return s
}

func (s *QueryUnionTaskListRequest) SetPageIndex(v int32) *QueryUnionTaskListRequest {
	s.PageIndex = &v
	return s
}

func (s *QueryUnionTaskListRequest) SetPageSize(v int32) *QueryUnionTaskListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryUnionTaskListRequest) SetProxyUserId(v int64) *QueryUnionTaskListRequest {
	s.ProxyUserId = &v
	return s
}

func (s *QueryUnionTaskListRequest) Validate() error {
	return dara.Validate(s)
}

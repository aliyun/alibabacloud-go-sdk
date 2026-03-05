// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUnionContentInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelId(v string) *QueryUnionContentInfoRequest
	GetChannelId() *string
	SetContentId(v string) *QueryUnionContentInfoRequest
	GetContentId() *string
	SetProxyUserId(v int64) *QueryUnionContentInfoRequest
	GetProxyUserId() *int64
}

type QueryUnionContentInfoRequest struct {
	// This parameter is required.
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// This parameter is required.
	ContentId *string `json:"ContentId,omitempty" xml:"ContentId,omitempty"`
	// This parameter is required.
	ProxyUserId *int64 `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
}

func (s QueryUnionContentInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryUnionContentInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryUnionContentInfoRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *QueryUnionContentInfoRequest) GetContentId() *string {
	return s.ContentId
}

func (s *QueryUnionContentInfoRequest) GetProxyUserId() *int64 {
	return s.ProxyUserId
}

func (s *QueryUnionContentInfoRequest) SetChannelId(v string) *QueryUnionContentInfoRequest {
	s.ChannelId = &v
	return s
}

func (s *QueryUnionContentInfoRequest) SetContentId(v string) *QueryUnionContentInfoRequest {
	s.ContentId = &v
	return s
}

func (s *QueryUnionContentInfoRequest) SetProxyUserId(v int64) *QueryUnionContentInfoRequest {
	s.ProxyUserId = &v
	return s
}

func (s *QueryUnionContentInfoRequest) Validate() error {
	return dara.Validate(s)
}

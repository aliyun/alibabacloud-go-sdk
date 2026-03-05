// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUnionChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelId(v string) *QueryUnionChannelRequest
	GetChannelId() *string
}

type QueryUnionChannelRequest struct {
	// This parameter is required.
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
}

func (s QueryUnionChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryUnionChannelRequest) GoString() string {
	return s.String()
}

func (s *QueryUnionChannelRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *QueryUnionChannelRequest) SetChannelId(v string) *QueryUnionChannelRequest {
	s.ChannelId = &v
	return s
}

func (s *QueryUnionChannelRequest) Validate() error {
	return dara.Validate(s)
}

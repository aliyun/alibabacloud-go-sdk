// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTaskBizTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelId(v string) *QueryTaskBizTypeRequest
	GetChannelId() *string
}

type QueryTaskBizTypeRequest struct {
	// This parameter is required.
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
}

func (s QueryTaskBizTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryTaskBizTypeRequest) GoString() string {
	return s.String()
}

func (s *QueryTaskBizTypeRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *QueryTaskBizTypeRequest) SetChannelId(v string) *QueryTaskBizTypeRequest {
	s.ChannelId = &v
	return s
}

func (s *QueryTaskBizTypeRequest) Validate() error {
	return dara.Validate(s)
}

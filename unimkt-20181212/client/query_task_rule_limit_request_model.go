// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTaskRuleLimitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelId(v string) *QueryTaskRuleLimitRequest
	GetChannelId() *string
}

type QueryTaskRuleLimitRequest struct {
	// This parameter is required.
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
}

func (s QueryTaskRuleLimitRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryTaskRuleLimitRequest) GoString() string {
	return s.String()
}

func (s *QueryTaskRuleLimitRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *QueryTaskRuleLimitRequest) SetChannelId(v string) *QueryTaskRuleLimitRequest {
	s.ChannelId = &v
	return s
}

func (s *QueryTaskRuleLimitRequest) Validate() error {
	return dara.Validate(s)
}

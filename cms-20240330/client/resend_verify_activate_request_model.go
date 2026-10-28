// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResendVerifyActivateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelType(v string) *ResendVerifyActivateRequest
	GetChannelType() *string
}

type ResendVerifyActivateRequest struct {
	// example:
	//
	// PHONE
	ChannelType *string `json:"channelType,omitempty" xml:"channelType,omitempty"`
}

func (s ResendVerifyActivateRequest) String() string {
	return dara.Prettify(s)
}

func (s ResendVerifyActivateRequest) GoString() string {
	return s.String()
}

func (s *ResendVerifyActivateRequest) GetChannelType() *string {
	return s.ChannelType
}

func (s *ResendVerifyActivateRequest) SetChannelType(v string) *ResendVerifyActivateRequest {
	s.ChannelType = &v
	return s
}

func (s *ResendVerifyActivateRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMyAgenticNotificationChannelAuthorizeUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannel(v string) *GetMyAgenticNotificationChannelAuthorizeUrlRequest
	GetChannel() *string
}

type GetMyAgenticNotificationChannelAuthorizeUrlRequest struct {
	// The notification channel type. Currently, only FEISHU (the Chinese version of Lark) is supported. Specifying other values results in a parameter error.
	//
	// This parameter is required.
	//
	// example:
	//
	// FEISHU
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
}

func (s GetMyAgenticNotificationChannelAuthorizeUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMyAgenticNotificationChannelAuthorizeUrlRequest) GoString() string {
	return s.String()
}

func (s *GetMyAgenticNotificationChannelAuthorizeUrlRequest) GetChannel() *string {
	return s.Channel
}

func (s *GetMyAgenticNotificationChannelAuthorizeUrlRequest) SetChannel(v string) *GetMyAgenticNotificationChannelAuthorizeUrlRequest {
	s.Channel = &v
	return s
}

func (s *GetMyAgenticNotificationChannelAuthorizeUrlRequest) Validate() error {
	return dara.Validate(s)
}

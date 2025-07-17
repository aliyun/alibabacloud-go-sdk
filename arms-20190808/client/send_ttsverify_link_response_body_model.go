// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendTTSVerifyLinkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsSuccess(v bool) *SendTTSVerifyLinkResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *SendTTSVerifyLinkResponseBody
	GetRequestId() *string
}

type SendTTSVerifyLinkResponseBody struct {
	// Indicates whether the text message was sent.
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 21E85B16-75A6-429A-9F65-8AAC9A54****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendTTSVerifyLinkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendTTSVerifyLinkResponseBody) GoString() string {
	return s.String()
}

func (s *SendTTSVerifyLinkResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *SendTTSVerifyLinkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendTTSVerifyLinkResponseBody) SetIsSuccess(v bool) *SendTTSVerifyLinkResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *SendTTSVerifyLinkResponseBody) SetRequestId(v string) *SendTTSVerifyLinkResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendTTSVerifyLinkResponseBody) Validate() error {
	return dara.Validate(s)
}

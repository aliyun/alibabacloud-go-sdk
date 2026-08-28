// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshPluginOAuthCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RefreshPluginOAuthCodeResponseBody
	GetCode() *string
	SetMessage(v string) *RefreshPluginOAuthCodeResponseBody
	GetMessage() *string
	SetRequestId(v string) *RefreshPluginOAuthCodeResponseBody
	GetRequestId() *string
}

type RefreshPluginOAuthCodeResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 566C6A32-A971-59F2-A9C6-9C73277BA0B8
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s RefreshPluginOAuthCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RefreshPluginOAuthCodeResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshPluginOAuthCodeResponseBody) GetCode() *string {
	return s.Code
}

func (s *RefreshPluginOAuthCodeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RefreshPluginOAuthCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RefreshPluginOAuthCodeResponseBody) SetCode(v string) *RefreshPluginOAuthCodeResponseBody {
	s.Code = &v
	return s
}

func (s *RefreshPluginOAuthCodeResponseBody) SetMessage(v string) *RefreshPluginOAuthCodeResponseBody {
	s.Message = &v
	return s
}

func (s *RefreshPluginOAuthCodeResponseBody) SetRequestId(v string) *RefreshPluginOAuthCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefreshPluginOAuthCodeResponseBody) Validate() error {
	return dara.Validate(s)
}

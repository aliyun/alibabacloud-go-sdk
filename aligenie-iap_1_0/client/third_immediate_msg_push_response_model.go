// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iThirdImmediateMsgPushResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ThirdImmediateMsgPushResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ThirdImmediateMsgPushResponse
	GetStatusCode() *int32
	SetBody(v *ThirdImmediateMsgPushResponseBody) *ThirdImmediateMsgPushResponse
	GetBody() *ThirdImmediateMsgPushResponseBody
}

type ThirdImmediateMsgPushResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ThirdImmediateMsgPushResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ThirdImmediateMsgPushResponse) String() string {
	return dara.Prettify(s)
}

func (s ThirdImmediateMsgPushResponse) GoString() string {
	return s.String()
}

func (s *ThirdImmediateMsgPushResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ThirdImmediateMsgPushResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ThirdImmediateMsgPushResponse) GetBody() *ThirdImmediateMsgPushResponseBody {
	return s.Body
}

func (s *ThirdImmediateMsgPushResponse) SetHeaders(v map[string]*string) *ThirdImmediateMsgPushResponse {
	s.Headers = v
	return s
}

func (s *ThirdImmediateMsgPushResponse) SetStatusCode(v int32) *ThirdImmediateMsgPushResponse {
	s.StatusCode = &v
	return s
}

func (s *ThirdImmediateMsgPushResponse) SetBody(v *ThirdImmediateMsgPushResponseBody) *ThirdImmediateMsgPushResponse {
	s.Body = v
	return s
}

func (s *ThirdImmediateMsgPushResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDingTalkOnlineTestResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DingTalkOnlineTestResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DingTalkOnlineTestResponse
	GetStatusCode() *int32
	SetBody(v *DingTalkOnlineTestResponseBody) *DingTalkOnlineTestResponse
	GetBody() *DingTalkOnlineTestResponseBody
}

type DingTalkOnlineTestResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DingTalkOnlineTestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DingTalkOnlineTestResponse) String() string {
	return dara.Prettify(s)
}

func (s DingTalkOnlineTestResponse) GoString() string {
	return s.String()
}

func (s *DingTalkOnlineTestResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DingTalkOnlineTestResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DingTalkOnlineTestResponse) GetBody() *DingTalkOnlineTestResponseBody {
	return s.Body
}

func (s *DingTalkOnlineTestResponse) SetHeaders(v map[string]*string) *DingTalkOnlineTestResponse {
	s.Headers = v
	return s
}

func (s *DingTalkOnlineTestResponse) SetStatusCode(v int32) *DingTalkOnlineTestResponse {
	s.StatusCode = &v
	return s
}

func (s *DingTalkOnlineTestResponse) SetBody(v *DingTalkOnlineTestResponseBody) *DingTalkOnlineTestResponse {
	s.Body = v
	return s
}

func (s *DingTalkOnlineTestResponse) Validate() error {
	return dara.Validate(s)
}

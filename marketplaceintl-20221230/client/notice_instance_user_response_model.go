// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNoticeInstanceUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *NoticeInstanceUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *NoticeInstanceUserResponse
	GetStatusCode() *int32
	SetBody(v *NoticeInstanceUserResponseBody) *NoticeInstanceUserResponse
	GetBody() *NoticeInstanceUserResponseBody
}

type NoticeInstanceUserResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *NoticeInstanceUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s NoticeInstanceUserResponse) String() string {
	return dara.Prettify(s)
}

func (s NoticeInstanceUserResponse) GoString() string {
	return s.String()
}

func (s *NoticeInstanceUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *NoticeInstanceUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *NoticeInstanceUserResponse) GetBody() *NoticeInstanceUserResponseBody {
	return s.Body
}

func (s *NoticeInstanceUserResponse) SetHeaders(v map[string]*string) *NoticeInstanceUserResponse {
	s.Headers = v
	return s
}

func (s *NoticeInstanceUserResponse) SetStatusCode(v int32) *NoticeInstanceUserResponse {
	s.StatusCode = &v
	return s
}

func (s *NoticeInstanceUserResponse) SetBody(v *NoticeInstanceUserResponseBody) *NoticeInstanceUserResponse {
	s.Body = v
	return s
}

func (s *NoticeInstanceUserResponse) Validate() error {
	return dara.Validate(s)
}

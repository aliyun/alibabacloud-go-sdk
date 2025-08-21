// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushNoticeToAndroidResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PushNoticeToAndroidResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PushNoticeToAndroidResponse
	GetStatusCode() *int32
	SetBody(v *PushNoticeToAndroidResponseBody) *PushNoticeToAndroidResponse
	GetBody() *PushNoticeToAndroidResponseBody
}

type PushNoticeToAndroidResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PushNoticeToAndroidResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PushNoticeToAndroidResponse) String() string {
	return dara.Prettify(s)
}

func (s PushNoticeToAndroidResponse) GoString() string {
	return s.String()
}

func (s *PushNoticeToAndroidResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PushNoticeToAndroidResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PushNoticeToAndroidResponse) GetBody() *PushNoticeToAndroidResponseBody {
	return s.Body
}

func (s *PushNoticeToAndroidResponse) SetHeaders(v map[string]*string) *PushNoticeToAndroidResponse {
	s.Headers = v
	return s
}

func (s *PushNoticeToAndroidResponse) SetStatusCode(v int32) *PushNoticeToAndroidResponse {
	s.StatusCode = &v
	return s
}

func (s *PushNoticeToAndroidResponse) SetBody(v *PushNoticeToAndroidResponseBody) *PushNoticeToAndroidResponse {
	s.Body = v
	return s
}

func (s *PushNoticeToAndroidResponse) Validate() error {
	return dara.Validate(s)
}

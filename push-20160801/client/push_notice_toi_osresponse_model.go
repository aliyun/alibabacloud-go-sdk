// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushNoticeToiOSResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PushNoticeToiOSResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PushNoticeToiOSResponse
	GetStatusCode() *int32
	SetBody(v *PushNoticeToiOSResponseBody) *PushNoticeToiOSResponse
	GetBody() *PushNoticeToiOSResponseBody
}

type PushNoticeToiOSResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PushNoticeToiOSResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PushNoticeToiOSResponse) String() string {
	return dara.Prettify(s)
}

func (s PushNoticeToiOSResponse) GoString() string {
	return s.String()
}

func (s *PushNoticeToiOSResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PushNoticeToiOSResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PushNoticeToiOSResponse) GetBody() *PushNoticeToiOSResponseBody {
	return s.Body
}

func (s *PushNoticeToiOSResponse) SetHeaders(v map[string]*string) *PushNoticeToiOSResponse {
	s.Headers = v
	return s
}

func (s *PushNoticeToiOSResponse) SetStatusCode(v int32) *PushNoticeToiOSResponse {
	s.StatusCode = &v
	return s
}

func (s *PushNoticeToiOSResponse) SetBody(v *PushNoticeToiOSResponseBody) *PushNoticeToiOSResponse {
	s.Body = v
	return s
}

func (s *PushNoticeToiOSResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

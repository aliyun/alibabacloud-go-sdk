// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnlineBizEntityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OnlineBizEntityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OnlineBizEntityResponse
	GetStatusCode() *int32
	SetBody(v *OnlineBizEntityResponseBody) *OnlineBizEntityResponse
	GetBody() *OnlineBizEntityResponseBody
}

type OnlineBizEntityResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnlineBizEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnlineBizEntityResponse) String() string {
	return dara.Prettify(s)
}

func (s OnlineBizEntityResponse) GoString() string {
	return s.String()
}

func (s *OnlineBizEntityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OnlineBizEntityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OnlineBizEntityResponse) GetBody() *OnlineBizEntityResponseBody {
	return s.Body
}

func (s *OnlineBizEntityResponse) SetHeaders(v map[string]*string) *OnlineBizEntityResponse {
	s.Headers = v
	return s
}

func (s *OnlineBizEntityResponse) SetStatusCode(v int32) *OnlineBizEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *OnlineBizEntityResponse) SetBody(v *OnlineBizEntityResponseBody) *OnlineBizEntityResponse {
	s.Body = v
	return s
}

func (s *OnlineBizEntityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

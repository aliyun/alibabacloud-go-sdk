// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PushFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PushFileResponse
	GetStatusCode() *int32
	SetBody(v *PushFileResponseBody) *PushFileResponse
	GetBody() *PushFileResponseBody
}

type PushFileResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PushFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PushFileResponse) String() string {
	return dara.Prettify(s)
}

func (s PushFileResponse) GoString() string {
	return s.String()
}

func (s *PushFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PushFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PushFileResponse) GetBody() *PushFileResponseBody {
	return s.Body
}

func (s *PushFileResponse) SetHeaders(v map[string]*string) *PushFileResponse {
	s.Headers = v
	return s
}

func (s *PushFileResponse) SetStatusCode(v int32) *PushFileResponse {
	s.StatusCode = &v
	return s
}

func (s *PushFileResponse) SetBody(v *PushFileResponseBody) *PushFileResponse {
	s.Body = v
	return s
}

func (s *PushFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

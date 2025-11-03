// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTestEventSourceConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TestEventSourceConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TestEventSourceConfigResponse
	GetStatusCode() *int32
	SetBody(v *TestEventSourceConfigResponseBody) *TestEventSourceConfigResponse
	GetBody() *TestEventSourceConfigResponseBody
}

type TestEventSourceConfigResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TestEventSourceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TestEventSourceConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s TestEventSourceConfigResponse) GoString() string {
	return s.String()
}

func (s *TestEventSourceConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TestEventSourceConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TestEventSourceConfigResponse) GetBody() *TestEventSourceConfigResponseBody {
	return s.Body
}

func (s *TestEventSourceConfigResponse) SetHeaders(v map[string]*string) *TestEventSourceConfigResponse {
	s.Headers = v
	return s
}

func (s *TestEventSourceConfigResponse) SetStatusCode(v int32) *TestEventSourceConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *TestEventSourceConfigResponse) SetBody(v *TestEventSourceConfigResponseBody) *TestEventSourceConfigResponse {
	s.Body = v
	return s
}

func (s *TestEventSourceConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

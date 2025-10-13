// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindNlbResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnbindNlbResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnbindNlbResponse
	GetStatusCode() *int32
	SetBody(v *UnbindNlbResponseBody) *UnbindNlbResponse
	GetBody() *UnbindNlbResponseBody
}

type UnbindNlbResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindNlbResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindNlbResponse) String() string {
	return dara.Prettify(s)
}

func (s UnbindNlbResponse) GoString() string {
	return s.String()
}

func (s *UnbindNlbResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnbindNlbResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnbindNlbResponse) GetBody() *UnbindNlbResponseBody {
	return s.Body
}

func (s *UnbindNlbResponse) SetHeaders(v map[string]*string) *UnbindNlbResponse {
	s.Headers = v
	return s
}

func (s *UnbindNlbResponse) SetStatusCode(v int32) *UnbindNlbResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindNlbResponse) SetBody(v *UnbindNlbResponseBody) *UnbindNlbResponse {
	s.Body = v
	return s
}

func (s *UnbindNlbResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

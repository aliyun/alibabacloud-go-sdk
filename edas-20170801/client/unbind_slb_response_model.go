// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindSlbResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnbindSlbResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnbindSlbResponse
	GetStatusCode() *int32
	SetBody(v *UnbindSlbResponseBody) *UnbindSlbResponse
	GetBody() *UnbindSlbResponseBody
}

type UnbindSlbResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindSlbResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindSlbResponse) String() string {
	return dara.Prettify(s)
}

func (s UnbindSlbResponse) GoString() string {
	return s.String()
}

func (s *UnbindSlbResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnbindSlbResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnbindSlbResponse) GetBody() *UnbindSlbResponseBody {
	return s.Body
}

func (s *UnbindSlbResponse) SetHeaders(v map[string]*string) *UnbindSlbResponse {
	s.Headers = v
	return s
}

func (s *UnbindSlbResponse) SetStatusCode(v int32) *UnbindSlbResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindSlbResponse) SetBody(v *UnbindSlbResponseBody) *UnbindSlbResponse {
	s.Body = v
	return s
}

func (s *UnbindSlbResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

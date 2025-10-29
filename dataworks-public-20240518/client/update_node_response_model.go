// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateNodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateNodeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateNodeResponseBody) *UpdateNodeResponse
	GetBody() *UpdateNodeResponseBody
}

type UpdateNodeResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateNodeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateNodeResponse) GoString() string {
	return s.String()
}

func (s *UpdateNodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateNodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateNodeResponse) GetBody() *UpdateNodeResponseBody {
	return s.Body
}

func (s *UpdateNodeResponse) SetHeaders(v map[string]*string) *UpdateNodeResponse {
	s.Headers = v
	return s
}

func (s *UpdateNodeResponse) SetStatusCode(v int32) *UpdateNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateNodeResponse) SetBody(v *UpdateNodeResponseBody) *UpdateNodeResponse {
	s.Body = v
	return s
}

func (s *UpdateNodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

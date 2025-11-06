// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateZnodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateZnodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateZnodeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateZnodeResponseBody) *UpdateZnodeResponse
	GetBody() *UpdateZnodeResponseBody
}

type UpdateZnodeResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateZnodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateZnodeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateZnodeResponse) GoString() string {
	return s.String()
}

func (s *UpdateZnodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateZnodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateZnodeResponse) GetBody() *UpdateZnodeResponseBody {
	return s.Body
}

func (s *UpdateZnodeResponse) SetHeaders(v map[string]*string) *UpdateZnodeResponse {
	s.Headers = v
	return s
}

func (s *UpdateZnodeResponse) SetStatusCode(v int32) *UpdateZnodeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateZnodeResponse) SetBody(v *UpdateZnodeResponseBody) *UpdateZnodeResponse {
	s.Body = v
	return s
}

func (s *UpdateZnodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

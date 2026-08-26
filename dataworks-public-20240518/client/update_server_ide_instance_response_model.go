// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServerIdeInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateServerIdeInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateServerIdeInstanceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateServerIdeInstanceResponseBody) *UpdateServerIdeInstanceResponse
	GetBody() *UpdateServerIdeInstanceResponseBody
}

type UpdateServerIdeInstanceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateServerIdeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateServerIdeInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateServerIdeInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpdateServerIdeInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateServerIdeInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateServerIdeInstanceResponse) GetBody() *UpdateServerIdeInstanceResponseBody {
	return s.Body
}

func (s *UpdateServerIdeInstanceResponse) SetHeaders(v map[string]*string) *UpdateServerIdeInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpdateServerIdeInstanceResponse) SetStatusCode(v int32) *UpdateServerIdeInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServerIdeInstanceResponse) SetBody(v *UpdateServerIdeInstanceResponseBody) *UpdateServerIdeInstanceResponse {
	s.Body = v
	return s
}

func (s *UpdateServerIdeInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

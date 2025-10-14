// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddBackendServersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddBackendServersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddBackendServersResponse
	GetStatusCode() *int32
	SetBody(v *AddBackendServersResponseBody) *AddBackendServersResponse
	GetBody() *AddBackendServersResponseBody
}

type AddBackendServersResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddBackendServersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddBackendServersResponse) String() string {
	return dara.Prettify(s)
}

func (s AddBackendServersResponse) GoString() string {
	return s.String()
}

func (s *AddBackendServersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddBackendServersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddBackendServersResponse) GetBody() *AddBackendServersResponseBody {
	return s.Body
}

func (s *AddBackendServersResponse) SetHeaders(v map[string]*string) *AddBackendServersResponse {
	s.Headers = v
	return s
}

func (s *AddBackendServersResponse) SetStatusCode(v int32) *AddBackendServersResponse {
	s.StatusCode = &v
	return s
}

func (s *AddBackendServersResponse) SetBody(v *AddBackendServersResponseBody) *AddBackendServersResponse {
	s.Body = v
	return s
}

func (s *AddBackendServersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

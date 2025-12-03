// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetBackendServersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetBackendServersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetBackendServersResponse
	GetStatusCode() *int32
	SetBody(v *SetBackendServersResponseBody) *SetBackendServersResponse
	GetBody() *SetBackendServersResponseBody
}

type SetBackendServersResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetBackendServersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetBackendServersResponse) String() string {
	return dara.Prettify(s)
}

func (s SetBackendServersResponse) GoString() string {
	return s.String()
}

func (s *SetBackendServersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetBackendServersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetBackendServersResponse) GetBody() *SetBackendServersResponseBody {
	return s.Body
}

func (s *SetBackendServersResponse) SetHeaders(v map[string]*string) *SetBackendServersResponse {
	s.Headers = v
	return s
}

func (s *SetBackendServersResponse) SetStatusCode(v int32) *SetBackendServersResponse {
	s.StatusCode = &v
	return s
}

func (s *SetBackendServersResponse) SetBody(v *SetBackendServersResponseBody) *SetBackendServersResponse {
	s.Body = v
	return s
}

func (s *SetBackendServersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

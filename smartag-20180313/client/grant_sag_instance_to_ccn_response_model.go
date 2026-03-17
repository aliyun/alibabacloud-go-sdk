// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantSagInstanceToCcnResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GrantSagInstanceToCcnResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GrantSagInstanceToCcnResponse
	GetStatusCode() *int32
	SetBody(v *GrantSagInstanceToCcnResponseBody) *GrantSagInstanceToCcnResponse
	GetBody() *GrantSagInstanceToCcnResponseBody
}

type GrantSagInstanceToCcnResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GrantSagInstanceToCcnResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GrantSagInstanceToCcnResponse) String() string {
	return dara.Prettify(s)
}

func (s GrantSagInstanceToCcnResponse) GoString() string {
	return s.String()
}

func (s *GrantSagInstanceToCcnResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GrantSagInstanceToCcnResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GrantSagInstanceToCcnResponse) GetBody() *GrantSagInstanceToCcnResponseBody {
	return s.Body
}

func (s *GrantSagInstanceToCcnResponse) SetHeaders(v map[string]*string) *GrantSagInstanceToCcnResponse {
	s.Headers = v
	return s
}

func (s *GrantSagInstanceToCcnResponse) SetStatusCode(v int32) *GrantSagInstanceToCcnResponse {
	s.StatusCode = &v
	return s
}

func (s *GrantSagInstanceToCcnResponse) SetBody(v *GrantSagInstanceToCcnResponseBody) *GrantSagInstanceToCcnResponse {
	s.Body = v
	return s
}

func (s *GrantSagInstanceToCcnResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

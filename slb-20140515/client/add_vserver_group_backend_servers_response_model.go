// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddVServerGroupBackendServersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddVServerGroupBackendServersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddVServerGroupBackendServersResponse
	GetStatusCode() *int32
	SetBody(v *AddVServerGroupBackendServersResponseBody) *AddVServerGroupBackendServersResponse
	GetBody() *AddVServerGroupBackendServersResponseBody
}

type AddVServerGroupBackendServersResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddVServerGroupBackendServersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddVServerGroupBackendServersResponse) String() string {
	return dara.Prettify(s)
}

func (s AddVServerGroupBackendServersResponse) GoString() string {
	return s.String()
}

func (s *AddVServerGroupBackendServersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddVServerGroupBackendServersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddVServerGroupBackendServersResponse) GetBody() *AddVServerGroupBackendServersResponseBody {
	return s.Body
}

func (s *AddVServerGroupBackendServersResponse) SetHeaders(v map[string]*string) *AddVServerGroupBackendServersResponse {
	s.Headers = v
	return s
}

func (s *AddVServerGroupBackendServersResponse) SetStatusCode(v int32) *AddVServerGroupBackendServersResponse {
	s.StatusCode = &v
	return s
}

func (s *AddVServerGroupBackendServersResponse) SetBody(v *AddVServerGroupBackendServersResponseBody) *AddVServerGroupBackendServersResponse {
	s.Body = v
	return s
}

func (s *AddVServerGroupBackendServersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

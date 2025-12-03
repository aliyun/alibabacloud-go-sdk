// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVServerGroupBackendServersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyVServerGroupBackendServersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyVServerGroupBackendServersResponse
	GetStatusCode() *int32
	SetBody(v *ModifyVServerGroupBackendServersResponseBody) *ModifyVServerGroupBackendServersResponse
	GetBody() *ModifyVServerGroupBackendServersResponseBody
}

type ModifyVServerGroupBackendServersResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyVServerGroupBackendServersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyVServerGroupBackendServersResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyVServerGroupBackendServersResponse) GoString() string {
	return s.String()
}

func (s *ModifyVServerGroupBackendServersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyVServerGroupBackendServersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyVServerGroupBackendServersResponse) GetBody() *ModifyVServerGroupBackendServersResponseBody {
	return s.Body
}

func (s *ModifyVServerGroupBackendServersResponse) SetHeaders(v map[string]*string) *ModifyVServerGroupBackendServersResponse {
	s.Headers = v
	return s
}

func (s *ModifyVServerGroupBackendServersResponse) SetStatusCode(v int32) *ModifyVServerGroupBackendServersResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyVServerGroupBackendServersResponse) SetBody(v *ModifyVServerGroupBackendServersResponseBody) *ModifyVServerGroupBackendServersResponse {
	s.Body = v
	return s
}

func (s *ModifyVServerGroupBackendServersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

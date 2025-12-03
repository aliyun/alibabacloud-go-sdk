// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveVServerGroupBackendServersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveVServerGroupBackendServersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveVServerGroupBackendServersResponse
	GetStatusCode() *int32
	SetBody(v *RemoveVServerGroupBackendServersResponseBody) *RemoveVServerGroupBackendServersResponse
	GetBody() *RemoveVServerGroupBackendServersResponseBody
}

type RemoveVServerGroupBackendServersResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveVServerGroupBackendServersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveVServerGroupBackendServersResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveVServerGroupBackendServersResponse) GoString() string {
	return s.String()
}

func (s *RemoveVServerGroupBackendServersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveVServerGroupBackendServersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveVServerGroupBackendServersResponse) GetBody() *RemoveVServerGroupBackendServersResponseBody {
	return s.Body
}

func (s *RemoveVServerGroupBackendServersResponse) SetHeaders(v map[string]*string) *RemoveVServerGroupBackendServersResponse {
	s.Headers = v
	return s
}

func (s *RemoveVServerGroupBackendServersResponse) SetStatusCode(v int32) *RemoveVServerGroupBackendServersResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveVServerGroupBackendServersResponse) SetBody(v *RemoveVServerGroupBackendServersResponseBody) *RemoveVServerGroupBackendServersResponse {
	s.Body = v
	return s
}

func (s *RemoveVServerGroupBackendServersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

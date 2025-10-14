// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveBackendServersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveBackendServersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveBackendServersResponse
	GetStatusCode() *int32
	SetBody(v *RemoveBackendServersResponseBody) *RemoveBackendServersResponse
	GetBody() *RemoveBackendServersResponseBody
}

type RemoveBackendServersResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveBackendServersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveBackendServersResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveBackendServersResponse) GoString() string {
	return s.String()
}

func (s *RemoveBackendServersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveBackendServersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveBackendServersResponse) GetBody() *RemoveBackendServersResponseBody {
	return s.Body
}

func (s *RemoveBackendServersResponse) SetHeaders(v map[string]*string) *RemoveBackendServersResponse {
	s.Headers = v
	return s
}

func (s *RemoveBackendServersResponse) SetStatusCode(v int32) *RemoveBackendServersResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveBackendServersResponse) SetBody(v *RemoveBackendServersResponseBody) *RemoveBackendServersResponse {
	s.Body = v
	return s
}

func (s *RemoveBackendServersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

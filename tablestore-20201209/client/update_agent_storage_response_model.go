// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentStorageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAgentStorageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAgentStorageResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAgentStorageResponseBody) *UpdateAgentStorageResponse
	GetBody() *UpdateAgentStorageResponseBody
}

type UpdateAgentStorageResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAgentStorageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAgentStorageResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentStorageResponse) GoString() string {
	return s.String()
}

func (s *UpdateAgentStorageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAgentStorageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAgentStorageResponse) GetBody() *UpdateAgentStorageResponseBody {
	return s.Body
}

func (s *UpdateAgentStorageResponse) SetHeaders(v map[string]*string) *UpdateAgentStorageResponse {
	s.Headers = v
	return s
}

func (s *UpdateAgentStorageResponse) SetStatusCode(v int32) *UpdateAgentStorageResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAgentStorageResponse) SetBody(v *UpdateAgentStorageResponseBody) *UpdateAgentStorageResponse {
	s.Body = v
	return s
}

func (s *UpdateAgentStorageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentStorageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAgentStorageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAgentStorageResponse
	GetStatusCode() *int32
	SetBody(v *CreateAgentStorageResponseBody) *CreateAgentStorageResponse
	GetBody() *CreateAgentStorageResponseBody
}

type CreateAgentStorageResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAgentStorageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAgentStorageResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentStorageResponse) GoString() string {
	return s.String()
}

func (s *CreateAgentStorageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAgentStorageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAgentStorageResponse) GetBody() *CreateAgentStorageResponseBody {
	return s.Body
}

func (s *CreateAgentStorageResponse) SetHeaders(v map[string]*string) *CreateAgentStorageResponse {
	s.Headers = v
	return s
}

func (s *CreateAgentStorageResponse) SetStatusCode(v int32) *CreateAgentStorageResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAgentStorageResponse) SetBody(v *CreateAgentStorageResponseBody) *CreateAgentStorageResponse {
	s.Body = v
	return s
}

func (s *CreateAgentStorageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

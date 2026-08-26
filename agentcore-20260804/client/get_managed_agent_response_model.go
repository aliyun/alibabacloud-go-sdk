// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetManagedAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetManagedAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetManagedAgentResponse
	GetStatusCode() *int32
	SetBody(v *GetManagedAgentResponseBody) *GetManagedAgentResponse
	GetBody() *GetManagedAgentResponseBody
}

type GetManagedAgentResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetManagedAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetManagedAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s GetManagedAgentResponse) GoString() string {
	return s.String()
}

func (s *GetManagedAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetManagedAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetManagedAgentResponse) GetBody() *GetManagedAgentResponseBody {
	return s.Body
}

func (s *GetManagedAgentResponse) SetHeaders(v map[string]*string) *GetManagedAgentResponse {
	s.Headers = v
	return s
}

func (s *GetManagedAgentResponse) SetStatusCode(v int32) *GetManagedAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetManagedAgentResponse) SetBody(v *GetManagedAgentResponseBody) *GetManagedAgentResponse {
	s.Body = v
	return s
}

func (s *GetManagedAgentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

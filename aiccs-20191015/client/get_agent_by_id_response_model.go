// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentByIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAgentByIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAgentByIdResponse
	GetStatusCode() *int32
	SetBody(v *GetAgentByIdResponseBody) *GetAgentByIdResponse
	GetBody() *GetAgentByIdResponseBody
}

type GetAgentByIdResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAgentByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgentByIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgentByIdResponse) GoString() string {
	return s.String()
}

func (s *GetAgentByIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAgentByIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAgentByIdResponse) GetBody() *GetAgentByIdResponseBody {
	return s.Body
}

func (s *GetAgentByIdResponse) SetHeaders(v map[string]*string) *GetAgentByIdResponse {
	s.Headers = v
	return s
}

func (s *GetAgentByIdResponse) SetStatusCode(v int32) *GetAgentByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgentByIdResponse) SetBody(v *GetAgentByIdResponseBody) *GetAgentByIdResponse {
	s.Body = v
	return s
}

func (s *GetAgentByIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

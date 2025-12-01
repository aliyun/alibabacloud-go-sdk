// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDBListFromAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDBListFromAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDBListFromAgentResponse
	GetStatusCode() *int32
	SetBody(v *GetDBListFromAgentResponseBody) *GetDBListFromAgentResponse
	GetBody() *GetDBListFromAgentResponseBody
}

type GetDBListFromAgentResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDBListFromAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDBListFromAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDBListFromAgentResponse) GoString() string {
	return s.String()
}

func (s *GetDBListFromAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDBListFromAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDBListFromAgentResponse) GetBody() *GetDBListFromAgentResponseBody {
	return s.Body
}

func (s *GetDBListFromAgentResponse) SetHeaders(v map[string]*string) *GetDBListFromAgentResponse {
	s.Headers = v
	return s
}

func (s *GetDBListFromAgentResponse) SetStatusCode(v int32) *GetDBListFromAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDBListFromAgentResponse) SetBody(v *GetDBListFromAgentResponseBody) *GetDBListFromAgentResponse {
	s.Body = v
	return s
}

func (s *GetDBListFromAgentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

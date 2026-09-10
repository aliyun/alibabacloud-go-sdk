// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLhmAgentStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLhmAgentStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLhmAgentStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetLhmAgentStatusResponseBody) *GetLhmAgentStatusResponse
	GetBody() *GetLhmAgentStatusResponseBody
}

type GetLhmAgentStatusResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLhmAgentStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLhmAgentStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLhmAgentStatusResponse) GoString() string {
	return s.String()
}

func (s *GetLhmAgentStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLhmAgentStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLhmAgentStatusResponse) GetBody() *GetLhmAgentStatusResponseBody {
	return s.Body
}

func (s *GetLhmAgentStatusResponse) SetHeaders(v map[string]*string) *GetLhmAgentStatusResponse {
	s.Headers = v
	return s
}

func (s *GetLhmAgentStatusResponse) SetStatusCode(v int32) *GetLhmAgentStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLhmAgentStatusResponse) SetBody(v *GetLhmAgentStatusResponseBody) *GetLhmAgentStatusResponse {
	s.Body = v
	return s
}

func (s *GetLhmAgentStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

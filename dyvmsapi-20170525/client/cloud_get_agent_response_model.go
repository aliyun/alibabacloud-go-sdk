// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudGetAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudGetAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudGetAgentResponse
	GetStatusCode() *int32
	SetBody(v *CloudGetAgentResponseBody) *CloudGetAgentResponse
	GetBody() *CloudGetAgentResponseBody
}

type CloudGetAgentResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudGetAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudGetAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudGetAgentResponse) GoString() string {
	return s.String()
}

func (s *CloudGetAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudGetAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudGetAgentResponse) GetBody() *CloudGetAgentResponseBody {
	return s.Body
}

func (s *CloudGetAgentResponse) SetHeaders(v map[string]*string) *CloudGetAgentResponse {
	s.Headers = v
	return s
}

func (s *CloudGetAgentResponse) SetStatusCode(v int32) *CloudGetAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudGetAgentResponse) SetBody(v *CloudGetAgentResponseBody) *CloudGetAgentResponse {
	s.Body = v
	return s
}

func (s *CloudGetAgentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

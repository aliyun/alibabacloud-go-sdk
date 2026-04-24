// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudCreateAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudCreateAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudCreateAgentResponse
	GetStatusCode() *int32
	SetBody(v *CloudCreateAgentResponseBody) *CloudCreateAgentResponse
	GetBody() *CloudCreateAgentResponseBody
}

type CloudCreateAgentResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudCreateAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudCreateAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudCreateAgentResponse) GoString() string {
	return s.String()
}

func (s *CloudCreateAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudCreateAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudCreateAgentResponse) GetBody() *CloudCreateAgentResponseBody {
	return s.Body
}

func (s *CloudCreateAgentResponse) SetHeaders(v map[string]*string) *CloudCreateAgentResponse {
	s.Headers = v
	return s
}

func (s *CloudCreateAgentResponse) SetStatusCode(v int32) *CloudCreateAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudCreateAgentResponse) SetBody(v *CloudCreateAgentResponseBody) *CloudCreateAgentResponse {
	s.Body = v
	return s
}

func (s *CloudCreateAgentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

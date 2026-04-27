// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudQueryAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudQueryAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudQueryAgentResponse
	GetStatusCode() *int32
	SetBody(v *CloudQueryAgentResponseBody) *CloudQueryAgentResponse
	GetBody() *CloudQueryAgentResponseBody
}

type CloudQueryAgentResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudQueryAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudQueryAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudQueryAgentResponse) GoString() string {
	return s.String()
}

func (s *CloudQueryAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudQueryAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudQueryAgentResponse) GetBody() *CloudQueryAgentResponseBody {
	return s.Body
}

func (s *CloudQueryAgentResponse) SetHeaders(v map[string]*string) *CloudQueryAgentResponse {
	s.Headers = v
	return s
}

func (s *CloudQueryAgentResponse) SetStatusCode(v int32) *CloudQueryAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudQueryAgentResponse) SetBody(v *CloudQueryAgentResponseBody) *CloudQueryAgentResponse {
	s.Body = v
	return s
}

func (s *CloudQueryAgentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

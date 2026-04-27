// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudUpdateAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudUpdateAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudUpdateAgentResponse
	GetStatusCode() *int32
	SetBody(v *CloudUpdateAgentResponseBody) *CloudUpdateAgentResponse
	GetBody() *CloudUpdateAgentResponseBody
}

type CloudUpdateAgentResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudUpdateAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudUpdateAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudUpdateAgentResponse) GoString() string {
	return s.String()
}

func (s *CloudUpdateAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudUpdateAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudUpdateAgentResponse) GetBody() *CloudUpdateAgentResponseBody {
	return s.Body
}

func (s *CloudUpdateAgentResponse) SetHeaders(v map[string]*string) *CloudUpdateAgentResponse {
	s.Headers = v
	return s
}

func (s *CloudUpdateAgentResponse) SetStatusCode(v int32) *CloudUpdateAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudUpdateAgentResponse) SetBody(v *CloudUpdateAgentResponseBody) *CloudUpdateAgentResponse {
	s.Body = v
	return s
}

func (s *CloudUpdateAgentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

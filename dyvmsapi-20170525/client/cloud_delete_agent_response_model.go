// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudDeleteAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudDeleteAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudDeleteAgentResponse
	GetStatusCode() *int32
	SetBody(v *CloudDeleteAgentResponseBody) *CloudDeleteAgentResponse
	GetBody() *CloudDeleteAgentResponseBody
}

type CloudDeleteAgentResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudDeleteAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudDeleteAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudDeleteAgentResponse) GoString() string {
	return s.String()
}

func (s *CloudDeleteAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudDeleteAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudDeleteAgentResponse) GetBody() *CloudDeleteAgentResponseBody {
	return s.Body
}

func (s *CloudDeleteAgentResponse) SetHeaders(v map[string]*string) *CloudDeleteAgentResponse {
	s.Headers = v
	return s
}

func (s *CloudDeleteAgentResponse) SetStatusCode(v int32) *CloudDeleteAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudDeleteAgentResponse) SetBody(v *CloudDeleteAgentResponseBody) *CloudDeleteAgentResponse {
	s.Body = v
	return s
}

func (s *CloudDeleteAgentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

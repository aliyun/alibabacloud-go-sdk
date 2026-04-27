// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudListOnlineAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudListOnlineAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudListOnlineAgentResponse
	GetStatusCode() *int32
	SetBody(v *CloudListOnlineAgentResponseBody) *CloudListOnlineAgentResponse
	GetBody() *CloudListOnlineAgentResponseBody
}

type CloudListOnlineAgentResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudListOnlineAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudListOnlineAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudListOnlineAgentResponse) GoString() string {
	return s.String()
}

func (s *CloudListOnlineAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudListOnlineAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudListOnlineAgentResponse) GetBody() *CloudListOnlineAgentResponseBody {
	return s.Body
}

func (s *CloudListOnlineAgentResponse) SetHeaders(v map[string]*string) *CloudListOnlineAgentResponse {
	s.Headers = v
	return s
}

func (s *CloudListOnlineAgentResponse) SetStatusCode(v int32) *CloudListOnlineAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudListOnlineAgentResponse) SetBody(v *CloudListOnlineAgentResponseBody) *CloudListOnlineAgentResponse {
	s.Body = v
	return s
}

func (s *CloudListOnlineAgentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

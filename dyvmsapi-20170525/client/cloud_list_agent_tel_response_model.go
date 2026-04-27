// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudListAgentTelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudListAgentTelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudListAgentTelResponse
	GetStatusCode() *int32
	SetBody(v *CloudListAgentTelResponseBody) *CloudListAgentTelResponse
	GetBody() *CloudListAgentTelResponseBody
}

type CloudListAgentTelResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudListAgentTelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudListAgentTelResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudListAgentTelResponse) GoString() string {
	return s.String()
}

func (s *CloudListAgentTelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudListAgentTelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudListAgentTelResponse) GetBody() *CloudListAgentTelResponseBody {
	return s.Body
}

func (s *CloudListAgentTelResponse) SetHeaders(v map[string]*string) *CloudListAgentTelResponse {
	s.Headers = v
	return s
}

func (s *CloudListAgentTelResponse) SetStatusCode(v int32) *CloudListAgentTelResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudListAgentTelResponse) SetBody(v *CloudListAgentTelResponseBody) *CloudListAgentTelResponse {
	s.Body = v
	return s
}

func (s *CloudListAgentTelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

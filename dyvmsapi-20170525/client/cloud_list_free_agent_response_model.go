// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudListFreeAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudListFreeAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudListFreeAgentResponse
	GetStatusCode() *int32
	SetBody(v *CloudListFreeAgentResponseBody) *CloudListFreeAgentResponse
	GetBody() *CloudListFreeAgentResponseBody
}

type CloudListFreeAgentResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudListFreeAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudListFreeAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudListFreeAgentResponse) GoString() string {
	return s.String()
}

func (s *CloudListFreeAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudListFreeAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudListFreeAgentResponse) GetBody() *CloudListFreeAgentResponseBody {
	return s.Body
}

func (s *CloudListFreeAgentResponse) SetHeaders(v map[string]*string) *CloudListFreeAgentResponse {
	s.Headers = v
	return s
}

func (s *CloudListFreeAgentResponse) SetStatusCode(v int32) *CloudListFreeAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudListFreeAgentResponse) SetBody(v *CloudListFreeAgentResponseBody) *CloudListFreeAgentResponse {
	s.Body = v
	return s
}

func (s *CloudListFreeAgentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

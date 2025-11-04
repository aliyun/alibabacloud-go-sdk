// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPublishedAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPublishedAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPublishedAgentResponse
	GetStatusCode() *int32
	SetBody(v *GetPublishedAgentResponseBody) *GetPublishedAgentResponse
	GetBody() *GetPublishedAgentResponseBody
}

type GetPublishedAgentResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPublishedAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPublishedAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPublishedAgentResponse) GoString() string {
	return s.String()
}

func (s *GetPublishedAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPublishedAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPublishedAgentResponse) GetBody() *GetPublishedAgentResponseBody {
	return s.Body
}

func (s *GetPublishedAgentResponse) SetHeaders(v map[string]*string) *GetPublishedAgentResponse {
	s.Headers = v
	return s
}

func (s *GetPublishedAgentResponse) SetStatusCode(v int32) *GetPublishedAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPublishedAgentResponse) SetBody(v *GetPublishedAgentResponseBody) *GetPublishedAgentResponse {
	s.Body = v
	return s
}

func (s *GetPublishedAgentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

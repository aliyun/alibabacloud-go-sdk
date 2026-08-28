// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAgentSpecVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitAgentSpecVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitAgentSpecVersionResponse
	GetStatusCode() *int32
	SetBody(v *SubmitAgentSpecVersionResponseBody) *SubmitAgentSpecVersionResponse
	GetBody() *SubmitAgentSpecVersionResponseBody
}

type SubmitAgentSpecVersionResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitAgentSpecVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitAgentSpecVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitAgentSpecVersionResponse) GoString() string {
	return s.String()
}

func (s *SubmitAgentSpecVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitAgentSpecVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitAgentSpecVersionResponse) GetBody() *SubmitAgentSpecVersionResponseBody {
	return s.Body
}

func (s *SubmitAgentSpecVersionResponse) SetHeaders(v map[string]*string) *SubmitAgentSpecVersionResponse {
	s.Headers = v
	return s
}

func (s *SubmitAgentSpecVersionResponse) SetStatusCode(v int32) *SubmitAgentSpecVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitAgentSpecVersionResponse) SetBody(v *SubmitAgentSpecVersionResponseBody) *SubmitAgentSpecVersionResponse {
	s.Body = v
	return s
}

func (s *SubmitAgentSpecVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

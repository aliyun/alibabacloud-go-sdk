// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationAgentRelationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateApplicationAgentRelationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateApplicationAgentRelationResponse
	GetStatusCode() *int32
	SetBody(v *CreateApplicationAgentRelationResponseBody) *CreateApplicationAgentRelationResponse
	GetBody() *CreateApplicationAgentRelationResponseBody
}

type CreateApplicationAgentRelationResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateApplicationAgentRelationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateApplicationAgentRelationResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationAgentRelationResponse) GoString() string {
	return s.String()
}

func (s *CreateApplicationAgentRelationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateApplicationAgentRelationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateApplicationAgentRelationResponse) GetBody() *CreateApplicationAgentRelationResponseBody {
	return s.Body
}

func (s *CreateApplicationAgentRelationResponse) SetHeaders(v map[string]*string) *CreateApplicationAgentRelationResponse {
	s.Headers = v
	return s
}

func (s *CreateApplicationAgentRelationResponse) SetStatusCode(v int32) *CreateApplicationAgentRelationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateApplicationAgentRelationResponse) SetBody(v *CreateApplicationAgentRelationResponseBody) *CreateApplicationAgentRelationResponse {
	s.Body = v
	return s
}

func (s *CreateApplicationAgentRelationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

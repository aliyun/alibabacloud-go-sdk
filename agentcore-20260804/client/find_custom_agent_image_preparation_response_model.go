// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFindCustomAgentImagePreparationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FindCustomAgentImagePreparationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FindCustomAgentImagePreparationResponse
	GetStatusCode() *int32
	SetBody(v *FindCustomAgentImagePreparationResponseBody) *FindCustomAgentImagePreparationResponse
	GetBody() *FindCustomAgentImagePreparationResponseBody
}

type FindCustomAgentImagePreparationResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FindCustomAgentImagePreparationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FindCustomAgentImagePreparationResponse) String() string {
	return dara.Prettify(s)
}

func (s FindCustomAgentImagePreparationResponse) GoString() string {
	return s.String()
}

func (s *FindCustomAgentImagePreparationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FindCustomAgentImagePreparationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FindCustomAgentImagePreparationResponse) GetBody() *FindCustomAgentImagePreparationResponseBody {
	return s.Body
}

func (s *FindCustomAgentImagePreparationResponse) SetHeaders(v map[string]*string) *FindCustomAgentImagePreparationResponse {
	s.Headers = v
	return s
}

func (s *FindCustomAgentImagePreparationResponse) SetStatusCode(v int32) *FindCustomAgentImagePreparationResponse {
	s.StatusCode = &v
	return s
}

func (s *FindCustomAgentImagePreparationResponse) SetBody(v *FindCustomAgentImagePreparationResponseBody) *FindCustomAgentImagePreparationResponse {
	s.Body = v
	return s
}

func (s *FindCustomAgentImagePreparationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

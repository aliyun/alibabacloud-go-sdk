// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomAgentImagePreparationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCustomAgentImagePreparationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCustomAgentImagePreparationResponse
	GetStatusCode() *int32
	SetBody(v *GetCustomAgentImagePreparationResponseBody) *GetCustomAgentImagePreparationResponse
	GetBody() *GetCustomAgentImagePreparationResponseBody
}

type GetCustomAgentImagePreparationResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCustomAgentImagePreparationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCustomAgentImagePreparationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCustomAgentImagePreparationResponse) GoString() string {
	return s.String()
}

func (s *GetCustomAgentImagePreparationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCustomAgentImagePreparationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCustomAgentImagePreparationResponse) GetBody() *GetCustomAgentImagePreparationResponseBody {
	return s.Body
}

func (s *GetCustomAgentImagePreparationResponse) SetHeaders(v map[string]*string) *GetCustomAgentImagePreparationResponse {
	s.Headers = v
	return s
}

func (s *GetCustomAgentImagePreparationResponse) SetStatusCode(v int32) *GetCustomAgentImagePreparationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCustomAgentImagePreparationResponse) SetBody(v *GetCustomAgentImagePreparationResponseBody) *GetCustomAgentImagePreparationResponse {
	s.Body = v
	return s
}

func (s *GetCustomAgentImagePreparationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

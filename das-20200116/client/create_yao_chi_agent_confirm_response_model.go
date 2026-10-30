// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateYaoChiAgentConfirmResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateYaoChiAgentConfirmResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateYaoChiAgentConfirmResponse
	GetStatusCode() *int32
	SetBody(v *CreateYaoChiAgentConfirmResponseBody) *CreateYaoChiAgentConfirmResponse
	GetBody() *CreateYaoChiAgentConfirmResponseBody
}

type CreateYaoChiAgentConfirmResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateYaoChiAgentConfirmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateYaoChiAgentConfirmResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateYaoChiAgentConfirmResponse) GoString() string {
	return s.String()
}

func (s *CreateYaoChiAgentConfirmResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateYaoChiAgentConfirmResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateYaoChiAgentConfirmResponse) GetBody() *CreateYaoChiAgentConfirmResponseBody {
	return s.Body
}

func (s *CreateYaoChiAgentConfirmResponse) SetHeaders(v map[string]*string) *CreateYaoChiAgentConfirmResponse {
	s.Headers = v
	return s
}

func (s *CreateYaoChiAgentConfirmResponse) SetStatusCode(v int32) *CreateYaoChiAgentConfirmResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateYaoChiAgentConfirmResponse) SetBody(v *CreateYaoChiAgentConfirmResponseBody) *CreateYaoChiAgentConfirmResponse {
	s.Body = v
	return s
}

func (s *CreateYaoChiAgentConfirmResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

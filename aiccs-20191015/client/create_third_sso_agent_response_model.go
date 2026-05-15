// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateThirdSsoAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateThirdSsoAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateThirdSsoAgentResponse
	GetStatusCode() *int32
	SetBody(v *CreateThirdSsoAgentResponseBody) *CreateThirdSsoAgentResponse
	GetBody() *CreateThirdSsoAgentResponseBody
}

type CreateThirdSsoAgentResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateThirdSsoAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateThirdSsoAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateThirdSsoAgentResponse) GoString() string {
	return s.String()
}

func (s *CreateThirdSsoAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateThirdSsoAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateThirdSsoAgentResponse) GetBody() *CreateThirdSsoAgentResponseBody {
	return s.Body
}

func (s *CreateThirdSsoAgentResponse) SetHeaders(v map[string]*string) *CreateThirdSsoAgentResponse {
	s.Headers = v
	return s
}

func (s *CreateThirdSsoAgentResponse) SetStatusCode(v int32) *CreateThirdSsoAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateThirdSsoAgentResponse) SetBody(v *CreateThirdSsoAgentResponseBody) *CreateThirdSsoAgentResponse {
	s.Body = v
	return s
}

func (s *CreateThirdSsoAgentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

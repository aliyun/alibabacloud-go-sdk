// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRulesResponse
	GetStatusCode() *int32
	SetBody(v *CreateRulesResponseBody) *CreateRulesResponse
	GetBody() *CreateRulesResponseBody
}

type CreateRulesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRulesResponse) GoString() string {
	return s.String()
}

func (s *CreateRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRulesResponse) GetBody() *CreateRulesResponseBody {
	return s.Body
}

func (s *CreateRulesResponse) SetHeaders(v map[string]*string) *CreateRulesResponse {
	s.Headers = v
	return s
}

func (s *CreateRulesResponse) SetStatusCode(v int32) *CreateRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRulesResponse) SetBody(v *CreateRulesResponseBody) *CreateRulesResponse {
	s.Body = v
	return s
}

func (s *CreateRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

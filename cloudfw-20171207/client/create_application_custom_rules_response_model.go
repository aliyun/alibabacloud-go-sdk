// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationCustomRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateApplicationCustomRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateApplicationCustomRulesResponse
	GetStatusCode() *int32
	SetBody(v *CreateApplicationCustomRulesResponseBody) *CreateApplicationCustomRulesResponse
	GetBody() *CreateApplicationCustomRulesResponseBody
}

type CreateApplicationCustomRulesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateApplicationCustomRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateApplicationCustomRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationCustomRulesResponse) GoString() string {
	return s.String()
}

func (s *CreateApplicationCustomRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateApplicationCustomRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateApplicationCustomRulesResponse) GetBody() *CreateApplicationCustomRulesResponseBody {
	return s.Body
}

func (s *CreateApplicationCustomRulesResponse) SetHeaders(v map[string]*string) *CreateApplicationCustomRulesResponse {
	s.Headers = v
	return s
}

func (s *CreateApplicationCustomRulesResponse) SetStatusCode(v int32) *CreateApplicationCustomRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateApplicationCustomRulesResponse) SetBody(v *CreateApplicationCustomRulesResponseBody) *CreateApplicationCustomRulesResponse {
	s.Body = v
	return s
}

func (s *CreateApplicationCustomRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

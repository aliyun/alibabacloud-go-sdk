// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationCustomRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateApplicationCustomRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateApplicationCustomRulesResponse
	GetStatusCode() *int32
	SetBody(v *UpdateApplicationCustomRulesResponseBody) *UpdateApplicationCustomRulesResponse
	GetBody() *UpdateApplicationCustomRulesResponseBody
}

type UpdateApplicationCustomRulesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateApplicationCustomRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateApplicationCustomRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationCustomRulesResponse) GoString() string {
	return s.String()
}

func (s *UpdateApplicationCustomRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateApplicationCustomRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateApplicationCustomRulesResponse) GetBody() *UpdateApplicationCustomRulesResponseBody {
	return s.Body
}

func (s *UpdateApplicationCustomRulesResponse) SetHeaders(v map[string]*string) *UpdateApplicationCustomRulesResponse {
	s.Headers = v
	return s
}

func (s *UpdateApplicationCustomRulesResponse) SetStatusCode(v int32) *UpdateApplicationCustomRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateApplicationCustomRulesResponse) SetBody(v *UpdateApplicationCustomRulesResponseBody) *UpdateApplicationCustomRulesResponse {
	s.Body = v
	return s
}

func (s *UpdateApplicationCustomRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActiveConfigRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ActiveConfigRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ActiveConfigRulesResponse
	GetStatusCode() *int32
	SetBody(v *ActiveConfigRulesResponseBody) *ActiveConfigRulesResponse
	GetBody() *ActiveConfigRulesResponseBody
}

type ActiveConfigRulesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ActiveConfigRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ActiveConfigRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ActiveConfigRulesResponse) GoString() string {
	return s.String()
}

func (s *ActiveConfigRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ActiveConfigRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ActiveConfigRulesResponse) GetBody() *ActiveConfigRulesResponseBody {
	return s.Body
}

func (s *ActiveConfigRulesResponse) SetHeaders(v map[string]*string) *ActiveConfigRulesResponse {
	s.Headers = v
	return s
}

func (s *ActiveConfigRulesResponse) SetStatusCode(v int32) *ActiveConfigRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ActiveConfigRulesResponse) SetBody(v *ActiveConfigRulesResponseBody) *ActiveConfigRulesResponse {
	s.Body = v
	return s
}

func (s *ActiveConfigRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

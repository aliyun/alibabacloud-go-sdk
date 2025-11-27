// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMaskingRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMaskingRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMaskingRulesResponse
	GetStatusCode() *int32
	SetBody(v *CreateMaskingRulesResponseBody) *CreateMaskingRulesResponse
	GetBody() *CreateMaskingRulesResponseBody
}

type CreateMaskingRulesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMaskingRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMaskingRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMaskingRulesResponse) GoString() string {
	return s.String()
}

func (s *CreateMaskingRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMaskingRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMaskingRulesResponse) GetBody() *CreateMaskingRulesResponseBody {
	return s.Body
}

func (s *CreateMaskingRulesResponse) SetHeaders(v map[string]*string) *CreateMaskingRulesResponse {
	s.Headers = v
	return s
}

func (s *CreateMaskingRulesResponse) SetStatusCode(v int32) *CreateMaskingRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMaskingRulesResponse) SetBody(v *CreateMaskingRulesResponseBody) *CreateMaskingRulesResponse {
	s.Body = v
	return s
}

func (s *CreateMaskingRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

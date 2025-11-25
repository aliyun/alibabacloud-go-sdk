// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWebCCRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateWebCCRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateWebCCRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateWebCCRuleResponseBody) *CreateWebCCRuleResponse
	GetBody() *CreateWebCCRuleResponseBody
}

type CreateWebCCRuleResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWebCCRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWebCCRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateWebCCRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateWebCCRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateWebCCRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateWebCCRuleResponse) GetBody() *CreateWebCCRuleResponseBody {
	return s.Body
}

func (s *CreateWebCCRuleResponse) SetHeaders(v map[string]*string) *CreateWebCCRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateWebCCRuleResponse) SetStatusCode(v int32) *CreateWebCCRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWebCCRuleResponse) SetBody(v *CreateWebCCRuleResponseBody) *CreateWebCCRuleResponse {
	s.Body = v
	return s
}

func (s *CreateWebCCRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

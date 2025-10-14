// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHttpResponseHeaderModificationRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateHttpResponseHeaderModificationRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateHttpResponseHeaderModificationRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateHttpResponseHeaderModificationRuleResponseBody) *CreateHttpResponseHeaderModificationRuleResponse
	GetBody() *CreateHttpResponseHeaderModificationRuleResponseBody
}

type CreateHttpResponseHeaderModificationRuleResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHttpResponseHeaderModificationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateHttpResponseHeaderModificationRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateHttpResponseHeaderModificationRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateHttpResponseHeaderModificationRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateHttpResponseHeaderModificationRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateHttpResponseHeaderModificationRuleResponse) GetBody() *CreateHttpResponseHeaderModificationRuleResponseBody {
	return s.Body
}

func (s *CreateHttpResponseHeaderModificationRuleResponse) SetHeaders(v map[string]*string) *CreateHttpResponseHeaderModificationRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateHttpResponseHeaderModificationRuleResponse) SetStatusCode(v int32) *CreateHttpResponseHeaderModificationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHttpResponseHeaderModificationRuleResponse) SetBody(v *CreateHttpResponseHeaderModificationRuleResponseBody) *CreateHttpResponseHeaderModificationRuleResponse {
	s.Body = v
	return s
}

func (s *CreateHttpResponseHeaderModificationRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

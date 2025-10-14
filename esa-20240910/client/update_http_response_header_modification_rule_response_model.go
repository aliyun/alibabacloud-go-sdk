// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHttpResponseHeaderModificationRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateHttpResponseHeaderModificationRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateHttpResponseHeaderModificationRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateHttpResponseHeaderModificationRuleResponseBody) *UpdateHttpResponseHeaderModificationRuleResponse
	GetBody() *UpdateHttpResponseHeaderModificationRuleResponseBody
}

type UpdateHttpResponseHeaderModificationRuleResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateHttpResponseHeaderModificationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateHttpResponseHeaderModificationRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpResponseHeaderModificationRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateHttpResponseHeaderModificationRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateHttpResponseHeaderModificationRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateHttpResponseHeaderModificationRuleResponse) GetBody() *UpdateHttpResponseHeaderModificationRuleResponseBody {
	return s.Body
}

func (s *UpdateHttpResponseHeaderModificationRuleResponse) SetHeaders(v map[string]*string) *UpdateHttpResponseHeaderModificationRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateHttpResponseHeaderModificationRuleResponse) SetStatusCode(v int32) *UpdateHttpResponseHeaderModificationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateHttpResponseHeaderModificationRuleResponse) SetBody(v *UpdateHttpResponseHeaderModificationRuleResponseBody) *UpdateHttpResponseHeaderModificationRuleResponse {
	s.Body = v
	return s
}

func (s *UpdateHttpResponseHeaderModificationRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

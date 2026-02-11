// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPostCustomizeRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PostCustomizeRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PostCustomizeRuleResponse
	GetStatusCode() *int32
	SetBody(v *PostCustomizeRuleResponseBody) *PostCustomizeRuleResponse
	GetBody() *PostCustomizeRuleResponseBody
}

type PostCustomizeRuleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PostCustomizeRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PostCustomizeRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s PostCustomizeRuleResponse) GoString() string {
	return s.String()
}

func (s *PostCustomizeRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PostCustomizeRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PostCustomizeRuleResponse) GetBody() *PostCustomizeRuleResponseBody {
	return s.Body
}

func (s *PostCustomizeRuleResponse) SetHeaders(v map[string]*string) *PostCustomizeRuleResponse {
	s.Headers = v
	return s
}

func (s *PostCustomizeRuleResponse) SetStatusCode(v int32) *PostCustomizeRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *PostCustomizeRuleResponse) SetBody(v *PostCustomizeRuleResponseBody) *PostCustomizeRuleResponse {
	s.Body = v
	return s
}

func (s *PostCustomizeRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

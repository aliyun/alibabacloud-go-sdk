// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWhiteRuleListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateWhiteRuleListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateWhiteRuleListResponse
	GetStatusCode() *int32
	SetBody(v *UpdateWhiteRuleListResponseBody) *UpdateWhiteRuleListResponse
	GetBody() *UpdateWhiteRuleListResponseBody
}

type UpdateWhiteRuleListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWhiteRuleListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWhiteRuleListResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateWhiteRuleListResponse) GoString() string {
	return s.String()
}

func (s *UpdateWhiteRuleListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateWhiteRuleListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateWhiteRuleListResponse) GetBody() *UpdateWhiteRuleListResponseBody {
	return s.Body
}

func (s *UpdateWhiteRuleListResponse) SetHeaders(v map[string]*string) *UpdateWhiteRuleListResponse {
	s.Headers = v
	return s
}

func (s *UpdateWhiteRuleListResponse) SetStatusCode(v int32) *UpdateWhiteRuleListResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWhiteRuleListResponse) SetBody(v *UpdateWhiteRuleListResponseBody) *UpdateWhiteRuleListResponse {
	s.Body = v
	return s
}

func (s *UpdateWhiteRuleListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

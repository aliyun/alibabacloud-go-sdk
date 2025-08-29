// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourceRuleItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateResourceRuleItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateResourceRuleItemResponse
	GetStatusCode() *int32
	SetBody(v *UpdateResourceRuleItemResponseBody) *UpdateResourceRuleItemResponse
	GetBody() *UpdateResourceRuleItemResponseBody
}

type UpdateResourceRuleItemResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateResourceRuleItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateResourceRuleItemResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceRuleItemResponse) GoString() string {
	return s.String()
}

func (s *UpdateResourceRuleItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateResourceRuleItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateResourceRuleItemResponse) GetBody() *UpdateResourceRuleItemResponseBody {
	return s.Body
}

func (s *UpdateResourceRuleItemResponse) SetHeaders(v map[string]*string) *UpdateResourceRuleItemResponse {
	s.Headers = v
	return s
}

func (s *UpdateResourceRuleItemResponse) SetStatusCode(v int32) *UpdateResourceRuleItemResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateResourceRuleItemResponse) SetBody(v *UpdateResourceRuleItemResponseBody) *UpdateResourceRuleItemResponse {
	s.Body = v
	return s
}

func (s *UpdateResourceRuleItemResponse) Validate() error {
	return dara.Validate(s)
}

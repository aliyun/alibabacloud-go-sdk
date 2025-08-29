// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourceRuleItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateResourceRuleItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateResourceRuleItemResponse
	GetStatusCode() *int32
	SetBody(v *CreateResourceRuleItemResponseBody) *CreateResourceRuleItemResponse
	GetBody() *CreateResourceRuleItemResponseBody
}

type CreateResourceRuleItemResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateResourceRuleItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateResourceRuleItemResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceRuleItemResponse) GoString() string {
	return s.String()
}

func (s *CreateResourceRuleItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateResourceRuleItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateResourceRuleItemResponse) GetBody() *CreateResourceRuleItemResponseBody {
	return s.Body
}

func (s *CreateResourceRuleItemResponse) SetHeaders(v map[string]*string) *CreateResourceRuleItemResponse {
	s.Headers = v
	return s
}

func (s *CreateResourceRuleItemResponse) SetStatusCode(v int32) *CreateResourceRuleItemResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateResourceRuleItemResponse) SetBody(v *CreateResourceRuleItemResponseBody) *CreateResourceRuleItemResponse {
	s.Body = v
	return s
}

func (s *CreateResourceRuleItemResponse) Validate() error {
	return dara.Validate(s)
}

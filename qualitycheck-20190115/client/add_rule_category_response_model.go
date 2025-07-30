// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRuleCategoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddRuleCategoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddRuleCategoryResponse
	GetStatusCode() *int32
	SetBody(v *AddRuleCategoryResponseBody) *AddRuleCategoryResponse
	GetBody() *AddRuleCategoryResponseBody
}

type AddRuleCategoryResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddRuleCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddRuleCategoryResponse) String() string {
	return dara.Prettify(s)
}

func (s AddRuleCategoryResponse) GoString() string {
	return s.String()
}

func (s *AddRuleCategoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddRuleCategoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddRuleCategoryResponse) GetBody() *AddRuleCategoryResponseBody {
	return s.Body
}

func (s *AddRuleCategoryResponse) SetHeaders(v map[string]*string) *AddRuleCategoryResponse {
	s.Headers = v
	return s
}

func (s *AddRuleCategoryResponse) SetStatusCode(v int32) *AddRuleCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *AddRuleCategoryResponse) SetBody(v *AddRuleCategoryResponseBody) *AddRuleCategoryResponse {
	s.Body = v
	return s
}

func (s *AddRuleCategoryResponse) Validate() error {
	return dara.Validate(s)
}

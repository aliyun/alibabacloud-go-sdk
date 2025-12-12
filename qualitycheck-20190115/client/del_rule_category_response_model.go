// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDelRuleCategoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DelRuleCategoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DelRuleCategoryResponse
	GetStatusCode() *int32
	SetBody(v *DelRuleCategoryResponseBody) *DelRuleCategoryResponse
	GetBody() *DelRuleCategoryResponseBody
}

type DelRuleCategoryResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DelRuleCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DelRuleCategoryResponse) String() string {
	return dara.Prettify(s)
}

func (s DelRuleCategoryResponse) GoString() string {
	return s.String()
}

func (s *DelRuleCategoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DelRuleCategoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DelRuleCategoryResponse) GetBody() *DelRuleCategoryResponseBody {
	return s.Body
}

func (s *DelRuleCategoryResponse) SetHeaders(v map[string]*string) *DelRuleCategoryResponse {
	s.Headers = v
	return s
}

func (s *DelRuleCategoryResponse) SetStatusCode(v int32) *DelRuleCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DelRuleCategoryResponse) SetBody(v *DelRuleCategoryResponseBody) *DelRuleCategoryResponse {
	s.Body = v
	return s
}

func (s *DelRuleCategoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

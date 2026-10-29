// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSdnsRuleCategoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSdnsRuleCategoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSdnsRuleCategoryResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSdnsRuleCategoryResponseBody) *UpdateSdnsRuleCategoryResponse
	GetBody() *UpdateSdnsRuleCategoryResponseBody
}

type UpdateSdnsRuleCategoryResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSdnsRuleCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSdnsRuleCategoryResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSdnsRuleCategoryResponse) GoString() string {
	return s.String()
}

func (s *UpdateSdnsRuleCategoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSdnsRuleCategoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSdnsRuleCategoryResponse) GetBody() *UpdateSdnsRuleCategoryResponseBody {
	return s.Body
}

func (s *UpdateSdnsRuleCategoryResponse) SetHeaders(v map[string]*string) *UpdateSdnsRuleCategoryResponse {
	s.Headers = v
	return s
}

func (s *UpdateSdnsRuleCategoryResponse) SetStatusCode(v int32) *UpdateSdnsRuleCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSdnsRuleCategoryResponse) SetBody(v *UpdateSdnsRuleCategoryResponseBody) *UpdateSdnsRuleCategoryResponse {
	s.Body = v
	return s
}

func (s *UpdateSdnsRuleCategoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

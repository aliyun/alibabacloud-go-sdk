// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSdnsRuleCategoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSdnsRuleCategoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSdnsRuleCategoryResponse
	GetStatusCode() *int32
	SetBody(v *CreateSdnsRuleCategoryResponseBody) *CreateSdnsRuleCategoryResponse
	GetBody() *CreateSdnsRuleCategoryResponseBody
}

type CreateSdnsRuleCategoryResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSdnsRuleCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSdnsRuleCategoryResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSdnsRuleCategoryResponse) GoString() string {
	return s.String()
}

func (s *CreateSdnsRuleCategoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSdnsRuleCategoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSdnsRuleCategoryResponse) GetBody() *CreateSdnsRuleCategoryResponseBody {
	return s.Body
}

func (s *CreateSdnsRuleCategoryResponse) SetHeaders(v map[string]*string) *CreateSdnsRuleCategoryResponse {
	s.Headers = v
	return s
}

func (s *CreateSdnsRuleCategoryResponse) SetStatusCode(v int32) *CreateSdnsRuleCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSdnsRuleCategoryResponse) SetBody(v *CreateSdnsRuleCategoryResponseBody) *CreateSdnsRuleCategoryResponse {
	s.Body = v
	return s
}

func (s *CreateSdnsRuleCategoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

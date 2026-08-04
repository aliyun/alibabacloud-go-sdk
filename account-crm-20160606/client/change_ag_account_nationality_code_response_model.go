// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeAgAccountNationalityCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChangeAgAccountNationalityCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChangeAgAccountNationalityCodeResponse
	GetStatusCode() *int32
	SetBody(v *ChangeAgAccountNationalityCodeResponseBody) *ChangeAgAccountNationalityCodeResponse
	GetBody() *ChangeAgAccountNationalityCodeResponseBody
}

type ChangeAgAccountNationalityCodeResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeAgAccountNationalityCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeAgAccountNationalityCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s ChangeAgAccountNationalityCodeResponse) GoString() string {
	return s.String()
}

func (s *ChangeAgAccountNationalityCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChangeAgAccountNationalityCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChangeAgAccountNationalityCodeResponse) GetBody() *ChangeAgAccountNationalityCodeResponseBody {
	return s.Body
}

func (s *ChangeAgAccountNationalityCodeResponse) SetHeaders(v map[string]*string) *ChangeAgAccountNationalityCodeResponse {
	s.Headers = v
	return s
}

func (s *ChangeAgAccountNationalityCodeResponse) SetStatusCode(v int32) *ChangeAgAccountNationalityCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeAgAccountNationalityCodeResponse) SetBody(v *ChangeAgAccountNationalityCodeResponseBody) *ChangeAgAccountNationalityCodeResponse {
	s.Body = v
	return s
}

func (s *ChangeAgAccountNationalityCodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

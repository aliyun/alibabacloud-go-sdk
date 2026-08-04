// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeAgSecurityMobileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChangeAgSecurityMobileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChangeAgSecurityMobileResponse
	GetStatusCode() *int32
	SetBody(v *ChangeAgSecurityMobileResponseBody) *ChangeAgSecurityMobileResponse
	GetBody() *ChangeAgSecurityMobileResponseBody
}

type ChangeAgSecurityMobileResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeAgSecurityMobileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeAgSecurityMobileResponse) String() string {
	return dara.Prettify(s)
}

func (s ChangeAgSecurityMobileResponse) GoString() string {
	return s.String()
}

func (s *ChangeAgSecurityMobileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChangeAgSecurityMobileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChangeAgSecurityMobileResponse) GetBody() *ChangeAgSecurityMobileResponseBody {
	return s.Body
}

func (s *ChangeAgSecurityMobileResponse) SetHeaders(v map[string]*string) *ChangeAgSecurityMobileResponse {
	s.Headers = v
	return s
}

func (s *ChangeAgSecurityMobileResponse) SetStatusCode(v int32) *ChangeAgSecurityMobileResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeAgSecurityMobileResponse) SetBody(v *ChangeAgSecurityMobileResponseBody) *ChangeAgSecurityMobileResponse {
	s.Body = v
	return s
}

func (s *ChangeAgSecurityMobileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

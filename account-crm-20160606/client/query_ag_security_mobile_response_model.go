// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAgSecurityMobileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryAgSecurityMobileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryAgSecurityMobileResponse
	GetStatusCode() *int32
	SetBody(v *QueryAgSecurityMobileResponseBody) *QueryAgSecurityMobileResponse
	GetBody() *QueryAgSecurityMobileResponseBody
}

type QueryAgSecurityMobileResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAgSecurityMobileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAgSecurityMobileResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryAgSecurityMobileResponse) GoString() string {
	return s.String()
}

func (s *QueryAgSecurityMobileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryAgSecurityMobileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryAgSecurityMobileResponse) GetBody() *QueryAgSecurityMobileResponseBody {
	return s.Body
}

func (s *QueryAgSecurityMobileResponse) SetHeaders(v map[string]*string) *QueryAgSecurityMobileResponse {
	s.Headers = v
	return s
}

func (s *QueryAgSecurityMobileResponse) SetStatusCode(v int32) *QueryAgSecurityMobileResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAgSecurityMobileResponse) SetBody(v *QueryAgSecurityMobileResponseBody) *QueryAgSecurityMobileResponse {
	s.Body = v
	return s
}

func (s *QueryAgSecurityMobileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

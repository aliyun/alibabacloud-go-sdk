// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeAgSecurityEmailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChangeAgSecurityEmailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChangeAgSecurityEmailResponse
	GetStatusCode() *int32
	SetBody(v *ChangeAgSecurityEmailResponseBody) *ChangeAgSecurityEmailResponse
	GetBody() *ChangeAgSecurityEmailResponseBody
}

type ChangeAgSecurityEmailResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeAgSecurityEmailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeAgSecurityEmailResponse) String() string {
	return dara.Prettify(s)
}

func (s ChangeAgSecurityEmailResponse) GoString() string {
	return s.String()
}

func (s *ChangeAgSecurityEmailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChangeAgSecurityEmailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChangeAgSecurityEmailResponse) GetBody() *ChangeAgSecurityEmailResponseBody {
	return s.Body
}

func (s *ChangeAgSecurityEmailResponse) SetHeaders(v map[string]*string) *ChangeAgSecurityEmailResponse {
	s.Headers = v
	return s
}

func (s *ChangeAgSecurityEmailResponse) SetStatusCode(v int32) *ChangeAgSecurityEmailResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeAgSecurityEmailResponse) SetBody(v *ChangeAgSecurityEmailResponseBody) *ChangeAgSecurityEmailResponse {
	s.Body = v
	return s
}

func (s *ChangeAgSecurityEmailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

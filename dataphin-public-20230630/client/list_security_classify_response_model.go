// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSecurityClassifyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSecurityClassifyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSecurityClassifyResponse
	GetStatusCode() *int32
	SetBody(v *ListSecurityClassifyResponseBody) *ListSecurityClassifyResponse
	GetBody() *ListSecurityClassifyResponseBody
}

type ListSecurityClassifyResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSecurityClassifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSecurityClassifyResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSecurityClassifyResponse) GoString() string {
	return s.String()
}

func (s *ListSecurityClassifyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSecurityClassifyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSecurityClassifyResponse) GetBody() *ListSecurityClassifyResponseBody {
	return s.Body
}

func (s *ListSecurityClassifyResponse) SetHeaders(v map[string]*string) *ListSecurityClassifyResponse {
	s.Headers = v
	return s
}

func (s *ListSecurityClassifyResponse) SetStatusCode(v int32) *ListSecurityClassifyResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSecurityClassifyResponse) SetBody(v *ListSecurityClassifyResponseBody) *ListSecurityClassifyResponse {
	s.Body = v
	return s
}

func (s *ListSecurityClassifyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

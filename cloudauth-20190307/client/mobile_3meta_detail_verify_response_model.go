// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMobile3MetaDetailVerifyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *Mobile3MetaDetailVerifyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *Mobile3MetaDetailVerifyResponse
	GetStatusCode() *int32
	SetBody(v *Mobile3MetaDetailVerifyResponseBody) *Mobile3MetaDetailVerifyResponse
	GetBody() *Mobile3MetaDetailVerifyResponseBody
}

type Mobile3MetaDetailVerifyResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Mobile3MetaDetailVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s Mobile3MetaDetailVerifyResponse) String() string {
	return dara.Prettify(s)
}

func (s Mobile3MetaDetailVerifyResponse) GoString() string {
	return s.String()
}

func (s *Mobile3MetaDetailVerifyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *Mobile3MetaDetailVerifyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *Mobile3MetaDetailVerifyResponse) GetBody() *Mobile3MetaDetailVerifyResponseBody {
	return s.Body
}

func (s *Mobile3MetaDetailVerifyResponse) SetHeaders(v map[string]*string) *Mobile3MetaDetailVerifyResponse {
	s.Headers = v
	return s
}

func (s *Mobile3MetaDetailVerifyResponse) SetStatusCode(v int32) *Mobile3MetaDetailVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *Mobile3MetaDetailVerifyResponse) SetBody(v *Mobile3MetaDetailVerifyResponseBody) *Mobile3MetaDetailVerifyResponse {
	s.Body = v
	return s
}

func (s *Mobile3MetaDetailVerifyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

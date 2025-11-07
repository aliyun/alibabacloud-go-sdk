// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMobile3MetaDetailStandardVerifyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *Mobile3MetaDetailStandardVerifyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *Mobile3MetaDetailStandardVerifyResponse
	GetStatusCode() *int32
	SetBody(v *Mobile3MetaDetailStandardVerifyResponseBody) *Mobile3MetaDetailStandardVerifyResponse
	GetBody() *Mobile3MetaDetailStandardVerifyResponseBody
}

type Mobile3MetaDetailStandardVerifyResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Mobile3MetaDetailStandardVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s Mobile3MetaDetailStandardVerifyResponse) String() string {
	return dara.Prettify(s)
}

func (s Mobile3MetaDetailStandardVerifyResponse) GoString() string {
	return s.String()
}

func (s *Mobile3MetaDetailStandardVerifyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *Mobile3MetaDetailStandardVerifyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *Mobile3MetaDetailStandardVerifyResponse) GetBody() *Mobile3MetaDetailStandardVerifyResponseBody {
	return s.Body
}

func (s *Mobile3MetaDetailStandardVerifyResponse) SetHeaders(v map[string]*string) *Mobile3MetaDetailStandardVerifyResponse {
	s.Headers = v
	return s
}

func (s *Mobile3MetaDetailStandardVerifyResponse) SetStatusCode(v int32) *Mobile3MetaDetailStandardVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *Mobile3MetaDetailStandardVerifyResponse) SetBody(v *Mobile3MetaDetailStandardVerifyResponseBody) *Mobile3MetaDetailStandardVerifyResponse {
	s.Body = v
	return s
}

func (s *Mobile3MetaDetailStandardVerifyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

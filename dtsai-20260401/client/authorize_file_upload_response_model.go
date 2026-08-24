// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeFileUploadResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AuthorizeFileUploadResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AuthorizeFileUploadResponse
	GetStatusCode() *int32
	SetBody(v *AuthorizeFileUploadResponseBody) *AuthorizeFileUploadResponse
	GetBody() *AuthorizeFileUploadResponseBody
}

type AuthorizeFileUploadResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AuthorizeFileUploadResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AuthorizeFileUploadResponse) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeFileUploadResponse) GoString() string {
	return s.String()
}

func (s *AuthorizeFileUploadResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AuthorizeFileUploadResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AuthorizeFileUploadResponse) GetBody() *AuthorizeFileUploadResponseBody {
	return s.Body
}

func (s *AuthorizeFileUploadResponse) SetHeaders(v map[string]*string) *AuthorizeFileUploadResponse {
	s.Headers = v
	return s
}

func (s *AuthorizeFileUploadResponse) SetStatusCode(v int32) *AuthorizeFileUploadResponse {
	s.StatusCode = &v
	return s
}

func (s *AuthorizeFileUploadResponse) SetBody(v *AuthorizeFileUploadResponseBody) *AuthorizeFileUploadResponse {
	s.Body = v
	return s
}

func (s *AuthorizeFileUploadResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

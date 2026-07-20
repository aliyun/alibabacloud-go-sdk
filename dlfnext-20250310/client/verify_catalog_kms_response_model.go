// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyCatalogKmsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VerifyCatalogKmsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VerifyCatalogKmsResponse
	GetStatusCode() *int32
	SetBody(v *VerifyCatalogKmsResponseBody) *VerifyCatalogKmsResponse
	GetBody() *VerifyCatalogKmsResponseBody
}

type VerifyCatalogKmsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyCatalogKmsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyCatalogKmsResponse) String() string {
	return dara.Prettify(s)
}

func (s VerifyCatalogKmsResponse) GoString() string {
	return s.String()
}

func (s *VerifyCatalogKmsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VerifyCatalogKmsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VerifyCatalogKmsResponse) GetBody() *VerifyCatalogKmsResponseBody {
	return s.Body
}

func (s *VerifyCatalogKmsResponse) SetHeaders(v map[string]*string) *VerifyCatalogKmsResponse {
	s.Headers = v
	return s
}

func (s *VerifyCatalogKmsResponse) SetStatusCode(v int32) *VerifyCatalogKmsResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyCatalogKmsResponse) SetBody(v *VerifyCatalogKmsResponseBody) *VerifyCatalogKmsResponse {
	s.Body = v
	return s
}

func (s *VerifyCatalogKmsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

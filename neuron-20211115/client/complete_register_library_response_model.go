// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompleteRegisterLibraryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CompleteRegisterLibraryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CompleteRegisterLibraryResponse
	GetStatusCode() *int32
	SetBody(v *CatalogCommonResult) *CompleteRegisterLibraryResponse
	GetBody() *CatalogCommonResult
}

type CompleteRegisterLibraryResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CatalogCommonResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CompleteRegisterLibraryResponse) String() string {
	return dara.Prettify(s)
}

func (s CompleteRegisterLibraryResponse) GoString() string {
	return s.String()
}

func (s *CompleteRegisterLibraryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CompleteRegisterLibraryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CompleteRegisterLibraryResponse) GetBody() *CatalogCommonResult {
	return s.Body
}

func (s *CompleteRegisterLibraryResponse) SetHeaders(v map[string]*string) *CompleteRegisterLibraryResponse {
	s.Headers = v
	return s
}

func (s *CompleteRegisterLibraryResponse) SetStatusCode(v int32) *CompleteRegisterLibraryResponse {
	s.StatusCode = &v
	return s
}

func (s *CompleteRegisterLibraryResponse) SetBody(v *CatalogCommonResult) *CompleteRegisterLibraryResponse {
	s.Body = v
	return s
}

func (s *CompleteRegisterLibraryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

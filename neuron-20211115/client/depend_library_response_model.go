// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDependLibraryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DependLibraryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DependLibraryResponse
	GetStatusCode() *int32
	SetBody(v *CatalogCommonResult) *DependLibraryResponse
	GetBody() *CatalogCommonResult
}

type DependLibraryResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CatalogCommonResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DependLibraryResponse) String() string {
	return dara.Prettify(s)
}

func (s DependLibraryResponse) GoString() string {
	return s.String()
}

func (s *DependLibraryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DependLibraryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DependLibraryResponse) GetBody() *CatalogCommonResult {
	return s.Body
}

func (s *DependLibraryResponse) SetHeaders(v map[string]*string) *DependLibraryResponse {
	s.Headers = v
	return s
}

func (s *DependLibraryResponse) SetStatusCode(v int32) *DependLibraryResponse {
	s.StatusCode = &v
	return s
}

func (s *DependLibraryResponse) SetBody(v *CatalogCommonResult) *DependLibraryResponse {
	s.Body = v
	return s
}

func (s *DependLibraryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

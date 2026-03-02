// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveDependLibraryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveDependLibraryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveDependLibraryResponse
	GetStatusCode() *int32
	SetBody(v *CatalogCommonResult) *RemoveDependLibraryResponse
	GetBody() *CatalogCommonResult
}

type RemoveDependLibraryResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CatalogCommonResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveDependLibraryResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveDependLibraryResponse) GoString() string {
	return s.String()
}

func (s *RemoveDependLibraryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveDependLibraryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveDependLibraryResponse) GetBody() *CatalogCommonResult {
	return s.Body
}

func (s *RemoveDependLibraryResponse) SetHeaders(v map[string]*string) *RemoveDependLibraryResponse {
	s.Headers = v
	return s
}

func (s *RemoveDependLibraryResponse) SetStatusCode(v int32) *RemoveDependLibraryResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveDependLibraryResponse) SetBody(v *CatalogCommonResult) *RemoveDependLibraryResponse {
	s.Body = v
	return s
}

func (s *RemoveDependLibraryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

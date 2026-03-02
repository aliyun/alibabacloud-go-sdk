// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpShelfLibraryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpShelfLibraryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpShelfLibraryResponse
	GetStatusCode() *int32
	SetBody(v *CatalogCommonResult) *UpShelfLibraryResponse
	GetBody() *CatalogCommonResult
}

type UpShelfLibraryResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CatalogCommonResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpShelfLibraryResponse) String() string {
	return dara.Prettify(s)
}

func (s UpShelfLibraryResponse) GoString() string {
	return s.String()
}

func (s *UpShelfLibraryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpShelfLibraryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpShelfLibraryResponse) GetBody() *CatalogCommonResult {
	return s.Body
}

func (s *UpShelfLibraryResponse) SetHeaders(v map[string]*string) *UpShelfLibraryResponse {
	s.Headers = v
	return s
}

func (s *UpShelfLibraryResponse) SetStatusCode(v int32) *UpShelfLibraryResponse {
	s.StatusCode = &v
	return s
}

func (s *UpShelfLibraryResponse) SetBody(v *CatalogCommonResult) *UpShelfLibraryResponse {
	s.Body = v
	return s
}

func (s *UpShelfLibraryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

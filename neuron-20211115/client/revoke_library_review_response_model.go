// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeLibraryReviewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevokeLibraryReviewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevokeLibraryReviewResponse
	GetStatusCode() *int32
	SetBody(v *CatalogCommonResult) *RevokeLibraryReviewResponse
	GetBody() *CatalogCommonResult
}

type RevokeLibraryReviewResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CatalogCommonResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeLibraryReviewResponse) String() string {
	return dara.Prettify(s)
}

func (s RevokeLibraryReviewResponse) GoString() string {
	return s.String()
}

func (s *RevokeLibraryReviewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevokeLibraryReviewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevokeLibraryReviewResponse) GetBody() *CatalogCommonResult {
	return s.Body
}

func (s *RevokeLibraryReviewResponse) SetHeaders(v map[string]*string) *RevokeLibraryReviewResponse {
	s.Headers = v
	return s
}

func (s *RevokeLibraryReviewResponse) SetStatusCode(v int32) *RevokeLibraryReviewResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeLibraryReviewResponse) SetBody(v *CatalogCommonResult) *RevokeLibraryReviewResponse {
	s.Body = v
	return s
}

func (s *RevokeLibraryReviewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

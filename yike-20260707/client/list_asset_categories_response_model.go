// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAssetCategoriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAssetCategoriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAssetCategoriesResponse
	GetStatusCode() *int32
	SetBody(v *ListAssetCategoriesResponseBody) *ListAssetCategoriesResponse
	GetBody() *ListAssetCategoriesResponseBody
}

type ListAssetCategoriesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAssetCategoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAssetCategoriesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAssetCategoriesResponse) GoString() string {
	return s.String()
}

func (s *ListAssetCategoriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAssetCategoriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAssetCategoriesResponse) GetBody() *ListAssetCategoriesResponseBody {
	return s.Body
}

func (s *ListAssetCategoriesResponse) SetHeaders(v map[string]*string) *ListAssetCategoriesResponse {
	s.Headers = v
	return s
}

func (s *ListAssetCategoriesResponse) SetStatusCode(v int32) *ListAssetCategoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAssetCategoriesResponse) SetBody(v *ListAssetCategoriesResponseBody) *ListAssetCategoriesResponse {
	s.Body = v
	return s
}

func (s *ListAssetCategoriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

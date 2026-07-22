// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAssetCategoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAssetCategoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAssetCategoryResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAssetCategoryResponseBody) *DeleteAssetCategoryResponse
	GetBody() *DeleteAssetCategoryResponseBody
}

type DeleteAssetCategoryResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAssetCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAssetCategoryResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAssetCategoryResponse) GoString() string {
	return s.String()
}

func (s *DeleteAssetCategoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAssetCategoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAssetCategoryResponse) GetBody() *DeleteAssetCategoryResponseBody {
	return s.Body
}

func (s *DeleteAssetCategoryResponse) SetHeaders(v map[string]*string) *DeleteAssetCategoryResponse {
	s.Headers = v
	return s
}

func (s *DeleteAssetCategoryResponse) SetStatusCode(v int32) *DeleteAssetCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAssetCategoryResponse) SetBody(v *DeleteAssetCategoryResponseBody) *DeleteAssetCategoryResponse {
	s.Body = v
	return s
}

func (s *DeleteAssetCategoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

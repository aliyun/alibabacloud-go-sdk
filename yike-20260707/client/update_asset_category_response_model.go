// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAssetCategoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAssetCategoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAssetCategoryResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAssetCategoryResponseBody) *UpdateAssetCategoryResponse
	GetBody() *UpdateAssetCategoryResponseBody
}

type UpdateAssetCategoryResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAssetCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAssetCategoryResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAssetCategoryResponse) GoString() string {
	return s.String()
}

func (s *UpdateAssetCategoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAssetCategoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAssetCategoryResponse) GetBody() *UpdateAssetCategoryResponseBody {
	return s.Body
}

func (s *UpdateAssetCategoryResponse) SetHeaders(v map[string]*string) *UpdateAssetCategoryResponse {
	s.Headers = v
	return s
}

func (s *UpdateAssetCategoryResponse) SetStatusCode(v int32) *UpdateAssetCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAssetCategoryResponse) SetBody(v *UpdateAssetCategoryResponseBody) *UpdateAssetCategoryResponse {
	s.Body = v
	return s
}

func (s *UpdateAssetCategoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

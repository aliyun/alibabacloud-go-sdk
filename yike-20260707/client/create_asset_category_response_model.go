// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAssetCategoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAssetCategoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAssetCategoryResponse
	GetStatusCode() *int32
	SetBody(v *CreateAssetCategoryResponseBody) *CreateAssetCategoryResponse
	GetBody() *CreateAssetCategoryResponseBody
}

type CreateAssetCategoryResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAssetCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAssetCategoryResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAssetCategoryResponse) GoString() string {
	return s.String()
}

func (s *CreateAssetCategoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAssetCategoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAssetCategoryResponse) GetBody() *CreateAssetCategoryResponseBody {
	return s.Body
}

func (s *CreateAssetCategoryResponse) SetHeaders(v map[string]*string) *CreateAssetCategoryResponse {
	s.Headers = v
	return s
}

func (s *CreateAssetCategoryResponse) SetStatusCode(v int32) *CreateAssetCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAssetCategoryResponse) SetBody(v *CreateAssetCategoryResponseBody) *CreateAssetCategoryResponse {
	s.Body = v
	return s
}

func (s *CreateAssetCategoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

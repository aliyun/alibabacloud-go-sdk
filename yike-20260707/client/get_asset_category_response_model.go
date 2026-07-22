// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAssetCategoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAssetCategoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAssetCategoryResponse
	GetStatusCode() *int32
	SetBody(v *GetAssetCategoryResponseBody) *GetAssetCategoryResponse
	GetBody() *GetAssetCategoryResponseBody
}

type GetAssetCategoryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAssetCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAssetCategoryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAssetCategoryResponse) GoString() string {
	return s.String()
}

func (s *GetAssetCategoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAssetCategoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAssetCategoryResponse) GetBody() *GetAssetCategoryResponseBody {
	return s.Body
}

func (s *GetAssetCategoryResponse) SetHeaders(v map[string]*string) *GetAssetCategoryResponse {
	s.Headers = v
	return s
}

func (s *GetAssetCategoryResponse) SetStatusCode(v int32) *GetAssetCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAssetCategoryResponse) SetBody(v *GetAssetCategoryResponseBody) *GetAssetCategoryResponse {
	s.Body = v
	return s
}

func (s *GetAssetCategoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

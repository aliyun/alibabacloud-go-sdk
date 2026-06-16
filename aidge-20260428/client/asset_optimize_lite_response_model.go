// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssetOptimizeLiteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssetOptimizeLiteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssetOptimizeLiteResponse
	GetStatusCode() *int32
	SetBody(v *AssetOptimizeLiteResponseBody) *AssetOptimizeLiteResponse
	GetBody() *AssetOptimizeLiteResponseBody
}

type AssetOptimizeLiteResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssetOptimizeLiteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssetOptimizeLiteResponse) String() string {
	return dara.Prettify(s)
}

func (s AssetOptimizeLiteResponse) GoString() string {
	return s.String()
}

func (s *AssetOptimizeLiteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssetOptimizeLiteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssetOptimizeLiteResponse) GetBody() *AssetOptimizeLiteResponseBody {
	return s.Body
}

func (s *AssetOptimizeLiteResponse) SetHeaders(v map[string]*string) *AssetOptimizeLiteResponse {
	s.Headers = v
	return s
}

func (s *AssetOptimizeLiteResponse) SetStatusCode(v int32) *AssetOptimizeLiteResponse {
	s.StatusCode = &v
	return s
}

func (s *AssetOptimizeLiteResponse) SetBody(v *AssetOptimizeLiteResponseBody) *AssetOptimizeLiteResponse {
	s.Body = v
	return s
}

func (s *AssetOptimizeLiteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

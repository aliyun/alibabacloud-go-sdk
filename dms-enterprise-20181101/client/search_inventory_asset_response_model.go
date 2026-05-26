// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchInventoryAssetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchInventoryAssetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchInventoryAssetResponse
	GetStatusCode() *int32
	SetBody(v *SearchInventoryAssetResponseBody) *SearchInventoryAssetResponse
	GetBody() *SearchInventoryAssetResponseBody
}

type SearchInventoryAssetResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchInventoryAssetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchInventoryAssetResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchInventoryAssetResponse) GoString() string {
	return s.String()
}

func (s *SearchInventoryAssetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchInventoryAssetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchInventoryAssetResponse) GetBody() *SearchInventoryAssetResponseBody {
	return s.Body
}

func (s *SearchInventoryAssetResponse) SetHeaders(v map[string]*string) *SearchInventoryAssetResponse {
	s.Headers = v
	return s
}

func (s *SearchInventoryAssetResponse) SetStatusCode(v int32) *SearchInventoryAssetResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchInventoryAssetResponse) SetBody(v *SearchInventoryAssetResponseBody) *SearchInventoryAssetResponse {
	s.Body = v
	return s
}

func (s *SearchInventoryAssetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

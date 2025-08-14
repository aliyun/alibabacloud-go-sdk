// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVodPackagingAssetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVodPackagingAssetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVodPackagingAssetResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVodPackagingAssetResponseBody) *DeleteVodPackagingAssetResponse
	GetBody() *DeleteVodPackagingAssetResponseBody
}

type DeleteVodPackagingAssetResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVodPackagingAssetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVodPackagingAssetResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVodPackagingAssetResponse) GoString() string {
	return s.String()
}

func (s *DeleteVodPackagingAssetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVodPackagingAssetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVodPackagingAssetResponse) GetBody() *DeleteVodPackagingAssetResponseBody {
	return s.Body
}

func (s *DeleteVodPackagingAssetResponse) SetHeaders(v map[string]*string) *DeleteVodPackagingAssetResponse {
	s.Headers = v
	return s
}

func (s *DeleteVodPackagingAssetResponse) SetStatusCode(v int32) *DeleteVodPackagingAssetResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVodPackagingAssetResponse) SetBody(v *DeleteVodPackagingAssetResponseBody) *DeleteVodPackagingAssetResponse {
	s.Body = v
	return s
}

func (s *DeleteVodPackagingAssetResponse) Validate() error {
	return dara.Validate(s)
}

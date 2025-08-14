// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCreateVodPackagingAssetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchCreateVodPackagingAssetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchCreateVodPackagingAssetResponse
	GetStatusCode() *int32
	SetBody(v *BatchCreateVodPackagingAssetResponseBody) *BatchCreateVodPackagingAssetResponse
	GetBody() *BatchCreateVodPackagingAssetResponseBody
}

type BatchCreateVodPackagingAssetResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchCreateVodPackagingAssetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchCreateVodPackagingAssetResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateVodPackagingAssetResponse) GoString() string {
	return s.String()
}

func (s *BatchCreateVodPackagingAssetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchCreateVodPackagingAssetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchCreateVodPackagingAssetResponse) GetBody() *BatchCreateVodPackagingAssetResponseBody {
	return s.Body
}

func (s *BatchCreateVodPackagingAssetResponse) SetHeaders(v map[string]*string) *BatchCreateVodPackagingAssetResponse {
	s.Headers = v
	return s
}

func (s *BatchCreateVodPackagingAssetResponse) SetStatusCode(v int32) *BatchCreateVodPackagingAssetResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchCreateVodPackagingAssetResponse) SetBody(v *BatchCreateVodPackagingAssetResponseBody) *BatchCreateVodPackagingAssetResponse {
	s.Body = v
	return s
}

func (s *BatchCreateVodPackagingAssetResponse) Validate() error {
	return dara.Validate(s)
}

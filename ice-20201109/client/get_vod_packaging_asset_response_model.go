// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVodPackagingAssetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVodPackagingAssetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVodPackagingAssetResponse
	GetStatusCode() *int32
	SetBody(v *GetVodPackagingAssetResponseBody) *GetVodPackagingAssetResponse
	GetBody() *GetVodPackagingAssetResponseBody
}

type GetVodPackagingAssetResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVodPackagingAssetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVodPackagingAssetResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVodPackagingAssetResponse) GoString() string {
	return s.String()
}

func (s *GetVodPackagingAssetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVodPackagingAssetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVodPackagingAssetResponse) GetBody() *GetVodPackagingAssetResponseBody {
	return s.Body
}

func (s *GetVodPackagingAssetResponse) SetHeaders(v map[string]*string) *GetVodPackagingAssetResponse {
	s.Headers = v
	return s
}

func (s *GetVodPackagingAssetResponse) SetStatusCode(v int32) *GetVodPackagingAssetResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVodPackagingAssetResponse) SetBody(v *GetVodPackagingAssetResponseBody) *GetVodPackagingAssetResponse {
	s.Body = v
	return s
}

func (s *GetVodPackagingAssetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

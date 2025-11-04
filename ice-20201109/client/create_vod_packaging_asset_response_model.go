// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVodPackagingAssetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVodPackagingAssetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVodPackagingAssetResponse
	GetStatusCode() *int32
	SetBody(v *CreateVodPackagingAssetResponseBody) *CreateVodPackagingAssetResponse
	GetBody() *CreateVodPackagingAssetResponseBody
}

type CreateVodPackagingAssetResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVodPackagingAssetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVodPackagingAssetResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVodPackagingAssetResponse) GoString() string {
	return s.String()
}

func (s *CreateVodPackagingAssetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVodPackagingAssetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVodPackagingAssetResponse) GetBody() *CreateVodPackagingAssetResponseBody {
	return s.Body
}

func (s *CreateVodPackagingAssetResponse) SetHeaders(v map[string]*string) *CreateVodPackagingAssetResponse {
	s.Headers = v
	return s
}

func (s *CreateVodPackagingAssetResponse) SetStatusCode(v int32) *CreateVodPackagingAssetResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVodPackagingAssetResponse) SetBody(v *CreateVodPackagingAssetResponseBody) *CreateVodPackagingAssetResponse {
	s.Body = v
	return s
}

func (s *CreateVodPackagingAssetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

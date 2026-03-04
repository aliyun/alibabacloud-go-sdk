// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateYikeAssetUploadResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateYikeAssetUploadResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateYikeAssetUploadResponse
	GetStatusCode() *int32
	SetBody(v *CreateYikeAssetUploadResponseBody) *CreateYikeAssetUploadResponse
	GetBody() *CreateYikeAssetUploadResponseBody
}

type CreateYikeAssetUploadResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateYikeAssetUploadResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateYikeAssetUploadResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateYikeAssetUploadResponse) GoString() string {
	return s.String()
}

func (s *CreateYikeAssetUploadResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateYikeAssetUploadResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateYikeAssetUploadResponse) GetBody() *CreateYikeAssetUploadResponseBody {
	return s.Body
}

func (s *CreateYikeAssetUploadResponse) SetHeaders(v map[string]*string) *CreateYikeAssetUploadResponse {
	s.Headers = v
	return s
}

func (s *CreateYikeAssetUploadResponse) SetStatusCode(v int32) *CreateYikeAssetUploadResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateYikeAssetUploadResponse) SetBody(v *CreateYikeAssetUploadResponseBody) *CreateYikeAssetUploadResponse {
	s.Body = v
	return s
}

func (s *CreateYikeAssetUploadResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

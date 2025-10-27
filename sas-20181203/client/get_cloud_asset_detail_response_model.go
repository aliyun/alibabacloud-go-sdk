// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCloudAssetDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCloudAssetDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCloudAssetDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetCloudAssetDetailResponseBody) *GetCloudAssetDetailResponse
	GetBody() *GetCloudAssetDetailResponseBody
}

type GetCloudAssetDetailResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCloudAssetDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCloudAssetDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCloudAssetDetailResponse) GoString() string {
	return s.String()
}

func (s *GetCloudAssetDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCloudAssetDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCloudAssetDetailResponse) GetBody() *GetCloudAssetDetailResponseBody {
	return s.Body
}

func (s *GetCloudAssetDetailResponse) SetHeaders(v map[string]*string) *GetCloudAssetDetailResponse {
	s.Headers = v
	return s
}

func (s *GetCloudAssetDetailResponse) SetStatusCode(v int32) *GetCloudAssetDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCloudAssetDetailResponse) SetBody(v *GetCloudAssetDetailResponseBody) *GetCloudAssetDetailResponse {
	s.Body = v
	return s
}

func (s *GetCloudAssetDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

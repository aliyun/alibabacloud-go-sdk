// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetYikeAssetMediaInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetYikeAssetMediaInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetYikeAssetMediaInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetYikeAssetMediaInfoResponseBody) *GetYikeAssetMediaInfoResponse
	GetBody() *GetYikeAssetMediaInfoResponseBody
}

type GetYikeAssetMediaInfoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetYikeAssetMediaInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetYikeAssetMediaInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetYikeAssetMediaInfoResponse) GoString() string {
	return s.String()
}

func (s *GetYikeAssetMediaInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetYikeAssetMediaInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetYikeAssetMediaInfoResponse) GetBody() *GetYikeAssetMediaInfoResponseBody {
	return s.Body
}

func (s *GetYikeAssetMediaInfoResponse) SetHeaders(v map[string]*string) *GetYikeAssetMediaInfoResponse {
	s.Headers = v
	return s
}

func (s *GetYikeAssetMediaInfoResponse) SetStatusCode(v int32) *GetYikeAssetMediaInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetYikeAssetMediaInfoResponse) SetBody(v *GetYikeAssetMediaInfoResponseBody) *GetYikeAssetMediaInfoResponse {
	s.Body = v
	return s
}

func (s *GetYikeAssetMediaInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

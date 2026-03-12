// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteYikeAssetMediaInfosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteYikeAssetMediaInfosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteYikeAssetMediaInfosResponse
	GetStatusCode() *int32
	SetBody(v *DeleteYikeAssetMediaInfosResponseBody) *DeleteYikeAssetMediaInfosResponse
	GetBody() *DeleteYikeAssetMediaInfosResponseBody
}

type DeleteYikeAssetMediaInfosResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteYikeAssetMediaInfosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteYikeAssetMediaInfosResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteYikeAssetMediaInfosResponse) GoString() string {
	return s.String()
}

func (s *DeleteYikeAssetMediaInfosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteYikeAssetMediaInfosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteYikeAssetMediaInfosResponse) GetBody() *DeleteYikeAssetMediaInfosResponseBody {
	return s.Body
}

func (s *DeleteYikeAssetMediaInfosResponse) SetHeaders(v map[string]*string) *DeleteYikeAssetMediaInfosResponse {
	s.Headers = v
	return s
}

func (s *DeleteYikeAssetMediaInfosResponse) SetStatusCode(v int32) *DeleteYikeAssetMediaInfosResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteYikeAssetMediaInfosResponse) SetBody(v *DeleteYikeAssetMediaInfosResponseBody) *DeleteYikeAssetMediaInfosResponse {
	s.Body = v
	return s
}

func (s *DeleteYikeAssetMediaInfosResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

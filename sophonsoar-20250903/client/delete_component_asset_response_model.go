// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteComponentAssetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteComponentAssetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteComponentAssetResponse
	GetStatusCode() *int32
	SetBody(v *DeleteComponentAssetResponseBody) *DeleteComponentAssetResponse
	GetBody() *DeleteComponentAssetResponseBody
}

type DeleteComponentAssetResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteComponentAssetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteComponentAssetResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteComponentAssetResponse) GoString() string {
	return s.String()
}

func (s *DeleteComponentAssetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteComponentAssetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteComponentAssetResponse) GetBody() *DeleteComponentAssetResponseBody {
	return s.Body
}

func (s *DeleteComponentAssetResponse) SetHeaders(v map[string]*string) *DeleteComponentAssetResponse {
	s.Headers = v
	return s
}

func (s *DeleteComponentAssetResponse) SetStatusCode(v int32) *DeleteComponentAssetResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteComponentAssetResponse) SetBody(v *DeleteComponentAssetResponseBody) *DeleteComponentAssetResponse {
	s.Body = v
	return s
}

func (s *DeleteComponentAssetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagMetaAssetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTagMetaAssetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTagMetaAssetResponse
	GetStatusCode() *int32
	SetBody(v *ListTagMetaAssetResponseBody) *ListTagMetaAssetResponse
	GetBody() *ListTagMetaAssetResponseBody
}

type ListTagMetaAssetResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagMetaAssetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTagMetaAssetResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTagMetaAssetResponse) GoString() string {
	return s.String()
}

func (s *ListTagMetaAssetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTagMetaAssetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTagMetaAssetResponse) GetBody() *ListTagMetaAssetResponseBody {
	return s.Body
}

func (s *ListTagMetaAssetResponse) SetHeaders(v map[string]*string) *ListTagMetaAssetResponse {
	s.Headers = v
	return s
}

func (s *ListTagMetaAssetResponse) SetStatusCode(v int32) *ListTagMetaAssetResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagMetaAssetResponse) SetBody(v *ListTagMetaAssetResponseBody) *ListTagMetaAssetResponse {
	s.Body = v
	return s
}

func (s *ListTagMetaAssetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

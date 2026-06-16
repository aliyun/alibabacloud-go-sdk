// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssetOptimizeProResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssetOptimizeProResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssetOptimizeProResponse
	GetStatusCode() *int32
	SetBody(v *AssetOptimizeProResponseBody) *AssetOptimizeProResponse
	GetBody() *AssetOptimizeProResponseBody
}

type AssetOptimizeProResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssetOptimizeProResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssetOptimizeProResponse) String() string {
	return dara.Prettify(s)
}

func (s AssetOptimizeProResponse) GoString() string {
	return s.String()
}

func (s *AssetOptimizeProResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssetOptimizeProResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssetOptimizeProResponse) GetBody() *AssetOptimizeProResponseBody {
	return s.Body
}

func (s *AssetOptimizeProResponse) SetHeaders(v map[string]*string) *AssetOptimizeProResponse {
	s.Headers = v
	return s
}

func (s *AssetOptimizeProResponse) SetStatusCode(v int32) *AssetOptimizeProResponse {
	s.StatusCode = &v
	return s
}

func (s *AssetOptimizeProResponse) SetBody(v *AssetOptimizeProResponseBody) *AssetOptimizeProResponse {
	s.Body = v
	return s
}

func (s *AssetOptimizeProResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

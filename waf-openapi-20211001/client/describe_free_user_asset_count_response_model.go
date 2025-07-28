// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFreeUserAssetCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFreeUserAssetCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFreeUserAssetCountResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFreeUserAssetCountResponseBody) *DescribeFreeUserAssetCountResponse
	GetBody() *DescribeFreeUserAssetCountResponseBody
}

type DescribeFreeUserAssetCountResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFreeUserAssetCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFreeUserAssetCountResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFreeUserAssetCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeFreeUserAssetCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFreeUserAssetCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFreeUserAssetCountResponse) GetBody() *DescribeFreeUserAssetCountResponseBody {
	return s.Body
}

func (s *DescribeFreeUserAssetCountResponse) SetHeaders(v map[string]*string) *DescribeFreeUserAssetCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeFreeUserAssetCountResponse) SetStatusCode(v int32) *DescribeFreeUserAssetCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFreeUserAssetCountResponse) SetBody(v *DescribeFreeUserAssetCountResponseBody) *DescribeFreeUserAssetCountResponse {
	s.Body = v
	return s
}

func (s *DescribeFreeUserAssetCountResponse) Validate() error {
	return dara.Validate(s)
}

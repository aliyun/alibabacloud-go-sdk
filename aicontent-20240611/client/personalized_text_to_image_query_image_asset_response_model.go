// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPersonalizedTextToImageQueryImageAssetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PersonalizedTextToImageQueryImageAssetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PersonalizedTextToImageQueryImageAssetResponse
	GetStatusCode() *int32
	SetBody(v interface{}) *PersonalizedTextToImageQueryImageAssetResponse
	GetBody() interface{}
}

type PersonalizedTextToImageQueryImageAssetResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       interface{}        `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PersonalizedTextToImageQueryImageAssetResponse) String() string {
	return dara.Prettify(s)
}

func (s PersonalizedTextToImageQueryImageAssetResponse) GoString() string {
	return s.String()
}

func (s *PersonalizedTextToImageQueryImageAssetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PersonalizedTextToImageQueryImageAssetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PersonalizedTextToImageQueryImageAssetResponse) GetBody() interface{} {
	return s.Body
}

func (s *PersonalizedTextToImageQueryImageAssetResponse) SetHeaders(v map[string]*string) *PersonalizedTextToImageQueryImageAssetResponse {
	s.Headers = v
	return s
}

func (s *PersonalizedTextToImageQueryImageAssetResponse) SetStatusCode(v int32) *PersonalizedTextToImageQueryImageAssetResponse {
	s.StatusCode = &v
	return s
}

func (s *PersonalizedTextToImageQueryImageAssetResponse) SetBody(v interface{}) *PersonalizedTextToImageQueryImageAssetResponse {
	s.Body = v
	return s
}

func (s *PersonalizedTextToImageQueryImageAssetResponse) Validate() error {
	return dara.Validate(s)
}

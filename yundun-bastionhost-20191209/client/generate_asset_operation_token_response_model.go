// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateAssetOperationTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateAssetOperationTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateAssetOperationTokenResponse
	GetStatusCode() *int32
	SetBody(v *GenerateAssetOperationTokenResponseBody) *GenerateAssetOperationTokenResponse
	GetBody() *GenerateAssetOperationTokenResponseBody
}

type GenerateAssetOperationTokenResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateAssetOperationTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateAssetOperationTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateAssetOperationTokenResponse) GoString() string {
	return s.String()
}

func (s *GenerateAssetOperationTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateAssetOperationTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateAssetOperationTokenResponse) GetBody() *GenerateAssetOperationTokenResponseBody {
	return s.Body
}

func (s *GenerateAssetOperationTokenResponse) SetHeaders(v map[string]*string) *GenerateAssetOperationTokenResponse {
	s.Headers = v
	return s
}

func (s *GenerateAssetOperationTokenResponse) SetStatusCode(v int32) *GenerateAssetOperationTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateAssetOperationTokenResponse) SetBody(v *GenerateAssetOperationTokenResponseBody) *GenerateAssetOperationTokenResponse {
	s.Body = v
	return s
}

func (s *GenerateAssetOperationTokenResponse) Validate() error {
	return dara.Validate(s)
}

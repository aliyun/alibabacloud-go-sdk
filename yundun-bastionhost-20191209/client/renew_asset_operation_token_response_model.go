// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewAssetOperationTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RenewAssetOperationTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RenewAssetOperationTokenResponse
	GetStatusCode() *int32
	SetBody(v *RenewAssetOperationTokenResponseBody) *RenewAssetOperationTokenResponse
	GetBody() *RenewAssetOperationTokenResponseBody
}

type RenewAssetOperationTokenResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenewAssetOperationTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenewAssetOperationTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s RenewAssetOperationTokenResponse) GoString() string {
	return s.String()
}

func (s *RenewAssetOperationTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RenewAssetOperationTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RenewAssetOperationTokenResponse) GetBody() *RenewAssetOperationTokenResponseBody {
	return s.Body
}

func (s *RenewAssetOperationTokenResponse) SetHeaders(v map[string]*string) *RenewAssetOperationTokenResponse {
	s.Headers = v
	return s
}

func (s *RenewAssetOperationTokenResponse) SetStatusCode(v int32) *RenewAssetOperationTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *RenewAssetOperationTokenResponse) SetBody(v *RenewAssetOperationTokenResponseBody) *RenewAssetOperationTokenResponse {
	s.Body = v
	return s
}

func (s *RenewAssetOperationTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

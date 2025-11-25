// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableSdlProtectedAssetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableSdlProtectedAssetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableSdlProtectedAssetResponse
	GetStatusCode() *int32
	SetBody(v *DisableSdlProtectedAssetResponseBody) *DisableSdlProtectedAssetResponse
	GetBody() *DisableSdlProtectedAssetResponseBody
}

type DisableSdlProtectedAssetResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableSdlProtectedAssetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableSdlProtectedAssetResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableSdlProtectedAssetResponse) GoString() string {
	return s.String()
}

func (s *DisableSdlProtectedAssetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableSdlProtectedAssetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableSdlProtectedAssetResponse) GetBody() *DisableSdlProtectedAssetResponseBody {
	return s.Body
}

func (s *DisableSdlProtectedAssetResponse) SetHeaders(v map[string]*string) *DisableSdlProtectedAssetResponse {
	s.Headers = v
	return s
}

func (s *DisableSdlProtectedAssetResponse) SetStatusCode(v int32) *DisableSdlProtectedAssetResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableSdlProtectedAssetResponse) SetBody(v *DisableSdlProtectedAssetResponseBody) *DisableSdlProtectedAssetResponse {
	s.Body = v
	return s
}

func (s *DisableSdlProtectedAssetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

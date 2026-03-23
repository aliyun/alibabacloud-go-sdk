// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterApAssetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RegisterApAssetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RegisterApAssetResponse
	GetStatusCode() *int32
	SetBody(v *RegisterApAssetResponseBody) *RegisterApAssetResponse
	GetBody() *RegisterApAssetResponseBody
}

type RegisterApAssetResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RegisterApAssetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RegisterApAssetResponse) String() string {
	return dara.Prettify(s)
}

func (s RegisterApAssetResponse) GoString() string {
	return s.String()
}

func (s *RegisterApAssetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RegisterApAssetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RegisterApAssetResponse) GetBody() *RegisterApAssetResponseBody {
	return s.Body
}

func (s *RegisterApAssetResponse) SetHeaders(v map[string]*string) *RegisterApAssetResponse {
	s.Headers = v
	return s
}

func (s *RegisterApAssetResponse) SetStatusCode(v int32) *RegisterApAssetResponse {
	s.StatusCode = &v
	return s
}

func (s *RegisterApAssetResponse) SetBody(v *RegisterApAssetResponseBody) *RegisterApAssetResponse {
	s.Body = v
	return s
}

func (s *RegisterApAssetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOtpConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteOtpConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteOtpConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteOtpConfigResponseBody) *DeleteOtpConfigResponse
	GetBody() *DeleteOtpConfigResponseBody
}

type DeleteOtpConfigResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteOtpConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteOtpConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteOtpConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteOtpConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteOtpConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteOtpConfigResponse) GetBody() *DeleteOtpConfigResponseBody {
	return s.Body
}

func (s *DeleteOtpConfigResponse) SetHeaders(v map[string]*string) *DeleteOtpConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteOtpConfigResponse) SetStatusCode(v int32) *DeleteOtpConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteOtpConfigResponse) SetBody(v *DeleteOtpConfigResponseBody) *DeleteOtpConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteOtpConfigResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyCheckCustomConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VerifyCheckCustomConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VerifyCheckCustomConfigResponse
	GetStatusCode() *int32
	SetBody(v *VerifyCheckCustomConfigResponseBody) *VerifyCheckCustomConfigResponse
	GetBody() *VerifyCheckCustomConfigResponseBody
}

type VerifyCheckCustomConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyCheckCustomConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyCheckCustomConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s VerifyCheckCustomConfigResponse) GoString() string {
	return s.String()
}

func (s *VerifyCheckCustomConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VerifyCheckCustomConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VerifyCheckCustomConfigResponse) GetBody() *VerifyCheckCustomConfigResponseBody {
	return s.Body
}

func (s *VerifyCheckCustomConfigResponse) SetHeaders(v map[string]*string) *VerifyCheckCustomConfigResponse {
	s.Headers = v
	return s
}

func (s *VerifyCheckCustomConfigResponse) SetStatusCode(v int32) *VerifyCheckCustomConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyCheckCustomConfigResponse) SetBody(v *VerifyCheckCustomConfigResponseBody) *VerifyCheckCustomConfigResponse {
	s.Body = v
	return s
}

func (s *VerifyCheckCustomConfigResponse) Validate() error {
	return dara.Validate(s)
}

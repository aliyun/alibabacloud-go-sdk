// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceTwoFactorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceTwoFactorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceTwoFactorResponse
	GetStatusCode() *int32
	SetBody(v *GetInstanceTwoFactorResponseBody) *GetInstanceTwoFactorResponse
	GetBody() *GetInstanceTwoFactorResponseBody
}

type GetInstanceTwoFactorResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceTwoFactorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceTwoFactorResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceTwoFactorResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceTwoFactorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceTwoFactorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceTwoFactorResponse) GetBody() *GetInstanceTwoFactorResponseBody {
	return s.Body
}

func (s *GetInstanceTwoFactorResponse) SetHeaders(v map[string]*string) *GetInstanceTwoFactorResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceTwoFactorResponse) SetStatusCode(v int32) *GetInstanceTwoFactorResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceTwoFactorResponse) SetBody(v *GetInstanceTwoFactorResponseBody) *GetInstanceTwoFactorResponse {
	s.Body = v
	return s
}

func (s *GetInstanceTwoFactorResponse) Validate() error {
	return dara.Validate(s)
}

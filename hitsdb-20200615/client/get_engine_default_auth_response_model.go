// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEngineDefaultAuthResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEngineDefaultAuthResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEngineDefaultAuthResponse
	GetStatusCode() *int32
	SetBody(v *GetEngineDefaultAuthResponseBody) *GetEngineDefaultAuthResponse
	GetBody() *GetEngineDefaultAuthResponseBody
}

type GetEngineDefaultAuthResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEngineDefaultAuthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEngineDefaultAuthResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEngineDefaultAuthResponse) GoString() string {
	return s.String()
}

func (s *GetEngineDefaultAuthResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEngineDefaultAuthResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEngineDefaultAuthResponse) GetBody() *GetEngineDefaultAuthResponseBody {
	return s.Body
}

func (s *GetEngineDefaultAuthResponse) SetHeaders(v map[string]*string) *GetEngineDefaultAuthResponse {
	s.Headers = v
	return s
}

func (s *GetEngineDefaultAuthResponse) SetStatusCode(v int32) *GetEngineDefaultAuthResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEngineDefaultAuthResponse) SetBody(v *GetEngineDefaultAuthResponseBody) *GetEngineDefaultAuthResponse {
	s.Body = v
	return s
}

func (s *GetEngineDefaultAuthResponse) Validate() error {
	return dara.Validate(s)
}

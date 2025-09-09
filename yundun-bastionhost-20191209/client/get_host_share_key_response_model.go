// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHostShareKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHostShareKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHostShareKeyResponse
	GetStatusCode() *int32
	SetBody(v *GetHostShareKeyResponseBody) *GetHostShareKeyResponse
	GetBody() *GetHostShareKeyResponseBody
}

type GetHostShareKeyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHostShareKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHostShareKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHostShareKeyResponse) GoString() string {
	return s.String()
}

func (s *GetHostShareKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHostShareKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHostShareKeyResponse) GetBody() *GetHostShareKeyResponseBody {
	return s.Body
}

func (s *GetHostShareKeyResponse) SetHeaders(v map[string]*string) *GetHostShareKeyResponse {
	s.Headers = v
	return s
}

func (s *GetHostShareKeyResponse) SetStatusCode(v int32) *GetHostShareKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHostShareKeyResponse) SetBody(v *GetHostShareKeyResponseBody) *GetHostShareKeyResponse {
	s.Body = v
	return s
}

func (s *GetHostShareKeyResponse) Validate() error {
	return dara.Validate(s)
}

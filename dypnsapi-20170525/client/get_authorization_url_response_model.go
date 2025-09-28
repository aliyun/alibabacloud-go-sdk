// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuthorizationUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAuthorizationUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAuthorizationUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetAuthorizationUrlResponseBody) *GetAuthorizationUrlResponse
	GetBody() *GetAuthorizationUrlResponseBody
}

type GetAuthorizationUrlResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAuthorizationUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAuthorizationUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAuthorizationUrlResponse) GoString() string {
	return s.String()
}

func (s *GetAuthorizationUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAuthorizationUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAuthorizationUrlResponse) GetBody() *GetAuthorizationUrlResponseBody {
	return s.Body
}

func (s *GetAuthorizationUrlResponse) SetHeaders(v map[string]*string) *GetAuthorizationUrlResponse {
	s.Headers = v
	return s
}

func (s *GetAuthorizationUrlResponse) SetStatusCode(v int32) *GetAuthorizationUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAuthorizationUrlResponse) SetBody(v *GetAuthorizationUrlResponseBody) *GetAuthorizationUrlResponse {
	s.Body = v
	return s
}

func (s *GetAuthorizationUrlResponse) Validate() error {
	return dara.Validate(s)
}

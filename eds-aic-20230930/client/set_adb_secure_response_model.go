// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAdbSecureResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetAdbSecureResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetAdbSecureResponse
	GetStatusCode() *int32
	SetBody(v *SetAdbSecureResponseBody) *SetAdbSecureResponse
	GetBody() *SetAdbSecureResponseBody
}

type SetAdbSecureResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetAdbSecureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetAdbSecureResponse) String() string {
	return dara.Prettify(s)
}

func (s SetAdbSecureResponse) GoString() string {
	return s.String()
}

func (s *SetAdbSecureResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetAdbSecureResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetAdbSecureResponse) GetBody() *SetAdbSecureResponseBody {
	return s.Body
}

func (s *SetAdbSecureResponse) SetHeaders(v map[string]*string) *SetAdbSecureResponse {
	s.Headers = v
	return s
}

func (s *SetAdbSecureResponse) SetStatusCode(v int32) *SetAdbSecureResponse {
	s.StatusCode = &v
	return s
}

func (s *SetAdbSecureResponse) SetBody(v *SetAdbSecureResponseBody) *SetAdbSecureResponse {
	s.Body = v
	return s
}

func (s *SetAdbSecureResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMobileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMobileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMobileResponse
	GetStatusCode() *int32
	SetBody(v *GetMobileResponseBody) *GetMobileResponse
	GetBody() *GetMobileResponseBody
}

type GetMobileResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMobileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMobileResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMobileResponse) GoString() string {
	return s.String()
}

func (s *GetMobileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMobileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMobileResponse) GetBody() *GetMobileResponseBody {
	return s.Body
}

func (s *GetMobileResponse) SetHeaders(v map[string]*string) *GetMobileResponse {
	s.Headers = v
	return s
}

func (s *GetMobileResponse) SetStatusCode(v int32) *GetMobileResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMobileResponse) SetBody(v *GetMobileResponseBody) *GetMobileResponse {
	s.Body = v
	return s
}

func (s *GetMobileResponse) Validate() error {
	return dara.Validate(s)
}

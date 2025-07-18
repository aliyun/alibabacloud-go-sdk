// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIdpConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetIdpConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetIdpConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetIdpConfigResponseBody) *GetIdpConfigResponse
	GetBody() *GetIdpConfigResponseBody
}

type GetIdpConfigResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetIdpConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIdpConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetIdpConfigResponse) GoString() string {
	return s.String()
}

func (s *GetIdpConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetIdpConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetIdpConfigResponse) GetBody() *GetIdpConfigResponseBody {
	return s.Body
}

func (s *GetIdpConfigResponse) SetHeaders(v map[string]*string) *GetIdpConfigResponse {
	s.Headers = v
	return s
}

func (s *GetIdpConfigResponse) SetStatusCode(v int32) *GetIdpConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIdpConfigResponse) SetBody(v *GetIdpConfigResponseBody) *GetIdpConfigResponse {
	s.Body = v
	return s
}

func (s *GetIdpConfigResponse) Validate() error {
	return dara.Validate(s)
}

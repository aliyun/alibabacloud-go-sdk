// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCDNStatisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCDNStatisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCDNStatisResponse
	GetStatusCode() *int32
	SetBody(v *GetCDNStatisResponseBody) *GetCDNStatisResponse
	GetBody() *GetCDNStatisResponseBody
}

type GetCDNStatisResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCDNStatisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCDNStatisResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCDNStatisResponse) GoString() string {
	return s.String()
}

func (s *GetCDNStatisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCDNStatisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCDNStatisResponse) GetBody() *GetCDNStatisResponseBody {
	return s.Body
}

func (s *GetCDNStatisResponse) SetHeaders(v map[string]*string) *GetCDNStatisResponse {
	s.Headers = v
	return s
}

func (s *GetCDNStatisResponse) SetStatusCode(v int32) *GetCDNStatisResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCDNStatisResponse) SetBody(v *GetCDNStatisResponseBody) *GetCDNStatisResponse {
	s.Body = v
	return s
}

func (s *GetCDNStatisResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCDNStatisSumResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCDNStatisSumResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCDNStatisSumResponse
	GetStatusCode() *int32
	SetBody(v *GetCDNStatisSumResponseBody) *GetCDNStatisSumResponse
	GetBody() *GetCDNStatisSumResponseBody
}

type GetCDNStatisSumResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCDNStatisSumResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCDNStatisSumResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCDNStatisSumResponse) GoString() string {
	return s.String()
}

func (s *GetCDNStatisSumResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCDNStatisSumResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCDNStatisSumResponse) GetBody() *GetCDNStatisSumResponseBody {
	return s.Body
}

func (s *GetCDNStatisSumResponse) SetHeaders(v map[string]*string) *GetCDNStatisSumResponse {
	s.Headers = v
	return s
}

func (s *GetCDNStatisSumResponse) SetStatusCode(v int32) *GetCDNStatisSumResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCDNStatisSumResponse) SetBody(v *GetCDNStatisSumResponseBody) *GetCDNStatisSumResponse {
	s.Body = v
	return s
}

func (s *GetCDNStatisSumResponse) Validate() error {
	return dara.Validate(s)
}

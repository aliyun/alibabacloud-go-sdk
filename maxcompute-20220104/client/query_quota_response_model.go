// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryQuotaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryQuotaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryQuotaResponse
	GetStatusCode() *int32
	SetBody(v *QueryQuotaResponseBody) *QueryQuotaResponse
	GetBody() *QueryQuotaResponseBody
}

type QueryQuotaResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryQuotaResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryQuotaResponse) GoString() string {
	return s.String()
}

func (s *QueryQuotaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryQuotaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryQuotaResponse) GetBody() *QueryQuotaResponseBody {
	return s.Body
}

func (s *QueryQuotaResponse) SetHeaders(v map[string]*string) *QueryQuotaResponse {
	s.Headers = v
	return s
}

func (s *QueryQuotaResponse) SetStatusCode(v int32) *QueryQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryQuotaResponse) SetBody(v *QueryQuotaResponseBody) *QueryQuotaResponse {
	s.Body = v
	return s
}

func (s *QueryQuotaResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAppQuotaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryAppQuotaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryAppQuotaResponse
	GetStatusCode() *int32
	SetBody(v *QueryAppQuotaResponseBody) *QueryAppQuotaResponse
	GetBody() *QueryAppQuotaResponseBody
}

type QueryAppQuotaResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAppQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAppQuotaResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryAppQuotaResponse) GoString() string {
	return s.String()
}

func (s *QueryAppQuotaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryAppQuotaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryAppQuotaResponse) GetBody() *QueryAppQuotaResponseBody {
	return s.Body
}

func (s *QueryAppQuotaResponse) SetHeaders(v map[string]*string) *QueryAppQuotaResponse {
	s.Headers = v
	return s
}

func (s *QueryAppQuotaResponse) SetStatusCode(v int32) *QueryAppQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAppQuotaResponse) SetBody(v *QueryAppQuotaResponseBody) *QueryAppQuotaResponse {
	s.Body = v
	return s
}

func (s *QueryAppQuotaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

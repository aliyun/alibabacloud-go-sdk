// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryIpcQuotaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryIpcQuotaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryIpcQuotaResponse
	GetStatusCode() *int32
	SetBody(v *QueryIpcQuotaResponseBody) *QueryIpcQuotaResponse
	GetBody() *QueryIpcQuotaResponseBody
}

type QueryIpcQuotaResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryIpcQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryIpcQuotaResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryIpcQuotaResponse) GoString() string {
	return s.String()
}

func (s *QueryIpcQuotaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryIpcQuotaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryIpcQuotaResponse) GetBody() *QueryIpcQuotaResponseBody {
	return s.Body
}

func (s *QueryIpcQuotaResponse) SetHeaders(v map[string]*string) *QueryIpcQuotaResponse {
	s.Headers = v
	return s
}

func (s *QueryIpcQuotaResponse) SetStatusCode(v int32) *QueryIpcQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryIpcQuotaResponse) SetBody(v *QueryIpcQuotaResponseBody) *QueryIpcQuotaResponse {
	s.Body = v
	return s
}

func (s *QueryIpcQuotaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQuotaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListQuotaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListQuotaResponse
	GetStatusCode() *int32
	SetBody(v *ListQuotaResponseBody) *ListQuotaResponse
	GetBody() *ListQuotaResponseBody
}

type ListQuotaResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListQuotaResponse) String() string {
	return dara.Prettify(s)
}

func (s ListQuotaResponse) GoString() string {
	return s.String()
}

func (s *ListQuotaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListQuotaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListQuotaResponse) GetBody() *ListQuotaResponseBody {
	return s.Body
}

func (s *ListQuotaResponse) SetHeaders(v map[string]*string) *ListQuotaResponse {
	s.Headers = v
	return s
}

func (s *ListQuotaResponse) SetStatusCode(v int32) *ListQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQuotaResponse) SetBody(v *ListQuotaResponseBody) *ListQuotaResponse {
	s.Body = v
	return s
}

func (s *ListQuotaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

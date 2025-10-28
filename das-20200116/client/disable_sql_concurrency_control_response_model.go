// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableSqlConcurrencyControlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableSqlConcurrencyControlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableSqlConcurrencyControlResponse
	GetStatusCode() *int32
	SetBody(v *DisableSqlConcurrencyControlResponseBody) *DisableSqlConcurrencyControlResponse
	GetBody() *DisableSqlConcurrencyControlResponseBody
}

type DisableSqlConcurrencyControlResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableSqlConcurrencyControlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableSqlConcurrencyControlResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableSqlConcurrencyControlResponse) GoString() string {
	return s.String()
}

func (s *DisableSqlConcurrencyControlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableSqlConcurrencyControlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableSqlConcurrencyControlResponse) GetBody() *DisableSqlConcurrencyControlResponseBody {
	return s.Body
}

func (s *DisableSqlConcurrencyControlResponse) SetHeaders(v map[string]*string) *DisableSqlConcurrencyControlResponse {
	s.Headers = v
	return s
}

func (s *DisableSqlConcurrencyControlResponse) SetStatusCode(v int32) *DisableSqlConcurrencyControlResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableSqlConcurrencyControlResponse) SetBody(v *DisableSqlConcurrencyControlResponseBody) *DisableSqlConcurrencyControlResponse {
	s.Body = v
	return s
}

func (s *DisableSqlConcurrencyControlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

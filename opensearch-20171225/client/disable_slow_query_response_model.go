// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableSlowQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableSlowQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableSlowQueryResponse
	GetStatusCode() *int32
	SetBody(v *DisableSlowQueryResponseBody) *DisableSlowQueryResponse
	GetBody() *DisableSlowQueryResponseBody
}

type DisableSlowQueryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableSlowQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableSlowQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableSlowQueryResponse) GoString() string {
	return s.String()
}

func (s *DisableSlowQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableSlowQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableSlowQueryResponse) GetBody() *DisableSlowQueryResponseBody {
	return s.Body
}

func (s *DisableSlowQueryResponse) SetHeaders(v map[string]*string) *DisableSlowQueryResponse {
	s.Headers = v
	return s
}

func (s *DisableSlowQueryResponse) SetStatusCode(v int32) *DisableSlowQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableSlowQueryResponse) SetBody(v *DisableSlowQueryResponseBody) *DisableSlowQueryResponse {
	s.Body = v
	return s
}

func (s *DisableSlowQueryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

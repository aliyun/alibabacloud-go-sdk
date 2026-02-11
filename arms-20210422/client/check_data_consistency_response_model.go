// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDataConsistencyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckDataConsistencyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckDataConsistencyResponse
	GetStatusCode() *int32
	SetBody(v *CheckDataConsistencyResponseBody) *CheckDataConsistencyResponse
	GetBody() *CheckDataConsistencyResponseBody
}

type CheckDataConsistencyResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckDataConsistencyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckDataConsistencyResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckDataConsistencyResponse) GoString() string {
	return s.String()
}

func (s *CheckDataConsistencyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckDataConsistencyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckDataConsistencyResponse) GetBody() *CheckDataConsistencyResponseBody {
	return s.Body
}

func (s *CheckDataConsistencyResponse) SetHeaders(v map[string]*string) *CheckDataConsistencyResponse {
	s.Headers = v
	return s
}

func (s *CheckDataConsistencyResponse) SetStatusCode(v int32) *CheckDataConsistencyResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckDataConsistencyResponse) SetBody(v *CheckDataConsistencyResponseBody) *CheckDataConsistencyResponse {
	s.Body = v
	return s
}

func (s *CheckDataConsistencyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

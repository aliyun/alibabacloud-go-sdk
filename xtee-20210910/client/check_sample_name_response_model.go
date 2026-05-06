// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckSampleNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckSampleNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckSampleNameResponse
	GetStatusCode() *int32
	SetBody(v *CheckSampleNameResponseBody) *CheckSampleNameResponse
	GetBody() *CheckSampleNameResponseBody
}

type CheckSampleNameResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckSampleNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckSampleNameResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckSampleNameResponse) GoString() string {
	return s.String()
}

func (s *CheckSampleNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckSampleNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckSampleNameResponse) GetBody() *CheckSampleNameResponseBody {
	return s.Body
}

func (s *CheckSampleNameResponse) SetHeaders(v map[string]*string) *CheckSampleNameResponse {
	s.Headers = v
	return s
}

func (s *CheckSampleNameResponse) SetStatusCode(v int32) *CheckSampleNameResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckSampleNameResponse) SetBody(v *CheckSampleNameResponseBody) *CheckSampleNameResponse {
	s.Body = v
	return s
}

func (s *CheckSampleNameResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

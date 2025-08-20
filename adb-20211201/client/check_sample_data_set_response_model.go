// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckSampleDataSetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckSampleDataSetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckSampleDataSetResponse
	GetStatusCode() *int32
	SetBody(v *CheckSampleDataSetResponseBody) *CheckSampleDataSetResponse
	GetBody() *CheckSampleDataSetResponseBody
}

type CheckSampleDataSetResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckSampleDataSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckSampleDataSetResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckSampleDataSetResponse) GoString() string {
	return s.String()
}

func (s *CheckSampleDataSetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckSampleDataSetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckSampleDataSetResponse) GetBody() *CheckSampleDataSetResponseBody {
	return s.Body
}

func (s *CheckSampleDataSetResponse) SetHeaders(v map[string]*string) *CheckSampleDataSetResponse {
	s.Headers = v
	return s
}

func (s *CheckSampleDataSetResponse) SetStatusCode(v int32) *CheckSampleDataSetResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckSampleDataSetResponse) SetBody(v *CheckSampleDataSetResponseBody) *CheckSampleDataSetResponse {
	s.Body = v
	return s
}

func (s *CheckSampleDataSetResponse) Validate() error {
	return dara.Validate(s)
}

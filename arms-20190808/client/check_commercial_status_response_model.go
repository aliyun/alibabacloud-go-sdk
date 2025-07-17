// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckCommercialStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckCommercialStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckCommercialStatusResponse
	GetStatusCode() *int32
	SetBody(v *CheckCommercialStatusResponseBody) *CheckCommercialStatusResponse
	GetBody() *CheckCommercialStatusResponseBody
}

type CheckCommercialStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckCommercialStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckCommercialStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckCommercialStatusResponse) GoString() string {
	return s.String()
}

func (s *CheckCommercialStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckCommercialStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckCommercialStatusResponse) GetBody() *CheckCommercialStatusResponseBody {
	return s.Body
}

func (s *CheckCommercialStatusResponse) SetHeaders(v map[string]*string) *CheckCommercialStatusResponse {
	s.Headers = v
	return s
}

func (s *CheckCommercialStatusResponse) SetStatusCode(v int32) *CheckCommercialStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckCommercialStatusResponse) SetBody(v *CheckCommercialStatusResponseBody) *CheckCommercialStatusResponse {
	s.Body = v
	return s
}

func (s *CheckCommercialStatusResponse) Validate() error {
	return dara.Validate(s)
}

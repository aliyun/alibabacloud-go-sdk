// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValuateApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ValuateApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ValuateApplicationResponse
	GetStatusCode() *int32
	SetBody(v *ValuateApplicationResponseBody) *ValuateApplicationResponse
	GetBody() *ValuateApplicationResponseBody
}

type ValuateApplicationResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ValuateApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ValuateApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s ValuateApplicationResponse) GoString() string {
	return s.String()
}

func (s *ValuateApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ValuateApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ValuateApplicationResponse) GetBody() *ValuateApplicationResponseBody {
	return s.Body
}

func (s *ValuateApplicationResponse) SetHeaders(v map[string]*string) *ValuateApplicationResponse {
	s.Headers = v
	return s
}

func (s *ValuateApplicationResponse) SetStatusCode(v int32) *ValuateApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *ValuateApplicationResponse) SetBody(v *ValuateApplicationResponseBody) *ValuateApplicationResponse {
	s.Body = v
	return s
}

func (s *ValuateApplicationResponse) Validate() error {
	return dara.Validate(s)
}

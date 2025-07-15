// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCallSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCallSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCallSummaryResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCallSummaryResponseBody) *UpdateCallSummaryResponse
	GetBody() *UpdateCallSummaryResponseBody
}

type UpdateCallSummaryResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCallSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCallSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCallSummaryResponse) GoString() string {
	return s.String()
}

func (s *UpdateCallSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCallSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCallSummaryResponse) GetBody() *UpdateCallSummaryResponseBody {
	return s.Body
}

func (s *UpdateCallSummaryResponse) SetHeaders(v map[string]*string) *UpdateCallSummaryResponse {
	s.Headers = v
	return s
}

func (s *UpdateCallSummaryResponse) SetStatusCode(v int32) *UpdateCallSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCallSummaryResponse) SetBody(v *UpdateCallSummaryResponseBody) *UpdateCallSummaryResponse {
	s.Body = v
	return s
}

func (s *UpdateCallSummaryResponse) Validate() error {
	return dara.Validate(s)
}

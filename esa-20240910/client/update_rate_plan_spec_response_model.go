// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRatePlanSpecResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRatePlanSpecResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRatePlanSpecResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRatePlanSpecResponseBody) *UpdateRatePlanSpecResponse
	GetBody() *UpdateRatePlanSpecResponseBody
}

type UpdateRatePlanSpecResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRatePlanSpecResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRatePlanSpecResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRatePlanSpecResponse) GoString() string {
	return s.String()
}

func (s *UpdateRatePlanSpecResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRatePlanSpecResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRatePlanSpecResponse) GetBody() *UpdateRatePlanSpecResponseBody {
	return s.Body
}

func (s *UpdateRatePlanSpecResponse) SetHeaders(v map[string]*string) *UpdateRatePlanSpecResponse {
	s.Headers = v
	return s
}

func (s *UpdateRatePlanSpecResponse) SetStatusCode(v int32) *UpdateRatePlanSpecResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRatePlanSpecResponse) SetBody(v *UpdateRatePlanSpecResponseBody) *UpdateRatePlanSpecResponse {
	s.Body = v
	return s
}

func (s *UpdateRatePlanSpecResponse) Validate() error {
	return dara.Validate(s)
}

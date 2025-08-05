// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateQuotaPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateQuotaPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateQuotaPlanResponse
	GetStatusCode() *int32
	SetBody(v *UpdateQuotaPlanResponseBody) *UpdateQuotaPlanResponse
	GetBody() *UpdateQuotaPlanResponseBody
}

type UpdateQuotaPlanResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateQuotaPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateQuotaPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateQuotaPlanResponse) GoString() string {
	return s.String()
}

func (s *UpdateQuotaPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateQuotaPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateQuotaPlanResponse) GetBody() *UpdateQuotaPlanResponseBody {
	return s.Body
}

func (s *UpdateQuotaPlanResponse) SetHeaders(v map[string]*string) *UpdateQuotaPlanResponse {
	s.Headers = v
	return s
}

func (s *UpdateQuotaPlanResponse) SetStatusCode(v int32) *UpdateQuotaPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateQuotaPlanResponse) SetBody(v *UpdateQuotaPlanResponseBody) *UpdateQuotaPlanResponse {
	s.Body = v
	return s
}

func (s *UpdateQuotaPlanResponse) Validate() error {
	return dara.Validate(s)
}

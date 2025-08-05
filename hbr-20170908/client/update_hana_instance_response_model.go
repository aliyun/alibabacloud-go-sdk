// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHanaInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateHanaInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateHanaInstanceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateHanaInstanceResponseBody) *UpdateHanaInstanceResponse
	GetBody() *UpdateHanaInstanceResponseBody
}

type UpdateHanaInstanceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateHanaInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateHanaInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateHanaInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpdateHanaInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateHanaInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateHanaInstanceResponse) GetBody() *UpdateHanaInstanceResponseBody {
	return s.Body
}

func (s *UpdateHanaInstanceResponse) SetHeaders(v map[string]*string) *UpdateHanaInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpdateHanaInstanceResponse) SetStatusCode(v int32) *UpdateHanaInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateHanaInstanceResponse) SetBody(v *UpdateHanaInstanceResponseBody) *UpdateHanaInstanceResponse {
	s.Body = v
	return s
}

func (s *UpdateHanaInstanceResponse) Validate() error {
	return dara.Validate(s)
}

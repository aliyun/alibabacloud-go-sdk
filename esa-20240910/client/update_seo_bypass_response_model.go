// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSeoBypassResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSeoBypassResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSeoBypassResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSeoBypassResponseBody) *UpdateSeoBypassResponse
	GetBody() *UpdateSeoBypassResponseBody
}

type UpdateSeoBypassResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSeoBypassResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSeoBypassResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSeoBypassResponse) GoString() string {
	return s.String()
}

func (s *UpdateSeoBypassResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSeoBypassResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSeoBypassResponse) GetBody() *UpdateSeoBypassResponseBody {
	return s.Body
}

func (s *UpdateSeoBypassResponse) SetHeaders(v map[string]*string) *UpdateSeoBypassResponse {
	s.Headers = v
	return s
}

func (s *UpdateSeoBypassResponse) SetStatusCode(v int32) *UpdateSeoBypassResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSeoBypassResponse) SetBody(v *UpdateSeoBypassResponseBody) *UpdateSeoBypassResponse {
	s.Body = v
	return s
}

func (s *UpdateSeoBypassResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdateStatusResponseBody) *UpdateStatusResponse
	GetBody() *UpdateStatusResponseBody
}

type UpdateStatusResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateStatusResponse) GetBody() *UpdateStatusResponseBody {
	return s.Body
}

func (s *UpdateStatusResponse) SetHeaders(v map[string]*string) *UpdateStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateStatusResponse) SetStatusCode(v int32) *UpdateStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateStatusResponse) SetBody(v *UpdateStatusResponseBody) *UpdateStatusResponse {
	s.Body = v
	return s
}

func (s *UpdateStatusResponse) Validate() error {
	return dara.Validate(s)
}

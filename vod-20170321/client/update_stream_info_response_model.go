// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStreamInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateStreamInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateStreamInfoResponse
	GetStatusCode() *int32
	SetBody(v *UpdateStreamInfoResponseBody) *UpdateStreamInfoResponse
	GetBody() *UpdateStreamInfoResponseBody
}

type UpdateStreamInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateStreamInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateStreamInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateStreamInfoResponse) GoString() string {
	return s.String()
}

func (s *UpdateStreamInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateStreamInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateStreamInfoResponse) GetBody() *UpdateStreamInfoResponseBody {
	return s.Body
}

func (s *UpdateStreamInfoResponse) SetHeaders(v map[string]*string) *UpdateStreamInfoResponse {
	s.Headers = v
	return s
}

func (s *UpdateStreamInfoResponse) SetStatusCode(v int32) *UpdateStreamInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateStreamInfoResponse) SetBody(v *UpdateStreamInfoResponseBody) *UpdateStreamInfoResponse {
	s.Body = v
	return s
}

func (s *UpdateStreamInfoResponse) Validate() error {
	return dara.Validate(s)
}

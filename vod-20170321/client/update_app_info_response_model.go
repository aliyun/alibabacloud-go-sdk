// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAppInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAppInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAppInfoResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAppInfoResponseBody) *UpdateAppInfoResponse
	GetBody() *UpdateAppInfoResponseBody
}

type UpdateAppInfoResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAppInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAppInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppInfoResponse) GoString() string {
	return s.String()
}

func (s *UpdateAppInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAppInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAppInfoResponse) GetBody() *UpdateAppInfoResponseBody {
	return s.Body
}

func (s *UpdateAppInfoResponse) SetHeaders(v map[string]*string) *UpdateAppInfoResponse {
	s.Headers = v
	return s
}

func (s *UpdateAppInfoResponse) SetStatusCode(v int32) *UpdateAppInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAppInfoResponse) SetBody(v *UpdateAppInfoResponseBody) *UpdateAppInfoResponse {
	s.Body = v
	return s
}

func (s *UpdateAppInfoResponse) Validate() error {
	return dara.Validate(s)
}

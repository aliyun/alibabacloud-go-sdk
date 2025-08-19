// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAppResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAppResponseBody) *UpdateAppResponse
	GetBody() *UpdateAppResponseBody
}

type UpdateAppResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAppResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppResponse) GoString() string {
	return s.String()
}

func (s *UpdateAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAppResponse) GetBody() *UpdateAppResponseBody {
	return s.Body
}

func (s *UpdateAppResponse) SetHeaders(v map[string]*string) *UpdateAppResponse {
	s.Headers = v
	return s
}

func (s *UpdateAppResponse) SetStatusCode(v int32) *UpdateAppResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAppResponse) SetBody(v *UpdateAppResponseBody) *UpdateAppResponse {
	s.Body = v
	return s
}

func (s *UpdateAppResponse) Validate() error {
	return dara.Validate(s)
}

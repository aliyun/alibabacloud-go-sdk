// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNodeRunModeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateNodeRunModeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateNodeRunModeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateNodeRunModeResponseBody) *UpdateNodeRunModeResponse
	GetBody() *UpdateNodeRunModeResponseBody
}

type UpdateNodeRunModeResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateNodeRunModeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateNodeRunModeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateNodeRunModeResponse) GoString() string {
	return s.String()
}

func (s *UpdateNodeRunModeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateNodeRunModeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateNodeRunModeResponse) GetBody() *UpdateNodeRunModeResponseBody {
	return s.Body
}

func (s *UpdateNodeRunModeResponse) SetHeaders(v map[string]*string) *UpdateNodeRunModeResponse {
	s.Headers = v
	return s
}

func (s *UpdateNodeRunModeResponse) SetStatusCode(v int32) *UpdateNodeRunModeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateNodeRunModeResponse) SetBody(v *UpdateNodeRunModeResponseBody) *UpdateNodeRunModeResponse {
	s.Body = v
	return s
}

func (s *UpdateNodeRunModeResponse) Validate() error {
	return dara.Validate(s)
}

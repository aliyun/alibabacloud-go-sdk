// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDIProjectConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDIProjectConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDIProjectConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDIProjectConfigResponseBody) *UpdateDIProjectConfigResponse
	GetBody() *UpdateDIProjectConfigResponseBody
}

type UpdateDIProjectConfigResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDIProjectConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDIProjectConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDIProjectConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateDIProjectConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDIProjectConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDIProjectConfigResponse) GetBody() *UpdateDIProjectConfigResponseBody {
	return s.Body
}

func (s *UpdateDIProjectConfigResponse) SetHeaders(v map[string]*string) *UpdateDIProjectConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateDIProjectConfigResponse) SetStatusCode(v int32) *UpdateDIProjectConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDIProjectConfigResponse) SetBody(v *UpdateDIProjectConfigResponseBody) *UpdateDIProjectConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateDIProjectConfigResponse) Validate() error {
	return dara.Validate(s)
}

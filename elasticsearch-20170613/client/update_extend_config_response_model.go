// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateExtendConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateExtendConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateExtendConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateExtendConfigResponseBody) *UpdateExtendConfigResponse
	GetBody() *UpdateExtendConfigResponseBody
}

type UpdateExtendConfigResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateExtendConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateExtendConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateExtendConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateExtendConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateExtendConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateExtendConfigResponse) GetBody() *UpdateExtendConfigResponseBody {
	return s.Body
}

func (s *UpdateExtendConfigResponse) SetHeaders(v map[string]*string) *UpdateExtendConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateExtendConfigResponse) SetStatusCode(v int32) *UpdateExtendConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateExtendConfigResponse) SetBody(v *UpdateExtendConfigResponseBody) *UpdateExtendConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateExtendConfigResponse) Validate() error {
	return dara.Validate(s)
}

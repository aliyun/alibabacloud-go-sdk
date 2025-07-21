// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationAccessPointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateApplicationAccessPointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateApplicationAccessPointResponse
	GetStatusCode() *int32
	SetBody(v *UpdateApplicationAccessPointResponseBody) *UpdateApplicationAccessPointResponse
	GetBody() *UpdateApplicationAccessPointResponseBody
}

type UpdateApplicationAccessPointResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateApplicationAccessPointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateApplicationAccessPointResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationAccessPointResponse) GoString() string {
	return s.String()
}

func (s *UpdateApplicationAccessPointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateApplicationAccessPointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateApplicationAccessPointResponse) GetBody() *UpdateApplicationAccessPointResponseBody {
	return s.Body
}

func (s *UpdateApplicationAccessPointResponse) SetHeaders(v map[string]*string) *UpdateApplicationAccessPointResponse {
	s.Headers = v
	return s
}

func (s *UpdateApplicationAccessPointResponse) SetStatusCode(v int32) *UpdateApplicationAccessPointResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateApplicationAccessPointResponse) SetBody(v *UpdateApplicationAccessPointResponseBody) *UpdateApplicationAccessPointResponse {
	s.Body = v
	return s
}

func (s *UpdateApplicationAccessPointResponse) Validate() error {
	return dara.Validate(s)
}

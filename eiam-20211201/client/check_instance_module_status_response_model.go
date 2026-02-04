// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckInstanceModuleStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckInstanceModuleStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckInstanceModuleStatusResponse
	GetStatusCode() *int32
	SetBody(v *CheckInstanceModuleStatusResponseBody) *CheckInstanceModuleStatusResponse
	GetBody() *CheckInstanceModuleStatusResponseBody
}

type CheckInstanceModuleStatusResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckInstanceModuleStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckInstanceModuleStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckInstanceModuleStatusResponse) GoString() string {
	return s.String()
}

func (s *CheckInstanceModuleStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckInstanceModuleStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckInstanceModuleStatusResponse) GetBody() *CheckInstanceModuleStatusResponseBody {
	return s.Body
}

func (s *CheckInstanceModuleStatusResponse) SetHeaders(v map[string]*string) *CheckInstanceModuleStatusResponse {
	s.Headers = v
	return s
}

func (s *CheckInstanceModuleStatusResponse) SetStatusCode(v int32) *CheckInstanceModuleStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckInstanceModuleStatusResponse) SetBody(v *CheckInstanceModuleStatusResponseBody) *CheckInstanceModuleStatusResponse {
	s.Body = v
	return s
}

func (s *CheckInstanceModuleStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

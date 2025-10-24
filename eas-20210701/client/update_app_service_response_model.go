// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAppServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAppServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAppServiceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAppServiceResponseBody) *UpdateAppServiceResponse
	GetBody() *UpdateAppServiceResponseBody
}

type UpdateAppServiceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAppServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAppServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppServiceResponse) GoString() string {
	return s.String()
}

func (s *UpdateAppServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAppServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAppServiceResponse) GetBody() *UpdateAppServiceResponseBody {
	return s.Body
}

func (s *UpdateAppServiceResponse) SetHeaders(v map[string]*string) *UpdateAppServiceResponse {
	s.Headers = v
	return s
}

func (s *UpdateAppServiceResponse) SetStatusCode(v int32) *UpdateAppServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAppServiceResponse) SetBody(v *UpdateAppServiceResponseBody) *UpdateAppServiceResponse {
	s.Body = v
	return s
}

func (s *UpdateAppServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

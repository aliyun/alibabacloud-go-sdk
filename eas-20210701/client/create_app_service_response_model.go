// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAppServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAppServiceResponse
	GetStatusCode() *int32
	SetBody(v *CreateAppServiceResponseBody) *CreateAppServiceResponse
	GetBody() *CreateAppServiceResponseBody
}

type CreateAppServiceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAppServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAppServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAppServiceResponse) GoString() string {
	return s.String()
}

func (s *CreateAppServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAppServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAppServiceResponse) GetBody() *CreateAppServiceResponseBody {
	return s.Body
}

func (s *CreateAppServiceResponse) SetHeaders(v map[string]*string) *CreateAppServiceResponse {
	s.Headers = v
	return s
}

func (s *CreateAppServiceResponse) SetStatusCode(v int32) *CreateAppServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppServiceResponse) SetBody(v *CreateAppServiceResponseBody) *CreateAppServiceResponse {
	s.Body = v
	return s
}

func (s *CreateAppServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

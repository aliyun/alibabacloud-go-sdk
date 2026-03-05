// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppTokenServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAppTokenServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAppTokenServiceResponse
	GetStatusCode() *int32
	SetBody(v *CreateAppTokenServiceResponseBody) *CreateAppTokenServiceResponse
	GetBody() *CreateAppTokenServiceResponseBody
}

type CreateAppTokenServiceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAppTokenServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAppTokenServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAppTokenServiceResponse) GoString() string {
	return s.String()
}

func (s *CreateAppTokenServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAppTokenServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAppTokenServiceResponse) GetBody() *CreateAppTokenServiceResponseBody {
	return s.Body
}

func (s *CreateAppTokenServiceResponse) SetHeaders(v map[string]*string) *CreateAppTokenServiceResponse {
	s.Headers = v
	return s
}

func (s *CreateAppTokenServiceResponse) SetStatusCode(v int32) *CreateAppTokenServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppTokenServiceResponse) SetBody(v *CreateAppTokenServiceResponseBody) *CreateAppTokenServiceResponse {
	s.Body = v
	return s
}

func (s *CreateAppTokenServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

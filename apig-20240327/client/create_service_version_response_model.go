// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateServiceVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateServiceVersionResponse
	GetStatusCode() *int32
	SetBody(v *CreateServiceVersionResponseBody) *CreateServiceVersionResponse
	GetBody() *CreateServiceVersionResponseBody
}

type CreateServiceVersionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateServiceVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServiceVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceVersionResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateServiceVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateServiceVersionResponse) GetBody() *CreateServiceVersionResponseBody {
	return s.Body
}

func (s *CreateServiceVersionResponse) SetHeaders(v map[string]*string) *CreateServiceVersionResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceVersionResponse) SetStatusCode(v int32) *CreateServiceVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceVersionResponse) SetBody(v *CreateServiceVersionResponseBody) *CreateServiceVersionResponse {
	s.Body = v
	return s
}

func (s *CreateServiceVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

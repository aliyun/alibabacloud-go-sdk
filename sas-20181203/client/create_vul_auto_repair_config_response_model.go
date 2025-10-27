// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVulAutoRepairConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVulAutoRepairConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVulAutoRepairConfigResponse
	GetStatusCode() *int32
	SetBody(v *CreateVulAutoRepairConfigResponseBody) *CreateVulAutoRepairConfigResponse
	GetBody() *CreateVulAutoRepairConfigResponseBody
}

type CreateVulAutoRepairConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVulAutoRepairConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVulAutoRepairConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVulAutoRepairConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateVulAutoRepairConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVulAutoRepairConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVulAutoRepairConfigResponse) GetBody() *CreateVulAutoRepairConfigResponseBody {
	return s.Body
}

func (s *CreateVulAutoRepairConfigResponse) SetHeaders(v map[string]*string) *CreateVulAutoRepairConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateVulAutoRepairConfigResponse) SetStatusCode(v int32) *CreateVulAutoRepairConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVulAutoRepairConfigResponse) SetBody(v *CreateVulAutoRepairConfigResponseBody) *CreateVulAutoRepairConfigResponse {
	s.Body = v
	return s
}

func (s *CreateVulAutoRepairConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

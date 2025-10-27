// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVulAutoRepairConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVulAutoRepairConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVulAutoRepairConfigResponse
	GetStatusCode() *int32
	SetBody(v *ListVulAutoRepairConfigResponseBody) *ListVulAutoRepairConfigResponse
	GetBody() *ListVulAutoRepairConfigResponseBody
}

type ListVulAutoRepairConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVulAutoRepairConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVulAutoRepairConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVulAutoRepairConfigResponse) GoString() string {
	return s.String()
}

func (s *ListVulAutoRepairConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVulAutoRepairConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVulAutoRepairConfigResponse) GetBody() *ListVulAutoRepairConfigResponseBody {
	return s.Body
}

func (s *ListVulAutoRepairConfigResponse) SetHeaders(v map[string]*string) *ListVulAutoRepairConfigResponse {
	s.Headers = v
	return s
}

func (s *ListVulAutoRepairConfigResponse) SetStatusCode(v int32) *ListVulAutoRepairConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVulAutoRepairConfigResponse) SetBody(v *ListVulAutoRepairConfigResponseBody) *ListVulAutoRepairConfigResponse {
	s.Body = v
	return s
}

func (s *ListVulAutoRepairConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

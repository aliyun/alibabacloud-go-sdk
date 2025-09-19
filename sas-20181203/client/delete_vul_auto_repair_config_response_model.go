// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVulAutoRepairConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVulAutoRepairConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVulAutoRepairConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVulAutoRepairConfigResponseBody) *DeleteVulAutoRepairConfigResponse
	GetBody() *DeleteVulAutoRepairConfigResponseBody
}

type DeleteVulAutoRepairConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVulAutoRepairConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVulAutoRepairConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVulAutoRepairConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteVulAutoRepairConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVulAutoRepairConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVulAutoRepairConfigResponse) GetBody() *DeleteVulAutoRepairConfigResponseBody {
	return s.Body
}

func (s *DeleteVulAutoRepairConfigResponse) SetHeaders(v map[string]*string) *DeleteVulAutoRepairConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteVulAutoRepairConfigResponse) SetStatusCode(v int32) *DeleteVulAutoRepairConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVulAutoRepairConfigResponse) SetBody(v *DeleteVulAutoRepairConfigResponseBody) *DeleteVulAutoRepairConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteVulAutoRepairConfigResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceInstanceAttributesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateServiceInstanceAttributesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateServiceInstanceAttributesResponse
	GetStatusCode() *int32
	SetBody(v *UpdateServiceInstanceAttributesResponseBody) *UpdateServiceInstanceAttributesResponse
	GetBody() *UpdateServiceInstanceAttributesResponseBody
}

type UpdateServiceInstanceAttributesResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateServiceInstanceAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateServiceInstanceAttributesResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceInstanceAttributesResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceInstanceAttributesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateServiceInstanceAttributesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateServiceInstanceAttributesResponse) GetBody() *UpdateServiceInstanceAttributesResponseBody {
	return s.Body
}

func (s *UpdateServiceInstanceAttributesResponse) SetHeaders(v map[string]*string) *UpdateServiceInstanceAttributesResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceInstanceAttributesResponse) SetStatusCode(v int32) *UpdateServiceInstanceAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServiceInstanceAttributesResponse) SetBody(v *UpdateServiceInstanceAttributesResponseBody) *UpdateServiceInstanceAttributesResponse {
	s.Body = v
	return s
}

func (s *UpdateServiceInstanceAttributesResponse) Validate() error {
	return dara.Validate(s)
}

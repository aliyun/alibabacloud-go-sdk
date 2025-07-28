// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceElasticVCUUpperLimitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateInstanceElasticVCUUpperLimitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateInstanceElasticVCUUpperLimitResponse
	GetStatusCode() *int32
	SetBody(v *UpdateInstanceElasticVCUUpperLimitResponseBody) *UpdateInstanceElasticVCUUpperLimitResponse
	GetBody() *UpdateInstanceElasticVCUUpperLimitResponseBody
}

type UpdateInstanceElasticVCUUpperLimitResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInstanceElasticVCUUpperLimitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInstanceElasticVCUUpperLimitResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceElasticVCUUpperLimitResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceElasticVCUUpperLimitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateInstanceElasticVCUUpperLimitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateInstanceElasticVCUUpperLimitResponse) GetBody() *UpdateInstanceElasticVCUUpperLimitResponseBody {
	return s.Body
}

func (s *UpdateInstanceElasticVCUUpperLimitResponse) SetHeaders(v map[string]*string) *UpdateInstanceElasticVCUUpperLimitResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceElasticVCUUpperLimitResponse) SetStatusCode(v int32) *UpdateInstanceElasticVCUUpperLimitResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceElasticVCUUpperLimitResponse) SetBody(v *UpdateInstanceElasticVCUUpperLimitResponseBody) *UpdateInstanceElasticVCUUpperLimitResponse {
	s.Body = v
	return s
}

func (s *UpdateInstanceElasticVCUUpperLimitResponse) Validate() error {
	return dara.Validate(s)
}

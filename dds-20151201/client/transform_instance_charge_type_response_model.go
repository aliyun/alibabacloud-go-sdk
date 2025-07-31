// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransformInstanceChargeTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TransformInstanceChargeTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TransformInstanceChargeTypeResponse
	GetStatusCode() *int32
	SetBody(v *TransformInstanceChargeTypeResponseBody) *TransformInstanceChargeTypeResponse
	GetBody() *TransformInstanceChargeTypeResponseBody
}

type TransformInstanceChargeTypeResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TransformInstanceChargeTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TransformInstanceChargeTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s TransformInstanceChargeTypeResponse) GoString() string {
	return s.String()
}

func (s *TransformInstanceChargeTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TransformInstanceChargeTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TransformInstanceChargeTypeResponse) GetBody() *TransformInstanceChargeTypeResponseBody {
	return s.Body
}

func (s *TransformInstanceChargeTypeResponse) SetHeaders(v map[string]*string) *TransformInstanceChargeTypeResponse {
	s.Headers = v
	return s
}

func (s *TransformInstanceChargeTypeResponse) SetStatusCode(v int32) *TransformInstanceChargeTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *TransformInstanceChargeTypeResponse) SetBody(v *TransformInstanceChargeTypeResponseBody) *TransformInstanceChargeTypeResponse {
	s.Body = v
	return s
}

func (s *TransformInstanceChargeTypeResponse) Validate() error {
	return dara.Validate(s)
}

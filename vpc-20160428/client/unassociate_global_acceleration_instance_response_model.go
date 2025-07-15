// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnassociateGlobalAccelerationInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnassociateGlobalAccelerationInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnassociateGlobalAccelerationInstanceResponse
	GetStatusCode() *int32
	SetBody(v *UnassociateGlobalAccelerationInstanceResponseBody) *UnassociateGlobalAccelerationInstanceResponse
	GetBody() *UnassociateGlobalAccelerationInstanceResponseBody
}

type UnassociateGlobalAccelerationInstanceResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnassociateGlobalAccelerationInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnassociateGlobalAccelerationInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s UnassociateGlobalAccelerationInstanceResponse) GoString() string {
	return s.String()
}

func (s *UnassociateGlobalAccelerationInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnassociateGlobalAccelerationInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnassociateGlobalAccelerationInstanceResponse) GetBody() *UnassociateGlobalAccelerationInstanceResponseBody {
	return s.Body
}

func (s *UnassociateGlobalAccelerationInstanceResponse) SetHeaders(v map[string]*string) *UnassociateGlobalAccelerationInstanceResponse {
	s.Headers = v
	return s
}

func (s *UnassociateGlobalAccelerationInstanceResponse) SetStatusCode(v int32) *UnassociateGlobalAccelerationInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *UnassociateGlobalAccelerationInstanceResponse) SetBody(v *UnassociateGlobalAccelerationInstanceResponseBody) *UnassociateGlobalAccelerationInstanceResponse {
	s.Body = v
	return s
}

func (s *UnassociateGlobalAccelerationInstanceResponse) Validate() error {
	return dara.Validate(s)
}

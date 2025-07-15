// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGlobalAccelerationInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteGlobalAccelerationInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteGlobalAccelerationInstanceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteGlobalAccelerationInstanceResponseBody) *DeleteGlobalAccelerationInstanceResponse
	GetBody() *DeleteGlobalAccelerationInstanceResponseBody
}

type DeleteGlobalAccelerationInstanceResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGlobalAccelerationInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGlobalAccelerationInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteGlobalAccelerationInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteGlobalAccelerationInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteGlobalAccelerationInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteGlobalAccelerationInstanceResponse) GetBody() *DeleteGlobalAccelerationInstanceResponseBody {
	return s.Body
}

func (s *DeleteGlobalAccelerationInstanceResponse) SetHeaders(v map[string]*string) *DeleteGlobalAccelerationInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteGlobalAccelerationInstanceResponse) SetStatusCode(v int32) *DeleteGlobalAccelerationInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGlobalAccelerationInstanceResponse) SetBody(v *DeleteGlobalAccelerationInstanceResponseBody) *DeleteGlobalAccelerationInstanceResponse {
	s.Body = v
	return s
}

func (s *DeleteGlobalAccelerationInstanceResponse) Validate() error {
	return dara.Validate(s)
}

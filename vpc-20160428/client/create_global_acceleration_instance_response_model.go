// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGlobalAccelerationInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateGlobalAccelerationInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateGlobalAccelerationInstanceResponse
	GetStatusCode() *int32
	SetBody(v *CreateGlobalAccelerationInstanceResponseBody) *CreateGlobalAccelerationInstanceResponse
	GetBody() *CreateGlobalAccelerationInstanceResponseBody
}

type CreateGlobalAccelerationInstanceResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGlobalAccelerationInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGlobalAccelerationInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateGlobalAccelerationInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateGlobalAccelerationInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateGlobalAccelerationInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateGlobalAccelerationInstanceResponse) GetBody() *CreateGlobalAccelerationInstanceResponseBody {
	return s.Body
}

func (s *CreateGlobalAccelerationInstanceResponse) SetHeaders(v map[string]*string) *CreateGlobalAccelerationInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateGlobalAccelerationInstanceResponse) SetStatusCode(v int32) *CreateGlobalAccelerationInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGlobalAccelerationInstanceResponse) SetBody(v *CreateGlobalAccelerationInstanceResponseBody) *CreateGlobalAccelerationInstanceResponse {
	s.Body = v
	return s
}

func (s *CreateGlobalAccelerationInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

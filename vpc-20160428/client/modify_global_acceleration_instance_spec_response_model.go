// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyGlobalAccelerationInstanceSpecResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyGlobalAccelerationInstanceSpecResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyGlobalAccelerationInstanceSpecResponse
	GetStatusCode() *int32
	SetBody(v *ModifyGlobalAccelerationInstanceSpecResponseBody) *ModifyGlobalAccelerationInstanceSpecResponse
	GetBody() *ModifyGlobalAccelerationInstanceSpecResponseBody
}

type ModifyGlobalAccelerationInstanceSpecResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyGlobalAccelerationInstanceSpecResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyGlobalAccelerationInstanceSpecResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyGlobalAccelerationInstanceSpecResponse) GoString() string {
	return s.String()
}

func (s *ModifyGlobalAccelerationInstanceSpecResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyGlobalAccelerationInstanceSpecResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyGlobalAccelerationInstanceSpecResponse) GetBody() *ModifyGlobalAccelerationInstanceSpecResponseBody {
	return s.Body
}

func (s *ModifyGlobalAccelerationInstanceSpecResponse) SetHeaders(v map[string]*string) *ModifyGlobalAccelerationInstanceSpecResponse {
	s.Headers = v
	return s
}

func (s *ModifyGlobalAccelerationInstanceSpecResponse) SetStatusCode(v int32) *ModifyGlobalAccelerationInstanceSpecResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyGlobalAccelerationInstanceSpecResponse) SetBody(v *ModifyGlobalAccelerationInstanceSpecResponseBody) *ModifyGlobalAccelerationInstanceSpecResponse {
	s.Body = v
	return s
}

func (s *ModifyGlobalAccelerationInstanceSpecResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

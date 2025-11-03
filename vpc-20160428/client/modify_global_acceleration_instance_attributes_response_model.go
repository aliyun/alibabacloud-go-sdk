// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyGlobalAccelerationInstanceAttributesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyGlobalAccelerationInstanceAttributesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyGlobalAccelerationInstanceAttributesResponse
	GetStatusCode() *int32
	SetBody(v *ModifyGlobalAccelerationInstanceAttributesResponseBody) *ModifyGlobalAccelerationInstanceAttributesResponse
	GetBody() *ModifyGlobalAccelerationInstanceAttributesResponseBody
}

type ModifyGlobalAccelerationInstanceAttributesResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyGlobalAccelerationInstanceAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyGlobalAccelerationInstanceAttributesResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyGlobalAccelerationInstanceAttributesResponse) GoString() string {
	return s.String()
}

func (s *ModifyGlobalAccelerationInstanceAttributesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyGlobalAccelerationInstanceAttributesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyGlobalAccelerationInstanceAttributesResponse) GetBody() *ModifyGlobalAccelerationInstanceAttributesResponseBody {
	return s.Body
}

func (s *ModifyGlobalAccelerationInstanceAttributesResponse) SetHeaders(v map[string]*string) *ModifyGlobalAccelerationInstanceAttributesResponse {
	s.Headers = v
	return s
}

func (s *ModifyGlobalAccelerationInstanceAttributesResponse) SetStatusCode(v int32) *ModifyGlobalAccelerationInstanceAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyGlobalAccelerationInstanceAttributesResponse) SetBody(v *ModifyGlobalAccelerationInstanceAttributesResponseBody) *ModifyGlobalAccelerationInstanceAttributesResponse {
	s.Body = v
	return s
}

func (s *ModifyGlobalAccelerationInstanceAttributesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

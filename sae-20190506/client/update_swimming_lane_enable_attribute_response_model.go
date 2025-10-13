// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSwimmingLaneEnableAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSwimmingLaneEnableAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSwimmingLaneEnableAttributeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSwimmingLaneEnableAttributeResponseBody) *UpdateSwimmingLaneEnableAttributeResponse
	GetBody() *UpdateSwimmingLaneEnableAttributeResponseBody
}

type UpdateSwimmingLaneEnableAttributeResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSwimmingLaneEnableAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSwimmingLaneEnableAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSwimmingLaneEnableAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateSwimmingLaneEnableAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSwimmingLaneEnableAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSwimmingLaneEnableAttributeResponse) GetBody() *UpdateSwimmingLaneEnableAttributeResponseBody {
	return s.Body
}

func (s *UpdateSwimmingLaneEnableAttributeResponse) SetHeaders(v map[string]*string) *UpdateSwimmingLaneEnableAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateSwimmingLaneEnableAttributeResponse) SetStatusCode(v int32) *UpdateSwimmingLaneEnableAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSwimmingLaneEnableAttributeResponse) SetBody(v *UpdateSwimmingLaneEnableAttributeResponseBody) *UpdateSwimmingLaneEnableAttributeResponse {
	s.Body = v
	return s
}

func (s *UpdateSwimmingLaneEnableAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

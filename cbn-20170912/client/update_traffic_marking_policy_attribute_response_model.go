// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTrafficMarkingPolicyAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTrafficMarkingPolicyAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTrafficMarkingPolicyAttributeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTrafficMarkingPolicyAttributeResponseBody) *UpdateTrafficMarkingPolicyAttributeResponse
	GetBody() *UpdateTrafficMarkingPolicyAttributeResponseBody
}

type UpdateTrafficMarkingPolicyAttributeResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTrafficMarkingPolicyAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTrafficMarkingPolicyAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTrafficMarkingPolicyAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateTrafficMarkingPolicyAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTrafficMarkingPolicyAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTrafficMarkingPolicyAttributeResponse) GetBody() *UpdateTrafficMarkingPolicyAttributeResponseBody {
	return s.Body
}

func (s *UpdateTrafficMarkingPolicyAttributeResponse) SetHeaders(v map[string]*string) *UpdateTrafficMarkingPolicyAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateTrafficMarkingPolicyAttributeResponse) SetStatusCode(v int32) *UpdateTrafficMarkingPolicyAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTrafficMarkingPolicyAttributeResponse) SetBody(v *UpdateTrafficMarkingPolicyAttributeResponseBody) *UpdateTrafficMarkingPolicyAttributeResponse {
	s.Body = v
	return s
}

func (s *UpdateTrafficMarkingPolicyAttributeResponse) Validate() error {
	return dara.Validate(s)
}

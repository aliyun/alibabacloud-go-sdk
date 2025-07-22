// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAutoScalingConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAutoScalingConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAutoScalingConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAutoScalingConfigResponseBody) *ModifyAutoScalingConfigResponse
	GetBody() *ModifyAutoScalingConfigResponseBody
}

type ModifyAutoScalingConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAutoScalingConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAutoScalingConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAutoScalingConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyAutoScalingConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAutoScalingConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAutoScalingConfigResponse) GetBody() *ModifyAutoScalingConfigResponseBody {
	return s.Body
}

func (s *ModifyAutoScalingConfigResponse) SetHeaders(v map[string]*string) *ModifyAutoScalingConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyAutoScalingConfigResponse) SetStatusCode(v int32) *ModifyAutoScalingConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAutoScalingConfigResponse) SetBody(v *ModifyAutoScalingConfigResponseBody) *ModifyAutoScalingConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyAutoScalingConfigResponse) Validate() error {
	return dara.Validate(s)
}

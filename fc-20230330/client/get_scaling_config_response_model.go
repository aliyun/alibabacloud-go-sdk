// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScalingConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetScalingConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetScalingConfigResponse
	GetStatusCode() *int32
	SetBody(v *ScalingConfigStatus) *GetScalingConfigResponse
	GetBody() *ScalingConfigStatus
}

type GetScalingConfigResponse struct {
	Headers    map[string]*string   `json:"headers" xml:"headers"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ScalingConfigStatus `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetScalingConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetScalingConfigResponse) GoString() string {
	return s.String()
}

func (s *GetScalingConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetScalingConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetScalingConfigResponse) GetBody() *ScalingConfigStatus {
	return s.Body
}

func (s *GetScalingConfigResponse) SetHeaders(v map[string]*string) *GetScalingConfigResponse {
	s.Headers = v
	return s
}

func (s *GetScalingConfigResponse) SetStatusCode(v int32) *GetScalingConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetScalingConfigResponse) SetBody(v *ScalingConfigStatus) *GetScalingConfigResponse {
	s.Body = v
	return s
}

func (s *GetScalingConfigResponse) Validate() error {
	return dara.Validate(s)
}

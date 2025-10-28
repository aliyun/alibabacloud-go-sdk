// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutScalingConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutScalingConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutScalingConfigResponse
	GetStatusCode() *int32
	SetBody(v *ScalingConfigStatus) *PutScalingConfigResponse
	GetBody() *ScalingConfigStatus
}

type PutScalingConfigResponse struct {
	Headers    map[string]*string   `json:"headers" xml:"headers"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ScalingConfigStatus `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutScalingConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s PutScalingConfigResponse) GoString() string {
	return s.String()
}

func (s *PutScalingConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutScalingConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutScalingConfigResponse) GetBody() *ScalingConfigStatus {
	return s.Body
}

func (s *PutScalingConfigResponse) SetHeaders(v map[string]*string) *PutScalingConfigResponse {
	s.Headers = v
	return s
}

func (s *PutScalingConfigResponse) SetStatusCode(v int32) *PutScalingConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *PutScalingConfigResponse) SetBody(v *ScalingConfigStatus) *PutScalingConfigResponse {
	s.Body = v
	return s
}

func (s *PutScalingConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

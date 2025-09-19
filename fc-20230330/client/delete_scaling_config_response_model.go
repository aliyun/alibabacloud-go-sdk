// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScalingConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteScalingConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteScalingConfigResponse
	GetStatusCode() *int32
}

type DeleteScalingConfigResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteScalingConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteScalingConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteScalingConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteScalingConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteScalingConfigResponse) SetHeaders(v map[string]*string) *DeleteScalingConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteScalingConfigResponse) SetStatusCode(v int32) *DeleteScalingConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteScalingConfigResponse) Validate() error {
	return dara.Validate(s)
}

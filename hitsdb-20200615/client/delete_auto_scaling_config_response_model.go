// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAutoScalingConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAutoScalingConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAutoScalingConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAutoScalingConfigResponseBody) *DeleteAutoScalingConfigResponse
	GetBody() *DeleteAutoScalingConfigResponseBody
}

type DeleteAutoScalingConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAutoScalingConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAutoScalingConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAutoScalingConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteAutoScalingConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAutoScalingConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAutoScalingConfigResponse) GetBody() *DeleteAutoScalingConfigResponseBody {
	return s.Body
}

func (s *DeleteAutoScalingConfigResponse) SetHeaders(v map[string]*string) *DeleteAutoScalingConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteAutoScalingConfigResponse) SetStatusCode(v int32) *DeleteAutoScalingConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAutoScalingConfigResponse) SetBody(v *DeleteAutoScalingConfigResponseBody) *DeleteAutoScalingConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteAutoScalingConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

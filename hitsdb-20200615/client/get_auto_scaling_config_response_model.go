// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAutoScalingConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAutoScalingConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAutoScalingConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetAutoScalingConfigResponseBody) *GetAutoScalingConfigResponse
	GetBody() *GetAutoScalingConfigResponseBody
}

type GetAutoScalingConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAutoScalingConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAutoScalingConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAutoScalingConfigResponse) GoString() string {
	return s.String()
}

func (s *GetAutoScalingConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAutoScalingConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAutoScalingConfigResponse) GetBody() *GetAutoScalingConfigResponseBody {
	return s.Body
}

func (s *GetAutoScalingConfigResponse) SetHeaders(v map[string]*string) *GetAutoScalingConfigResponse {
	s.Headers = v
	return s
}

func (s *GetAutoScalingConfigResponse) SetStatusCode(v int32) *GetAutoScalingConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAutoScalingConfigResponse) SetBody(v *GetAutoScalingConfigResponseBody) *GetAutoScalingConfigResponse {
	s.Body = v
	return s
}

func (s *GetAutoScalingConfigResponse) Validate() error {
	return dara.Validate(s)
}

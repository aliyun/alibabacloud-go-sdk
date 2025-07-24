// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAutoScalingActivityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAutoScalingActivityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAutoScalingActivityResponse
	GetStatusCode() *int32
	SetBody(v *GetAutoScalingActivityResponseBody) *GetAutoScalingActivityResponse
	GetBody() *GetAutoScalingActivityResponseBody
}

type GetAutoScalingActivityResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAutoScalingActivityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAutoScalingActivityResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAutoScalingActivityResponse) GoString() string {
	return s.String()
}

func (s *GetAutoScalingActivityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAutoScalingActivityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAutoScalingActivityResponse) GetBody() *GetAutoScalingActivityResponseBody {
	return s.Body
}

func (s *GetAutoScalingActivityResponse) SetHeaders(v map[string]*string) *GetAutoScalingActivityResponse {
	s.Headers = v
	return s
}

func (s *GetAutoScalingActivityResponse) SetStatusCode(v int32) *GetAutoScalingActivityResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAutoScalingActivityResponse) SetBody(v *GetAutoScalingActivityResponseBody) *GetAutoScalingActivityResponse {
	s.Body = v
	return s
}

func (s *GetAutoScalingActivityResponse) Validate() error {
	return dara.Validate(s)
}

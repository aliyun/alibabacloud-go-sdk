// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetPolicyAccessTimeRangeConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetPolicyAccessTimeRangeConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetPolicyAccessTimeRangeConfigResponse
	GetStatusCode() *int32
	SetBody(v *SetPolicyAccessTimeRangeConfigResponseBody) *SetPolicyAccessTimeRangeConfigResponse
	GetBody() *SetPolicyAccessTimeRangeConfigResponseBody
}

type SetPolicyAccessTimeRangeConfigResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetPolicyAccessTimeRangeConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetPolicyAccessTimeRangeConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s SetPolicyAccessTimeRangeConfigResponse) GoString() string {
	return s.String()
}

func (s *SetPolicyAccessTimeRangeConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetPolicyAccessTimeRangeConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetPolicyAccessTimeRangeConfigResponse) GetBody() *SetPolicyAccessTimeRangeConfigResponseBody {
	return s.Body
}

func (s *SetPolicyAccessTimeRangeConfigResponse) SetHeaders(v map[string]*string) *SetPolicyAccessTimeRangeConfigResponse {
	s.Headers = v
	return s
}

func (s *SetPolicyAccessTimeRangeConfigResponse) SetStatusCode(v int32) *SetPolicyAccessTimeRangeConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SetPolicyAccessTimeRangeConfigResponse) SetBody(v *SetPolicyAccessTimeRangeConfigResponseBody) *SetPolicyAccessTimeRangeConfigResponse {
	s.Body = v
	return s
}

func (s *SetPolicyAccessTimeRangeConfigResponse) Validate() error {
	return dara.Validate(s)
}

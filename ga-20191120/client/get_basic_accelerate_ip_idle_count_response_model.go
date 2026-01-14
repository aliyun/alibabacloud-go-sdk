// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBasicAccelerateIpIdleCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBasicAccelerateIpIdleCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBasicAccelerateIpIdleCountResponse
	GetStatusCode() *int32
	SetBody(v *GetBasicAccelerateIpIdleCountResponseBody) *GetBasicAccelerateIpIdleCountResponse
	GetBody() *GetBasicAccelerateIpIdleCountResponseBody
}

type GetBasicAccelerateIpIdleCountResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBasicAccelerateIpIdleCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBasicAccelerateIpIdleCountResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBasicAccelerateIpIdleCountResponse) GoString() string {
	return s.String()
}

func (s *GetBasicAccelerateIpIdleCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBasicAccelerateIpIdleCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBasicAccelerateIpIdleCountResponse) GetBody() *GetBasicAccelerateIpIdleCountResponseBody {
	return s.Body
}

func (s *GetBasicAccelerateIpIdleCountResponse) SetHeaders(v map[string]*string) *GetBasicAccelerateIpIdleCountResponse {
	s.Headers = v
	return s
}

func (s *GetBasicAccelerateIpIdleCountResponse) SetStatusCode(v int32) *GetBasicAccelerateIpIdleCountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBasicAccelerateIpIdleCountResponse) SetBody(v *GetBasicAccelerateIpIdleCountResponseBody) *GetBasicAccelerateIpIdleCountResponse {
	s.Body = v
	return s
}

func (s *GetBasicAccelerateIpIdleCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

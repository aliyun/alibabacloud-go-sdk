// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBasicAccelerateIpResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBasicAccelerateIpResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBasicAccelerateIpResponse
	GetStatusCode() *int32
	SetBody(v *GetBasicAccelerateIpResponseBody) *GetBasicAccelerateIpResponse
	GetBody() *GetBasicAccelerateIpResponseBody
}

type GetBasicAccelerateIpResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBasicAccelerateIpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBasicAccelerateIpResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBasicAccelerateIpResponse) GoString() string {
	return s.String()
}

func (s *GetBasicAccelerateIpResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBasicAccelerateIpResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBasicAccelerateIpResponse) GetBody() *GetBasicAccelerateIpResponseBody {
	return s.Body
}

func (s *GetBasicAccelerateIpResponse) SetHeaders(v map[string]*string) *GetBasicAccelerateIpResponse {
	s.Headers = v
	return s
}

func (s *GetBasicAccelerateIpResponse) SetStatusCode(v int32) *GetBasicAccelerateIpResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBasicAccelerateIpResponse) SetBody(v *GetBasicAccelerateIpResponseBody) *GetBasicAccelerateIpResponse {
	s.Body = v
	return s
}

func (s *GetBasicAccelerateIpResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

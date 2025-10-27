// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddIdcProbeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddIdcProbeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddIdcProbeResponse
	GetStatusCode() *int32
	SetBody(v *AddIdcProbeResponseBody) *AddIdcProbeResponse
	GetBody() *AddIdcProbeResponseBody
}

type AddIdcProbeResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddIdcProbeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddIdcProbeResponse) String() string {
	return dara.Prettify(s)
}

func (s AddIdcProbeResponse) GoString() string {
	return s.String()
}

func (s *AddIdcProbeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddIdcProbeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddIdcProbeResponse) GetBody() *AddIdcProbeResponseBody {
	return s.Body
}

func (s *AddIdcProbeResponse) SetHeaders(v map[string]*string) *AddIdcProbeResponse {
	s.Headers = v
	return s
}

func (s *AddIdcProbeResponse) SetStatusCode(v int32) *AddIdcProbeResponse {
	s.StatusCode = &v
	return s
}

func (s *AddIdcProbeResponse) SetBody(v *AddIdcProbeResponseBody) *AddIdcProbeResponse {
	s.Body = v
	return s
}

func (s *AddIdcProbeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SwitchServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SwitchServiceResponse
	GetStatusCode() *int32
	SetBody(v *SwitchServiceResponseBody) *SwitchServiceResponse
	GetBody() *SwitchServiceResponseBody
}

type SwitchServiceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SwitchServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SwitchServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s SwitchServiceResponse) GoString() string {
	return s.String()
}

func (s *SwitchServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SwitchServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SwitchServiceResponse) GetBody() *SwitchServiceResponseBody {
	return s.Body
}

func (s *SwitchServiceResponse) SetHeaders(v map[string]*string) *SwitchServiceResponse {
	s.Headers = v
	return s
}

func (s *SwitchServiceResponse) SetStatusCode(v int32) *SwitchServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *SwitchServiceResponse) SetBody(v *SwitchServiceResponseBody) *SwitchServiceResponse {
	s.Body = v
	return s
}

func (s *SwitchServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

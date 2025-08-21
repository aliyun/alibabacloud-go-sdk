// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopSnatIpForSnatEntryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopSnatIpForSnatEntryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopSnatIpForSnatEntryResponse
	GetStatusCode() *int32
	SetBody(v *StopSnatIpForSnatEntryResponseBody) *StopSnatIpForSnatEntryResponse
	GetBody() *StopSnatIpForSnatEntryResponseBody
}

type StopSnatIpForSnatEntryResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopSnatIpForSnatEntryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopSnatIpForSnatEntryResponse) String() string {
	return dara.Prettify(s)
}

func (s StopSnatIpForSnatEntryResponse) GoString() string {
	return s.String()
}

func (s *StopSnatIpForSnatEntryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopSnatIpForSnatEntryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopSnatIpForSnatEntryResponse) GetBody() *StopSnatIpForSnatEntryResponseBody {
	return s.Body
}

func (s *StopSnatIpForSnatEntryResponse) SetHeaders(v map[string]*string) *StopSnatIpForSnatEntryResponse {
	s.Headers = v
	return s
}

func (s *StopSnatIpForSnatEntryResponse) SetStatusCode(v int32) *StopSnatIpForSnatEntryResponse {
	s.StatusCode = &v
	return s
}

func (s *StopSnatIpForSnatEntryResponse) SetBody(v *StopSnatIpForSnatEntryResponseBody) *StopSnatIpForSnatEntryResponse {
	s.Body = v
	return s
}

func (s *StopSnatIpForSnatEntryResponse) Validate() error {
	return dara.Validate(s)
}

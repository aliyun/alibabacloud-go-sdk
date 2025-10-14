// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartSnatIpForSnatEntryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartSnatIpForSnatEntryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartSnatIpForSnatEntryResponse
	GetStatusCode() *int32
	SetBody(v *StartSnatIpForSnatEntryResponseBody) *StartSnatIpForSnatEntryResponse
	GetBody() *StartSnatIpForSnatEntryResponseBody
}

type StartSnatIpForSnatEntryResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartSnatIpForSnatEntryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartSnatIpForSnatEntryResponse) String() string {
	return dara.Prettify(s)
}

func (s StartSnatIpForSnatEntryResponse) GoString() string {
	return s.String()
}

func (s *StartSnatIpForSnatEntryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartSnatIpForSnatEntryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartSnatIpForSnatEntryResponse) GetBody() *StartSnatIpForSnatEntryResponseBody {
	return s.Body
}

func (s *StartSnatIpForSnatEntryResponse) SetHeaders(v map[string]*string) *StartSnatIpForSnatEntryResponse {
	s.Headers = v
	return s
}

func (s *StartSnatIpForSnatEntryResponse) SetStatusCode(v int32) *StartSnatIpForSnatEntryResponse {
	s.StatusCode = &v
	return s
}

func (s *StartSnatIpForSnatEntryResponse) SetBody(v *StartSnatIpForSnatEntryResponseBody) *StartSnatIpForSnatEntryResponse {
	s.Body = v
	return s
}

func (s *StartSnatIpForSnatEntryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSnatIpForSnatEntryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddSnatIpForSnatEntryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddSnatIpForSnatEntryResponse
	GetStatusCode() *int32
	SetBody(v *AddSnatIpForSnatEntryResponseBody) *AddSnatIpForSnatEntryResponse
	GetBody() *AddSnatIpForSnatEntryResponseBody
}

type AddSnatIpForSnatEntryResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddSnatIpForSnatEntryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddSnatIpForSnatEntryResponse) String() string {
	return dara.Prettify(s)
}

func (s AddSnatIpForSnatEntryResponse) GoString() string {
	return s.String()
}

func (s *AddSnatIpForSnatEntryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddSnatIpForSnatEntryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddSnatIpForSnatEntryResponse) GetBody() *AddSnatIpForSnatEntryResponseBody {
	return s.Body
}

func (s *AddSnatIpForSnatEntryResponse) SetHeaders(v map[string]*string) *AddSnatIpForSnatEntryResponse {
	s.Headers = v
	return s
}

func (s *AddSnatIpForSnatEntryResponse) SetStatusCode(v int32) *AddSnatIpForSnatEntryResponse {
	s.StatusCode = &v
	return s
}

func (s *AddSnatIpForSnatEntryResponse) SetBody(v *AddSnatIpForSnatEntryResponseBody) *AddSnatIpForSnatEntryResponse {
	s.Body = v
	return s
}

func (s *AddSnatIpForSnatEntryResponse) Validate() error {
	return dara.Validate(s)
}

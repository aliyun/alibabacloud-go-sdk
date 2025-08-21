// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSnatIpForSnatEntryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSnatIpForSnatEntryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSnatIpForSnatEntryResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSnatIpForSnatEntryResponseBody) *DeleteSnatIpForSnatEntryResponse
	GetBody() *DeleteSnatIpForSnatEntryResponseBody
}

type DeleteSnatIpForSnatEntryResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSnatIpForSnatEntryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSnatIpForSnatEntryResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSnatIpForSnatEntryResponse) GoString() string {
	return s.String()
}

func (s *DeleteSnatIpForSnatEntryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSnatIpForSnatEntryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSnatIpForSnatEntryResponse) GetBody() *DeleteSnatIpForSnatEntryResponseBody {
	return s.Body
}

func (s *DeleteSnatIpForSnatEntryResponse) SetHeaders(v map[string]*string) *DeleteSnatIpForSnatEntryResponse {
	s.Headers = v
	return s
}

func (s *DeleteSnatIpForSnatEntryResponse) SetStatusCode(v int32) *DeleteSnatIpForSnatEntryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSnatIpForSnatEntryResponse) SetBody(v *DeleteSnatIpForSnatEntryResponseBody) *DeleteSnatIpForSnatEntryResponse {
	s.Body = v
	return s
}

func (s *DeleteSnatIpForSnatEntryResponse) Validate() error {
	return dara.Validate(s)
}

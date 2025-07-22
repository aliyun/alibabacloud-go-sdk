// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSnatEntryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSnatEntryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSnatEntryResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSnatEntryResponseBody) *DeleteSnatEntryResponse
	GetBody() *DeleteSnatEntryResponseBody
}

type DeleteSnatEntryResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSnatEntryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSnatEntryResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSnatEntryResponse) GoString() string {
	return s.String()
}

func (s *DeleteSnatEntryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSnatEntryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSnatEntryResponse) GetBody() *DeleteSnatEntryResponseBody {
	return s.Body
}

func (s *DeleteSnatEntryResponse) SetHeaders(v map[string]*string) *DeleteSnatEntryResponse {
	s.Headers = v
	return s
}

func (s *DeleteSnatEntryResponse) SetStatusCode(v int32) *DeleteSnatEntryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSnatEntryResponse) SetBody(v *DeleteSnatEntryResponseBody) *DeleteSnatEntryResponse {
	s.Body = v
	return s
}

func (s *DeleteSnatEntryResponse) Validate() error {
	return dara.Validate(s)
}

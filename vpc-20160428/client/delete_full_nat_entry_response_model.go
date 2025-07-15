// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFullNatEntryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteFullNatEntryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteFullNatEntryResponse
	GetStatusCode() *int32
	SetBody(v *DeleteFullNatEntryResponseBody) *DeleteFullNatEntryResponse
	GetBody() *DeleteFullNatEntryResponseBody
}

type DeleteFullNatEntryResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFullNatEntryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFullNatEntryResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteFullNatEntryResponse) GoString() string {
	return s.String()
}

func (s *DeleteFullNatEntryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteFullNatEntryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteFullNatEntryResponse) GetBody() *DeleteFullNatEntryResponseBody {
	return s.Body
}

func (s *DeleteFullNatEntryResponse) SetHeaders(v map[string]*string) *DeleteFullNatEntryResponse {
	s.Headers = v
	return s
}

func (s *DeleteFullNatEntryResponse) SetStatusCode(v int32) *DeleteFullNatEntryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFullNatEntryResponse) SetBody(v *DeleteFullNatEntryResponseBody) *DeleteFullNatEntryResponse {
	s.Body = v
	return s
}

func (s *DeleteFullNatEntryResponse) Validate() error {
	return dara.Validate(s)
}

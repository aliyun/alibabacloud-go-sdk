// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDnatEntryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDnatEntryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDnatEntryResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDnatEntryResponseBody) *DeleteDnatEntryResponse
	GetBody() *DeleteDnatEntryResponseBody
}

type DeleteDnatEntryResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDnatEntryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDnatEntryResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDnatEntryResponse) GoString() string {
	return s.String()
}

func (s *DeleteDnatEntryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDnatEntryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDnatEntryResponse) GetBody() *DeleteDnatEntryResponseBody {
	return s.Body
}

func (s *DeleteDnatEntryResponse) SetHeaders(v map[string]*string) *DeleteDnatEntryResponse {
	s.Headers = v
	return s
}

func (s *DeleteDnatEntryResponse) SetStatusCode(v int32) *DeleteDnatEntryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDnatEntryResponse) SetBody(v *DeleteDnatEntryResponseBody) *DeleteDnatEntryResponse {
	s.Body = v
	return s
}

func (s *DeleteDnatEntryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

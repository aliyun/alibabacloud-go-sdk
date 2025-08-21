// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEnsRouteEntryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEnsRouteEntryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEnsRouteEntryResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEnsRouteEntryResponseBody) *DeleteEnsRouteEntryResponse
	GetBody() *DeleteEnsRouteEntryResponseBody
}

type DeleteEnsRouteEntryResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEnsRouteEntryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEnsRouteEntryResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEnsRouteEntryResponse) GoString() string {
	return s.String()
}

func (s *DeleteEnsRouteEntryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEnsRouteEntryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEnsRouteEntryResponse) GetBody() *DeleteEnsRouteEntryResponseBody {
	return s.Body
}

func (s *DeleteEnsRouteEntryResponse) SetHeaders(v map[string]*string) *DeleteEnsRouteEntryResponse {
	s.Headers = v
	return s
}

func (s *DeleteEnsRouteEntryResponse) SetStatusCode(v int32) *DeleteEnsRouteEntryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEnsRouteEntryResponse) SetBody(v *DeleteEnsRouteEntryResponseBody) *DeleteEnsRouteEntryResponse {
	s.Body = v
	return s
}

func (s *DeleteEnsRouteEntryResponse) Validate() error {
	return dara.Validate(s)
}

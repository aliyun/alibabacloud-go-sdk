// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDentryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDentryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDentryResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDentryResponseBody) *DeleteDentryResponse
	GetBody() *DeleteDentryResponseBody
}

type DeleteDentryResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDentryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDentryResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDentryResponse) GoString() string {
	return s.String()
}

func (s *DeleteDentryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDentryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDentryResponse) GetBody() *DeleteDentryResponseBody {
	return s.Body
}

func (s *DeleteDentryResponse) SetHeaders(v map[string]*string) *DeleteDentryResponse {
	s.Headers = v
	return s
}

func (s *DeleteDentryResponse) SetStatusCode(v int32) *DeleteDentryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDentryResponse) SetBody(v *DeleteDentryResponseBody) *DeleteDentryResponse {
	s.Body = v
	return s
}

func (s *DeleteDentryResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSuspEventNodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSuspEventNodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSuspEventNodeResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSuspEventNodeResponseBody) *DeleteSuspEventNodeResponse
	GetBody() *DeleteSuspEventNodeResponseBody
}

type DeleteSuspEventNodeResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSuspEventNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSuspEventNodeResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSuspEventNodeResponse) GoString() string {
	return s.String()
}

func (s *DeleteSuspEventNodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSuspEventNodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSuspEventNodeResponse) GetBody() *DeleteSuspEventNodeResponseBody {
	return s.Body
}

func (s *DeleteSuspEventNodeResponse) SetHeaders(v map[string]*string) *DeleteSuspEventNodeResponse {
	s.Headers = v
	return s
}

func (s *DeleteSuspEventNodeResponse) SetStatusCode(v int32) *DeleteSuspEventNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSuspEventNodeResponse) SetBody(v *DeleteSuspEventNodeResponseBody) *DeleteSuspEventNodeResponse {
	s.Body = v
	return s
}

func (s *DeleteSuspEventNodeResponse) Validate() error {
	return dara.Validate(s)
}

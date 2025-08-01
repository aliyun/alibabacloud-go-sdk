// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteServiceSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteServiceSourceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteServiceSourceResponseBody) *DeleteServiceSourceResponse
	GetBody() *DeleteServiceSourceResponseBody
}

type DeleteServiceSourceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteServiceSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteServiceSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceSourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteServiceSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteServiceSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteServiceSourceResponse) GetBody() *DeleteServiceSourceResponseBody {
	return s.Body
}

func (s *DeleteServiceSourceResponse) SetHeaders(v map[string]*string) *DeleteServiceSourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteServiceSourceResponse) SetStatusCode(v int32) *DeleteServiceSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServiceSourceResponse) SetBody(v *DeleteServiceSourceResponseBody) *DeleteServiceSourceResponse {
	s.Body = v
	return s
}

func (s *DeleteServiceSourceResponse) Validate() error {
	return dara.Validate(s)
}

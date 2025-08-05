// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHanaInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteHanaInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteHanaInstanceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteHanaInstanceResponseBody) *DeleteHanaInstanceResponse
	GetBody() *DeleteHanaInstanceResponseBody
}

type DeleteHanaInstanceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHanaInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHanaInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteHanaInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteHanaInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteHanaInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteHanaInstanceResponse) GetBody() *DeleteHanaInstanceResponseBody {
	return s.Body
}

func (s *DeleteHanaInstanceResponse) SetHeaders(v map[string]*string) *DeleteHanaInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteHanaInstanceResponse) SetStatusCode(v int32) *DeleteHanaInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHanaInstanceResponse) SetBody(v *DeleteHanaInstanceResponseBody) *DeleteHanaInstanceResponse {
	s.Body = v
	return s
}

func (s *DeleteHanaInstanceResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSubnetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSubnetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSubnetResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSubnetResponseBody) *DeleteSubnetResponse
	GetBody() *DeleteSubnetResponseBody
}

type DeleteSubnetResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSubnetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSubnetResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSubnetResponse) GoString() string {
	return s.String()
}

func (s *DeleteSubnetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSubnetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSubnetResponse) GetBody() *DeleteSubnetResponseBody {
	return s.Body
}

func (s *DeleteSubnetResponse) SetHeaders(v map[string]*string) *DeleteSubnetResponse {
	s.Headers = v
	return s
}

func (s *DeleteSubnetResponse) SetStatusCode(v int32) *DeleteSubnetResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSubnetResponse) SetBody(v *DeleteSubnetResponseBody) *DeleteSubnetResponse {
	s.Body = v
	return s
}

func (s *DeleteSubnetResponse) Validate() error {
	return dara.Validate(s)
}

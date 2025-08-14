// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFieldResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteFieldResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteFieldResponse
	GetStatusCode() *int32
	SetBody(v *DeleteFieldResponseBody) *DeleteFieldResponse
	GetBody() *DeleteFieldResponseBody
}

type DeleteFieldResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFieldResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFieldResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteFieldResponse) GoString() string {
	return s.String()
}

func (s *DeleteFieldResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteFieldResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteFieldResponse) GetBody() *DeleteFieldResponseBody {
	return s.Body
}

func (s *DeleteFieldResponse) SetHeaders(v map[string]*string) *DeleteFieldResponse {
	s.Headers = v
	return s
}

func (s *DeleteFieldResponse) SetStatusCode(v int32) *DeleteFieldResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFieldResponse) SetBody(v *DeleteFieldResponseBody) *DeleteFieldResponse {
	s.Body = v
	return s
}

func (s *DeleteFieldResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFromMetaCategoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteFromMetaCategoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteFromMetaCategoryResponse
	GetStatusCode() *int32
	SetBody(v *DeleteFromMetaCategoryResponseBody) *DeleteFromMetaCategoryResponse
	GetBody() *DeleteFromMetaCategoryResponseBody
}

type DeleteFromMetaCategoryResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFromMetaCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFromMetaCategoryResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteFromMetaCategoryResponse) GoString() string {
	return s.String()
}

func (s *DeleteFromMetaCategoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteFromMetaCategoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteFromMetaCategoryResponse) GetBody() *DeleteFromMetaCategoryResponseBody {
	return s.Body
}

func (s *DeleteFromMetaCategoryResponse) SetHeaders(v map[string]*string) *DeleteFromMetaCategoryResponse {
	s.Headers = v
	return s
}

func (s *DeleteFromMetaCategoryResponse) SetStatusCode(v int32) *DeleteFromMetaCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFromMetaCategoryResponse) SetBody(v *DeleteFromMetaCategoryResponseBody) *DeleteFromMetaCategoryResponse {
	s.Body = v
	return s
}

func (s *DeleteFromMetaCategoryResponse) Validate() error {
	return dara.Validate(s)
}

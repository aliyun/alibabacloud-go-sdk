// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVectorIndexResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVectorIndexResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVectorIndexResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVectorIndexResponseBody) *DeleteVectorIndexResponse
	GetBody() *DeleteVectorIndexResponseBody
}

type DeleteVectorIndexResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVectorIndexResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVectorIndexResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVectorIndexResponse) GoString() string {
	return s.String()
}

func (s *DeleteVectorIndexResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVectorIndexResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVectorIndexResponse) GetBody() *DeleteVectorIndexResponseBody {
	return s.Body
}

func (s *DeleteVectorIndexResponse) SetHeaders(v map[string]*string) *DeleteVectorIndexResponse {
	s.Headers = v
	return s
}

func (s *DeleteVectorIndexResponse) SetStatusCode(v int32) *DeleteVectorIndexResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVectorIndexResponse) SetBody(v *DeleteVectorIndexResponseBody) *DeleteVectorIndexResponse {
	s.Body = v
	return s
}

func (s *DeleteVectorIndexResponse) Validate() error {
	return dara.Validate(s)
}

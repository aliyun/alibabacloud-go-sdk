// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEditingProjectMaterialsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEditingProjectMaterialsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEditingProjectMaterialsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEditingProjectMaterialsResponseBody) *DeleteEditingProjectMaterialsResponse
	GetBody() *DeleteEditingProjectMaterialsResponseBody
}

type DeleteEditingProjectMaterialsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEditingProjectMaterialsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEditingProjectMaterialsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEditingProjectMaterialsResponse) GoString() string {
	return s.String()
}

func (s *DeleteEditingProjectMaterialsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEditingProjectMaterialsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEditingProjectMaterialsResponse) GetBody() *DeleteEditingProjectMaterialsResponseBody {
	return s.Body
}

func (s *DeleteEditingProjectMaterialsResponse) SetHeaders(v map[string]*string) *DeleteEditingProjectMaterialsResponse {
	s.Headers = v
	return s
}

func (s *DeleteEditingProjectMaterialsResponse) SetStatusCode(v int32) *DeleteEditingProjectMaterialsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEditingProjectMaterialsResponse) SetBody(v *DeleteEditingProjectMaterialsResponseBody) *DeleteEditingProjectMaterialsResponse {
	s.Body = v
	return s
}

func (s *DeleteEditingProjectMaterialsResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetEditingProjectMaterialsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetEditingProjectMaterialsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetEditingProjectMaterialsResponse
	GetStatusCode() *int32
	SetBody(v *SetEditingProjectMaterialsResponseBody) *SetEditingProjectMaterialsResponse
	GetBody() *SetEditingProjectMaterialsResponseBody
}

type SetEditingProjectMaterialsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetEditingProjectMaterialsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetEditingProjectMaterialsResponse) String() string {
	return dara.Prettify(s)
}

func (s SetEditingProjectMaterialsResponse) GoString() string {
	return s.String()
}

func (s *SetEditingProjectMaterialsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetEditingProjectMaterialsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetEditingProjectMaterialsResponse) GetBody() *SetEditingProjectMaterialsResponseBody {
	return s.Body
}

func (s *SetEditingProjectMaterialsResponse) SetHeaders(v map[string]*string) *SetEditingProjectMaterialsResponse {
	s.Headers = v
	return s
}

func (s *SetEditingProjectMaterialsResponse) SetStatusCode(v int32) *SetEditingProjectMaterialsResponse {
	s.StatusCode = &v
	return s
}

func (s *SetEditingProjectMaterialsResponse) SetBody(v *SetEditingProjectMaterialsResponseBody) *SetEditingProjectMaterialsResponse {
	s.Body = v
	return s
}

func (s *SetEditingProjectMaterialsResponse) Validate() error {
	return dara.Validate(s)
}

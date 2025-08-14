// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEditingProjectMaterialsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEditingProjectMaterialsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEditingProjectMaterialsResponse
	GetStatusCode() *int32
	SetBody(v *GetEditingProjectMaterialsResponseBody) *GetEditingProjectMaterialsResponse
	GetBody() *GetEditingProjectMaterialsResponseBody
}

type GetEditingProjectMaterialsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEditingProjectMaterialsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEditingProjectMaterialsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEditingProjectMaterialsResponse) GoString() string {
	return s.String()
}

func (s *GetEditingProjectMaterialsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEditingProjectMaterialsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEditingProjectMaterialsResponse) GetBody() *GetEditingProjectMaterialsResponseBody {
	return s.Body
}

func (s *GetEditingProjectMaterialsResponse) SetHeaders(v map[string]*string) *GetEditingProjectMaterialsResponse {
	s.Headers = v
	return s
}

func (s *GetEditingProjectMaterialsResponse) SetStatusCode(v int32) *GetEditingProjectMaterialsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEditingProjectMaterialsResponse) SetBody(v *GetEditingProjectMaterialsResponseBody) *GetEditingProjectMaterialsResponse {
	s.Body = v
	return s
}

func (s *GetEditingProjectMaterialsResponse) Validate() error {
	return dara.Validate(s)
}

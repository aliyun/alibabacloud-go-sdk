// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddEditingProjectMaterialsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddEditingProjectMaterialsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddEditingProjectMaterialsResponse
	GetStatusCode() *int32
	SetBody(v *AddEditingProjectMaterialsResponseBody) *AddEditingProjectMaterialsResponse
	GetBody() *AddEditingProjectMaterialsResponseBody
}

type AddEditingProjectMaterialsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddEditingProjectMaterialsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddEditingProjectMaterialsResponse) String() string {
	return dara.Prettify(s)
}

func (s AddEditingProjectMaterialsResponse) GoString() string {
	return s.String()
}

func (s *AddEditingProjectMaterialsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddEditingProjectMaterialsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddEditingProjectMaterialsResponse) GetBody() *AddEditingProjectMaterialsResponseBody {
	return s.Body
}

func (s *AddEditingProjectMaterialsResponse) SetHeaders(v map[string]*string) *AddEditingProjectMaterialsResponse {
	s.Headers = v
	return s
}

func (s *AddEditingProjectMaterialsResponse) SetStatusCode(v int32) *AddEditingProjectMaterialsResponse {
	s.StatusCode = &v
	return s
}

func (s *AddEditingProjectMaterialsResponse) SetBody(v *AddEditingProjectMaterialsResponseBody) *AddEditingProjectMaterialsResponse {
	s.Body = v
	return s
}

func (s *AddEditingProjectMaterialsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

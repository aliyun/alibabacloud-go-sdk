// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProductAsAdminResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetProductAsAdminResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetProductAsAdminResponse
	GetStatusCode() *int32
	SetBody(v *GetProductAsAdminResponseBody) *GetProductAsAdminResponse
	GetBody() *GetProductAsAdminResponseBody
}

type GetProductAsAdminResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProductAsAdminResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProductAsAdminResponse) String() string {
	return dara.Prettify(s)
}

func (s GetProductAsAdminResponse) GoString() string {
	return s.String()
}

func (s *GetProductAsAdminResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetProductAsAdminResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetProductAsAdminResponse) GetBody() *GetProductAsAdminResponseBody {
	return s.Body
}

func (s *GetProductAsAdminResponse) SetHeaders(v map[string]*string) *GetProductAsAdminResponse {
	s.Headers = v
	return s
}

func (s *GetProductAsAdminResponse) SetStatusCode(v int32) *GetProductAsAdminResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProductAsAdminResponse) SetBody(v *GetProductAsAdminResponseBody) *GetProductAsAdminResponse {
	s.Body = v
	return s
}

func (s *GetProductAsAdminResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

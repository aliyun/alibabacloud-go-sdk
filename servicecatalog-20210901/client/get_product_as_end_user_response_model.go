// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProductAsEndUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetProductAsEndUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetProductAsEndUserResponse
	GetStatusCode() *int32
	SetBody(v *GetProductAsEndUserResponseBody) *GetProductAsEndUserResponse
	GetBody() *GetProductAsEndUserResponseBody
}

type GetProductAsEndUserResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProductAsEndUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProductAsEndUserResponse) String() string {
	return dara.Prettify(s)
}

func (s GetProductAsEndUserResponse) GoString() string {
	return s.String()
}

func (s *GetProductAsEndUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetProductAsEndUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetProductAsEndUserResponse) GetBody() *GetProductAsEndUserResponseBody {
	return s.Body
}

func (s *GetProductAsEndUserResponse) SetHeaders(v map[string]*string) *GetProductAsEndUserResponse {
	s.Headers = v
	return s
}

func (s *GetProductAsEndUserResponse) SetStatusCode(v int32) *GetProductAsEndUserResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProductAsEndUserResponse) SetBody(v *GetProductAsEndUserResponseBody) *GetProductAsEndUserResponse {
	s.Body = v
	return s
}

func (s *GetProductAsEndUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

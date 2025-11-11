// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMaterialByIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMaterialByIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMaterialByIdResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMaterialByIdResponseBody) *DeleteMaterialByIdResponse
	GetBody() *DeleteMaterialByIdResponseBody
}

type DeleteMaterialByIdResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMaterialByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMaterialByIdResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMaterialByIdResponse) GoString() string {
	return s.String()
}

func (s *DeleteMaterialByIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMaterialByIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMaterialByIdResponse) GetBody() *DeleteMaterialByIdResponseBody {
	return s.Body
}

func (s *DeleteMaterialByIdResponse) SetHeaders(v map[string]*string) *DeleteMaterialByIdResponse {
	s.Headers = v
	return s
}

func (s *DeleteMaterialByIdResponse) SetStatusCode(v int32) *DeleteMaterialByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMaterialByIdResponse) SetBody(v *DeleteMaterialByIdResponseBody) *DeleteMaterialByIdResponse {
	s.Body = v
	return s
}

func (s *DeleteMaterialByIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

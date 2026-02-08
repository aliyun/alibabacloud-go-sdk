// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMaterialDirectoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMaterialDirectoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMaterialDirectoryResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMaterialDirectoryResponseBody) *DeleteMaterialDirectoryResponse
	GetBody() *DeleteMaterialDirectoryResponseBody
}

type DeleteMaterialDirectoryResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMaterialDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMaterialDirectoryResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMaterialDirectoryResponse) GoString() string {
	return s.String()
}

func (s *DeleteMaterialDirectoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMaterialDirectoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMaterialDirectoryResponse) GetBody() *DeleteMaterialDirectoryResponseBody {
	return s.Body
}

func (s *DeleteMaterialDirectoryResponse) SetHeaders(v map[string]*string) *DeleteMaterialDirectoryResponse {
	s.Headers = v
	return s
}

func (s *DeleteMaterialDirectoryResponse) SetStatusCode(v int32) *DeleteMaterialDirectoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMaterialDirectoryResponse) SetBody(v *DeleteMaterialDirectoryResponseBody) *DeleteMaterialDirectoryResponse {
	s.Body = v
	return s
}

func (s *DeleteMaterialDirectoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

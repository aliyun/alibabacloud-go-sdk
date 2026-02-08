// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMaterialTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMaterialTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMaterialTaskResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMaterialTaskResponseBody) *DeleteMaterialTaskResponse
	GetBody() *DeleteMaterialTaskResponseBody
}

type DeleteMaterialTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMaterialTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMaterialTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMaterialTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteMaterialTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMaterialTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMaterialTaskResponse) GetBody() *DeleteMaterialTaskResponseBody {
	return s.Body
}

func (s *DeleteMaterialTaskResponse) SetHeaders(v map[string]*string) *DeleteMaterialTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteMaterialTaskResponse) SetStatusCode(v int32) *DeleteMaterialTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMaterialTaskResponse) SetBody(v *DeleteMaterialTaskResponseBody) *DeleteMaterialTaskResponse {
	s.Body = v
	return s
}

func (s *DeleteMaterialTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

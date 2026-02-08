// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMaterialDirectoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMaterialDirectoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMaterialDirectoryResponse
	GetStatusCode() *int32
	SetBody(v *CreateMaterialDirectoryResponseBody) *CreateMaterialDirectoryResponse
	GetBody() *CreateMaterialDirectoryResponseBody
}

type CreateMaterialDirectoryResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMaterialDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMaterialDirectoryResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMaterialDirectoryResponse) GoString() string {
	return s.String()
}

func (s *CreateMaterialDirectoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMaterialDirectoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMaterialDirectoryResponse) GetBody() *CreateMaterialDirectoryResponseBody {
	return s.Body
}

func (s *CreateMaterialDirectoryResponse) SetHeaders(v map[string]*string) *CreateMaterialDirectoryResponse {
	s.Headers = v
	return s
}

func (s *CreateMaterialDirectoryResponse) SetStatusCode(v int32) *CreateMaterialDirectoryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMaterialDirectoryResponse) SetBody(v *CreateMaterialDirectoryResponseBody) *CreateMaterialDirectoryResponse {
	s.Body = v
	return s
}

func (s *CreateMaterialDirectoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

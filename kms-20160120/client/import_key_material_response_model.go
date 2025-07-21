// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportKeyMaterialResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImportKeyMaterialResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImportKeyMaterialResponse
	GetStatusCode() *int32
	SetBody(v *ImportKeyMaterialResponseBody) *ImportKeyMaterialResponse
	GetBody() *ImportKeyMaterialResponseBody
}

type ImportKeyMaterialResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportKeyMaterialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportKeyMaterialResponse) String() string {
	return dara.Prettify(s)
}

func (s ImportKeyMaterialResponse) GoString() string {
	return s.String()
}

func (s *ImportKeyMaterialResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImportKeyMaterialResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImportKeyMaterialResponse) GetBody() *ImportKeyMaterialResponseBody {
	return s.Body
}

func (s *ImportKeyMaterialResponse) SetHeaders(v map[string]*string) *ImportKeyMaterialResponse {
	s.Headers = v
	return s
}

func (s *ImportKeyMaterialResponse) SetStatusCode(v int32) *ImportKeyMaterialResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportKeyMaterialResponse) SetBody(v *ImportKeyMaterialResponseBody) *ImportKeyMaterialResponse {
	s.Body = v
	return s
}

func (s *ImportKeyMaterialResponse) Validate() error {
	return dara.Validate(s)
}

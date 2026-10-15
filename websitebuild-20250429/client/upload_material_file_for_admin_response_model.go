// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadMaterialFileForAdminResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadMaterialFileForAdminResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadMaterialFileForAdminResponse
	GetStatusCode() *int32
	SetBody(v *UploadMaterialFileForAdminResponseBody) *UploadMaterialFileForAdminResponse
	GetBody() *UploadMaterialFileForAdminResponseBody
}

type UploadMaterialFileForAdminResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadMaterialFileForAdminResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadMaterialFileForAdminResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadMaterialFileForAdminResponse) GoString() string {
	return s.String()
}

func (s *UploadMaterialFileForAdminResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadMaterialFileForAdminResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadMaterialFileForAdminResponse) GetBody() *UploadMaterialFileForAdminResponseBody {
	return s.Body
}

func (s *UploadMaterialFileForAdminResponse) SetHeaders(v map[string]*string) *UploadMaterialFileForAdminResponse {
	s.Headers = v
	return s
}

func (s *UploadMaterialFileForAdminResponse) SetStatusCode(v int32) *UploadMaterialFileForAdminResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadMaterialFileForAdminResponse) SetBody(v *UploadMaterialFileForAdminResponseBody) *UploadMaterialFileForAdminResponse {
	s.Body = v
	return s
}

func (s *UploadMaterialFileForAdminResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

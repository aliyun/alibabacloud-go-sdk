// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadMaterialFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadMaterialFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadMaterialFileResponse
	GetStatusCode() *int32
	SetBody(v *UploadMaterialFileResponseBody) *UploadMaterialFileResponse
	GetBody() *UploadMaterialFileResponseBody
}

type UploadMaterialFileResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadMaterialFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadMaterialFileResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadMaterialFileResponse) GoString() string {
	return s.String()
}

func (s *UploadMaterialFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadMaterialFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadMaterialFileResponse) GetBody() *UploadMaterialFileResponseBody {
	return s.Body
}

func (s *UploadMaterialFileResponse) SetHeaders(v map[string]*string) *UploadMaterialFileResponse {
	s.Headers = v
	return s
}

func (s *UploadMaterialFileResponse) SetStatusCode(v int32) *UploadMaterialFileResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadMaterialFileResponse) SetBody(v *UploadMaterialFileResponseBody) *UploadMaterialFileResponse {
	s.Body = v
	return s
}

func (s *UploadMaterialFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

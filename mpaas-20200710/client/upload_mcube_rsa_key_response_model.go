// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadMcubeRsaKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadMcubeRsaKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadMcubeRsaKeyResponse
	GetStatusCode() *int32
	SetBody(v *UploadMcubeRsaKeyResponseBody) *UploadMcubeRsaKeyResponse
	GetBody() *UploadMcubeRsaKeyResponseBody
}

type UploadMcubeRsaKeyResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadMcubeRsaKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadMcubeRsaKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadMcubeRsaKeyResponse) GoString() string {
	return s.String()
}

func (s *UploadMcubeRsaKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadMcubeRsaKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadMcubeRsaKeyResponse) GetBody() *UploadMcubeRsaKeyResponseBody {
	return s.Body
}

func (s *UploadMcubeRsaKeyResponse) SetHeaders(v map[string]*string) *UploadMcubeRsaKeyResponse {
	s.Headers = v
	return s
}

func (s *UploadMcubeRsaKeyResponse) SetStatusCode(v int32) *UploadMcubeRsaKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadMcubeRsaKeyResponse) SetBody(v *UploadMcubeRsaKeyResponseBody) *UploadMcubeRsaKeyResponse {
	s.Body = v
	return s
}

func (s *UploadMcubeRsaKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMcpCodePackageUploadUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMcpCodePackageUploadUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMcpCodePackageUploadUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetMcpCodePackageUploadUrlResponseBody) *GetMcpCodePackageUploadUrlResponse
	GetBody() *GetMcpCodePackageUploadUrlResponseBody
}

type GetMcpCodePackageUploadUrlResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMcpCodePackageUploadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMcpCodePackageUploadUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMcpCodePackageUploadUrlResponse) GoString() string {
	return s.String()
}

func (s *GetMcpCodePackageUploadUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMcpCodePackageUploadUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMcpCodePackageUploadUrlResponse) GetBody() *GetMcpCodePackageUploadUrlResponseBody {
	return s.Body
}

func (s *GetMcpCodePackageUploadUrlResponse) SetHeaders(v map[string]*string) *GetMcpCodePackageUploadUrlResponse {
	s.Headers = v
	return s
}

func (s *GetMcpCodePackageUploadUrlResponse) SetStatusCode(v int32) *GetMcpCodePackageUploadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMcpCodePackageUploadUrlResponse) SetBody(v *GetMcpCodePackageUploadUrlResponseBody) *GetMcpCodePackageUploadUrlResponse {
	s.Body = v
	return s
}

func (s *GetMcpCodePackageUploadUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

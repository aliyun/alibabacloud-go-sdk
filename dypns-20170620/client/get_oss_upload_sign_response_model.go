// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOssUploadSignResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOssUploadSignResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOssUploadSignResponse
	GetStatusCode() *int32
	SetBody(v *GetOssUploadSignResponseBody) *GetOssUploadSignResponse
	GetBody() *GetOssUploadSignResponseBody
}

type GetOssUploadSignResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOssUploadSignResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOssUploadSignResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOssUploadSignResponse) GoString() string {
	return s.String()
}

func (s *GetOssUploadSignResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOssUploadSignResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOssUploadSignResponse) GetBody() *GetOssUploadSignResponseBody {
	return s.Body
}

func (s *GetOssUploadSignResponse) SetHeaders(v map[string]*string) *GetOssUploadSignResponse {
	s.Headers = v
	return s
}

func (s *GetOssUploadSignResponse) SetStatusCode(v int32) *GetOssUploadSignResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOssUploadSignResponse) SetBody(v *GetOssUploadSignResponseBody) *GetOssUploadSignResponse {
	s.Body = v
	return s
}

func (s *GetOssUploadSignResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

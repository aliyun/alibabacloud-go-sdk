// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOssUploadTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOssUploadTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOssUploadTokenResponse
	GetStatusCode() *int32
	SetBody(v *GetOssUploadTokenResult) *GetOssUploadTokenResponse
	GetBody() *GetOssUploadTokenResult
}

type GetOssUploadTokenResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOssUploadTokenResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOssUploadTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOssUploadTokenResponse) GoString() string {
	return s.String()
}

func (s *GetOssUploadTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOssUploadTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOssUploadTokenResponse) GetBody() *GetOssUploadTokenResult {
	return s.Body
}

func (s *GetOssUploadTokenResponse) SetHeaders(v map[string]*string) *GetOssUploadTokenResponse {
	s.Headers = v
	return s
}

func (s *GetOssUploadTokenResponse) SetStatusCode(v int32) *GetOssUploadTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOssUploadTokenResponse) SetBody(v *GetOssUploadTokenResult) *GetOssUploadTokenResponse {
	s.Body = v
	return s
}

func (s *GetOssUploadTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

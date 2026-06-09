// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOssUploadPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOssUploadPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOssUploadPolicyResponse
	GetStatusCode() *int32
	SetBody(v *GetOssUploadPolicyResponseBody) *GetOssUploadPolicyResponse
	GetBody() *GetOssUploadPolicyResponseBody
}

type GetOssUploadPolicyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOssUploadPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOssUploadPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOssUploadPolicyResponse) GoString() string {
	return s.String()
}

func (s *GetOssUploadPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOssUploadPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOssUploadPolicyResponse) GetBody() *GetOssUploadPolicyResponseBody {
	return s.Body
}

func (s *GetOssUploadPolicyResponse) SetHeaders(v map[string]*string) *GetOssUploadPolicyResponse {
	s.Headers = v
	return s
}

func (s *GetOssUploadPolicyResponse) SetStatusCode(v int32) *GetOssUploadPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOssUploadPolicyResponse) SetBody(v *GetOssUploadPolicyResponseBody) *GetOssUploadPolicyResponse {
	s.Body = v
	return s
}

func (s *GetOssUploadPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

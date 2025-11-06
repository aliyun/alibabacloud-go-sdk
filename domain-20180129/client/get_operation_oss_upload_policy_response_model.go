// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOperationOssUploadPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOperationOssUploadPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOperationOssUploadPolicyResponse
	GetStatusCode() *int32
	SetBody(v *GetOperationOssUploadPolicyResponseBody) *GetOperationOssUploadPolicyResponse
	GetBody() *GetOperationOssUploadPolicyResponseBody
}

type GetOperationOssUploadPolicyResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOperationOssUploadPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOperationOssUploadPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOperationOssUploadPolicyResponse) GoString() string {
	return s.String()
}

func (s *GetOperationOssUploadPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOperationOssUploadPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOperationOssUploadPolicyResponse) GetBody() *GetOperationOssUploadPolicyResponseBody {
	return s.Body
}

func (s *GetOperationOssUploadPolicyResponse) SetHeaders(v map[string]*string) *GetOperationOssUploadPolicyResponse {
	s.Headers = v
	return s
}

func (s *GetOperationOssUploadPolicyResponse) SetStatusCode(v int32) *GetOperationOssUploadPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOperationOssUploadPolicyResponse) SetBody(v *GetOperationOssUploadPolicyResponseBody) *GetOperationOssUploadPolicyResponse {
	s.Body = v
	return s
}

func (s *GetOperationOssUploadPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateUploadFilePolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateUploadFilePolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateUploadFilePolicyResponse
	GetStatusCode() *int32
	SetBody(v *GenerateUploadFilePolicyResponseBody) *GenerateUploadFilePolicyResponse
	GetBody() *GenerateUploadFilePolicyResponseBody
}

type GenerateUploadFilePolicyResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateUploadFilePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateUploadFilePolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateUploadFilePolicyResponse) GoString() string {
	return s.String()
}

func (s *GenerateUploadFilePolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateUploadFilePolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateUploadFilePolicyResponse) GetBody() *GenerateUploadFilePolicyResponseBody {
	return s.Body
}

func (s *GenerateUploadFilePolicyResponse) SetHeaders(v map[string]*string) *GenerateUploadFilePolicyResponse {
	s.Headers = v
	return s
}

func (s *GenerateUploadFilePolicyResponse) SetStatusCode(v int32) *GenerateUploadFilePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateUploadFilePolicyResponse) SetBody(v *GenerateUploadFilePolicyResponseBody) *GenerateUploadFilePolicyResponse {
	s.Body = v
	return s
}

func (s *GenerateUploadFilePolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

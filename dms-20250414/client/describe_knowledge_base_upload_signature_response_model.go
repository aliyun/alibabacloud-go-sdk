// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKnowledgeBaseUploadSignatureResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeKnowledgeBaseUploadSignatureResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeKnowledgeBaseUploadSignatureResponse
	GetStatusCode() *int32
	SetBody(v *DescribeKnowledgeBaseUploadSignatureResponseBody) *DescribeKnowledgeBaseUploadSignatureResponse
	GetBody() *DescribeKnowledgeBaseUploadSignatureResponseBody
}

type DescribeKnowledgeBaseUploadSignatureResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeKnowledgeBaseUploadSignatureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeKnowledgeBaseUploadSignatureResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeKnowledgeBaseUploadSignatureResponse) GoString() string {
	return s.String()
}

func (s *DescribeKnowledgeBaseUploadSignatureResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeKnowledgeBaseUploadSignatureResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeKnowledgeBaseUploadSignatureResponse) GetBody() *DescribeKnowledgeBaseUploadSignatureResponseBody {
	return s.Body
}

func (s *DescribeKnowledgeBaseUploadSignatureResponse) SetHeaders(v map[string]*string) *DescribeKnowledgeBaseUploadSignatureResponse {
	s.Headers = v
	return s
}

func (s *DescribeKnowledgeBaseUploadSignatureResponse) SetStatusCode(v int32) *DescribeKnowledgeBaseUploadSignatureResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeKnowledgeBaseUploadSignatureResponse) SetBody(v *DescribeKnowledgeBaseUploadSignatureResponseBody) *DescribeKnowledgeBaseUploadSignatureResponse {
	s.Body = v
	return s
}

func (s *DescribeKnowledgeBaseUploadSignatureResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

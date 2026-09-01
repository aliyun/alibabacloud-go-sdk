// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSkillFileUploadSignatureResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSkillFileUploadSignatureResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSkillFileUploadSignatureResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSkillFileUploadSignatureResponseBody) *DescribeSkillFileUploadSignatureResponse
	GetBody() *DescribeSkillFileUploadSignatureResponseBody
}

type DescribeSkillFileUploadSignatureResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSkillFileUploadSignatureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSkillFileUploadSignatureResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSkillFileUploadSignatureResponse) GoString() string {
	return s.String()
}

func (s *DescribeSkillFileUploadSignatureResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSkillFileUploadSignatureResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSkillFileUploadSignatureResponse) GetBody() *DescribeSkillFileUploadSignatureResponseBody {
	return s.Body
}

func (s *DescribeSkillFileUploadSignatureResponse) SetHeaders(v map[string]*string) *DescribeSkillFileUploadSignatureResponse {
	s.Headers = v
	return s
}

func (s *DescribeSkillFileUploadSignatureResponse) SetStatusCode(v int32) *DescribeSkillFileUploadSignatureResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSkillFileUploadSignatureResponse) SetBody(v *DescribeSkillFileUploadSignatureResponseBody) *DescribeSkillFileUploadSignatureResponse {
	s.Body = v
	return s
}

func (s *DescribeSkillFileUploadSignatureResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

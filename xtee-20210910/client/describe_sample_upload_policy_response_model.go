// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSampleUploadPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSampleUploadPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSampleUploadPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSampleUploadPolicyResponseBody) *DescribeSampleUploadPolicyResponse
	GetBody() *DescribeSampleUploadPolicyResponseBody
}

type DescribeSampleUploadPolicyResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSampleUploadPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSampleUploadPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSampleUploadPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeSampleUploadPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSampleUploadPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSampleUploadPolicyResponse) GetBody() *DescribeSampleUploadPolicyResponseBody {
	return s.Body
}

func (s *DescribeSampleUploadPolicyResponse) SetHeaders(v map[string]*string) *DescribeSampleUploadPolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribeSampleUploadPolicyResponse) SetStatusCode(v int32) *DescribeSampleUploadPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSampleUploadPolicyResponse) SetBody(v *DescribeSampleUploadPolicyResponseBody) *DescribeSampleUploadPolicyResponse {
	s.Body = v
	return s
}

func (s *DescribeSampleUploadPolicyResponse) Validate() error {
	return dara.Validate(s)
}

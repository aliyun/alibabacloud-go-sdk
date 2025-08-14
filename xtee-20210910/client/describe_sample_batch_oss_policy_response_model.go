// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSampleBatchOssPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSampleBatchOssPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSampleBatchOssPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSampleBatchOssPolicyResponseBody) *DescribeSampleBatchOssPolicyResponse
	GetBody() *DescribeSampleBatchOssPolicyResponseBody
}

type DescribeSampleBatchOssPolicyResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSampleBatchOssPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSampleBatchOssPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSampleBatchOssPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeSampleBatchOssPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSampleBatchOssPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSampleBatchOssPolicyResponse) GetBody() *DescribeSampleBatchOssPolicyResponseBody {
	return s.Body
}

func (s *DescribeSampleBatchOssPolicyResponse) SetHeaders(v map[string]*string) *DescribeSampleBatchOssPolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribeSampleBatchOssPolicyResponse) SetStatusCode(v int32) *DescribeSampleBatchOssPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSampleBatchOssPolicyResponse) SetBody(v *DescribeSampleBatchOssPolicyResponseBody) *DescribeSampleBatchOssPolicyResponse {
	s.Body = v
	return s
}

func (s *DescribeSampleBatchOssPolicyResponse) Validate() error {
	return dara.Validate(s)
}

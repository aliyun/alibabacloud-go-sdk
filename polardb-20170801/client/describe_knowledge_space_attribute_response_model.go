// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKnowledgeSpaceAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeKnowledgeSpaceAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeKnowledgeSpaceAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeKnowledgeSpaceAttributeResponseBody) *DescribeKnowledgeSpaceAttributeResponse
	GetBody() *DescribeKnowledgeSpaceAttributeResponseBody
}

type DescribeKnowledgeSpaceAttributeResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeKnowledgeSpaceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeKnowledgeSpaceAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeKnowledgeSpaceAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeKnowledgeSpaceAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeKnowledgeSpaceAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeKnowledgeSpaceAttributeResponse) GetBody() *DescribeKnowledgeSpaceAttributeResponseBody {
	return s.Body
}

func (s *DescribeKnowledgeSpaceAttributeResponse) SetHeaders(v map[string]*string) *DescribeKnowledgeSpaceAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeKnowledgeSpaceAttributeResponse) SetStatusCode(v int32) *DescribeKnowledgeSpaceAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeKnowledgeSpaceAttributeResponse) SetBody(v *DescribeKnowledgeSpaceAttributeResponseBody) *DescribeKnowledgeSpaceAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeKnowledgeSpaceAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

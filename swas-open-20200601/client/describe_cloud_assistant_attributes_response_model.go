// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudAssistantAttributesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCloudAssistantAttributesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCloudAssistantAttributesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCloudAssistantAttributesResponseBody) *DescribeCloudAssistantAttributesResponse
	GetBody() *DescribeCloudAssistantAttributesResponseBody
}

type DescribeCloudAssistantAttributesResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCloudAssistantAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCloudAssistantAttributesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudAssistantAttributesResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudAssistantAttributesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCloudAssistantAttributesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCloudAssistantAttributesResponse) GetBody() *DescribeCloudAssistantAttributesResponseBody {
	return s.Body
}

func (s *DescribeCloudAssistantAttributesResponse) SetHeaders(v map[string]*string) *DescribeCloudAssistantAttributesResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudAssistantAttributesResponse) SetStatusCode(v int32) *DescribeCloudAssistantAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudAssistantAttributesResponse) SetBody(v *DescribeCloudAssistantAttributesResponseBody) *DescribeCloudAssistantAttributesResponse {
	s.Body = v
	return s
}

func (s *DescribeCloudAssistantAttributesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

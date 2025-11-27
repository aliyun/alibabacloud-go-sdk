// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCCloudAssistantStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRCCloudAssistantStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRCCloudAssistantStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRCCloudAssistantStatusResponseBody) *DescribeRCCloudAssistantStatusResponse
	GetBody() *DescribeRCCloudAssistantStatusResponseBody
}

type DescribeRCCloudAssistantStatusResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRCCloudAssistantStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRCCloudAssistantStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCCloudAssistantStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeRCCloudAssistantStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRCCloudAssistantStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRCCloudAssistantStatusResponse) GetBody() *DescribeRCCloudAssistantStatusResponseBody {
	return s.Body
}

func (s *DescribeRCCloudAssistantStatusResponse) SetHeaders(v map[string]*string) *DescribeRCCloudAssistantStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeRCCloudAssistantStatusResponse) SetStatusCode(v int32) *DescribeRCCloudAssistantStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRCCloudAssistantStatusResponse) SetBody(v *DescribeRCCloudAssistantStatusResponseBody) *DescribeRCCloudAssistantStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeRCCloudAssistantStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

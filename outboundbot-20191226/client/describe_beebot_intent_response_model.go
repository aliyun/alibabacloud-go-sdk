// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBeebotIntentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBeebotIntentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBeebotIntentResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBeebotIntentResponseBody) *DescribeBeebotIntentResponse
	GetBody() *DescribeBeebotIntentResponseBody
}

type DescribeBeebotIntentResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBeebotIntentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBeebotIntentResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBeebotIntentResponse) GoString() string {
	return s.String()
}

func (s *DescribeBeebotIntentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBeebotIntentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBeebotIntentResponse) GetBody() *DescribeBeebotIntentResponseBody {
	return s.Body
}

func (s *DescribeBeebotIntentResponse) SetHeaders(v map[string]*string) *DescribeBeebotIntentResponse {
	s.Headers = v
	return s
}

func (s *DescribeBeebotIntentResponse) SetStatusCode(v int32) *DescribeBeebotIntentResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBeebotIntentResponse) SetBody(v *DescribeBeebotIntentResponseBody) *DescribeBeebotIntentResponse {
	s.Body = v
	return s
}

func (s *DescribeBeebotIntentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

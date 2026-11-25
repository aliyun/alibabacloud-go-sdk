// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAiEngineEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAiEngineEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAiEngineEventResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAiEngineEventResponseBody) *DescribeAiEngineEventResponse
	GetBody() *DescribeAiEngineEventResponseBody
}

type DescribeAiEngineEventResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAiEngineEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAiEngineEventResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAiEngineEventResponse) GoString() string {
	return s.String()
}

func (s *DescribeAiEngineEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAiEngineEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAiEngineEventResponse) GetBody() *DescribeAiEngineEventResponseBody {
	return s.Body
}

func (s *DescribeAiEngineEventResponse) SetHeaders(v map[string]*string) *DescribeAiEngineEventResponse {
	s.Headers = v
	return s
}

func (s *DescribeAiEngineEventResponse) SetStatusCode(v int32) *DescribeAiEngineEventResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAiEngineEventResponse) SetBody(v *DescribeAiEngineEventResponseBody) *DescribeAiEngineEventResponse {
	s.Body = v
	return s
}

func (s *DescribeAiEngineEventResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

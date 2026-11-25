// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAiEngineEventNodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAiEngineEventNodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAiEngineEventNodesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAiEngineEventNodesResponseBody) *DescribeAiEngineEventNodesResponse
	GetBody() *DescribeAiEngineEventNodesResponseBody
}

type DescribeAiEngineEventNodesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAiEngineEventNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAiEngineEventNodesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAiEngineEventNodesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAiEngineEventNodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAiEngineEventNodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAiEngineEventNodesResponse) GetBody() *DescribeAiEngineEventNodesResponseBody {
	return s.Body
}

func (s *DescribeAiEngineEventNodesResponse) SetHeaders(v map[string]*string) *DescribeAiEngineEventNodesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAiEngineEventNodesResponse) SetStatusCode(v int32) *DescribeAiEngineEventNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAiEngineEventNodesResponse) SetBody(v *DescribeAiEngineEventNodesResponseBody) *DescribeAiEngineEventNodesResponse {
	s.Body = v
	return s
}

func (s *DescribeAiEngineEventNodesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

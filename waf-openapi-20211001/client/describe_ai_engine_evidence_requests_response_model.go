// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAiEngineEvidenceRequestsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAiEngineEvidenceRequestsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAiEngineEvidenceRequestsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAiEngineEvidenceRequestsResponseBody) *DescribeAiEngineEvidenceRequestsResponse
	GetBody() *DescribeAiEngineEvidenceRequestsResponseBody
}

type DescribeAiEngineEvidenceRequestsResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAiEngineEvidenceRequestsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAiEngineEvidenceRequestsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAiEngineEvidenceRequestsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAiEngineEvidenceRequestsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAiEngineEvidenceRequestsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAiEngineEvidenceRequestsResponse) GetBody() *DescribeAiEngineEvidenceRequestsResponseBody {
	return s.Body
}

func (s *DescribeAiEngineEvidenceRequestsResponse) SetHeaders(v map[string]*string) *DescribeAiEngineEvidenceRequestsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAiEngineEvidenceRequestsResponse) SetStatusCode(v int32) *DescribeAiEngineEvidenceRequestsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestsResponse) SetBody(v *DescribeAiEngineEvidenceRequestsResponseBody) *DescribeAiEngineEvidenceRequestsResponse {
	s.Body = v
	return s
}

func (s *DescribeAiEngineEvidenceRequestsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

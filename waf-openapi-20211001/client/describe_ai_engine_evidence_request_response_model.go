// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAiEngineEvidenceRequestResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAiEngineEvidenceRequestResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAiEngineEvidenceRequestResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAiEngineEvidenceRequestResponseBody) *DescribeAiEngineEvidenceRequestResponse
	GetBody() *DescribeAiEngineEvidenceRequestResponseBody
}

type DescribeAiEngineEvidenceRequestResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAiEngineEvidenceRequestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAiEngineEvidenceRequestResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAiEngineEvidenceRequestResponse) GoString() string {
	return s.String()
}

func (s *DescribeAiEngineEvidenceRequestResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAiEngineEvidenceRequestResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAiEngineEvidenceRequestResponse) GetBody() *DescribeAiEngineEvidenceRequestResponseBody {
	return s.Body
}

func (s *DescribeAiEngineEvidenceRequestResponse) SetHeaders(v map[string]*string) *DescribeAiEngineEvidenceRequestResponse {
	s.Headers = v
	return s
}

func (s *DescribeAiEngineEvidenceRequestResponse) SetStatusCode(v int32) *DescribeAiEngineEvidenceRequestResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAiEngineEvidenceRequestResponse) SetBody(v *DescribeAiEngineEvidenceRequestResponseBody) *DescribeAiEngineEvidenceRequestResponse {
	s.Body = v
	return s
}

func (s *DescribeAiEngineEvidenceRequestResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

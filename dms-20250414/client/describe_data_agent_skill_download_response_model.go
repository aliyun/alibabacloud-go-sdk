// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataAgentSkillDownloadResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDataAgentSkillDownloadResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDataAgentSkillDownloadResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDataAgentSkillDownloadResponseBody) *DescribeDataAgentSkillDownloadResponse
	GetBody() *DescribeDataAgentSkillDownloadResponseBody
}

type DescribeDataAgentSkillDownloadResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDataAgentSkillDownloadResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDataAgentSkillDownloadResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataAgentSkillDownloadResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataAgentSkillDownloadResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDataAgentSkillDownloadResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDataAgentSkillDownloadResponse) GetBody() *DescribeDataAgentSkillDownloadResponseBody {
	return s.Body
}

func (s *DescribeDataAgentSkillDownloadResponse) SetHeaders(v map[string]*string) *DescribeDataAgentSkillDownloadResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataAgentSkillDownloadResponse) SetStatusCode(v int32) *DescribeDataAgentSkillDownloadResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataAgentSkillDownloadResponse) SetBody(v *DescribeDataAgentSkillDownloadResponseBody) *DescribeDataAgentSkillDownloadResponse {
	s.Body = v
	return s
}

func (s *DescribeDataAgentSkillDownloadResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

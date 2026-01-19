// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAIDBClusterTaskLogFilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAIDBClusterTaskLogFilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAIDBClusterTaskLogFilesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAIDBClusterTaskLogFilesResponseBody) *DescribeAIDBClusterTaskLogFilesResponse
	GetBody() *DescribeAIDBClusterTaskLogFilesResponseBody
}

type DescribeAIDBClusterTaskLogFilesResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAIDBClusterTaskLogFilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAIDBClusterTaskLogFilesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIDBClusterTaskLogFilesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAIDBClusterTaskLogFilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAIDBClusterTaskLogFilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAIDBClusterTaskLogFilesResponse) GetBody() *DescribeAIDBClusterTaskLogFilesResponseBody {
	return s.Body
}

func (s *DescribeAIDBClusterTaskLogFilesResponse) SetHeaders(v map[string]*string) *DescribeAIDBClusterTaskLogFilesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAIDBClusterTaskLogFilesResponse) SetStatusCode(v int32) *DescribeAIDBClusterTaskLogFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAIDBClusterTaskLogFilesResponse) SetBody(v *DescribeAIDBClusterTaskLogFilesResponseBody) *DescribeAIDBClusterTaskLogFilesResponse {
	s.Body = v
	return s
}

func (s *DescribeAIDBClusterTaskLogFilesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

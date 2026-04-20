// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAIDBClusterDatasetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAIDBClusterDatasetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAIDBClusterDatasetsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAIDBClusterDatasetsResponseBody) *DescribeAIDBClusterDatasetsResponse
	GetBody() *DescribeAIDBClusterDatasetsResponseBody
}

type DescribeAIDBClusterDatasetsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAIDBClusterDatasetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAIDBClusterDatasetsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIDBClusterDatasetsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAIDBClusterDatasetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAIDBClusterDatasetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAIDBClusterDatasetsResponse) GetBody() *DescribeAIDBClusterDatasetsResponseBody {
	return s.Body
}

func (s *DescribeAIDBClusterDatasetsResponse) SetHeaders(v map[string]*string) *DescribeAIDBClusterDatasetsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAIDBClusterDatasetsResponse) SetStatusCode(v int32) *DescribeAIDBClusterDatasetsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAIDBClusterDatasetsResponse) SetBody(v *DescribeAIDBClusterDatasetsResponseBody) *DescribeAIDBClusterDatasetsResponse {
	s.Body = v
	return s
}

func (s *DescribeAIDBClusterDatasetsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

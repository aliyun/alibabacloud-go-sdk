// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkInstanceRelationListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNetworkInstanceRelationListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNetworkInstanceRelationListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNetworkInstanceRelationListResponseBody) *DescribeNetworkInstanceRelationListResponse
	GetBody() *DescribeNetworkInstanceRelationListResponseBody
}

type DescribeNetworkInstanceRelationListResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNetworkInstanceRelationListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNetworkInstanceRelationListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInstanceRelationListResponse) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInstanceRelationListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNetworkInstanceRelationListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNetworkInstanceRelationListResponse) GetBody() *DescribeNetworkInstanceRelationListResponseBody {
	return s.Body
}

func (s *DescribeNetworkInstanceRelationListResponse) SetHeaders(v map[string]*string) *DescribeNetworkInstanceRelationListResponse {
	s.Headers = v
	return s
}

func (s *DescribeNetworkInstanceRelationListResponse) SetStatusCode(v int32) *DescribeNetworkInstanceRelationListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNetworkInstanceRelationListResponse) SetBody(v *DescribeNetworkInstanceRelationListResponseBody) *DescribeNetworkInstanceRelationListResponse {
	s.Body = v
	return s
}

func (s *DescribeNetworkInstanceRelationListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

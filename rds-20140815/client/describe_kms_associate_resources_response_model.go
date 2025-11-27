// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKmsAssociateResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeKmsAssociateResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeKmsAssociateResourcesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeKmsAssociateResourcesResponseBody) *DescribeKmsAssociateResourcesResponse
	GetBody() *DescribeKmsAssociateResourcesResponseBody
}

type DescribeKmsAssociateResourcesResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeKmsAssociateResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeKmsAssociateResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeKmsAssociateResourcesResponse) GoString() string {
	return s.String()
}

func (s *DescribeKmsAssociateResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeKmsAssociateResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeKmsAssociateResourcesResponse) GetBody() *DescribeKmsAssociateResourcesResponseBody {
	return s.Body
}

func (s *DescribeKmsAssociateResourcesResponse) SetHeaders(v map[string]*string) *DescribeKmsAssociateResourcesResponse {
	s.Headers = v
	return s
}

func (s *DescribeKmsAssociateResourcesResponse) SetStatusCode(v int32) *DescribeKmsAssociateResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeKmsAssociateResourcesResponse) SetBody(v *DescribeKmsAssociateResourcesResponseBody) *DescribeKmsAssociateResourcesResponse {
	s.Body = v
	return s
}

func (s *DescribeKmsAssociateResourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

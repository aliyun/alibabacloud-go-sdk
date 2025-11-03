// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEcGrantRelationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEcGrantRelationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEcGrantRelationResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEcGrantRelationResponseBody) *DescribeEcGrantRelationResponse
	GetBody() *DescribeEcGrantRelationResponseBody
}

type DescribeEcGrantRelationResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEcGrantRelationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEcGrantRelationResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEcGrantRelationResponse) GoString() string {
	return s.String()
}

func (s *DescribeEcGrantRelationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEcGrantRelationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEcGrantRelationResponse) GetBody() *DescribeEcGrantRelationResponseBody {
	return s.Body
}

func (s *DescribeEcGrantRelationResponse) SetHeaders(v map[string]*string) *DescribeEcGrantRelationResponse {
	s.Headers = v
	return s
}

func (s *DescribeEcGrantRelationResponse) SetStatusCode(v int32) *DescribeEcGrantRelationResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEcGrantRelationResponse) SetBody(v *DescribeEcGrantRelationResponseBody) *DescribeEcGrantRelationResponse {
	s.Body = v
	return s
}

func (s *DescribeEcGrantRelationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

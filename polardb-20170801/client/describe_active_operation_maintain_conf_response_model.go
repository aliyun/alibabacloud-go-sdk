// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeActiveOperationMaintainConfResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeActiveOperationMaintainConfResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeActiveOperationMaintainConfResponse
	GetStatusCode() *int32
	SetBody(v *DescribeActiveOperationMaintainConfResponseBody) *DescribeActiveOperationMaintainConfResponse
	GetBody() *DescribeActiveOperationMaintainConfResponseBody
}

type DescribeActiveOperationMaintainConfResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeActiveOperationMaintainConfResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeActiveOperationMaintainConfResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeActiveOperationMaintainConfResponse) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationMaintainConfResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeActiveOperationMaintainConfResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeActiveOperationMaintainConfResponse) GetBody() *DescribeActiveOperationMaintainConfResponseBody {
	return s.Body
}

func (s *DescribeActiveOperationMaintainConfResponse) SetHeaders(v map[string]*string) *DescribeActiveOperationMaintainConfResponse {
	s.Headers = v
	return s
}

func (s *DescribeActiveOperationMaintainConfResponse) SetStatusCode(v int32) *DescribeActiveOperationMaintainConfResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeActiveOperationMaintainConfResponse) SetBody(v *DescribeActiveOperationMaintainConfResponseBody) *DescribeActiveOperationMaintainConfResponse {
	s.Body = v
	return s
}

func (s *DescribeActiveOperationMaintainConfResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

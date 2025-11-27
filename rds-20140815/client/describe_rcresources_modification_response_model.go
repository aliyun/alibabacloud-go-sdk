// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCResourcesModificationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRCResourcesModificationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRCResourcesModificationResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRCResourcesModificationResponseBody) *DescribeRCResourcesModificationResponse
	GetBody() *DescribeRCResourcesModificationResponseBody
}

type DescribeRCResourcesModificationResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRCResourcesModificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRCResourcesModificationResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCResourcesModificationResponse) GoString() string {
	return s.String()
}

func (s *DescribeRCResourcesModificationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRCResourcesModificationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRCResourcesModificationResponse) GetBody() *DescribeRCResourcesModificationResponseBody {
	return s.Body
}

func (s *DescribeRCResourcesModificationResponse) SetHeaders(v map[string]*string) *DescribeRCResourcesModificationResponse {
	s.Headers = v
	return s
}

func (s *DescribeRCResourcesModificationResponse) SetStatusCode(v int32) *DescribeRCResourcesModificationResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRCResourcesModificationResponse) SetBody(v *DescribeRCResourcesModificationResponseBody) *DescribeRCResourcesModificationResponse {
	s.Body = v
	return s
}

func (s *DescribeRCResourcesModificationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

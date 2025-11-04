// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNotificationTypesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNotificationTypesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNotificationTypesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNotificationTypesResponseBody) *DescribeNotificationTypesResponse
	GetBody() *DescribeNotificationTypesResponseBody
}

type DescribeNotificationTypesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNotificationTypesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNotificationTypesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNotificationTypesResponse) GoString() string {
	return s.String()
}

func (s *DescribeNotificationTypesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNotificationTypesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNotificationTypesResponse) GetBody() *DescribeNotificationTypesResponseBody {
	return s.Body
}

func (s *DescribeNotificationTypesResponse) SetHeaders(v map[string]*string) *DescribeNotificationTypesResponse {
	s.Headers = v
	return s
}

func (s *DescribeNotificationTypesResponse) SetStatusCode(v int32) *DescribeNotificationTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNotificationTypesResponse) SetBody(v *DescribeNotificationTypesResponseBody) *DescribeNotificationTypesResponse {
	s.Body = v
	return s
}

func (s *DescribeNotificationTypesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

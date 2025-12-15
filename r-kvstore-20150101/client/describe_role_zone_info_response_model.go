// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRoleZoneInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRoleZoneInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRoleZoneInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRoleZoneInfoResponseBody) *DescribeRoleZoneInfoResponse
	GetBody() *DescribeRoleZoneInfoResponseBody
}

type DescribeRoleZoneInfoResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRoleZoneInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRoleZoneInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRoleZoneInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeRoleZoneInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRoleZoneInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRoleZoneInfoResponse) GetBody() *DescribeRoleZoneInfoResponseBody {
	return s.Body
}

func (s *DescribeRoleZoneInfoResponse) SetHeaders(v map[string]*string) *DescribeRoleZoneInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeRoleZoneInfoResponse) SetStatusCode(v int32) *DescribeRoleZoneInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRoleZoneInfoResponse) SetBody(v *DescribeRoleZoneInfoResponseBody) *DescribeRoleZoneInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeRoleZoneInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContactListByContactGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeContactListByContactGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeContactListByContactGroupResponse
	GetStatusCode() *int32
	SetBody(v *DescribeContactListByContactGroupResponseBody) *DescribeContactListByContactGroupResponse
	GetBody() *DescribeContactListByContactGroupResponseBody
}

type DescribeContactListByContactGroupResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeContactListByContactGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeContactListByContactGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeContactListByContactGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeContactListByContactGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeContactListByContactGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeContactListByContactGroupResponse) GetBody() *DescribeContactListByContactGroupResponseBody {
	return s.Body
}

func (s *DescribeContactListByContactGroupResponse) SetHeaders(v map[string]*string) *DescribeContactListByContactGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeContactListByContactGroupResponse) SetStatusCode(v int32) *DescribeContactListByContactGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeContactListByContactGroupResponse) SetBody(v *DescribeContactListByContactGroupResponseBody) *DescribeContactListByContactGroupResponse {
	s.Body = v
	return s
}

func (s *DescribeContactListByContactGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

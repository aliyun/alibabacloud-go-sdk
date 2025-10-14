// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGtmAvailableAlertGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGtmAvailableAlertGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGtmAvailableAlertGroupResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGtmAvailableAlertGroupResponseBody) *DescribeGtmAvailableAlertGroupResponse
	GetBody() *DescribeGtmAvailableAlertGroupResponseBody
}

type DescribeGtmAvailableAlertGroupResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGtmAvailableAlertGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGtmAvailableAlertGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmAvailableAlertGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeGtmAvailableAlertGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGtmAvailableAlertGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGtmAvailableAlertGroupResponse) GetBody() *DescribeGtmAvailableAlertGroupResponseBody {
	return s.Body
}

func (s *DescribeGtmAvailableAlertGroupResponse) SetHeaders(v map[string]*string) *DescribeGtmAvailableAlertGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeGtmAvailableAlertGroupResponse) SetStatusCode(v int32) *DescribeGtmAvailableAlertGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGtmAvailableAlertGroupResponse) SetBody(v *DescribeGtmAvailableAlertGroupResponseBody) *DescribeGtmAvailableAlertGroupResponse {
	s.Body = v
	return s
}

func (s *DescribeGtmAvailableAlertGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnsGtmAvailableAlertGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDnsGtmAvailableAlertGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDnsGtmAvailableAlertGroupResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDnsGtmAvailableAlertGroupResponseBody) *DescribeDnsGtmAvailableAlertGroupResponse
	GetBody() *DescribeDnsGtmAvailableAlertGroupResponseBody
}

type DescribeDnsGtmAvailableAlertGroupResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDnsGtmAvailableAlertGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDnsGtmAvailableAlertGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmAvailableAlertGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAvailableAlertGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDnsGtmAvailableAlertGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDnsGtmAvailableAlertGroupResponse) GetBody() *DescribeDnsGtmAvailableAlertGroupResponseBody {
	return s.Body
}

func (s *DescribeDnsGtmAvailableAlertGroupResponse) SetHeaders(v map[string]*string) *DescribeDnsGtmAvailableAlertGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeDnsGtmAvailableAlertGroupResponse) SetStatusCode(v int32) *DescribeDnsGtmAvailableAlertGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDnsGtmAvailableAlertGroupResponse) SetBody(v *DescribeDnsGtmAvailableAlertGroupResponseBody) *DescribeDnsGtmAvailableAlertGroupResponse {
	s.Body = v
	return s
}

func (s *DescribeDnsGtmAvailableAlertGroupResponse) Validate() error {
	return dara.Validate(s)
}

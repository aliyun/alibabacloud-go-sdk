// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnsGtmAvailableAlertGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAvailableAlertGroup(v string) *DescribeDnsGtmAvailableAlertGroupResponseBody
	GetAvailableAlertGroup() *string
	SetRequestId(v string) *DescribeDnsGtmAvailableAlertGroupResponseBody
	GetRequestId() *string
}

type DescribeDnsGtmAvailableAlertGroupResponseBody struct {
	// The available alert contact groups.
	//
	// example:
	//
	// [\\"测试报警组\\",\\"研发报警组\\"]
	AvailableAlertGroup *string `json:"AvailableAlertGroup,omitempty" xml:"AvailableAlertGroup,omitempty"`
	// The unique request ID.
	//
	// example:
	//
	// 6856BCF6-11D6-4D7E-AC53-FD579933522B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDnsGtmAvailableAlertGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmAvailableAlertGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAvailableAlertGroupResponseBody) GetAvailableAlertGroup() *string {
	return s.AvailableAlertGroup
}

func (s *DescribeDnsGtmAvailableAlertGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDnsGtmAvailableAlertGroupResponseBody) SetAvailableAlertGroup(v string) *DescribeDnsGtmAvailableAlertGroupResponseBody {
	s.AvailableAlertGroup = &v
	return s
}

func (s *DescribeDnsGtmAvailableAlertGroupResponseBody) SetRequestId(v string) *DescribeDnsGtmAvailableAlertGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDnsGtmAvailableAlertGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

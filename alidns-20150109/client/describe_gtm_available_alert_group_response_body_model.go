// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGtmAvailableAlertGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAvailableAlertGroup(v string) *DescribeGtmAvailableAlertGroupResponseBody
	GetAvailableAlertGroup() *string
	SetRequestId(v string) *DescribeGtmAvailableAlertGroupResponseBody
	GetRequestId() *string
}

type DescribeGtmAvailableAlertGroupResponseBody struct {
	// The available alert groups of the GTM instance.
	AvailableAlertGroup *string `json:"AvailableAlertGroup,omitempty" xml:"AvailableAlertGroup,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 29D0F8F8-5499-4F6C-9FDC-1EE13BF55925
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeGtmAvailableAlertGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmAvailableAlertGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGtmAvailableAlertGroupResponseBody) GetAvailableAlertGroup() *string {
	return s.AvailableAlertGroup
}

func (s *DescribeGtmAvailableAlertGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGtmAvailableAlertGroupResponseBody) SetAvailableAlertGroup(v string) *DescribeGtmAvailableAlertGroupResponseBody {
	s.AvailableAlertGroup = &v
	return s
}

func (s *DescribeGtmAvailableAlertGroupResponseBody) SetRequestId(v string) *DescribeGtmAvailableAlertGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGtmAvailableAlertGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

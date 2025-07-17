// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAlertConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeAlertConfigurationResponseBody
	GetRequestId() *string
	SetScaleStatuses(v []*string) *DescribeAlertConfigurationResponseBody
	GetScaleStatuses() []*string
}

type DescribeAlertConfigurationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the scaling activities that trigger text message, internal message, or email-based notifications.
	ScaleStatuses []*string `json:"ScaleStatuses,omitempty" xml:"ScaleStatuses,omitempty" type:"Repeated"`
}

func (s DescribeAlertConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAlertConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAlertConfigurationResponseBody) GetScaleStatuses() []*string {
	return s.ScaleStatuses
}

func (s *DescribeAlertConfigurationResponseBody) SetRequestId(v string) *DescribeAlertConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAlertConfigurationResponseBody) SetScaleStatuses(v []*string) *DescribeAlertConfigurationResponseBody {
	s.ScaleStatuses = v
	return s
}

func (s *DescribeAlertConfigurationResponseBody) Validate() error {
	return dara.Validate(s)
}

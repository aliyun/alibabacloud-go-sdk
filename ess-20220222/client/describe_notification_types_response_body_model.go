// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNotificationTypesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNotificationTypes(v []*string) *DescribeNotificationTypesResponseBody
	GetNotificationTypes() []*string
	SetRequestId(v string) *DescribeNotificationTypesResponseBody
	GetRequestId() *string
}

type DescribeNotificationTypesResponseBody struct {
	// The types of the notifications.
	//
	// 	- AUTOSCALING:SCALE_OUT_SUCCESS: The scale-out activity succeeds.
	//
	// 	- AUTOSCALING:SCALE_IN_SUCCESS: The scale-in activity succeeds.
	//
	// 	- AUTOSCALING:SCALE_OUT_ERROR: The scale-out activity fails.
	//
	// 	- AUTOSCALING:SCALE_IN_ERROR: The scale-in activity fails.
	//
	// 	- AUTOSCALING:SCALE_REJECT: The request for scaling activities is rejected.
	//
	// 	- AUTOSCALING:SCALE_OUT_START: The scale-out activity starts.
	//
	// 	- AUTOSCALING:SCALE_IN_START: The scale-in activity starts.
	//
	// 	- AUTOSCALING:SCHEDULE_TASK_EXPIRING: Auto Scaling sends a notification when a scheduled task is about to expire.
	NotificationTypes []*string `json:"NotificationTypes,omitempty" xml:"NotificationTypes,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeNotificationTypesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNotificationTypesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNotificationTypesResponseBody) GetNotificationTypes() []*string {
	return s.NotificationTypes
}

func (s *DescribeNotificationTypesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNotificationTypesResponseBody) SetNotificationTypes(v []*string) *DescribeNotificationTypesResponseBody {
	s.NotificationTypes = v
	return s
}

func (s *DescribeNotificationTypesResponseBody) SetRequestId(v string) *DescribeNotificationTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNotificationTypesResponseBody) Validate() error {
	return dara.Validate(s)
}

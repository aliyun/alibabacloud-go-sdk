// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNotificationRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListNotificationRecordsRequest
	GetInstanceId() *string
	SetNotificationKeys(v string) *ListNotificationRecordsRequest
	GetNotificationKeys() *string
}

type ListNotificationRecordsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ["job-468a63a2-****-****-****-b1ecf726d4be","job-532a63a2-****-****-****-bcccc26d4cc"]
	NotificationKeys *string `json:"NotificationKeys,omitempty" xml:"NotificationKeys,omitempty"`
}

func (s ListNotificationRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNotificationRecordsRequest) GoString() string {
	return s.String()
}

func (s *ListNotificationRecordsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListNotificationRecordsRequest) GetNotificationKeys() *string {
	return s.NotificationKeys
}

func (s *ListNotificationRecordsRequest) SetInstanceId(v string) *ListNotificationRecordsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListNotificationRecordsRequest) SetNotificationKeys(v string) *ListNotificationRecordsRequest {
	s.NotificationKeys = &v
	return s
}

func (s *ListNotificationRecordsRequest) Validate() error {
	return dara.Validate(s)
}

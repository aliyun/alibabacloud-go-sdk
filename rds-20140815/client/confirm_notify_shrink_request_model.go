// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfirmNotifyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfirmor(v int64) *ConfirmNotifyShrinkRequest
	GetConfirmor() *int64
	SetNotifyIdListShrink(v string) *ConfirmNotifyShrinkRequest
	GetNotifyIdListShrink() *string
}

type ConfirmNotifyShrinkRequest struct {
	// The ID of the Alibaba Cloud account that is used to confirm the notification. You can set this parameter to **0**, which indicates that the notification is confirmed by the system.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	Confirmor *int64 `json:"Confirmor,omitempty" xml:"Confirmor,omitempty"`
	// The notification IDs.
	//
	// This parameter is required.
	//
	// if can be null:
	// false
	NotifyIdListShrink *string `json:"NotifyIdList,omitempty" xml:"NotifyIdList,omitempty"`
}

func (s ConfirmNotifyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfirmNotifyShrinkRequest) GoString() string {
	return s.String()
}

func (s *ConfirmNotifyShrinkRequest) GetConfirmor() *int64 {
	return s.Confirmor
}

func (s *ConfirmNotifyShrinkRequest) GetNotifyIdListShrink() *string {
	return s.NotifyIdListShrink
}

func (s *ConfirmNotifyShrinkRequest) SetConfirmor(v int64) *ConfirmNotifyShrinkRequest {
	s.Confirmor = &v
	return s
}

func (s *ConfirmNotifyShrinkRequest) SetNotifyIdListShrink(v string) *ConfirmNotifyShrinkRequest {
	s.NotifyIdListShrink = &v
	return s
}

func (s *ConfirmNotifyShrinkRequest) Validate() error {
	return dara.Validate(s)
}

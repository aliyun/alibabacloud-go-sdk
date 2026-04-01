// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfirmNotifyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfirmor(v int64) *ConfirmNotifyRequest
	GetConfirmor() *int64
	SetNotifyIdList(v []*int64) *ConfirmNotifyRequest
	GetNotifyIdList() []*int64
}

type ConfirmNotifyRequest struct {
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
	NotifyIdList []*int64 `json:"NotifyIdList,omitempty" xml:"NotifyIdList,omitempty" type:"Repeated"`
}

func (s ConfirmNotifyRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfirmNotifyRequest) GoString() string {
	return s.String()
}

func (s *ConfirmNotifyRequest) GetConfirmor() *int64 {
	return s.Confirmor
}

func (s *ConfirmNotifyRequest) GetNotifyIdList() []*int64 {
	return s.NotifyIdList
}

func (s *ConfirmNotifyRequest) SetConfirmor(v int64) *ConfirmNotifyRequest {
	s.Confirmor = &v
	return s
}

func (s *ConfirmNotifyRequest) SetNotifyIdList(v []*int64) *ConfirmNotifyRequest {
	s.NotifyIdList = v
	return s
}

func (s *ConfirmNotifyRequest) Validate() error {
	return dara.Validate(s)
}

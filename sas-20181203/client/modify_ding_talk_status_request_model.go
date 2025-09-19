// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDingTalkStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIds(v string) *ModifyDingTalkStatusRequest
	GetIds() *string
	SetStatus(v int32) *ModifyDingTalkStatusRequest
	GetStatus() *int32
}

type ModifyDingTalkStatusRequest struct {
	// The notification IDs of DingTalk chatbots. Separate multiple IDs with commas (,).
	//
	// >  You can call the [DescribeDingTalk](~~DescribeDingTalk~~) operation to query the notification IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2259
	Ids *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	// The notification status of a DingTalk chatbot. Valid values:
	//
	// 	- **0**: disabled
	//
	// 	- **1**: enabled
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyDingTalkStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDingTalkStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyDingTalkStatusRequest) GetIds() *string {
	return s.Ids
}

func (s *ModifyDingTalkStatusRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ModifyDingTalkStatusRequest) SetIds(v string) *ModifyDingTalkStatusRequest {
	s.Ids = &v
	return s
}

func (s *ModifyDingTalkStatusRequest) SetStatus(v int32) *ModifyDingTalkStatusRequest {
	s.Status = &v
	return s
}

func (s *ModifyDingTalkStatusRequest) Validate() error {
	return dara.Validate(s)
}

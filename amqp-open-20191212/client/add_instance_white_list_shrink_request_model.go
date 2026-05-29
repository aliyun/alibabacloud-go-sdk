// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddInstanceWhiteListShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *AddInstanceWhiteListShrinkRequest
	GetInstanceId() *string
	SetWhiteListItemShrink(v string) *AddInstanceWhiteListShrinkRequest
	GetWhiteListItemShrink() *string
	SetWhiteListType(v int32) *AddInstanceWhiteListShrinkRequest
	GetWhiteListType() *int32
}

type AddInstanceWhiteListShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// rabbitmq-cn-xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 172.0.0.20/30
	WhiteListItemShrink *string `json:"WhiteListItem,omitempty" xml:"WhiteListItem,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	WhiteListType *int32 `json:"WhiteListType,omitempty" xml:"WhiteListType,omitempty"`
}

func (s AddInstanceWhiteListShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddInstanceWhiteListShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddInstanceWhiteListShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AddInstanceWhiteListShrinkRequest) GetWhiteListItemShrink() *string {
	return s.WhiteListItemShrink
}

func (s *AddInstanceWhiteListShrinkRequest) GetWhiteListType() *int32 {
	return s.WhiteListType
}

func (s *AddInstanceWhiteListShrinkRequest) SetInstanceId(v string) *AddInstanceWhiteListShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *AddInstanceWhiteListShrinkRequest) SetWhiteListItemShrink(v string) *AddInstanceWhiteListShrinkRequest {
	s.WhiteListItemShrink = &v
	return s
}

func (s *AddInstanceWhiteListShrinkRequest) SetWhiteListType(v int32) *AddInstanceWhiteListShrinkRequest {
	s.WhiteListType = &v
	return s
}

func (s *AddInstanceWhiteListShrinkRequest) Validate() error {
	return dara.Validate(s)
}

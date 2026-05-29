// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddInstanceWhiteListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *AddInstanceWhiteListRequest
	GetInstanceId() *string
	SetWhiteListItem(v []*string) *AddInstanceWhiteListRequest
	GetWhiteListItem() []*string
	SetWhiteListType(v int32) *AddInstanceWhiteListRequest
	GetWhiteListType() *int32
}

type AddInstanceWhiteListRequest struct {
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
	WhiteListItem []*string `json:"WhiteListItem,omitempty" xml:"WhiteListItem,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	WhiteListType *int32 `json:"WhiteListType,omitempty" xml:"WhiteListType,omitempty"`
}

func (s AddInstanceWhiteListRequest) String() string {
	return dara.Prettify(s)
}

func (s AddInstanceWhiteListRequest) GoString() string {
	return s.String()
}

func (s *AddInstanceWhiteListRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AddInstanceWhiteListRequest) GetWhiteListItem() []*string {
	return s.WhiteListItem
}

func (s *AddInstanceWhiteListRequest) GetWhiteListType() *int32 {
	return s.WhiteListType
}

func (s *AddInstanceWhiteListRequest) SetInstanceId(v string) *AddInstanceWhiteListRequest {
	s.InstanceId = &v
	return s
}

func (s *AddInstanceWhiteListRequest) SetWhiteListItem(v []*string) *AddInstanceWhiteListRequest {
	s.WhiteListItem = v
	return s
}

func (s *AddInstanceWhiteListRequest) SetWhiteListType(v int32) *AddInstanceWhiteListRequest {
	s.WhiteListType = &v
	return s
}

func (s *AddInstanceWhiteListRequest) Validate() error {
	return dara.Validate(s)
}

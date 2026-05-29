// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceWhiteListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListInstanceWhiteListRequest
	GetInstanceId() *string
	SetWhiteListType(v int32) *ListInstanceWhiteListRequest
	GetWhiteListType() *int32
}

type ListInstanceWhiteListRequest struct {
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
	// 2
	WhiteListType *int32 `json:"whiteListType,omitempty" xml:"whiteListType,omitempty"`
}

func (s ListInstanceWhiteListRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceWhiteListRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceWhiteListRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstanceWhiteListRequest) GetWhiteListType() *int32 {
	return s.WhiteListType
}

func (s *ListInstanceWhiteListRequest) SetInstanceId(v string) *ListInstanceWhiteListRequest {
	s.InstanceId = &v
	return s
}

func (s *ListInstanceWhiteListRequest) SetWhiteListType(v int32) *ListInstanceWhiteListRequest {
	s.WhiteListType = &v
	return s
}

func (s *ListInstanceWhiteListRequest) Validate() error {
	return dara.Validate(s)
}

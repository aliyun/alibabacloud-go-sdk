// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendLikeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *SendLikeRequest
	GetAppId() *string
	SetBroadCastType(v int32) *SendLikeRequest
	GetBroadCastType() *int32
	SetCount(v string) *SendLikeRequest
	GetCount() *string
	SetGroupId(v string) *SendLikeRequest
	GetGroupId() *string
	SetOperatorUserId(v string) *SendLikeRequest
	GetOperatorUserId() *string
}

type SendLikeRequest struct {
	// The ID of the interactive messaging application.
	//
	// This parameter is required.
	//
	// example:
	//
	// VKL3***
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The mode in which system messages are broadcasted. Valid values:
	//
	// 	- 0: specifies that system messages are not broadcasted. This is the default value.
	//
	// 	- 1: specifies that system messages are broadcasted to specified users.
	//
	// 	- 2: specifies that system messages are broadcasted to the message group.
	//
	// example:
	//
	// 2
	BroadCastType *int32 `json:"BroadCastType,omitempty" xml:"BroadCastType,omitempty"`
	// The number of likes.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	Count *string `json:"Count,omitempty" xml:"Count,omitempty"`
	// The ID of the message group.
	//
	// This parameter is required.
	//
	// example:
	//
	// AE35-****-T95F
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the user who performs the operation.
	//
	// example:
	//
	// de1**a0
	OperatorUserId *string `json:"OperatorUserId,omitempty" xml:"OperatorUserId,omitempty"`
}

func (s SendLikeRequest) String() string {
	return dara.Prettify(s)
}

func (s SendLikeRequest) GoString() string {
	return s.String()
}

func (s *SendLikeRequest) GetAppId() *string {
	return s.AppId
}

func (s *SendLikeRequest) GetBroadCastType() *int32 {
	return s.BroadCastType
}

func (s *SendLikeRequest) GetCount() *string {
	return s.Count
}

func (s *SendLikeRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *SendLikeRequest) GetOperatorUserId() *string {
	return s.OperatorUserId
}

func (s *SendLikeRequest) SetAppId(v string) *SendLikeRequest {
	s.AppId = &v
	return s
}

func (s *SendLikeRequest) SetBroadCastType(v int32) *SendLikeRequest {
	s.BroadCastType = &v
	return s
}

func (s *SendLikeRequest) SetCount(v string) *SendLikeRequest {
	s.Count = &v
	return s
}

func (s *SendLikeRequest) SetGroupId(v string) *SendLikeRequest {
	s.GroupId = &v
	return s
}

func (s *SendLikeRequest) SetOperatorUserId(v string) *SendLikeRequest {
	s.OperatorUserId = &v
	return s
}

func (s *SendLikeRequest) Validate() error {
	return dara.Validate(s)
}

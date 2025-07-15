// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMuteGroupUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListMuteGroupUserRequest
	GetAppId() *string
	SetGroupId(v string) *ListMuteGroupUserRequest
	GetGroupId() *string
	SetOperatorUserId(v string) *ListMuteGroupUserRequest
	GetOperatorUserId() *string
	SetPageNum(v int32) *ListMuteGroupUserRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListMuteGroupUserRequest
	GetPageSize() *int32
}

type ListMuteGroupUserRequest struct {
	// The ID of the interactive messaging application.
	//
	// This parameter is required.
	//
	// example:
	//
	// VKL3***
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the messaging group.
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
	// The page number. Default value: 1. Valid values: 1 to 100000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page. Default value: 20. Valid values: 1 to 50.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListMuteGroupUserRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMuteGroupUserRequest) GoString() string {
	return s.String()
}

func (s *ListMuteGroupUserRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListMuteGroupUserRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ListMuteGroupUserRequest) GetOperatorUserId() *string {
	return s.OperatorUserId
}

func (s *ListMuteGroupUserRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListMuteGroupUserRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMuteGroupUserRequest) SetAppId(v string) *ListMuteGroupUserRequest {
	s.AppId = &v
	return s
}

func (s *ListMuteGroupUserRequest) SetGroupId(v string) *ListMuteGroupUserRequest {
	s.GroupId = &v
	return s
}

func (s *ListMuteGroupUserRequest) SetOperatorUserId(v string) *ListMuteGroupUserRequest {
	s.OperatorUserId = &v
	return s
}

func (s *ListMuteGroupUserRequest) SetPageNum(v int32) *ListMuteGroupUserRequest {
	s.PageNum = &v
	return s
}

func (s *ListMuteGroupUserRequest) SetPageSize(v int32) *ListMuteGroupUserRequest {
	s.PageSize = &v
	return s
}

func (s *ListMuteGroupUserRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySmsAppIcpRecordShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppIcpRecordIdListShrink(v string) *QuerySmsAppIcpRecordShrinkRequest
	GetAppIcpRecordIdListShrink() *string
	SetOwnerId(v int64) *QuerySmsAppIcpRecordShrinkRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *QuerySmsAppIcpRecordShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QuerySmsAppIcpRecordShrinkRequest
	GetResourceOwnerId() *int64
}

type QuerySmsAppIcpRecordShrinkRequest struct {
	// app-icp备案实体id列表
	//
	// This parameter is required.
	AppIcpRecordIdListShrink *string `json:"AppIcpRecordIdList,omitempty" xml:"AppIcpRecordIdList,omitempty"`
	OwnerId                  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount     *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId          *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QuerySmsAppIcpRecordShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySmsAppIcpRecordShrinkRequest) GoString() string {
	return s.String()
}

func (s *QuerySmsAppIcpRecordShrinkRequest) GetAppIcpRecordIdListShrink() *string {
	return s.AppIcpRecordIdListShrink
}

func (s *QuerySmsAppIcpRecordShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QuerySmsAppIcpRecordShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QuerySmsAppIcpRecordShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QuerySmsAppIcpRecordShrinkRequest) SetAppIcpRecordIdListShrink(v string) *QuerySmsAppIcpRecordShrinkRequest {
	s.AppIcpRecordIdListShrink = &v
	return s
}

func (s *QuerySmsAppIcpRecordShrinkRequest) SetOwnerId(v int64) *QuerySmsAppIcpRecordShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *QuerySmsAppIcpRecordShrinkRequest) SetResourceOwnerAccount(v string) *QuerySmsAppIcpRecordShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QuerySmsAppIcpRecordShrinkRequest) SetResourceOwnerId(v int64) *QuerySmsAppIcpRecordShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QuerySmsAppIcpRecordShrinkRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySmsAppIcpRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppIcpRecordIdList(v []*int64) *QuerySmsAppIcpRecordRequest
	GetAppIcpRecordIdList() []*int64
	SetOwnerId(v int64) *QuerySmsAppIcpRecordRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *QuerySmsAppIcpRecordRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QuerySmsAppIcpRecordRequest
	GetResourceOwnerId() *int64
}

type QuerySmsAppIcpRecordRequest struct {
	// app-icp备案实体id列表
	//
	// This parameter is required.
	AppIcpRecordIdList   []*int64 `json:"AppIcpRecordIdList,omitempty" xml:"AppIcpRecordIdList,omitempty" type:"Repeated"`
	OwnerId              *int64   `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string  `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64   `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QuerySmsAppIcpRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySmsAppIcpRecordRequest) GoString() string {
	return s.String()
}

func (s *QuerySmsAppIcpRecordRequest) GetAppIcpRecordIdList() []*int64 {
	return s.AppIcpRecordIdList
}

func (s *QuerySmsAppIcpRecordRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QuerySmsAppIcpRecordRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QuerySmsAppIcpRecordRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QuerySmsAppIcpRecordRequest) SetAppIcpRecordIdList(v []*int64) *QuerySmsAppIcpRecordRequest {
	s.AppIcpRecordIdList = v
	return s
}

func (s *QuerySmsAppIcpRecordRequest) SetOwnerId(v int64) *QuerySmsAppIcpRecordRequest {
	s.OwnerId = &v
	return s
}

func (s *QuerySmsAppIcpRecordRequest) SetResourceOwnerAccount(v string) *QuerySmsAppIcpRecordRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QuerySmsAppIcpRecordRequest) SetResourceOwnerId(v int64) *QuerySmsAppIcpRecordRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QuerySmsAppIcpRecordRequest) Validate() error {
	return dara.Validate(s)
}

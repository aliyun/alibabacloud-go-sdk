// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartInstanceNodeListShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *RestartInstanceNodeListShrinkRequest
	GetDBClusterId() *string
	SetNodeListShrink(v string) *RestartInstanceNodeListShrinkRequest
	GetNodeListShrink() *string
	SetOwnerAccount(v string) *RestartInstanceNodeListShrinkRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RestartInstanceNodeListShrinkRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *RestartInstanceNodeListShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *RestartInstanceNodeListShrinkRequest
	GetPageSize() *int32
	SetRegionId(v string) *RestartInstanceNodeListShrinkRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *RestartInstanceNodeListShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RestartInstanceNodeListShrinkRequest
	GetResourceOwnerId() *int64
	SetRestartTime(v string) *RestartInstanceNodeListShrinkRequest
	GetRestartTime() *string
}

type RestartInstanceNodeListShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cc-bp108z124a8o7****
	DBClusterId    *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	NodeListShrink *string `json:"NodeList,omitempty" xml:"NodeList,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// 2023-03-22T00:00:50Z
	RestartTime *string `json:"RestartTime,omitempty" xml:"RestartTime,omitempty"`
}

func (s RestartInstanceNodeListShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RestartInstanceNodeListShrinkRequest) GoString() string {
	return s.String()
}

func (s *RestartInstanceNodeListShrinkRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *RestartInstanceNodeListShrinkRequest) GetNodeListShrink() *string {
	return s.NodeListShrink
}

func (s *RestartInstanceNodeListShrinkRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RestartInstanceNodeListShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RestartInstanceNodeListShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *RestartInstanceNodeListShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *RestartInstanceNodeListShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RestartInstanceNodeListShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RestartInstanceNodeListShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RestartInstanceNodeListShrinkRequest) GetRestartTime() *string {
	return s.RestartTime
}

func (s *RestartInstanceNodeListShrinkRequest) SetDBClusterId(v string) *RestartInstanceNodeListShrinkRequest {
	s.DBClusterId = &v
	return s
}

func (s *RestartInstanceNodeListShrinkRequest) SetNodeListShrink(v string) *RestartInstanceNodeListShrinkRequest {
	s.NodeListShrink = &v
	return s
}

func (s *RestartInstanceNodeListShrinkRequest) SetOwnerAccount(v string) *RestartInstanceNodeListShrinkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RestartInstanceNodeListShrinkRequest) SetOwnerId(v int64) *RestartInstanceNodeListShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *RestartInstanceNodeListShrinkRequest) SetPageNumber(v int32) *RestartInstanceNodeListShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *RestartInstanceNodeListShrinkRequest) SetPageSize(v int32) *RestartInstanceNodeListShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *RestartInstanceNodeListShrinkRequest) SetRegionId(v string) *RestartInstanceNodeListShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *RestartInstanceNodeListShrinkRequest) SetResourceOwnerAccount(v string) *RestartInstanceNodeListShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RestartInstanceNodeListShrinkRequest) SetResourceOwnerId(v int64) *RestartInstanceNodeListShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RestartInstanceNodeListShrinkRequest) SetRestartTime(v string) *RestartInstanceNodeListShrinkRequest {
	s.RestartTime = &v
	return s
}

func (s *RestartInstanceNodeListShrinkRequest) Validate() error {
	return dara.Validate(s)
}

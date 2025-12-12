// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartInstanceNodeListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *RestartInstanceNodeListRequest
	GetDBClusterId() *string
	SetNodeList(v []*string) *RestartInstanceNodeListRequest
	GetNodeList() []*string
	SetOwnerAccount(v string) *RestartInstanceNodeListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RestartInstanceNodeListRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *RestartInstanceNodeListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *RestartInstanceNodeListRequest
	GetPageSize() *int32
	SetRegionId(v string) *RestartInstanceNodeListRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *RestartInstanceNodeListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RestartInstanceNodeListRequest
	GetResourceOwnerId() *int64
	SetRestartTime(v string) *RestartInstanceNodeListRequest
	GetRestartTime() *string
}

type RestartInstanceNodeListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cc-bp108z124a8o7****
	DBClusterId  *string   `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	NodeList     []*string `json:"NodeList,omitempty" xml:"NodeList,omitempty" type:"Repeated"`
	OwnerAccount *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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

func (s RestartInstanceNodeListRequest) String() string {
	return dara.Prettify(s)
}

func (s RestartInstanceNodeListRequest) GoString() string {
	return s.String()
}

func (s *RestartInstanceNodeListRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *RestartInstanceNodeListRequest) GetNodeList() []*string {
	return s.NodeList
}

func (s *RestartInstanceNodeListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RestartInstanceNodeListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RestartInstanceNodeListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *RestartInstanceNodeListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *RestartInstanceNodeListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RestartInstanceNodeListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RestartInstanceNodeListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RestartInstanceNodeListRequest) GetRestartTime() *string {
	return s.RestartTime
}

func (s *RestartInstanceNodeListRequest) SetDBClusterId(v string) *RestartInstanceNodeListRequest {
	s.DBClusterId = &v
	return s
}

func (s *RestartInstanceNodeListRequest) SetNodeList(v []*string) *RestartInstanceNodeListRequest {
	s.NodeList = v
	return s
}

func (s *RestartInstanceNodeListRequest) SetOwnerAccount(v string) *RestartInstanceNodeListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RestartInstanceNodeListRequest) SetOwnerId(v int64) *RestartInstanceNodeListRequest {
	s.OwnerId = &v
	return s
}

func (s *RestartInstanceNodeListRequest) SetPageNumber(v int32) *RestartInstanceNodeListRequest {
	s.PageNumber = &v
	return s
}

func (s *RestartInstanceNodeListRequest) SetPageSize(v int32) *RestartInstanceNodeListRequest {
	s.PageSize = &v
	return s
}

func (s *RestartInstanceNodeListRequest) SetRegionId(v string) *RestartInstanceNodeListRequest {
	s.RegionId = &v
	return s
}

func (s *RestartInstanceNodeListRequest) SetResourceOwnerAccount(v string) *RestartInstanceNodeListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RestartInstanceNodeListRequest) SetResourceOwnerId(v int64) *RestartInstanceNodeListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RestartInstanceNodeListRequest) SetRestartTime(v string) *RestartInstanceNodeListRequest {
	s.RestartTime = &v
	return s
}

func (s *RestartInstanceNodeListRequest) Validate() error {
	return dara.Validate(s)
}

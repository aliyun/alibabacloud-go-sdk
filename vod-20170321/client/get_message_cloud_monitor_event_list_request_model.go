// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMessageCloudMonitorEventListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *GetMessageCloudMonitorEventListRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *GetMessageCloudMonitorEventListRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *GetMessageCloudMonitorEventListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *GetMessageCloudMonitorEventListRequest
	GetResourceOwnerId() *string
	SetResourceRealOwnerId(v int64) *GetMessageCloudMonitorEventListRequest
	GetResourceRealOwnerId() *int64
}

type GetMessageCloudMonitorEventListRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceRealOwnerId  *int64  `json:"ResourceRealOwnerId,omitempty" xml:"ResourceRealOwnerId,omitempty"`
}

func (s GetMessageCloudMonitorEventListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMessageCloudMonitorEventListRequest) GoString() string {
	return s.String()
}

func (s *GetMessageCloudMonitorEventListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetMessageCloudMonitorEventListRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GetMessageCloudMonitorEventListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetMessageCloudMonitorEventListRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *GetMessageCloudMonitorEventListRequest) GetResourceRealOwnerId() *int64 {
	return s.ResourceRealOwnerId
}

func (s *GetMessageCloudMonitorEventListRequest) SetOwnerAccount(v string) *GetMessageCloudMonitorEventListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetMessageCloudMonitorEventListRequest) SetOwnerId(v string) *GetMessageCloudMonitorEventListRequest {
	s.OwnerId = &v
	return s
}

func (s *GetMessageCloudMonitorEventListRequest) SetResourceOwnerAccount(v string) *GetMessageCloudMonitorEventListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetMessageCloudMonitorEventListRequest) SetResourceOwnerId(v string) *GetMessageCloudMonitorEventListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetMessageCloudMonitorEventListRequest) SetResourceRealOwnerId(v int64) *GetMessageCloudMonitorEventListRequest {
	s.ResourceRealOwnerId = &v
	return s
}

func (s *GetMessageCloudMonitorEventListRequest) Validate() error {
	return dara.Validate(s)
}

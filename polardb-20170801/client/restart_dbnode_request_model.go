// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartDBNodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBNodeId(v string) *RestartDBNodeRequest
	GetDBNodeId() *string
	SetOwnerAccount(v string) *RestartDBNodeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RestartDBNodeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *RestartDBNodeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RestartDBNodeRequest
	GetResourceOwnerId() *int64
}

type RestartDBNodeRequest struct {
	// The ID of the node.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/185342.html) operation to query the details of all clusters that belong to your Alibaba Cloud account, such as cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// pi-*************
	DBNodeId             *string `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s RestartDBNodeRequest) String() string {
	return dara.Prettify(s)
}

func (s RestartDBNodeRequest) GoString() string {
	return s.String()
}

func (s *RestartDBNodeRequest) GetDBNodeId() *string {
	return s.DBNodeId
}

func (s *RestartDBNodeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RestartDBNodeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RestartDBNodeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RestartDBNodeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RestartDBNodeRequest) SetDBNodeId(v string) *RestartDBNodeRequest {
	s.DBNodeId = &v
	return s
}

func (s *RestartDBNodeRequest) SetOwnerAccount(v string) *RestartDBNodeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RestartDBNodeRequest) SetOwnerId(v int64) *RestartDBNodeRequest {
	s.OwnerId = &v
	return s
}

func (s *RestartDBNodeRequest) SetResourceOwnerAccount(v string) *RestartDBNodeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RestartDBNodeRequest) SetResourceOwnerId(v int64) *RestartDBNodeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RestartDBNodeRequest) Validate() error {
	return dara.Validate(s)
}

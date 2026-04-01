// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartDBInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RestartDBInstanceRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *RestartDBInstanceRequest
	GetDBInstanceId() *string
	SetNodeId(v string) *RestartDBInstanceRequest
	GetNodeId() *string
	SetOwnerAccount(v string) *RestartDBInstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RestartDBInstanceRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *RestartDBInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RestartDBInstanceRequest
	GetResourceOwnerId() *int64
}

type RestartDBInstanceRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID. You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/610396.html) operation to query the ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The node ID, which can be used to restart a specified node. You can call the [DescribeDBInstanceHAConfig](https://help.aliyun.com/document_detail/610434.html) operation to obtain the node ID.
	//
	// > : The secondary instance restart feature is supported for RDS instances that run SQL Server EE on RDS Cluster Edition. For more information, see [Restart a secondary database](https://help.aliyun.com/document_detail/2411880.html).
	//
	// example:
	//
	// 2614****
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s RestartDBInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RestartDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *RestartDBInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RestartDBInstanceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *RestartDBInstanceRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *RestartDBInstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RestartDBInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RestartDBInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RestartDBInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RestartDBInstanceRequest) SetClientToken(v string) *RestartDBInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *RestartDBInstanceRequest) SetDBInstanceId(v string) *RestartDBInstanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *RestartDBInstanceRequest) SetNodeId(v string) *RestartDBInstanceRequest {
	s.NodeId = &v
	return s
}

func (s *RestartDBInstanceRequest) SetOwnerAccount(v string) *RestartDBInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RestartDBInstanceRequest) SetOwnerId(v int64) *RestartDBInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *RestartDBInstanceRequest) SetResourceOwnerAccount(v string) *RestartDBInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RestartDBInstanceRequest) SetResourceOwnerId(v int64) *RestartDBInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RestartDBInstanceRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKillProcessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *KillProcessRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *KillProcessRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *KillProcessRequest
	GetOwnerId() *int64
	SetProcessId(v string) *KillProcessRequest
	GetProcessId() *string
	SetResourceOwnerAccount(v string) *KillProcessRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *KillProcessRequest
	GetResourceOwnerId() *int64
}

type KillProcessRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp1xxxxxxxx47
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The unique ID of the query. You can call the [DescribeProcessList](https://help.aliyun.com/document_detail/190092.html) operation to obtain the unique ID of a query.
	//
	// example:
	//
	// 202011191048151921681492420315100****
	ProcessId            *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s KillProcessRequest) String() string {
	return dara.Prettify(s)
}

func (s KillProcessRequest) GoString() string {
	return s.String()
}

func (s *KillProcessRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *KillProcessRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *KillProcessRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *KillProcessRequest) GetProcessId() *string {
	return s.ProcessId
}

func (s *KillProcessRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *KillProcessRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *KillProcessRequest) SetDBClusterId(v string) *KillProcessRequest {
	s.DBClusterId = &v
	return s
}

func (s *KillProcessRequest) SetOwnerAccount(v string) *KillProcessRequest {
	s.OwnerAccount = &v
	return s
}

func (s *KillProcessRequest) SetOwnerId(v int64) *KillProcessRequest {
	s.OwnerId = &v
	return s
}

func (s *KillProcessRequest) SetProcessId(v string) *KillProcessRequest {
	s.ProcessId = &v
	return s
}

func (s *KillProcessRequest) SetResourceOwnerAccount(v string) *KillProcessRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *KillProcessRequest) SetResourceOwnerId(v int64) *KillProcessRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *KillProcessRequest) Validate() error {
	return dara.Validate(s)
}

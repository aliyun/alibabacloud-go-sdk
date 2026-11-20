// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAxgGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v int64) *DeleteAxgGroupRequest
	GetGroupId() *int64
	SetOwnerId(v int64) *DeleteAxgGroupRequest
	GetOwnerId() *int64
	SetPoolKey(v string) *DeleteAxgGroupRequest
	GetPoolKey() *string
	SetResourceOwnerAccount(v string) *DeleteAxgGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteAxgGroupRequest
	GetResourceOwnerId() *int64
}

type DeleteAxgGroupRequest struct {
	// The G group ID (GID).
	//
	// - Log on to the Phone Number Privacy Protection console. On the [Number Pool Management](https://dyplsnext.console.aliyun.com/number-pool) page, filter for number pools of the **AXG Privacy Number*	- type and click **G Number Group Management*	- to view the IDs of all G groups.
	//
	// - If you created the G group by calling the [CreateAxgGroup](https://help.aliyun.com/document_detail/110250.html) API, use the returned **GroupID*	- as the value for this **GroupId*	- parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 200001234****
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The key of the number pool.
	//
	// You can obtain the **key of the number pool*	- on the [Number Pool Management](https://dyplsnext.console.aliyun.com/number-pool) page in the Phone Number Privacy Protection console.
	//
	// This parameter is required.
	//
	// example:
	//
	// FC2235****
	PoolKey              *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteAxgGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAxgGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteAxgGroupRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *DeleteAxgGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteAxgGroupRequest) GetPoolKey() *string {
	return s.PoolKey
}

func (s *DeleteAxgGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteAxgGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteAxgGroupRequest) SetGroupId(v int64) *DeleteAxgGroupRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteAxgGroupRequest) SetOwnerId(v int64) *DeleteAxgGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteAxgGroupRequest) SetPoolKey(v string) *DeleteAxgGroupRequest {
	s.PoolKey = &v
	return s
}

func (s *DeleteAxgGroupRequest) SetResourceOwnerAccount(v string) *DeleteAxgGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteAxgGroupRequest) SetResourceOwnerId(v int64) *DeleteAxgGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteAxgGroupRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClusterStateChangeReason interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ClusterStateChangeReason
	GetCode() *string
	SetMessage(v string) *ClusterStateChangeReason
	GetMessage() *string
}

type ClusterStateChangeReason struct {
	// 状态码。取值范围：
	//
	// - UserRequest：用户请求。
	//
	// - OutOfStock：请求的ECS实例类型无库存。
	//
	// - NotAuthorized：无权限。
	//
	// - QuotaExceeded：Quota超出。
	//
	// - OperationDenied：操作被拒绝。
	//
	// - AccountException：账号异常。
	//
	// - NodeFailure：ECS节点异常。
	//
	// - BootstrapFailure：引导失败。
	//
	// - ValidationFail：业务逻辑校验失败。
	//
	// - ServiceFailure：依赖的其他服务失败。
	//
	// - InternalError：内部错误。
	//
	// example:
	//
	// OutOfStock
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 状态变化信息。
	//
	// example:
	//
	// The requested resource is sold out in the specified zone, try other types of resources or other regions and zones.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s ClusterStateChangeReason) String() string {
	return dara.Prettify(s)
}

func (s ClusterStateChangeReason) GoString() string {
	return s.String()
}

func (s *ClusterStateChangeReason) GetCode() *string {
	return s.Code
}

func (s *ClusterStateChangeReason) GetMessage() *string {
	return s.Message
}

func (s *ClusterStateChangeReason) SetCode(v string) *ClusterStateChangeReason {
	s.Code = &v
	return s
}

func (s *ClusterStateChangeReason) SetMessage(v string) *ClusterStateChangeReason {
	s.Message = &v
	return s
}

func (s *ClusterStateChangeReason) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteControlPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclUuid(v string) *DeleteControlPolicyRequest
	GetAclUuid() *string
	SetDirection(v string) *DeleteControlPolicyRequest
	GetDirection() *string
	SetLang(v string) *DeleteControlPolicyRequest
	GetLang() *string
	SetSourceIp(v string) *DeleteControlPolicyRequest
	GetSourceIp() *string
}

type DeleteControlPolicyRequest struct {
	// The unique ID of the access control policy.
	//
	// To delete an access control policy, you must provide the unique ID of the policy. You can call the [DescribeControlPolicy](https://help.aliyun.com/document_detail/138866.html) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 00281255-d220-4db1-8f4f-c4df221ad84c
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// The traffic direction controlled by the access control policy.
	//
	// Valid values:
	//
	// - **in**: inbound traffic.
	//
	// - **out**: outbound traffic.
	//
	// example:
	//
	// in
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The language of the request and response.
	//
	// Valid values:
	//
	// - **zh*	- (default): Chinese
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Deprecated
	//
	// The source IP address of the traffic.
	//
	// example:
	//
	// 192.0.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DeleteControlPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteControlPolicyRequest) GetAclUuid() *string {
	return s.AclUuid
}

func (s *DeleteControlPolicyRequest) GetDirection() *string {
	return s.Direction
}

func (s *DeleteControlPolicyRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteControlPolicyRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DeleteControlPolicyRequest) SetAclUuid(v string) *DeleteControlPolicyRequest {
	s.AclUuid = &v
	return s
}

func (s *DeleteControlPolicyRequest) SetDirection(v string) *DeleteControlPolicyRequest {
	s.Direction = &v
	return s
}

func (s *DeleteControlPolicyRequest) SetLang(v string) *DeleteControlPolicyRequest {
	s.Lang = &v
	return s
}

func (s *DeleteControlPolicyRequest) SetSourceIp(v string) *DeleteControlPolicyRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteControlPolicyRequest) Validate() error {
	return dara.Validate(s)
}

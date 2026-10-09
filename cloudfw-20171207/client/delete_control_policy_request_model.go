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
	SetClientToken(v string) *DeleteControlPolicyRequest
	GetClientToken() *string
	SetDirection(v string) *DeleteControlPolicyRequest
	GetDirection() *string
	SetDryRun(v bool) *DeleteControlPolicyRequest
	GetDryRun() *bool
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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token. Make sure that the token is unique among different requests. The token must be a string that is case-sensitive and matches the regular expression [0-9a-zA-Z-_]{1,64}. We recommend that you use a UUID. The server ensures idempotence within the validity period of 600 seconds. If you send a repeated request with the same client token and the same business parameters, the server returns the same response as the first request.
	//
	// example:
	//
	// dedeaxfedxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The traffic direction controlled by the access control policy.
	//
	// Valid values:
	//
	// - **in**: inbound traffic
	//
	// - **out**: outbound traffic
	//
	// example:
	//
	// in
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// Specifies whether to only precheck the request. If you set this parameter to true, the system only performs prechecks on parameter validity, identity permissions, resource existence, quota limits, and dependencies. The system does not create, update, or delete actual resources, trigger actual asynchronous traffic diversion tasks, or generate downstream side effects such as billing, notifications, or callbacks. If the precheck is successful, the response includes DryRun=true, which distinguishes it from the response of an actual call.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The language of the request and response.
	//
	// Valid values:
	//
	// - **zh*	- (default): Chinese
	//
	// - **en**: English
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

func (s *DeleteControlPolicyRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteControlPolicyRequest) GetDirection() *string {
	return s.Direction
}

func (s *DeleteControlPolicyRequest) GetDryRun() *bool {
	return s.DryRun
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

func (s *DeleteControlPolicyRequest) SetClientToken(v string) *DeleteControlPolicyRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteControlPolicyRequest) SetDirection(v string) *DeleteControlPolicyRequest {
	s.Direction = &v
	return s
}

func (s *DeleteControlPolicyRequest) SetDryRun(v bool) *DeleteControlPolicyRequest {
	s.DryRun = &v
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

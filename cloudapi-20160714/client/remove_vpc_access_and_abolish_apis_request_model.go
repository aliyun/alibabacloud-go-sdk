// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveVpcAccessAndAbolishApisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *RemoveVpcAccessAndAbolishApisRequest
	GetInstanceId() *string
	SetNeedBatchWork(v bool) *RemoveVpcAccessAndAbolishApisRequest
	GetNeedBatchWork() *bool
	SetPort(v int32) *RemoveVpcAccessAndAbolishApisRequest
	GetPort() *int32
	SetSecurityToken(v string) *RemoveVpcAccessAndAbolishApisRequest
	GetSecurityToken() *string
	SetVpcId(v string) *RemoveVpcAccessAndAbolishApisRequest
	GetVpcId() *string
}

type RemoveVpcAccessAndAbolishApisRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// i-uf6iaale3gfef9t9cb41
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// true
	NeedBatchWork *bool `json:"NeedBatchWork,omitempty" xml:"NeedBatchWork,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 8080
	Port          *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp1iw82phcgkvupgfv0o8
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s RemoveVpcAccessAndAbolishApisRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveVpcAccessAndAbolishApisRequest) GoString() string {
	return s.String()
}

func (s *RemoveVpcAccessAndAbolishApisRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RemoveVpcAccessAndAbolishApisRequest) GetNeedBatchWork() *bool {
	return s.NeedBatchWork
}

func (s *RemoveVpcAccessAndAbolishApisRequest) GetPort() *int32 {
	return s.Port
}

func (s *RemoveVpcAccessAndAbolishApisRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *RemoveVpcAccessAndAbolishApisRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *RemoveVpcAccessAndAbolishApisRequest) SetInstanceId(v string) *RemoveVpcAccessAndAbolishApisRequest {
	s.InstanceId = &v
	return s
}

func (s *RemoveVpcAccessAndAbolishApisRequest) SetNeedBatchWork(v bool) *RemoveVpcAccessAndAbolishApisRequest {
	s.NeedBatchWork = &v
	return s
}

func (s *RemoveVpcAccessAndAbolishApisRequest) SetPort(v int32) *RemoveVpcAccessAndAbolishApisRequest {
	s.Port = &v
	return s
}

func (s *RemoveVpcAccessAndAbolishApisRequest) SetSecurityToken(v string) *RemoveVpcAccessAndAbolishApisRequest {
	s.SecurityToken = &v
	return s
}

func (s *RemoveVpcAccessAndAbolishApisRequest) SetVpcId(v string) *RemoveVpcAccessAndAbolishApisRequest {
	s.VpcId = &v
	return s
}

func (s *RemoveVpcAccessAndAbolishApisRequest) Validate() error {
	return dara.Validate(s)
}

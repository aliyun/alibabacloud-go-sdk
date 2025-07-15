// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeployApiTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOperationUid(v string) *DescribeDeployApiTaskRequest
	GetOperationUid() *string
	SetSecurityToken(v string) *DescribeDeployApiTaskRequest
	GetSecurityToken() *string
}

type DescribeDeployApiTaskRequest struct {
	// The task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 51524fb8f12846d694d0a1de9a0cf274
	OperationUid  *string `json:"OperationUid,omitempty" xml:"OperationUid,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeDeployApiTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeployApiTaskRequest) GoString() string {
	return s.String()
}

func (s *DescribeDeployApiTaskRequest) GetOperationUid() *string {
	return s.OperationUid
}

func (s *DescribeDeployApiTaskRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeDeployApiTaskRequest) SetOperationUid(v string) *DescribeDeployApiTaskRequest {
	s.OperationUid = &v
	return s
}

func (s *DescribeDeployApiTaskRequest) SetSecurityToken(v string) *DescribeDeployApiTaskRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeDeployApiTaskRequest) Validate() error {
	return dara.Validate(s)
}

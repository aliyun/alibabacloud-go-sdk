// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUpdateVpcInfoTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOperationUid(v string) *DescribeUpdateVpcInfoTaskRequest
	GetOperationUid() *string
	SetSecurityToken(v string) *DescribeUpdateVpcInfoTaskRequest
	GetSecurityToken() *string
}

type DescribeUpdateVpcInfoTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 7d341787517a47afaaef9cc1bdb7acce
	OperationUid  *string `json:"OperationUid,omitempty" xml:"OperationUid,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeUpdateVpcInfoTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUpdateVpcInfoTaskRequest) GoString() string {
	return s.String()
}

func (s *DescribeUpdateVpcInfoTaskRequest) GetOperationUid() *string {
	return s.OperationUid
}

func (s *DescribeUpdateVpcInfoTaskRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeUpdateVpcInfoTaskRequest) SetOperationUid(v string) *DescribeUpdateVpcInfoTaskRequest {
	s.OperationUid = &v
	return s
}

func (s *DescribeUpdateVpcInfoTaskRequest) SetSecurityToken(v string) *DescribeUpdateVpcInfoTaskRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeUpdateVpcInfoTaskRequest) Validate() error {
	return dara.Validate(s)
}

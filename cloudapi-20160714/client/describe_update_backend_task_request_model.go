// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUpdateBackendTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOperationUid(v string) *DescribeUpdateBackendTaskRequest
	GetOperationUid() *string
	SetSecurityToken(v string) *DescribeUpdateBackendTaskRequest
	GetSecurityToken() *string
}

type DescribeUpdateBackendTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4e13c3e0c44c4a4ebb5231264eeb9bc1
	OperationUid  *string `json:"OperationUid,omitempty" xml:"OperationUid,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeUpdateBackendTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUpdateBackendTaskRequest) GoString() string {
	return s.String()
}

func (s *DescribeUpdateBackendTaskRequest) GetOperationUid() *string {
	return s.OperationUid
}

func (s *DescribeUpdateBackendTaskRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeUpdateBackendTaskRequest) SetOperationUid(v string) *DescribeUpdateBackendTaskRequest {
	s.OperationUid = &v
	return s
}

func (s *DescribeUpdateBackendTaskRequest) SetSecurityToken(v string) *DescribeUpdateBackendTaskRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeUpdateBackendTaskRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAbolishApiTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOperationUid(v string) *DescribeAbolishApiTaskRequest
	GetOperationUid() *string
	SetSecurityToken(v string) *DescribeAbolishApiTaskRequest
	GetSecurityToken() *string
}

type DescribeAbolishApiTaskRequest struct {
	// The ID of the unpublishing operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// fc54****dd4c4***ad7edd7****39401
	OperationUid  *string `json:"OperationUid,omitempty" xml:"OperationUid,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeAbolishApiTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAbolishApiTaskRequest) GoString() string {
	return s.String()
}

func (s *DescribeAbolishApiTaskRequest) GetOperationUid() *string {
	return s.OperationUid
}

func (s *DescribeAbolishApiTaskRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeAbolishApiTaskRequest) SetOperationUid(v string) *DescribeAbolishApiTaskRequest {
	s.OperationUid = &v
	return s
}

func (s *DescribeAbolishApiTaskRequest) SetSecurityToken(v string) *DescribeAbolishApiTaskRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeAbolishApiTaskRequest) Validate() error {
	return dara.Validate(s)
}

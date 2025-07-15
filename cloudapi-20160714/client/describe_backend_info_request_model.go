// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackendInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackendId(v string) *DescribeBackendInfoRequest
	GetBackendId() *string
	SetSecurityToken(v string) *DescribeBackendInfoRequest
	GetSecurityToken() *string
}

type DescribeBackendInfoRequest struct {
	// The ID of the backend service.
	//
	// example:
	//
	// 739b68751c0b4e899e04d0c92b6d0be7
	BackendId     *string `json:"BackendId,omitempty" xml:"BackendId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeBackendInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackendInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackendInfoRequest) GetBackendId() *string {
	return s.BackendId
}

func (s *DescribeBackendInfoRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeBackendInfoRequest) SetBackendId(v string) *DescribeBackendInfoRequest {
	s.BackendId = &v
	return s
}

func (s *DescribeBackendInfoRequest) SetSecurityToken(v string) *DescribeBackendInfoRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeBackendInfoRequest) Validate() error {
	return dara.Validate(s)
}

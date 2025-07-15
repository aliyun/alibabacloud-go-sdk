// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSignaturesByApiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiId(v string) *DescribeSignaturesByApiRequest
	GetApiId() *string
	SetGroupId(v string) *DescribeSignaturesByApiRequest
	GetGroupId() *string
	SetSecurityToken(v string) *DescribeSignaturesByApiRequest
	GetSecurityToken() *string
	SetStageName(v string) *DescribeSignaturesByApiRequest
	GetStageName() *string
}

type DescribeSignaturesByApiRequest struct {
	// The ID of the API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3b81fd160f5645e097cc8855d75a1cf6
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The ID of the group to which the API belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0009db9c828549768a200320714b8930
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The runtime environment. Valid values:
	//
	// 	- **RELEASE**
	//
	// 	- **TEST**
	//
	// This parameter is required.
	//
	// example:
	//
	// TEST
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s DescribeSignaturesByApiRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSignaturesByApiRequest) GoString() string {
	return s.String()
}

func (s *DescribeSignaturesByApiRequest) GetApiId() *string {
	return s.ApiId
}

func (s *DescribeSignaturesByApiRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeSignaturesByApiRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeSignaturesByApiRequest) GetStageName() *string {
	return s.StageName
}

func (s *DescribeSignaturesByApiRequest) SetApiId(v string) *DescribeSignaturesByApiRequest {
	s.ApiId = &v
	return s
}

func (s *DescribeSignaturesByApiRequest) SetGroupId(v string) *DescribeSignaturesByApiRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeSignaturesByApiRequest) SetSecurityToken(v string) *DescribeSignaturesByApiRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeSignaturesByApiRequest) SetStageName(v string) *DescribeSignaturesByApiRequest {
	s.StageName = &v
	return s
}

func (s *DescribeSignaturesByApiRequest) Validate() error {
	return dara.Validate(s)
}

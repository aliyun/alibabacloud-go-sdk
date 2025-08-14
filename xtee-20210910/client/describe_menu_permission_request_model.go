// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMenuPermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeMenuPermissionRequest
	GetLang() *string
	SetPermissionType(v string) *DescribeMenuPermissionRequest
	GetPermissionType() *string
	SetRegId(v string) *DescribeMenuPermissionRequest
	GetRegId() *string
}

type DescribeMenuPermissionRequest struct {
	// Sets the language type for requests and responses, with a default value of **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Permission type
	//
	// example:
	//
	// MENU
	PermissionType *string `json:"permissionType,omitempty" xml:"permissionType,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeMenuPermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMenuPermissionRequest) GoString() string {
	return s.String()
}

func (s *DescribeMenuPermissionRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeMenuPermissionRequest) GetPermissionType() *string {
	return s.PermissionType
}

func (s *DescribeMenuPermissionRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeMenuPermissionRequest) SetLang(v string) *DescribeMenuPermissionRequest {
	s.Lang = &v
	return s
}

func (s *DescribeMenuPermissionRequest) SetPermissionType(v string) *DescribeMenuPermissionRequest {
	s.PermissionType = &v
	return s
}

func (s *DescribeMenuPermissionRequest) SetRegId(v string) *DescribeMenuPermissionRequest {
	s.RegId = &v
	return s
}

func (s *DescribeMenuPermissionRequest) Validate() error {
	return dara.Validate(s)
}

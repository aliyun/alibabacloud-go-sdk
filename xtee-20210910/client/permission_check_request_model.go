// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPermissionCheckRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *PermissionCheckRequest
	GetLang() *string
	SetRegId(v string) *PermissionCheckRequest
	GetRegId() *string
}

type PermissionCheckRequest struct {
	// Sets the language type for requests and received messages, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Region code.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"RegId,omitempty" xml:"RegId,omitempty"`
}

func (s PermissionCheckRequest) String() string {
	return dara.Prettify(s)
}

func (s PermissionCheckRequest) GoString() string {
	return s.String()
}

func (s *PermissionCheckRequest) GetLang() *string {
	return s.Lang
}

func (s *PermissionCheckRequest) GetRegId() *string {
	return s.RegId
}

func (s *PermissionCheckRequest) SetLang(v string) *PermissionCheckRequest {
	s.Lang = &v
	return s
}

func (s *PermissionCheckRequest) SetRegId(v string) *PermissionCheckRequest {
	s.RegId = &v
	return s
}

func (s *PermissionCheckRequest) Validate() error {
	return dara.Validate(s)
}

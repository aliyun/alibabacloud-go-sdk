// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAclCheckRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclType(v string) *CreateAclCheckRequest
	GetAclType() *string
	SetCheckNames(v []*string) *CreateAclCheckRequest
	GetCheckNames() []*string
	SetLang(v string) *CreateAclCheckRequest
	GetLang() *string
}

type CreateAclCheckRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// Internet
	AclType    *string   `json:"AclType,omitempty" xml:"AclType,omitempty"`
	CheckNames []*string `json:"CheckNames,omitempty" xml:"CheckNames,omitempty" type:"Repeated"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s CreateAclCheckRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAclCheckRequest) GoString() string {
	return s.String()
}

func (s *CreateAclCheckRequest) GetAclType() *string {
	return s.AclType
}

func (s *CreateAclCheckRequest) GetCheckNames() []*string {
	return s.CheckNames
}

func (s *CreateAclCheckRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateAclCheckRequest) SetAclType(v string) *CreateAclCheckRequest {
	s.AclType = &v
	return s
}

func (s *CreateAclCheckRequest) SetCheckNames(v []*string) *CreateAclCheckRequest {
	s.CheckNames = v
	return s
}

func (s *CreateAclCheckRequest) SetLang(v string) *CreateAclCheckRequest {
	s.Lang = &v
	return s
}

func (s *CreateAclCheckRequest) Validate() error {
	return dara.Validate(s)
}

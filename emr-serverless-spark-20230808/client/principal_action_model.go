// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPrincipalAction interface {
	dara.Model
	String() string
	GoString() string
	SetActionArn(v string) *PrincipalAction
	GetActionArn() *string
	SetPrincipalArn(v string) *PrincipalAction
	GetPrincipalArn() *string
}

type PrincipalAction struct {
	// example:
	//
	// acs:emr::workspaceId:action/create_queue
	ActionArn *string `json:"actionArn,omitempty" xml:"actionArn,omitempty"`
	// example:
	//
	// acs:emr::workspaceId:user/237593691541622267
	PrincipalArn *string `json:"principalArn,omitempty" xml:"principalArn,omitempty"`
}

func (s PrincipalAction) String() string {
	return dara.Prettify(s)
}

func (s PrincipalAction) GoString() string {
	return s.String()
}

func (s *PrincipalAction) GetActionArn() *string {
	return s.ActionArn
}

func (s *PrincipalAction) GetPrincipalArn() *string {
	return s.PrincipalArn
}

func (s *PrincipalAction) SetActionArn(v string) *PrincipalAction {
	s.ActionArn = &v
	return s
}

func (s *PrincipalAction) SetPrincipalArn(v string) *PrincipalAction {
	s.PrincipalArn = &v
	return s
}

func (s *PrincipalAction) Validate() error {
	return dara.Validate(s)
}

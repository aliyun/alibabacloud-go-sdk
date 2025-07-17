// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckServiceLinkedRoleResult interface {
	dara.Model
	String() string
	GoString() string
	SetExisted(v bool) *CheckServiceLinkedRoleResult
	GetExisted() *bool
}

type CheckServiceLinkedRoleResult struct {
	Existed *bool `json:"existed,omitempty" xml:"existed,omitempty"`
}

func (s CheckServiceLinkedRoleResult) String() string {
	return dara.Prettify(s)
}

func (s CheckServiceLinkedRoleResult) GoString() string {
	return s.String()
}

func (s *CheckServiceLinkedRoleResult) GetExisted() *bool {
	return s.Existed
}

func (s *CheckServiceLinkedRoleResult) SetExisted(v bool) *CheckServiceLinkedRoleResult {
	s.Existed = &v
	return s
}

func (s *CheckServiceLinkedRoleResult) Validate() error {
	return dara.Validate(s)
}

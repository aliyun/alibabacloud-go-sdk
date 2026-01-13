// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iManageAccountResult interface {
	dara.Model
	String() string
	GoString() string
	SetSuccess(v bool) *ManageAccountResult
	GetSuccess() *bool
}

type ManageAccountResult struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ManageAccountResult) String() string {
	return dara.Prettify(s)
}

func (s ManageAccountResult) GoString() string {
	return s.String()
}

func (s *ManageAccountResult) GetSuccess() *bool {
	return s.Success
}

func (s *ManageAccountResult) SetSuccess(v bool) *ManageAccountResult {
	s.Success = &v
	return s
}

func (s *ManageAccountResult) Validate() error {
	return dara.Validate(s)
}

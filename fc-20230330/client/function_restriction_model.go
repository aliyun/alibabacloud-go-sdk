// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFunctionRestriction interface {
	dara.Model
	String() string
	GoString() string
	SetDisable(v bool) *FunctionRestriction
	GetDisable() *bool
	SetLastModifiedTime(v string) *FunctionRestriction
	GetLastModifiedTime() *string
	SetReason(v string) *FunctionRestriction
	GetReason() *string
}

type FunctionRestriction struct {
	Disable          *bool   `json:"disable,omitempty" xml:"disable,omitempty"`
	LastModifiedTime *string `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	Reason           *string `json:"reason,omitempty" xml:"reason,omitempty"`
}

func (s FunctionRestriction) String() string {
	return dara.Prettify(s)
}

func (s FunctionRestriction) GoString() string {
	return s.String()
}

func (s *FunctionRestriction) GetDisable() *bool {
	return s.Disable
}

func (s *FunctionRestriction) GetLastModifiedTime() *string {
	return s.LastModifiedTime
}

func (s *FunctionRestriction) GetReason() *string {
	return s.Reason
}

func (s *FunctionRestriction) SetDisable(v bool) *FunctionRestriction {
	s.Disable = &v
	return s
}

func (s *FunctionRestriction) SetLastModifiedTime(v string) *FunctionRestriction {
	s.LastModifiedTime = &v
	return s
}

func (s *FunctionRestriction) SetReason(v string) *FunctionRestriction {
	s.Reason = &v
	return s
}

func (s *FunctionRestriction) Validate() error {
	return dara.Validate(s)
}

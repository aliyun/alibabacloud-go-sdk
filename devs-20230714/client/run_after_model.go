// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunAfter interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *RunAfter
	GetName() *string
}

type RunAfter struct {
	// example:
	//
	// task-1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s RunAfter) String() string {
	return dara.Prettify(s)
}

func (s RunAfter) GoString() string {
	return s.String()
}

func (s *RunAfter) GetName() *string {
	return s.Name
}

func (s *RunAfter) SetName(v string) *RunAfter {
	s.Name = &v
	return s
}

func (s *RunAfter) Validate() error {
	return dara.Validate(s)
}

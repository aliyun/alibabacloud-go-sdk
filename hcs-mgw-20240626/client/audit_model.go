// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAudit interface {
	dara.Model
	String() string
	GoString() string
	SetLogMode(v string) *Audit
	GetLogMode() *string
}

type Audit struct {
	// example:
	//
	// off
	LogMode *string `json:"LogMode,omitempty" xml:"LogMode,omitempty"`
}

func (s Audit) String() string {
	return dara.Prettify(s)
}

func (s Audit) GoString() string {
	return s.String()
}

func (s *Audit) GetLogMode() *string {
	return s.LogMode
}

func (s *Audit) SetLogMode(v string) *Audit {
	s.LogMode = &v
	return s
}

func (s *Audit) Validate() error {
	return dara.Validate(s)
}

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
	// Specifies whether to push migration logs to Simple Log Service (SLS). To push migration logs to SLS, you must create the AliyunOSSImportSlsAuditRole role and grant the required permissions to the role. Valid values: off, basic, and detail. off: does not push migration logs. basic: pushes only migration failure logs. detail: pushes migration success logs and migration failure logs.
	//
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

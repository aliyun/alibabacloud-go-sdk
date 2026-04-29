// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUnknownThreatDetectProcessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProcessId(v string) *UpdateUnknownThreatDetectProcessRequest
	GetProcessId() *string
	SetRemark(v string) *UpdateUnknownThreatDetectProcessRequest
	GetRemark() *string
}

type UpdateUnknownThreatDetectProcessRequest struct {
	// example:
	//
	// 2026011210040602108912721603151374234
	ProcessId *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	// example:
	//
	// remark
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s UpdateUnknownThreatDetectProcessRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUnknownThreatDetectProcessRequest) GoString() string {
	return s.String()
}

func (s *UpdateUnknownThreatDetectProcessRequest) GetProcessId() *string {
	return s.ProcessId
}

func (s *UpdateUnknownThreatDetectProcessRequest) GetRemark() *string {
	return s.Remark
}

func (s *UpdateUnknownThreatDetectProcessRequest) SetProcessId(v string) *UpdateUnknownThreatDetectProcessRequest {
	s.ProcessId = &v
	return s
}

func (s *UpdateUnknownThreatDetectProcessRequest) SetRemark(v string) *UpdateUnknownThreatDetectProcessRequest {
	s.Remark = &v
	return s
}

func (s *UpdateUnknownThreatDetectProcessRequest) Validate() error {
	return dara.Validate(s)
}

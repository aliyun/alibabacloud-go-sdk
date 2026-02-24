// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDigitalEmployeeSkillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVersion(v string) *GetDigitalEmployeeSkillRequest
	GetVersion() *string
}

type GetDigitalEmployeeSkillRequest struct {
	// example:
	//
	// 1770386951147366810
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetDigitalEmployeeSkillRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDigitalEmployeeSkillRequest) GoString() string {
	return s.String()
}

func (s *GetDigitalEmployeeSkillRequest) GetVersion() *string {
	return s.Version
}

func (s *GetDigitalEmployeeSkillRequest) SetVersion(v string) *GetDigitalEmployeeSkillRequest {
	s.Version = &v
	return s
}

func (s *GetDigitalEmployeeSkillRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDigitalEmployeeSkillRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteDigitalEmployeeSkillRequest struct {
}

func (s DeleteDigitalEmployeeSkillRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDigitalEmployeeSkillRequest) GoString() string {
	return s.String()
}

func (s *DeleteDigitalEmployeeSkillRequest) Validate() error {
	return dara.Validate(s)
}

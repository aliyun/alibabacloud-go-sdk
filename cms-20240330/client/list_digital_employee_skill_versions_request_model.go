// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDigitalEmployeeSkillVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
}

type ListDigitalEmployeeSkillVersionsRequest struct {
}

func (s ListDigitalEmployeeSkillVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDigitalEmployeeSkillVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListDigitalEmployeeSkillVersionsRequest) Validate() error {
	return dara.Validate(s)
}

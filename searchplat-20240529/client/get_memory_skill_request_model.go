// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMemorySkillRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetMemorySkillRequest struct {
}

func (s GetMemorySkillRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMemorySkillRequest) GoString() string {
	return s.String()
}

func (s *GetMemorySkillRequest) Validate() error {
	return dara.Validate(s)
}

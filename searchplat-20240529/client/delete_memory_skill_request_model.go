// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMemorySkillRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteMemorySkillRequest struct {
}

func (s DeleteMemorySkillRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMemorySkillRequest) GoString() string {
	return s.String()
}

func (s *DeleteMemorySkillRequest) Validate() error {
	return dara.Validate(s)
}

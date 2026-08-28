// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSkillRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteSkillRequest struct {
}

func (s DeleteSkillRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSkillRequest) GoString() string {
	return s.String()
}

func (s *DeleteSkillRequest) Validate() error {
	return dara.Validate(s)
}

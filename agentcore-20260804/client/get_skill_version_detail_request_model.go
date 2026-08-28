// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillVersionDetailRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetSkillVersionDetailRequest struct {
}

func (s GetSkillVersionDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSkillVersionDetailRequest) GoString() string {
	return s.String()
}

func (s *GetSkillVersionDetailRequest) Validate() error {
	return dara.Validate(s)
}

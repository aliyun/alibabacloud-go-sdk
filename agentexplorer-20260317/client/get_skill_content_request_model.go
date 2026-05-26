// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillContentRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetSkillContentRequest struct {
}

func (s GetSkillContentRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSkillContentRequest) GoString() string {
	return s.String()
}

func (s *GetSkillContentRequest) Validate() error {
	return dara.Validate(s)
}

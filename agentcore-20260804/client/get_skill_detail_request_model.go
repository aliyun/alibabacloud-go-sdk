// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillDetailRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetSkillDetailRequest struct {
}

func (s GetSkillDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSkillDetailRequest) GoString() string {
	return s.String()
}

func (s *GetSkillDetailRequest) Validate() error {
	return dara.Validate(s)
}

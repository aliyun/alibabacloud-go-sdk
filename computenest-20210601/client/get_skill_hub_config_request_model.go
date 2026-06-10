// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillHubConfigRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetSkillHubConfigRequest struct {
}

func (s GetSkillHubConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSkillHubConfigRequest) GoString() string {
	return s.String()
}

func (s *GetSkillHubConfigRequest) Validate() error {
	return dara.Validate(s)
}

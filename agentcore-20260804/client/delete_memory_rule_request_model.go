// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMemoryRuleRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteMemoryRuleRequest struct {
}

func (s DeleteMemoryRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMemoryRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteMemoryRuleRequest) Validate() error {
	return dara.Validate(s)
}

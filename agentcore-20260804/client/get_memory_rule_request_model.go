// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMemoryRuleRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetMemoryRuleRequest struct {
}

func (s GetMemoryRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryRuleRequest) GoString() string {
	return s.String()
}

func (s *GetMemoryRuleRequest) Validate() error {
	return dara.Validate(s)
}

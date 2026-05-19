// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFileProtectClientRuleFileTypeRequest interface {
	dara.Model
	String() string
	GoString() string
}

type ListFileProtectClientRuleFileTypeRequest struct {
}

func (s ListFileProtectClientRuleFileTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFileProtectClientRuleFileTypeRequest) GoString() string {
	return s.String()
}

func (s *ListFileProtectClientRuleFileTypeRequest) Validate() error {
	return dara.Validate(s)
}

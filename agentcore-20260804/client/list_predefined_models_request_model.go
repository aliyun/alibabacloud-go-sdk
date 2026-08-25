// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPredefinedModelsRequest interface {
	dara.Model
	String() string
	GoString() string
}

type ListPredefinedModelsRequest struct {
}

func (s ListPredefinedModelsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPredefinedModelsRequest) GoString() string {
	return s.String()
}

func (s *ListPredefinedModelsRequest) Validate() error {
	return dara.Validate(s)
}

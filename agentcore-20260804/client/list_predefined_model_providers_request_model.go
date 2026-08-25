// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPredefinedModelProvidersRequest interface {
	dara.Model
	String() string
	GoString() string
}

type ListPredefinedModelProvidersRequest struct {
}

func (s ListPredefinedModelProvidersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPredefinedModelProvidersRequest) GoString() string {
	return s.String()
}

func (s *ListPredefinedModelProvidersRequest) Validate() error {
	return dara.Validate(s)
}

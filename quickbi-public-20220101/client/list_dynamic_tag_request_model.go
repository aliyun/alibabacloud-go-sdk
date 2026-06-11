// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDynamicTagRequest interface {
	dara.Model
	String() string
	GoString() string
}

type ListDynamicTagRequest struct {
}

func (s ListDynamicTagRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDynamicTagRequest) GoString() string {
	return s.String()
}

func (s *ListDynamicTagRequest) Validate() error {
	return dara.Validate(s)
}

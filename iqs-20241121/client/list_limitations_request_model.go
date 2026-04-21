// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLimitationsRequest interface {
	dara.Model
	String() string
	GoString() string
}

type ListLimitationsRequest struct {
}

func (s ListLimitationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLimitationsRequest) GoString() string {
	return s.String()
}

func (s *ListLimitationsRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPermissionsRequest interface {
	dara.Model
	String() string
	GoString() string
}

type ListPermissionsRequest struct {
}

func (s ListPermissionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPermissionsRequest) GoString() string {
	return s.String()
}

func (s *ListPermissionsRequest) Validate() error {
	return dara.Validate(s)
}

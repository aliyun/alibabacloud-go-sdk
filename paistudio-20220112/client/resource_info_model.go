// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResourceInfo interface {
	dara.Model
	String() string
	GoString() string
}

type ResourceInfo struct {
}

func (s ResourceInfo) String() string {
	return dara.Prettify(s)
}

func (s ResourceInfo) GoString() string {
	return s.String()
}

func (s *ResourceInfo) Validate() error {
	return dara.Validate(s)
}

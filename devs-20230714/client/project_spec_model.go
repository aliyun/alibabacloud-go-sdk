// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProjectSpec interface {
	dara.Model
	String() string
	GoString() string
}

type ProjectSpec struct {
}

func (s ProjectSpec) String() string {
	return dara.Prettify(s)
}

func (s ProjectSpec) GoString() string {
	return s.String()
}

func (s *ProjectSpec) Validate() error {
	return dara.Validate(s)
}

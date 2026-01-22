// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iItemValues interface {
	dara.Model
	String() string
	GoString() string
}

type ItemValues struct {
}

func (s ItemValues) String() string {
	return dara.Prettify(s)
}

func (s ItemValues) GoString() string {
	return s.String()
}

func (s *ItemValues) Validate() error {
	return dara.Validate(s)
}

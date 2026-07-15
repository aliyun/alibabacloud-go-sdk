// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePluginClassRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeletePluginClassRequest struct {
}

func (s DeletePluginClassRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePluginClassRequest) GoString() string {
	return s.String()
}

func (s *DeletePluginClassRequest) Validate() error {
	return dara.Validate(s)
}

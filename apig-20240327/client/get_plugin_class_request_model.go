// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPluginClassRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetPluginClassRequest struct {
}

func (s GetPluginClassRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPluginClassRequest) GoString() string {
	return s.String()
}

func (s *GetPluginClassRequest) Validate() error {
	return dara.Validate(s)
}

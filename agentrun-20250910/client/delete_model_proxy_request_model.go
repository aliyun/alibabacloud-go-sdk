// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteModelProxyRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteModelProxyRequest struct {
}

func (s DeleteModelProxyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteModelProxyRequest) GoString() string {
	return s.String()
}

func (s *DeleteModelProxyRequest) Validate() error {
	return dara.Validate(s)
}

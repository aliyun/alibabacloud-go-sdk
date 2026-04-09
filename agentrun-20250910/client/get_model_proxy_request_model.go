// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModelProxyRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetModelProxyRequest struct {
}

func (s GetModelProxyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetModelProxyRequest) GoString() string {
	return s.String()
}

func (s *GetModelProxyRequest) Validate() error {
	return dara.Validate(s)
}

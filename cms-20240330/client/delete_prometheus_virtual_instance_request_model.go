// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePrometheusVirtualInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeletePrometheusVirtualInstanceRequest struct {
}

func (s DeletePrometheusVirtualInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePrometheusVirtualInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeletePrometheusVirtualInstanceRequest) Validate() error {
	return dara.Validate(s)
}

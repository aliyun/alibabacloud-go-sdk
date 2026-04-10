// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePrometheusInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeletePrometheusInstanceRequest struct {
}

func (s DeletePrometheusInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePrometheusInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeletePrometheusInstanceRequest) Validate() error {
	return dara.Validate(s)
}

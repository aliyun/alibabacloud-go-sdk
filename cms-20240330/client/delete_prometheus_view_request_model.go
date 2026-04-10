// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePrometheusViewRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeletePrometheusViewRequest struct {
}

func (s DeletePrometheusViewRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePrometheusViewRequest) GoString() string {
	return s.String()
}

func (s *DeletePrometheusViewRequest) Validate() error {
	return dara.Validate(s)
}

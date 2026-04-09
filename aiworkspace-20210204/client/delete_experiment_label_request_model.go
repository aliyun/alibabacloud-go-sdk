// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExperimentLabelRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteExperimentLabelRequest struct {
}

func (s DeleteExperimentLabelRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteExperimentLabelRequest) GoString() string {
	return s.String()
}

func (s *DeleteExperimentLabelRequest) Validate() error {
	return dara.Validate(s)
}

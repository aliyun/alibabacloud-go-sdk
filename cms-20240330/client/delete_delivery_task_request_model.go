// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDeliveryTaskRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteDeliveryTaskRequest struct {
}

func (s DeleteDeliveryTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDeliveryTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteDeliveryTaskRequest) Validate() error {
	return dara.Validate(s)
}

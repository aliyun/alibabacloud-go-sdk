// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeliveryTaskRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetDeliveryTaskRequest struct {
}

func (s GetDeliveryTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDeliveryTaskRequest) GoString() string {
	return s.String()
}

func (s *GetDeliveryTaskRequest) Validate() error {
	return dara.Validate(s)
}

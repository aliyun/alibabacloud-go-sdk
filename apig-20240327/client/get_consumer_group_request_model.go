// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConsumerGroupRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetConsumerGroupRequest struct {
}

func (s GetConsumerGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s GetConsumerGroupRequest) GoString() string {
	return s.String()
}

func (s *GetConsumerGroupRequest) Validate() error {
	return dara.Validate(s)
}

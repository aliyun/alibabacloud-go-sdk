// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTriggerMmsTimerRequest interface {
	dara.Model
	String() string
	GoString() string
}

type TriggerMmsTimerRequest struct {
}

func (s TriggerMmsTimerRequest) String() string {
	return dara.Prettify(s)
}

func (s TriggerMmsTimerRequest) GoString() string {
	return s.String()
}

func (s *TriggerMmsTimerRequest) Validate() error {
	return dara.Validate(s)
}

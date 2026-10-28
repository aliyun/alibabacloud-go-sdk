// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOncallScheduleRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetOncallScheduleRequest struct {
}

func (s GetOncallScheduleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOncallScheduleRequest) GoString() string {
	return s.String()
}

func (s *GetOncallScheduleRequest) Validate() error {
	return dara.Validate(s)
}

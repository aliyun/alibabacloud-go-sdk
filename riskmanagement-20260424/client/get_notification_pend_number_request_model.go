// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNotificationPendNumberRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetNotificationPendNumberRequest struct {
}

func (s GetNotificationPendNumberRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNotificationPendNumberRequest) GoString() string {
	return s.String()
}

func (s *GetNotificationPendNumberRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNotificationContactsRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetNotificationContactsRequest struct {
}

func (s GetNotificationContactsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNotificationContactsRequest) GoString() string {
	return s.String()
}

func (s *GetNotificationContactsRequest) Validate() error {
	return dara.Validate(s)
}

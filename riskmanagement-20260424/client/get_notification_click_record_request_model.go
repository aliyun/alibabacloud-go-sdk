// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNotificationClickRecordRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetNotificationClickRecordRequest struct {
}

func (s GetNotificationClickRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNotificationClickRecordRequest) GoString() string {
	return s.String()
}

func (s *GetNotificationClickRecordRequest) Validate() error {
	return dara.Validate(s)
}

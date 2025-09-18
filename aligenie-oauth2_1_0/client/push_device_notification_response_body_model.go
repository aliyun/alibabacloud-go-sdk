// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushDeviceNotificationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PushDeviceNotificationResponseBody
	GetRequestId() *string
}

type PushDeviceNotificationResponseBody struct {
	// example:
	//
	// 908FA068-529C-0C20-8DB5-63B0EF7CFF1F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s PushDeviceNotificationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PushDeviceNotificationResponseBody) GoString() string {
	return s.String()
}

func (s *PushDeviceNotificationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PushDeviceNotificationResponseBody) SetRequestId(v string) *PushDeviceNotificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *PushDeviceNotificationResponseBody) Validate() error {
	return dara.Validate(s)
}

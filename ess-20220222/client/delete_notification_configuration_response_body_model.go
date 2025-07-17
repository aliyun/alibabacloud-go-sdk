// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNotificationConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteNotificationConfigurationResponseBody
	GetRequestId() *string
}

type DeleteNotificationConfigurationResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteNotificationConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteNotificationConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNotificationConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteNotificationConfigurationResponseBody) SetRequestId(v string) *DeleteNotificationConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNotificationConfigurationResponseBody) Validate() error {
	return dara.Validate(s)
}

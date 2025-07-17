// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNotificationConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyNotificationConfigurationResponseBody
	GetRequestId() *string
}

type ModifyNotificationConfigurationResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyNotificationConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyNotificationConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyNotificationConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyNotificationConfigurationResponseBody) SetRequestId(v string) *ModifyNotificationConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyNotificationConfigurationResponseBody) Validate() error {
	return dara.Validate(s)
}

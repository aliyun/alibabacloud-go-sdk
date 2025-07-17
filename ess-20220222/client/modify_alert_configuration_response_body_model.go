// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAlertConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAlertConfigurationResponseBody
	GetRequestId() *string
}

type ModifyAlertConfigurationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 73469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAlertConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAlertConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAlertConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAlertConfigurationResponseBody) SetRequestId(v string) *ModifyAlertConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAlertConfigurationResponseBody) Validate() error {
	return dara.Validate(s)
}

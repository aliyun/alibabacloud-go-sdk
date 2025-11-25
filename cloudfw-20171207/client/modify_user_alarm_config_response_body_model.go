// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUserAlarmConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyUserAlarmConfigResponseBody
	GetRequestId() *string
}

type ModifyUserAlarmConfigResponseBody struct {
	// example:
	//
	// 3B168A0F-A43D-5FD0-8059-B51BDD6E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyUserAlarmConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyUserAlarmConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyUserAlarmConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyUserAlarmConfigResponseBody) SetRequestId(v string) *ModifyUserAlarmConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyUserAlarmConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

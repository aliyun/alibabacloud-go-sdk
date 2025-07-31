// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchDBInstanceHAResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SwitchDBInstanceHAResponseBody
	GetRequestId() *string
}

type SwitchDBInstanceHAResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 26BD4E5F-BDB4-47BA-B232-413AA78CFA8F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SwitchDBInstanceHAResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SwitchDBInstanceHAResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchDBInstanceHAResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SwitchDBInstanceHAResponseBody) SetRequestId(v string) *SwitchDBInstanceHAResponseBody {
	s.RequestId = &v
	return s
}

func (s *SwitchDBInstanceHAResponseBody) Validate() error {
	return dara.Validate(s)
}

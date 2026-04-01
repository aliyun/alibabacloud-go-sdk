// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchDBInstanceNetTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNewConnectionString(v string) *SwitchDBInstanceNetTypeResponseBody
	GetNewConnectionString() *string
	SetOldConnectionString(v string) *SwitchDBInstanceNetTypeResponseBody
	GetOldConnectionString() *string
	SetRequestId(v string) *SwitchDBInstanceNetTypeResponseBody
	GetRequestId() *string
}

type SwitchDBInstanceNetTypeResponseBody struct {
	// The endpoint that is used to connect to the instance after the switch of endpoints.
	//
	// example:
	//
	// new**********.mysql.rds.aliyuncs.com
	NewConnectionString *string `json:"NewConnectionString,omitempty" xml:"NewConnectionString,omitempty"`
	// The endpoint that is used to connect to the instance before the switch of endpoints.
	//
	// example:
	//
	// rm-bp1**************.mysql.rds.aliyuncs.com
	OldConnectionString *string `json:"OldConnectionString,omitempty" xml:"OldConnectionString,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 65BDA532-28AF-4122-AA39-B382721EEE64
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SwitchDBInstanceNetTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SwitchDBInstanceNetTypeResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchDBInstanceNetTypeResponseBody) GetNewConnectionString() *string {
	return s.NewConnectionString
}

func (s *SwitchDBInstanceNetTypeResponseBody) GetOldConnectionString() *string {
	return s.OldConnectionString
}

func (s *SwitchDBInstanceNetTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SwitchDBInstanceNetTypeResponseBody) SetNewConnectionString(v string) *SwitchDBInstanceNetTypeResponseBody {
	s.NewConnectionString = &v
	return s
}

func (s *SwitchDBInstanceNetTypeResponseBody) SetOldConnectionString(v string) *SwitchDBInstanceNetTypeResponseBody {
	s.OldConnectionString = &v
	return s
}

func (s *SwitchDBInstanceNetTypeResponseBody) SetRequestId(v string) *SwitchDBInstanceNetTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *SwitchDBInstanceNetTypeResponseBody) Validate() error {
	return dara.Validate(s)
}

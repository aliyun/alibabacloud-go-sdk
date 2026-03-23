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
	NewConnectionString *string `json:"NewConnectionString,omitempty" xml:"NewConnectionString,omitempty"`
	OldConnectionString *string `json:"OldConnectionString,omitempty" xml:"OldConnectionString,omitempty"`
	RequestId           *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

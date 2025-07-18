// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchDBInstanceNetTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SwitchDBInstanceNetTypeResponseBody
	GetRequestId() *string
}

type SwitchDBInstanceNetTypeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// FA67B751-2A2D-470C-850B-D6B93699D35C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SwitchDBInstanceNetTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SwitchDBInstanceNetTypeResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchDBInstanceNetTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SwitchDBInstanceNetTypeResponseBody) SetRequestId(v string) *SwitchDBInstanceNetTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *SwitchDBInstanceNetTypeResponseBody) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchDBInstanceVpcResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SwitchDBInstanceVpcResponseBody
	GetRequestId() *string
}

type SwitchDBInstanceVpcResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SwitchDBInstanceVpcResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SwitchDBInstanceVpcResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchDBInstanceVpcResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SwitchDBInstanceVpcResponseBody) SetRequestId(v string) *SwitchDBInstanceVpcResponseBody {
	s.RequestId = &v
	return s
}

func (s *SwitchDBInstanceVpcResponseBody) Validate() error {
	return dara.Validate(s)
}

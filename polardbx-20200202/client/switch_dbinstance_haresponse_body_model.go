// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchDBInstanceHAResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *SwitchDBInstanceHAResponseBody
	GetMessage() *string
	SetRequestId(v string) *SwitchDBInstanceHAResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SwitchDBInstanceHAResponseBody
	GetSuccess() *bool
}

type SwitchDBInstanceHAResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SwitchDBInstanceHAResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SwitchDBInstanceHAResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchDBInstanceHAResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SwitchDBInstanceHAResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SwitchDBInstanceHAResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SwitchDBInstanceHAResponseBody) SetMessage(v string) *SwitchDBInstanceHAResponseBody {
	s.Message = &v
	return s
}

func (s *SwitchDBInstanceHAResponseBody) SetRequestId(v string) *SwitchDBInstanceHAResponseBody {
	s.RequestId = &v
	return s
}

func (s *SwitchDBInstanceHAResponseBody) SetSuccess(v bool) *SwitchDBInstanceHAResponseBody {
	s.Success = &v
	return s
}

func (s *SwitchDBInstanceHAResponseBody) Validate() error {
	return dara.Validate(s)
}

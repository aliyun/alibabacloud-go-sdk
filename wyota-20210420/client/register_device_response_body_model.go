// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterDeviceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RegisterDeviceResponseBody
	GetCode() *string
	SetData(v *RegisterDeviceResponseBodyData) *RegisterDeviceResponseBody
	GetData() *RegisterDeviceResponseBodyData
	SetMessage(v string) *RegisterDeviceResponseBody
	GetMessage() *string
	SetRequestId(v string) *RegisterDeviceResponseBody
	GetRequestId() *string
}

type RegisterDeviceResponseBody struct {
	Code      *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *RegisterDeviceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RegisterDeviceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RegisterDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterDeviceResponseBody) GetCode() *string {
	return s.Code
}

func (s *RegisterDeviceResponseBody) GetData() *RegisterDeviceResponseBodyData {
	return s.Data
}

func (s *RegisterDeviceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RegisterDeviceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RegisterDeviceResponseBody) SetCode(v string) *RegisterDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *RegisterDeviceResponseBody) SetData(v *RegisterDeviceResponseBodyData) *RegisterDeviceResponseBody {
	s.Data = v
	return s
}

func (s *RegisterDeviceResponseBody) SetMessage(v string) *RegisterDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *RegisterDeviceResponseBody) SetRequestId(v string) *RegisterDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RegisterDeviceResponseBody) Validate() error {
	return dara.Validate(s)
}

type RegisterDeviceResponseBodyData struct {
	NewUpgrade *bool   `json:"NewUpgrade,omitempty" xml:"NewUpgrade,omitempty"`
	Uuid       *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s RegisterDeviceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RegisterDeviceResponseBodyData) GoString() string {
	return s.String()
}

func (s *RegisterDeviceResponseBodyData) GetNewUpgrade() *bool {
	return s.NewUpgrade
}

func (s *RegisterDeviceResponseBodyData) GetUuid() *string {
	return s.Uuid
}

func (s *RegisterDeviceResponseBodyData) SetNewUpgrade(v bool) *RegisterDeviceResponseBodyData {
	s.NewUpgrade = &v
	return s
}

func (s *RegisterDeviceResponseBodyData) SetUuid(v string) *RegisterDeviceResponseBodyData {
	s.Uuid = &v
	return s
}

func (s *RegisterDeviceResponseBodyData) Validate() error {
	return dara.Validate(s)
}

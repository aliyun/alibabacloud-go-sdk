// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceOtaAutoStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDeviceOtaAutoStatusResponseBody
	GetCode() *string
	SetData(v *GetDeviceOtaAutoStatusResponseBodyData) *GetDeviceOtaAutoStatusResponseBody
	GetData() *GetDeviceOtaAutoStatusResponseBodyData
	SetMessage(v string) *GetDeviceOtaAutoStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDeviceOtaAutoStatusResponseBody
	GetRequestId() *string
}

type GetDeviceOtaAutoStatusResponseBody struct {
	Code      *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetDeviceOtaAutoStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDeviceOtaAutoStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceOtaAutoStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeviceOtaAutoStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDeviceOtaAutoStatusResponseBody) GetData() *GetDeviceOtaAutoStatusResponseBodyData {
	return s.Data
}

func (s *GetDeviceOtaAutoStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDeviceOtaAutoStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDeviceOtaAutoStatusResponseBody) SetCode(v string) *GetDeviceOtaAutoStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetDeviceOtaAutoStatusResponseBody) SetData(v *GetDeviceOtaAutoStatusResponseBodyData) *GetDeviceOtaAutoStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetDeviceOtaAutoStatusResponseBody) SetMessage(v string) *GetDeviceOtaAutoStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetDeviceOtaAutoStatusResponseBody) SetRequestId(v string) *GetDeviceOtaAutoStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeviceOtaAutoStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDeviceOtaAutoStatusResponseBodyData struct {
	AutoUpdate             *int32  `json:"AutoUpdate,omitempty" xml:"AutoUpdate,omitempty"`
	AutoUpdateTimeSchedule *string `json:"AutoUpdateTimeSchedule,omitempty" xml:"AutoUpdateTimeSchedule,omitempty"`
	ForceUpgrade           *int32  `json:"ForceUpgrade,omitempty" xml:"ForceUpgrade,omitempty"`
	Status                 *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetDeviceOtaAutoStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceOtaAutoStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDeviceOtaAutoStatusResponseBodyData) GetAutoUpdate() *int32 {
	return s.AutoUpdate
}

func (s *GetDeviceOtaAutoStatusResponseBodyData) GetAutoUpdateTimeSchedule() *string {
	return s.AutoUpdateTimeSchedule
}

func (s *GetDeviceOtaAutoStatusResponseBodyData) GetForceUpgrade() *int32 {
	return s.ForceUpgrade
}

func (s *GetDeviceOtaAutoStatusResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *GetDeviceOtaAutoStatusResponseBodyData) SetAutoUpdate(v int32) *GetDeviceOtaAutoStatusResponseBodyData {
	s.AutoUpdate = &v
	return s
}

func (s *GetDeviceOtaAutoStatusResponseBodyData) SetAutoUpdateTimeSchedule(v string) *GetDeviceOtaAutoStatusResponseBodyData {
	s.AutoUpdateTimeSchedule = &v
	return s
}

func (s *GetDeviceOtaAutoStatusResponseBodyData) SetForceUpgrade(v int32) *GetDeviceOtaAutoStatusResponseBodyData {
	s.ForceUpgrade = &v
	return s
}

func (s *GetDeviceOtaAutoStatusResponseBodyData) SetStatus(v int32) *GetDeviceOtaAutoStatusResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetDeviceOtaAutoStatusResponseBodyData) Validate() error {
	return dara.Validate(s)
}

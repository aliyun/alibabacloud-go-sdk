// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAndDoVoipCallForHotelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CheckAndDoVoipCallForHotelResponseBody
	GetCode() *int32
	SetMessage(v string) *CheckAndDoVoipCallForHotelResponseBody
	GetMessage() *string
	SetRequestId(v string) *CheckAndDoVoipCallForHotelResponseBody
	GetRequestId() *string
	SetResult(v *CheckAndDoVoipCallForHotelResponseBodyResult) *CheckAndDoVoipCallForHotelResponseBody
	GetResult() *CheckAndDoVoipCallForHotelResponseBodyResult
}

type CheckAndDoVoipCallForHotelResponseBody struct {
	Code      *int32                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CheckAndDoVoipCallForHotelResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CheckAndDoVoipCallForHotelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckAndDoVoipCallForHotelResponseBody) GoString() string {
	return s.String()
}

func (s *CheckAndDoVoipCallForHotelResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CheckAndDoVoipCallForHotelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CheckAndDoVoipCallForHotelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckAndDoVoipCallForHotelResponseBody) GetResult() *CheckAndDoVoipCallForHotelResponseBodyResult {
	return s.Result
}

func (s *CheckAndDoVoipCallForHotelResponseBody) SetCode(v int32) *CheckAndDoVoipCallForHotelResponseBody {
	s.Code = &v
	return s
}

func (s *CheckAndDoVoipCallForHotelResponseBody) SetMessage(v string) *CheckAndDoVoipCallForHotelResponseBody {
	s.Message = &v
	return s
}

func (s *CheckAndDoVoipCallForHotelResponseBody) SetRequestId(v string) *CheckAndDoVoipCallForHotelResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckAndDoVoipCallForHotelResponseBody) SetResult(v *CheckAndDoVoipCallForHotelResponseBodyResult) *CheckAndDoVoipCallForHotelResponseBody {
	s.Result = v
	return s
}

func (s *CheckAndDoVoipCallForHotelResponseBody) Validate() error {
	return dara.Validate(s)
}

type CheckAndDoVoipCallForHotelResponseBodyResult struct {
	DeviceTargets   *CheckAndDoVoipCallForHotelResponseBodyResultDeviceTargets   `json:"DeviceTargets,omitempty" xml:"DeviceTargets,omitempty" type:"Struct"`
	IsStartCall     *bool                                                        `json:"IsStartCall,omitempty" xml:"IsStartCall,omitempty"`
	Passed          *bool                                                        `json:"Passed,omitempty" xml:"Passed,omitempty"`
	StartCallResult *CheckAndDoVoipCallForHotelResponseBodyResultStartCallResult `json:"StartCallResult,omitempty" xml:"StartCallResult,omitempty" type:"Struct"`
}

func (s CheckAndDoVoipCallForHotelResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CheckAndDoVoipCallForHotelResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CheckAndDoVoipCallForHotelResponseBodyResult) GetDeviceTargets() *CheckAndDoVoipCallForHotelResponseBodyResultDeviceTargets {
	return s.DeviceTargets
}

func (s *CheckAndDoVoipCallForHotelResponseBodyResult) GetIsStartCall() *bool {
	return s.IsStartCall
}

func (s *CheckAndDoVoipCallForHotelResponseBodyResult) GetPassed() *bool {
	return s.Passed
}

func (s *CheckAndDoVoipCallForHotelResponseBodyResult) GetStartCallResult() *CheckAndDoVoipCallForHotelResponseBodyResultStartCallResult {
	return s.StartCallResult
}

func (s *CheckAndDoVoipCallForHotelResponseBodyResult) SetDeviceTargets(v *CheckAndDoVoipCallForHotelResponseBodyResultDeviceTargets) *CheckAndDoVoipCallForHotelResponseBodyResult {
	s.DeviceTargets = v
	return s
}

func (s *CheckAndDoVoipCallForHotelResponseBodyResult) SetIsStartCall(v bool) *CheckAndDoVoipCallForHotelResponseBodyResult {
	s.IsStartCall = &v
	return s
}

func (s *CheckAndDoVoipCallForHotelResponseBodyResult) SetPassed(v bool) *CheckAndDoVoipCallForHotelResponseBodyResult {
	s.Passed = &v
	return s
}

func (s *CheckAndDoVoipCallForHotelResponseBodyResult) SetStartCallResult(v *CheckAndDoVoipCallForHotelResponseBodyResultStartCallResult) *CheckAndDoVoipCallForHotelResponseBodyResult {
	s.StartCallResult = v
	return s
}

func (s *CheckAndDoVoipCallForHotelResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type CheckAndDoVoipCallForHotelResponseBodyResultDeviceTargets struct {
	Code *int32                                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*CheckAndDoVoipCallForHotelResponseBodyResultDeviceTargetsData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Msg  *string                                                          `json:"Msg,omitempty" xml:"Msg,omitempty"`
}

func (s CheckAndDoVoipCallForHotelResponseBodyResultDeviceTargets) String() string {
	return dara.Prettify(s)
}

func (s CheckAndDoVoipCallForHotelResponseBodyResultDeviceTargets) GoString() string {
	return s.String()
}

func (s *CheckAndDoVoipCallForHotelResponseBodyResultDeviceTargets) GetCode() *int32 {
	return s.Code
}

func (s *CheckAndDoVoipCallForHotelResponseBodyResultDeviceTargets) GetData() []*CheckAndDoVoipCallForHotelResponseBodyResultDeviceTargetsData {
	return s.Data
}

func (s *CheckAndDoVoipCallForHotelResponseBodyResultDeviceTargets) GetMsg() *string {
	return s.Msg
}

func (s *CheckAndDoVoipCallForHotelResponseBodyResultDeviceTargets) SetCode(v int32) *CheckAndDoVoipCallForHotelResponseBodyResultDeviceTargets {
	s.Code = &v
	return s
}

func (s *CheckAndDoVoipCallForHotelResponseBodyResultDeviceTargets) SetData(v []*CheckAndDoVoipCallForHotelResponseBodyResultDeviceTargetsData) *CheckAndDoVoipCallForHotelResponseBodyResultDeviceTargets {
	s.Data = v
	return s
}

func (s *CheckAndDoVoipCallForHotelResponseBodyResultDeviceTargets) SetMsg(v string) *CheckAndDoVoipCallForHotelResponseBodyResultDeviceTargets {
	s.Msg = &v
	return s
}

func (s *CheckAndDoVoipCallForHotelResponseBodyResultDeviceTargets) Validate() error {
	return dara.Validate(s)
}

type CheckAndDoVoipCallForHotelResponseBodyResultDeviceTargetsData struct {
	DeviceIcon *string `json:"DeviceIcon,omitempty" xml:"DeviceIcon,omitempty"`
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	Online     *bool   `json:"Online,omitempty" xml:"Online,omitempty"`
	Uuid       *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s CheckAndDoVoipCallForHotelResponseBodyResultDeviceTargetsData) String() string {
	return dara.Prettify(s)
}

func (s CheckAndDoVoipCallForHotelResponseBodyResultDeviceTargetsData) GoString() string {
	return s.String()
}

func (s *CheckAndDoVoipCallForHotelResponseBodyResultDeviceTargetsData) GetDeviceIcon() *string {
	return s.DeviceIcon
}

func (s *CheckAndDoVoipCallForHotelResponseBodyResultDeviceTargetsData) GetDeviceName() *string {
	return s.DeviceName
}

func (s *CheckAndDoVoipCallForHotelResponseBodyResultDeviceTargetsData) GetDeviceType() *string {
	return s.DeviceType
}

func (s *CheckAndDoVoipCallForHotelResponseBodyResultDeviceTargetsData) GetOnline() *bool {
	return s.Online
}

func (s *CheckAndDoVoipCallForHotelResponseBodyResultDeviceTargetsData) GetUuid() *string {
	return s.Uuid
}

func (s *CheckAndDoVoipCallForHotelResponseBodyResultDeviceTargetsData) SetDeviceIcon(v string) *CheckAndDoVoipCallForHotelResponseBodyResultDeviceTargetsData {
	s.DeviceIcon = &v
	return s
}

func (s *CheckAndDoVoipCallForHotelResponseBodyResultDeviceTargetsData) SetDeviceName(v string) *CheckAndDoVoipCallForHotelResponseBodyResultDeviceTargetsData {
	s.DeviceName = &v
	return s
}

func (s *CheckAndDoVoipCallForHotelResponseBodyResultDeviceTargetsData) SetDeviceType(v string) *CheckAndDoVoipCallForHotelResponseBodyResultDeviceTargetsData {
	s.DeviceType = &v
	return s
}

func (s *CheckAndDoVoipCallForHotelResponseBodyResultDeviceTargetsData) SetOnline(v bool) *CheckAndDoVoipCallForHotelResponseBodyResultDeviceTargetsData {
	s.Online = &v
	return s
}

func (s *CheckAndDoVoipCallForHotelResponseBodyResultDeviceTargetsData) SetUuid(v string) *CheckAndDoVoipCallForHotelResponseBodyResultDeviceTargetsData {
	s.Uuid = &v
	return s
}

func (s *CheckAndDoVoipCallForHotelResponseBodyResultDeviceTargetsData) Validate() error {
	return dara.Validate(s)
}

type CheckAndDoVoipCallForHotelResponseBodyResultStartCallResult struct {
	Message  *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RetCode  *int32  `json:"RetCode,omitempty" xml:"RetCode,omitempty"`
	RetValue *string `json:"RetValue,omitempty" xml:"RetValue,omitempty"`
	Success  *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId  *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s CheckAndDoVoipCallForHotelResponseBodyResultStartCallResult) String() string {
	return dara.Prettify(s)
}

func (s CheckAndDoVoipCallForHotelResponseBodyResultStartCallResult) GoString() string {
	return s.String()
}

func (s *CheckAndDoVoipCallForHotelResponseBodyResultStartCallResult) GetMessage() *string {
	return s.Message
}

func (s *CheckAndDoVoipCallForHotelResponseBodyResultStartCallResult) GetRetCode() *int32 {
	return s.RetCode
}

func (s *CheckAndDoVoipCallForHotelResponseBodyResultStartCallResult) GetRetValue() *string {
	return s.RetValue
}

func (s *CheckAndDoVoipCallForHotelResponseBodyResultStartCallResult) GetSuccess() *bool {
	return s.Success
}

func (s *CheckAndDoVoipCallForHotelResponseBodyResultStartCallResult) GetTraceId() *string {
	return s.TraceId
}

func (s *CheckAndDoVoipCallForHotelResponseBodyResultStartCallResult) SetMessage(v string) *CheckAndDoVoipCallForHotelResponseBodyResultStartCallResult {
	s.Message = &v
	return s
}

func (s *CheckAndDoVoipCallForHotelResponseBodyResultStartCallResult) SetRetCode(v int32) *CheckAndDoVoipCallForHotelResponseBodyResultStartCallResult {
	s.RetCode = &v
	return s
}

func (s *CheckAndDoVoipCallForHotelResponseBodyResultStartCallResult) SetRetValue(v string) *CheckAndDoVoipCallForHotelResponseBodyResultStartCallResult {
	s.RetValue = &v
	return s
}

func (s *CheckAndDoVoipCallForHotelResponseBodyResultStartCallResult) SetSuccess(v bool) *CheckAndDoVoipCallForHotelResponseBodyResultStartCallResult {
	s.Success = &v
	return s
}

func (s *CheckAndDoVoipCallForHotelResponseBodyResultStartCallResult) SetTraceId(v string) *CheckAndDoVoipCallForHotelResponseBodyResultStartCallResult {
	s.TraceId = &v
	return s
}

func (s *CheckAndDoVoipCallForHotelResponseBodyResultStartCallResult) Validate() error {
	return dara.Validate(s)
}

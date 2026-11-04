// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryVisionDeviceInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryVisionDeviceInfoResponseBody
	GetCode() *string
	SetData(v *QueryVisionDeviceInfoResponseBodyData) *QueryVisionDeviceInfoResponseBody
	GetData() *QueryVisionDeviceInfoResponseBodyData
	SetErrorMessage(v string) *QueryVisionDeviceInfoResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *QueryVisionDeviceInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryVisionDeviceInfoResponseBody
	GetSuccess() *bool
}

type QueryVisionDeviceInfoResponseBody struct {
	Code         *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QueryVisionDeviceInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryVisionDeviceInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryVisionDeviceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryVisionDeviceInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryVisionDeviceInfoResponseBody) GetData() *QueryVisionDeviceInfoResponseBodyData {
	return s.Data
}

func (s *QueryVisionDeviceInfoResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *QueryVisionDeviceInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryVisionDeviceInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryVisionDeviceInfoResponseBody) SetCode(v string) *QueryVisionDeviceInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryVisionDeviceInfoResponseBody) SetData(v *QueryVisionDeviceInfoResponseBodyData) *QueryVisionDeviceInfoResponseBody {
	s.Data = v
	return s
}

func (s *QueryVisionDeviceInfoResponseBody) SetErrorMessage(v string) *QueryVisionDeviceInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryVisionDeviceInfoResponseBody) SetRequestId(v string) *QueryVisionDeviceInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryVisionDeviceInfoResponseBody) SetSuccess(v bool) *QueryVisionDeviceInfoResponseBody {
	s.Success = &v
	return s
}

func (s *QueryVisionDeviceInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryVisionDeviceInfoResponseBodyData struct {
	Description    *string                                              `json:"Description,omitempty" xml:"Description,omitempty"`
	DeviceType     *int32                                               `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	GbDeviceInfo   *QueryVisionDeviceInfoResponseBodyDataGbDeviceInfo   `json:"GbDeviceInfo,omitempty" xml:"GbDeviceInfo,omitempty" type:"Struct"`
	RtmpDeviceInfo *QueryVisionDeviceInfoResponseBodyDataRtmpDeviceInfo `json:"RtmpDeviceInfo,omitempty" xml:"RtmpDeviceInfo,omitempty" type:"Struct"`
}

func (s QueryVisionDeviceInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryVisionDeviceInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryVisionDeviceInfoResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *QueryVisionDeviceInfoResponseBodyData) GetDeviceType() *int32 {
	return s.DeviceType
}

func (s *QueryVisionDeviceInfoResponseBodyData) GetGbDeviceInfo() *QueryVisionDeviceInfoResponseBodyDataGbDeviceInfo {
	return s.GbDeviceInfo
}

func (s *QueryVisionDeviceInfoResponseBodyData) GetRtmpDeviceInfo() *QueryVisionDeviceInfoResponseBodyDataRtmpDeviceInfo {
	return s.RtmpDeviceInfo
}

func (s *QueryVisionDeviceInfoResponseBodyData) SetDescription(v string) *QueryVisionDeviceInfoResponseBodyData {
	s.Description = &v
	return s
}

func (s *QueryVisionDeviceInfoResponseBodyData) SetDeviceType(v int32) *QueryVisionDeviceInfoResponseBodyData {
	s.DeviceType = &v
	return s
}

func (s *QueryVisionDeviceInfoResponseBodyData) SetGbDeviceInfo(v *QueryVisionDeviceInfoResponseBodyDataGbDeviceInfo) *QueryVisionDeviceInfoResponseBodyData {
	s.GbDeviceInfo = v
	return s
}

func (s *QueryVisionDeviceInfoResponseBodyData) SetRtmpDeviceInfo(v *QueryVisionDeviceInfoResponseBodyDataRtmpDeviceInfo) *QueryVisionDeviceInfoResponseBodyData {
	s.RtmpDeviceInfo = v
	return s
}

func (s *QueryVisionDeviceInfoResponseBodyData) Validate() error {
	if s.GbDeviceInfo != nil {
		if err := s.GbDeviceInfo.Validate(); err != nil {
			return err
		}
	}
	if s.RtmpDeviceInfo != nil {
		if err := s.RtmpDeviceInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryVisionDeviceInfoResponseBodyDataGbDeviceInfo struct {
	DeviceProtocol *int32  `json:"DeviceProtocol,omitempty" xml:"DeviceProtocol,omitempty"`
	GbId           *string `json:"GbId,omitempty" xml:"GbId,omitempty"`
	NetProtocol    *int32  `json:"NetProtocol,omitempty" xml:"NetProtocol,omitempty"`
	NickName       *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	Password       *string `json:"Password,omitempty" xml:"Password,omitempty"`
	SubProductKey  *string `json:"SubProductKey,omitempty" xml:"SubProductKey,omitempty"`
}

func (s QueryVisionDeviceInfoResponseBodyDataGbDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryVisionDeviceInfoResponseBodyDataGbDeviceInfo) GoString() string {
	return s.String()
}

func (s *QueryVisionDeviceInfoResponseBodyDataGbDeviceInfo) GetDeviceProtocol() *int32 {
	return s.DeviceProtocol
}

func (s *QueryVisionDeviceInfoResponseBodyDataGbDeviceInfo) GetGbId() *string {
	return s.GbId
}

func (s *QueryVisionDeviceInfoResponseBodyDataGbDeviceInfo) GetNetProtocol() *int32 {
	return s.NetProtocol
}

func (s *QueryVisionDeviceInfoResponseBodyDataGbDeviceInfo) GetNickName() *string {
	return s.NickName
}

func (s *QueryVisionDeviceInfoResponseBodyDataGbDeviceInfo) GetPassword() *string {
	return s.Password
}

func (s *QueryVisionDeviceInfoResponseBodyDataGbDeviceInfo) GetSubProductKey() *string {
	return s.SubProductKey
}

func (s *QueryVisionDeviceInfoResponseBodyDataGbDeviceInfo) SetDeviceProtocol(v int32) *QueryVisionDeviceInfoResponseBodyDataGbDeviceInfo {
	s.DeviceProtocol = &v
	return s
}

func (s *QueryVisionDeviceInfoResponseBodyDataGbDeviceInfo) SetGbId(v string) *QueryVisionDeviceInfoResponseBodyDataGbDeviceInfo {
	s.GbId = &v
	return s
}

func (s *QueryVisionDeviceInfoResponseBodyDataGbDeviceInfo) SetNetProtocol(v int32) *QueryVisionDeviceInfoResponseBodyDataGbDeviceInfo {
	s.NetProtocol = &v
	return s
}

func (s *QueryVisionDeviceInfoResponseBodyDataGbDeviceInfo) SetNickName(v string) *QueryVisionDeviceInfoResponseBodyDataGbDeviceInfo {
	s.NickName = &v
	return s
}

func (s *QueryVisionDeviceInfoResponseBodyDataGbDeviceInfo) SetPassword(v string) *QueryVisionDeviceInfoResponseBodyDataGbDeviceInfo {
	s.Password = &v
	return s
}

func (s *QueryVisionDeviceInfoResponseBodyDataGbDeviceInfo) SetSubProductKey(v string) *QueryVisionDeviceInfoResponseBodyDataGbDeviceInfo {
	s.SubProductKey = &v
	return s
}

func (s *QueryVisionDeviceInfoResponseBodyDataGbDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type QueryVisionDeviceInfoResponseBodyDataRtmpDeviceInfo struct {
	PullAuthKey       *string `json:"PullAuthKey,omitempty" xml:"PullAuthKey,omitempty"`
	PullKeyExpireTime *int32  `json:"PullKeyExpireTime,omitempty" xml:"PullKeyExpireTime,omitempty"`
	PushAuthKey       *string `json:"PushAuthKey,omitempty" xml:"PushAuthKey,omitempty"`
	PushKeyExpireTime *int32  `json:"PushKeyExpireTime,omitempty" xml:"PushKeyExpireTime,omitempty"`
	PushUrlSample     *string `json:"PushUrlSample,omitempty" xml:"PushUrlSample,omitempty"`
	StreamName        *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	StreamStatus      *int32  `json:"StreamStatus,omitempty" xml:"StreamStatus,omitempty"`
}

func (s QueryVisionDeviceInfoResponseBodyDataRtmpDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryVisionDeviceInfoResponseBodyDataRtmpDeviceInfo) GoString() string {
	return s.String()
}

func (s *QueryVisionDeviceInfoResponseBodyDataRtmpDeviceInfo) GetPullAuthKey() *string {
	return s.PullAuthKey
}

func (s *QueryVisionDeviceInfoResponseBodyDataRtmpDeviceInfo) GetPullKeyExpireTime() *int32 {
	return s.PullKeyExpireTime
}

func (s *QueryVisionDeviceInfoResponseBodyDataRtmpDeviceInfo) GetPushAuthKey() *string {
	return s.PushAuthKey
}

func (s *QueryVisionDeviceInfoResponseBodyDataRtmpDeviceInfo) GetPushKeyExpireTime() *int32 {
	return s.PushKeyExpireTime
}

func (s *QueryVisionDeviceInfoResponseBodyDataRtmpDeviceInfo) GetPushUrlSample() *string {
	return s.PushUrlSample
}

func (s *QueryVisionDeviceInfoResponseBodyDataRtmpDeviceInfo) GetStreamName() *string {
	return s.StreamName
}

func (s *QueryVisionDeviceInfoResponseBodyDataRtmpDeviceInfo) GetStreamStatus() *int32 {
	return s.StreamStatus
}

func (s *QueryVisionDeviceInfoResponseBodyDataRtmpDeviceInfo) SetPullAuthKey(v string) *QueryVisionDeviceInfoResponseBodyDataRtmpDeviceInfo {
	s.PullAuthKey = &v
	return s
}

func (s *QueryVisionDeviceInfoResponseBodyDataRtmpDeviceInfo) SetPullKeyExpireTime(v int32) *QueryVisionDeviceInfoResponseBodyDataRtmpDeviceInfo {
	s.PullKeyExpireTime = &v
	return s
}

func (s *QueryVisionDeviceInfoResponseBodyDataRtmpDeviceInfo) SetPushAuthKey(v string) *QueryVisionDeviceInfoResponseBodyDataRtmpDeviceInfo {
	s.PushAuthKey = &v
	return s
}

func (s *QueryVisionDeviceInfoResponseBodyDataRtmpDeviceInfo) SetPushKeyExpireTime(v int32) *QueryVisionDeviceInfoResponseBodyDataRtmpDeviceInfo {
	s.PushKeyExpireTime = &v
	return s
}

func (s *QueryVisionDeviceInfoResponseBodyDataRtmpDeviceInfo) SetPushUrlSample(v string) *QueryVisionDeviceInfoResponseBodyDataRtmpDeviceInfo {
	s.PushUrlSample = &v
	return s
}

func (s *QueryVisionDeviceInfoResponseBodyDataRtmpDeviceInfo) SetStreamName(v string) *QueryVisionDeviceInfoResponseBodyDataRtmpDeviceInfo {
	s.StreamName = &v
	return s
}

func (s *QueryVisionDeviceInfoResponseBodyDataRtmpDeviceInfo) SetStreamStatus(v int32) *QueryVisionDeviceInfoResponseBodyDataRtmpDeviceInfo {
	s.StreamStatus = &v
	return s
}

func (s *QueryVisionDeviceInfoResponseBodyDataRtmpDeviceInfo) Validate() error {
	return dara.Validate(s)
}

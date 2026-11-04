// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchQueryVisionDeviceInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BatchQueryVisionDeviceInfoResponseBody
	GetCode() *string
	SetData(v *BatchQueryVisionDeviceInfoResponseBodyData) *BatchQueryVisionDeviceInfoResponseBody
	GetData() *BatchQueryVisionDeviceInfoResponseBodyData
	SetErrorMessage(v string) *BatchQueryVisionDeviceInfoResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *BatchQueryVisionDeviceInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BatchQueryVisionDeviceInfoResponseBody
	GetSuccess() *bool
}

type BatchQueryVisionDeviceInfoResponseBody struct {
	Code         *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *BatchQueryVisionDeviceInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                     `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchQueryVisionDeviceInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchQueryVisionDeviceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *BatchQueryVisionDeviceInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *BatchQueryVisionDeviceInfoResponseBody) GetData() *BatchQueryVisionDeviceInfoResponseBodyData {
	return s.Data
}

func (s *BatchQueryVisionDeviceInfoResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *BatchQueryVisionDeviceInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchQueryVisionDeviceInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BatchQueryVisionDeviceInfoResponseBody) SetCode(v string) *BatchQueryVisionDeviceInfoResponseBody {
	s.Code = &v
	return s
}

func (s *BatchQueryVisionDeviceInfoResponseBody) SetData(v *BatchQueryVisionDeviceInfoResponseBodyData) *BatchQueryVisionDeviceInfoResponseBody {
	s.Data = v
	return s
}

func (s *BatchQueryVisionDeviceInfoResponseBody) SetErrorMessage(v string) *BatchQueryVisionDeviceInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *BatchQueryVisionDeviceInfoResponseBody) SetRequestId(v string) *BatchQueryVisionDeviceInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchQueryVisionDeviceInfoResponseBody) SetSuccess(v bool) *BatchQueryVisionDeviceInfoResponseBody {
	s.Success = &v
	return s
}

func (s *BatchQueryVisionDeviceInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchQueryVisionDeviceInfoResponseBodyData struct {
	DeviceInfoList []*BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoList `json:"DeviceInfoList,omitempty" xml:"DeviceInfoList,omitempty" type:"Repeated"`
}

func (s BatchQueryVisionDeviceInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BatchQueryVisionDeviceInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *BatchQueryVisionDeviceInfoResponseBodyData) GetDeviceInfoList() []*BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoList {
	return s.DeviceInfoList
}

func (s *BatchQueryVisionDeviceInfoResponseBodyData) SetDeviceInfoList(v []*BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoList) *BatchQueryVisionDeviceInfoResponseBodyData {
	s.DeviceInfoList = v
	return s
}

func (s *BatchQueryVisionDeviceInfoResponseBodyData) Validate() error {
	if s.DeviceInfoList != nil {
		for _, item := range s.DeviceInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoList struct {
	Description    *string                                                                 `json:"Description,omitempty" xml:"Description,omitempty"`
	DeviceType     *int32                                                                  `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	GbDeviceInfo   *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListGbDeviceInfo   `json:"GbDeviceInfo,omitempty" xml:"GbDeviceInfo,omitempty" type:"Struct"`
	IotId          *string                                                                 `json:"IotId,omitempty" xml:"IotId,omitempty"`
	RtmpDeviceInfo *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListRtmpDeviceInfo `json:"RtmpDeviceInfo,omitempty" xml:"RtmpDeviceInfo,omitempty" type:"Struct"`
}

func (s BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoList) String() string {
	return dara.Prettify(s)
}

func (s BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoList) GoString() string {
	return s.String()
}

func (s *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoList) GetDescription() *string {
	return s.Description
}

func (s *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoList) GetDeviceType() *int32 {
	return s.DeviceType
}

func (s *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoList) GetGbDeviceInfo() *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListGbDeviceInfo {
	return s.GbDeviceInfo
}

func (s *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoList) GetIotId() *string {
	return s.IotId
}

func (s *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoList) GetRtmpDeviceInfo() *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListRtmpDeviceInfo {
	return s.RtmpDeviceInfo
}

func (s *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoList) SetDescription(v string) *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoList {
	s.Description = &v
	return s
}

func (s *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoList) SetDeviceType(v int32) *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoList {
	s.DeviceType = &v
	return s
}

func (s *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoList) SetGbDeviceInfo(v *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListGbDeviceInfo) *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoList {
	s.GbDeviceInfo = v
	return s
}

func (s *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoList) SetIotId(v string) *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoList {
	s.IotId = &v
	return s
}

func (s *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoList) SetRtmpDeviceInfo(v *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListRtmpDeviceInfo) *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoList {
	s.RtmpDeviceInfo = v
	return s
}

func (s *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoList) Validate() error {
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

type BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListGbDeviceInfo struct {
	DeviceProtocol *int32  `json:"DeviceProtocol,omitempty" xml:"DeviceProtocol,omitempty"`
	GbId           *string `json:"GbId,omitempty" xml:"GbId,omitempty"`
	NetProtocol    *int32  `json:"NetProtocol,omitempty" xml:"NetProtocol,omitempty"`
	NickName       *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	Password       *string `json:"Password,omitempty" xml:"Password,omitempty"`
	SubProductKey  *string `json:"SubProductKey,omitempty" xml:"SubProductKey,omitempty"`
}

func (s BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListGbDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListGbDeviceInfo) GoString() string {
	return s.String()
}

func (s *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListGbDeviceInfo) GetDeviceProtocol() *int32 {
	return s.DeviceProtocol
}

func (s *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListGbDeviceInfo) GetGbId() *string {
	return s.GbId
}

func (s *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListGbDeviceInfo) GetNetProtocol() *int32 {
	return s.NetProtocol
}

func (s *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListGbDeviceInfo) GetNickName() *string {
	return s.NickName
}

func (s *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListGbDeviceInfo) GetPassword() *string {
	return s.Password
}

func (s *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListGbDeviceInfo) GetSubProductKey() *string {
	return s.SubProductKey
}

func (s *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListGbDeviceInfo) SetDeviceProtocol(v int32) *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListGbDeviceInfo {
	s.DeviceProtocol = &v
	return s
}

func (s *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListGbDeviceInfo) SetGbId(v string) *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListGbDeviceInfo {
	s.GbId = &v
	return s
}

func (s *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListGbDeviceInfo) SetNetProtocol(v int32) *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListGbDeviceInfo {
	s.NetProtocol = &v
	return s
}

func (s *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListGbDeviceInfo) SetNickName(v string) *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListGbDeviceInfo {
	s.NickName = &v
	return s
}

func (s *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListGbDeviceInfo) SetPassword(v string) *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListGbDeviceInfo {
	s.Password = &v
	return s
}

func (s *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListGbDeviceInfo) SetSubProductKey(v string) *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListGbDeviceInfo {
	s.SubProductKey = &v
	return s
}

func (s *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListGbDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListRtmpDeviceInfo struct {
	PullAuthKey       *string `json:"PullAuthKey,omitempty" xml:"PullAuthKey,omitempty"`
	PullKeyExpireTime *int32  `json:"PullKeyExpireTime,omitempty" xml:"PullKeyExpireTime,omitempty"`
	PushAuthKey       *string `json:"PushAuthKey,omitempty" xml:"PushAuthKey,omitempty"`
	PushKeyExpireTime *int32  `json:"PushKeyExpireTime,omitempty" xml:"PushKeyExpireTime,omitempty"`
	PushUrlSample     *string `json:"PushUrlSample,omitempty" xml:"PushUrlSample,omitempty"`
	StreamName        *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	StreamStatus      *int32  `json:"StreamStatus,omitempty" xml:"StreamStatus,omitempty"`
}

func (s BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListRtmpDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListRtmpDeviceInfo) GoString() string {
	return s.String()
}

func (s *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListRtmpDeviceInfo) GetPullAuthKey() *string {
	return s.PullAuthKey
}

func (s *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListRtmpDeviceInfo) GetPullKeyExpireTime() *int32 {
	return s.PullKeyExpireTime
}

func (s *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListRtmpDeviceInfo) GetPushAuthKey() *string {
	return s.PushAuthKey
}

func (s *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListRtmpDeviceInfo) GetPushKeyExpireTime() *int32 {
	return s.PushKeyExpireTime
}

func (s *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListRtmpDeviceInfo) GetPushUrlSample() *string {
	return s.PushUrlSample
}

func (s *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListRtmpDeviceInfo) GetStreamName() *string {
	return s.StreamName
}

func (s *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListRtmpDeviceInfo) GetStreamStatus() *int32 {
	return s.StreamStatus
}

func (s *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListRtmpDeviceInfo) SetPullAuthKey(v string) *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListRtmpDeviceInfo {
	s.PullAuthKey = &v
	return s
}

func (s *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListRtmpDeviceInfo) SetPullKeyExpireTime(v int32) *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListRtmpDeviceInfo {
	s.PullKeyExpireTime = &v
	return s
}

func (s *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListRtmpDeviceInfo) SetPushAuthKey(v string) *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListRtmpDeviceInfo {
	s.PushAuthKey = &v
	return s
}

func (s *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListRtmpDeviceInfo) SetPushKeyExpireTime(v int32) *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListRtmpDeviceInfo {
	s.PushKeyExpireTime = &v
	return s
}

func (s *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListRtmpDeviceInfo) SetPushUrlSample(v string) *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListRtmpDeviceInfo {
	s.PushUrlSample = &v
	return s
}

func (s *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListRtmpDeviceInfo) SetStreamName(v string) *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListRtmpDeviceInfo {
	s.StreamName = &v
	return s
}

func (s *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListRtmpDeviceInfo) SetStreamStatus(v int32) *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListRtmpDeviceInfo {
	s.StreamStatus = &v
	return s
}

func (s *BatchQueryVisionDeviceInfoResponseBodyDataDeviceInfoListRtmpDeviceInfo) Validate() error {
	return dara.Validate(s)
}

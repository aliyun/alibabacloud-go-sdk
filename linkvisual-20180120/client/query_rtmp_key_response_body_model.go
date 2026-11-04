// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRtmpKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryRtmpKeyResponseBody
	GetCode() *string
	SetData(v *QueryRtmpKeyResponseBodyData) *QueryRtmpKeyResponseBody
	GetData() *QueryRtmpKeyResponseBodyData
	SetErrorMessage(v string) *QueryRtmpKeyResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *QueryRtmpKeyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryRtmpKeyResponseBody
	GetSuccess() *bool
}

type QueryRtmpKeyResponseBody struct {
	Code         *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QueryRtmpKeyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                       `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryRtmpKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryRtmpKeyResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRtmpKeyResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryRtmpKeyResponseBody) GetData() *QueryRtmpKeyResponseBodyData {
	return s.Data
}

func (s *QueryRtmpKeyResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *QueryRtmpKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryRtmpKeyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryRtmpKeyResponseBody) SetCode(v string) *QueryRtmpKeyResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRtmpKeyResponseBody) SetData(v *QueryRtmpKeyResponseBodyData) *QueryRtmpKeyResponseBody {
	s.Data = v
	return s
}

func (s *QueryRtmpKeyResponseBody) SetErrorMessage(v string) *QueryRtmpKeyResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryRtmpKeyResponseBody) SetRequestId(v string) *QueryRtmpKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRtmpKeyResponseBody) SetSuccess(v bool) *QueryRtmpKeyResponseBody {
	s.Success = &v
	return s
}

func (s *QueryRtmpKeyResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryRtmpKeyResponseBodyData struct {
	PullAuthKey       *string `json:"PullAuthKey,omitempty" xml:"PullAuthKey,omitempty"`
	PullKeyExpireTime *int32  `json:"PullKeyExpireTime,omitempty" xml:"PullKeyExpireTime,omitempty"`
	PushAuthKey       *string `json:"PushAuthKey,omitempty" xml:"PushAuthKey,omitempty"`
	PushKeyExpireTime *int32  `json:"PushKeyExpireTime,omitempty" xml:"PushKeyExpireTime,omitempty"`
	StreamName        *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s QueryRtmpKeyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryRtmpKeyResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryRtmpKeyResponseBodyData) GetPullAuthKey() *string {
	return s.PullAuthKey
}

func (s *QueryRtmpKeyResponseBodyData) GetPullKeyExpireTime() *int32 {
	return s.PullKeyExpireTime
}

func (s *QueryRtmpKeyResponseBodyData) GetPushAuthKey() *string {
	return s.PushAuthKey
}

func (s *QueryRtmpKeyResponseBodyData) GetPushKeyExpireTime() *int32 {
	return s.PushKeyExpireTime
}

func (s *QueryRtmpKeyResponseBodyData) GetStreamName() *string {
	return s.StreamName
}

func (s *QueryRtmpKeyResponseBodyData) SetPullAuthKey(v string) *QueryRtmpKeyResponseBodyData {
	s.PullAuthKey = &v
	return s
}

func (s *QueryRtmpKeyResponseBodyData) SetPullKeyExpireTime(v int32) *QueryRtmpKeyResponseBodyData {
	s.PullKeyExpireTime = &v
	return s
}

func (s *QueryRtmpKeyResponseBodyData) SetPushAuthKey(v string) *QueryRtmpKeyResponseBodyData {
	s.PushAuthKey = &v
	return s
}

func (s *QueryRtmpKeyResponseBodyData) SetPushKeyExpireTime(v int32) *QueryRtmpKeyResponseBodyData {
	s.PushKeyExpireTime = &v
	return s
}

func (s *QueryRtmpKeyResponseBodyData) SetStreamName(v string) *QueryRtmpKeyResponseBodyData {
	s.StreamName = &v
	return s
}

func (s *QueryRtmpKeyResponseBodyData) Validate() error {
	return dara.Validate(s)
}

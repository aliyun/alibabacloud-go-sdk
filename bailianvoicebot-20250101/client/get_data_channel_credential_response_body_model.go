// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataChannelCredentialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDataChannelCredentialResponseBody
	GetCode() *string
	SetData(v *GetDataChannelCredentialResponseBodyData) *GetDataChannelCredentialResponseBody
	GetData() *GetDataChannelCredentialResponseBodyData
	SetHttpStatusCode(v int32) *GetDataChannelCredentialResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetDataChannelCredentialResponseBody
	GetMessage() *string
	SetParams(v []*string) *GetDataChannelCredentialResponseBody
	GetParams() []*string
	SetRequestId(v string) *GetDataChannelCredentialResponseBody
	GetRequestId() *string
}

type GetDataChannelCredentialResponseBody struct {
	// example:
	//
	// OK
	Code *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetDataChannelCredentialResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params  []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// D771A1B6-3D5F-174A-BEE1-98CE1000D337
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDataChannelCredentialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataChannelCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataChannelCredentialResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDataChannelCredentialResponseBody) GetData() *GetDataChannelCredentialResponseBodyData {
	return s.Data
}

func (s *GetDataChannelCredentialResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetDataChannelCredentialResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDataChannelCredentialResponseBody) GetParams() []*string {
	return s.Params
}

func (s *GetDataChannelCredentialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataChannelCredentialResponseBody) SetCode(v string) *GetDataChannelCredentialResponseBody {
	s.Code = &v
	return s
}

func (s *GetDataChannelCredentialResponseBody) SetData(v *GetDataChannelCredentialResponseBodyData) *GetDataChannelCredentialResponseBody {
	s.Data = v
	return s
}

func (s *GetDataChannelCredentialResponseBody) SetHttpStatusCode(v int32) *GetDataChannelCredentialResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDataChannelCredentialResponseBody) SetMessage(v string) *GetDataChannelCredentialResponseBody {
	s.Message = &v
	return s
}

func (s *GetDataChannelCredentialResponseBody) SetParams(v []*string) *GetDataChannelCredentialResponseBody {
	s.Params = v
	return s
}

func (s *GetDataChannelCredentialResponseBody) SetRequestId(v string) *GetDataChannelCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataChannelCredentialResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDataChannelCredentialResponseBodyData struct {
	// example:
	//
	// 26c2f022-b6c0-4ab0-9019-6e1a42dc5582
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// example:
	//
	// device-3i5x4234f2j4w55e
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// example:
	//
	// mqtt-cn-ux146hgtt04.mqtt.aliyuncs.com
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// example:
	//
	// 1745004535507
	ExpirationTime *int64 `json:"ExpirationTime,omitempty" xml:"ExpirationTime,omitempty"`
	// example:
	//
	// ***
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// Topic
	//
	// example:
	//
	// datachannel-prepub-a/dc52807f0eff4b9b8224d06c7f240c07
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// example:
	//
	// Token|LTAI5tRYzHUYYi4XstgMCsL4|mqtt-cn-ux146hgtt04
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s GetDataChannelCredentialResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDataChannelCredentialResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDataChannelCredentialResponseBodyData) GetClientId() *string {
	return s.ClientId
}

func (s *GetDataChannelCredentialResponseBodyData) GetDeviceId() *string {
	return s.DeviceId
}

func (s *GetDataChannelCredentialResponseBodyData) GetEndpoint() *string {
	return s.Endpoint
}

func (s *GetDataChannelCredentialResponseBodyData) GetExpirationTime() *int64 {
	return s.ExpirationTime
}

func (s *GetDataChannelCredentialResponseBodyData) GetPassword() *string {
	return s.Password
}

func (s *GetDataChannelCredentialResponseBodyData) GetTopic() *string {
	return s.Topic
}

func (s *GetDataChannelCredentialResponseBodyData) GetUserName() *string {
	return s.UserName
}

func (s *GetDataChannelCredentialResponseBodyData) SetClientId(v string) *GetDataChannelCredentialResponseBodyData {
	s.ClientId = &v
	return s
}

func (s *GetDataChannelCredentialResponseBodyData) SetDeviceId(v string) *GetDataChannelCredentialResponseBodyData {
	s.DeviceId = &v
	return s
}

func (s *GetDataChannelCredentialResponseBodyData) SetEndpoint(v string) *GetDataChannelCredentialResponseBodyData {
	s.Endpoint = &v
	return s
}

func (s *GetDataChannelCredentialResponseBodyData) SetExpirationTime(v int64) *GetDataChannelCredentialResponseBodyData {
	s.ExpirationTime = &v
	return s
}

func (s *GetDataChannelCredentialResponseBodyData) SetPassword(v string) *GetDataChannelCredentialResponseBodyData {
	s.Password = &v
	return s
}

func (s *GetDataChannelCredentialResponseBodyData) SetTopic(v string) *GetDataChannelCredentialResponseBodyData {
	s.Topic = &v
	return s
}

func (s *GetDataChannelCredentialResponseBodyData) SetUserName(v string) *GetDataChannelCredentialResponseBodyData {
	s.UserName = &v
	return s
}

func (s *GetDataChannelCredentialResponseBodyData) Validate() error {
	return dara.Validate(s)
}

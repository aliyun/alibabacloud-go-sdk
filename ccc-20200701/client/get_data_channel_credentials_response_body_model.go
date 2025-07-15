// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataChannelCredentialsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDataChannelCredentialsResponseBody
	GetCode() *string
	SetData(v *GetDataChannelCredentialsResponseBodyData) *GetDataChannelCredentialsResponseBody
	GetData() *GetDataChannelCredentialsResponseBodyData
	SetHttpStatusCode(v int32) *GetDataChannelCredentialsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetDataChannelCredentialsResponseBody
	GetMessage() *string
	SetParams(v []*string) *GetDataChannelCredentialsResponseBody
	GetParams() []*string
	SetRequestId(v string) *GetDataChannelCredentialsResponseBody
	GetRequestId() *string
}

type GetDataChannelCredentialsResponseBody struct {
	Code           *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GetDataChannelCredentialsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                                     `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	Params         []*string                                  `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	RequestId      *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDataChannelCredentialsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataChannelCredentialsResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataChannelCredentialsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDataChannelCredentialsResponseBody) GetData() *GetDataChannelCredentialsResponseBodyData {
	return s.Data
}

func (s *GetDataChannelCredentialsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetDataChannelCredentialsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDataChannelCredentialsResponseBody) GetParams() []*string {
	return s.Params
}

func (s *GetDataChannelCredentialsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataChannelCredentialsResponseBody) SetCode(v string) *GetDataChannelCredentialsResponseBody {
	s.Code = &v
	return s
}

func (s *GetDataChannelCredentialsResponseBody) SetData(v *GetDataChannelCredentialsResponseBodyData) *GetDataChannelCredentialsResponseBody {
	s.Data = v
	return s
}

func (s *GetDataChannelCredentialsResponseBody) SetHttpStatusCode(v int32) *GetDataChannelCredentialsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDataChannelCredentialsResponseBody) SetMessage(v string) *GetDataChannelCredentialsResponseBody {
	s.Message = &v
	return s
}

func (s *GetDataChannelCredentialsResponseBody) SetParams(v []*string) *GetDataChannelCredentialsResponseBody {
	s.Params = v
	return s
}

func (s *GetDataChannelCredentialsResponseBody) SetRequestId(v string) *GetDataChannelCredentialsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataChannelCredentialsResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDataChannelCredentialsResponseBodyData struct {
	ClientId    *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	DeviceId    *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	Endpoint    *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	ExpiredTime *int64  `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	Password    *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Topic       *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	UserName    *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s GetDataChannelCredentialsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDataChannelCredentialsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDataChannelCredentialsResponseBodyData) GetClientId() *string {
	return s.ClientId
}

func (s *GetDataChannelCredentialsResponseBodyData) GetDeviceId() *string {
	return s.DeviceId
}

func (s *GetDataChannelCredentialsResponseBodyData) GetEndpoint() *string {
	return s.Endpoint
}

func (s *GetDataChannelCredentialsResponseBodyData) GetExpiredTime() *int64 {
	return s.ExpiredTime
}

func (s *GetDataChannelCredentialsResponseBodyData) GetPassword() *string {
	return s.Password
}

func (s *GetDataChannelCredentialsResponseBodyData) GetTopic() *string {
	return s.Topic
}

func (s *GetDataChannelCredentialsResponseBodyData) GetUserName() *string {
	return s.UserName
}

func (s *GetDataChannelCredentialsResponseBodyData) SetClientId(v string) *GetDataChannelCredentialsResponseBodyData {
	s.ClientId = &v
	return s
}

func (s *GetDataChannelCredentialsResponseBodyData) SetDeviceId(v string) *GetDataChannelCredentialsResponseBodyData {
	s.DeviceId = &v
	return s
}

func (s *GetDataChannelCredentialsResponseBodyData) SetEndpoint(v string) *GetDataChannelCredentialsResponseBodyData {
	s.Endpoint = &v
	return s
}

func (s *GetDataChannelCredentialsResponseBodyData) SetExpiredTime(v int64) *GetDataChannelCredentialsResponseBodyData {
	s.ExpiredTime = &v
	return s
}

func (s *GetDataChannelCredentialsResponseBodyData) SetPassword(v string) *GetDataChannelCredentialsResponseBodyData {
	s.Password = &v
	return s
}

func (s *GetDataChannelCredentialsResponseBodyData) SetTopic(v string) *GetDataChannelCredentialsResponseBodyData {
	s.Topic = &v
	return s
}

func (s *GetDataChannelCredentialsResponseBodyData) SetUserName(v string) *GetDataChannelCredentialsResponseBodyData {
	s.UserName = &v
	return s
}

func (s *GetDataChannelCredentialsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

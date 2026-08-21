// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMessageCallbackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessageCallback(v *GetMessageCallbackResponseBodyMessageCallback) *GetMessageCallbackResponseBody
	GetMessageCallback() *GetMessageCallbackResponseBodyMessageCallback
	SetRequestId(v string) *GetMessageCallbackResponseBody
	GetRequestId() *string
}

type GetMessageCallbackResponseBody struct {
	// The event notification configuration.
	MessageCallback *GetMessageCallbackResponseBodyMessageCallback `json:"MessageCallback,omitempty" xml:"MessageCallback,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 272A222A-F7F7-4A3E-****-F531574F1234
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMessageCallbackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMessageCallbackResponseBody) GoString() string {
	return s.String()
}

func (s *GetMessageCallbackResponseBody) GetMessageCallback() *GetMessageCallbackResponseBodyMessageCallback {
	return s.MessageCallback
}

func (s *GetMessageCallbackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMessageCallbackResponseBody) SetMessageCallback(v *GetMessageCallbackResponseBodyMessageCallback) *GetMessageCallbackResponseBody {
	s.MessageCallback = v
	return s
}

func (s *GetMessageCallbackResponseBody) SetRequestId(v string) *GetMessageCallbackResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMessageCallbackResponseBody) Validate() error {
	if s.MessageCallback != nil {
		if err := s.MessageCallback.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMessageCallbackResponseBodyMessageCallback struct {
	// The application ID.
	//
	// example:
	//
	// app-1000000
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The authentication key when the callback method is set to HTTP.
	//
	// example:
	//
	// 12345678abc
	AuthKey *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	// The callback authentication switch when the callback method is set to HTTP. Valid values:
	//
	// - **on**: enabled.
	//
	// - **off**: disabled.
	//
	// example:
	//
	// on
	AuthSwitch *string `json:"AuthSwitch,omitempty" xml:"AuthSwitch,omitempty"`
	// The callback method. Valid values:
	//
	// - **HTTP**
	//
	// - **MNS**
	//
	// example:
	//
	// HTTP
	CallbackType *string `json:"CallbackType,omitempty" xml:"CallbackType,omitempty"`
	// The callback URL when the callback method is set to HTTP.
	//
	// example:
	//
	// http://test.com/test
	CallbackURL *string `json:"CallbackURL,omitempty" xml:"CallbackURL,omitempty"`
	// The callback event types.
	//
	// example:
	//
	// FileUploadComplete,StreamTranscodeComplete,TranscodeComplete,SnapshotComplete,AIComplete,AddLiveRecordVideoComplete,CreateAuditComplete,UploadByURLComplete,ProduceMediaComplete,LiveRecordVideoComposeStart,ImageUploadComplete,VideoAnalysisComplete
	EventTypeList *string `json:"EventTypeList,omitempty" xml:"EventTypeList,omitempty"`
	// The public endpoint of the MSMQ when the callback method is set to MNS.
	//
	// example:
	//
	// http://1234567.mns.cn-shanghai-internal.aliyuncs.com/
	MnsEndpoint *string `json:"MnsEndpoint,omitempty" xml:"MnsEndpoint,omitempty"`
	// The name of the MSMQ when the callback method is set to MNS.
	//
	// example:
	//
	// vodcallback
	MnsQueueName *string `json:"MnsQueueName,omitempty" xml:"MnsQueueName,omitempty"`
}

func (s GetMessageCallbackResponseBodyMessageCallback) String() string {
	return dara.Prettify(s)
}

func (s GetMessageCallbackResponseBodyMessageCallback) GoString() string {
	return s.String()
}

func (s *GetMessageCallbackResponseBodyMessageCallback) GetAppId() *string {
	return s.AppId
}

func (s *GetMessageCallbackResponseBodyMessageCallback) GetAuthKey() *string {
	return s.AuthKey
}

func (s *GetMessageCallbackResponseBodyMessageCallback) GetAuthSwitch() *string {
	return s.AuthSwitch
}

func (s *GetMessageCallbackResponseBodyMessageCallback) GetCallbackType() *string {
	return s.CallbackType
}

func (s *GetMessageCallbackResponseBodyMessageCallback) GetCallbackURL() *string {
	return s.CallbackURL
}

func (s *GetMessageCallbackResponseBodyMessageCallback) GetEventTypeList() *string {
	return s.EventTypeList
}

func (s *GetMessageCallbackResponseBodyMessageCallback) GetMnsEndpoint() *string {
	return s.MnsEndpoint
}

func (s *GetMessageCallbackResponseBodyMessageCallback) GetMnsQueueName() *string {
	return s.MnsQueueName
}

func (s *GetMessageCallbackResponseBodyMessageCallback) SetAppId(v string) *GetMessageCallbackResponseBodyMessageCallback {
	s.AppId = &v
	return s
}

func (s *GetMessageCallbackResponseBodyMessageCallback) SetAuthKey(v string) *GetMessageCallbackResponseBodyMessageCallback {
	s.AuthKey = &v
	return s
}

func (s *GetMessageCallbackResponseBodyMessageCallback) SetAuthSwitch(v string) *GetMessageCallbackResponseBodyMessageCallback {
	s.AuthSwitch = &v
	return s
}

func (s *GetMessageCallbackResponseBodyMessageCallback) SetCallbackType(v string) *GetMessageCallbackResponseBodyMessageCallback {
	s.CallbackType = &v
	return s
}

func (s *GetMessageCallbackResponseBodyMessageCallback) SetCallbackURL(v string) *GetMessageCallbackResponseBodyMessageCallback {
	s.CallbackURL = &v
	return s
}

func (s *GetMessageCallbackResponseBodyMessageCallback) SetEventTypeList(v string) *GetMessageCallbackResponseBodyMessageCallback {
	s.EventTypeList = &v
	return s
}

func (s *GetMessageCallbackResponseBodyMessageCallback) SetMnsEndpoint(v string) *GetMessageCallbackResponseBodyMessageCallback {
	s.MnsEndpoint = &v
	return s
}

func (s *GetMessageCallbackResponseBodyMessageCallback) SetMnsQueueName(v string) *GetMessageCallbackResponseBodyMessageCallback {
	s.MnsQueueName = &v
	return s
}

func (s *GetMessageCallbackResponseBodyMessageCallback) Validate() error {
	return dara.Validate(s)
}

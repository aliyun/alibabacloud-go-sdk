// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveMessageAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeLiveMessageAppResponseBody
	GetAppId() *string
	SetAppKey(v string) *DescribeLiveMessageAppResponseBody
	GetAppKey() *string
	SetAppName(v string) *DescribeLiveMessageAppResponseBody
	GetAppName() *string
	SetAppSign(v string) *DescribeLiveMessageAppResponseBody
	GetAppSign() *string
	SetAuditType(v int32) *DescribeLiveMessageAppResponseBody
	GetAuditType() *int32
	SetAuditUrl(v string) *DescribeLiveMessageAppResponseBody
	GetAuditUrl() *string
	SetCreateTime(v int64) *DescribeLiveMessageAppResponseBody
	GetCreateTime() *int64
	SetDataCenter(v string) *DescribeLiveMessageAppResponseBody
	GetDataCenter() *string
	SetDisable(v bool) *DescribeLiveMessageAppResponseBody
	GetDisable() *bool
	SetEventCallbackUrl(v string) *DescribeLiveMessageAppResponseBody
	GetEventCallbackUrl() *string
	SetModifyTime(v int64) *DescribeLiveMessageAppResponseBody
	GetModifyTime() *int64
	SetMsgLifeCycle(v int32) *DescribeLiveMessageAppResponseBody
	GetMsgLifeCycle() *int32
	SetRequestId(v string) *DescribeLiveMessageAppResponseBody
	GetRequestId() *string
}

type DescribeLiveMessageAppResponseBody struct {
	// The ID of the interactive messaging application.
	//
	// example:
	//
	// demo
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The AppKey of the interactive messaging application. It is used to authorize operations related to the application ID.
	//
	// example:
	//
	// **********************************
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// The name of the interactive messaging application.
	//
	// example:
	//
	// testApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The signature of the interactive messaging application. It is required by the interactive messaging SDK.
	//
	// example:
	//
	// **************************************************************************
	AppSign *string `json:"AppSign,omitempty" xml:"AppSign,omitempty"`
	// The content moderation method. Valid values:
	//
	// 	- 0: Content moderation is disabled.
	//
	// 	- 1: Built-in content moderation is used.
	//
	// 	- 2: Custom content moderation is used.
	//
	// example:
	//
	// 2
	AuditType *int32 `json:"AuditType,omitempty" xml:"AuditType,omitempty"`
	// The URL for content moderation. This parameter is returned when the value of AuditType is 2.
	//
	// example:
	//
	// http://example.aliyundoc.com/exampleaudit
	AuditUrl *string `json:"AuditUrl,omitempty" xml:"AuditUrl,omitempty"`
	// The time when the interactive messaging application was created. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1698305471
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The data center.
	//
	// example:
	//
	// cn-shanghai
	DataCenter *string `json:"DataCenter,omitempty" xml:"DataCenter,omitempty"`
	// Indicates whether the interactive messaging application is disabled.
	//
	// example:
	//
	// false
	Disable *bool `json:"Disable,omitempty" xml:"Disable,omitempty"`
	// The callback URL for events such as user logon, logoff, joining a group, and leaving a group. An empty value indicates that callbacks are disabled.
	//
	// example:
	//
	// http://example.aliyundoc.com/examplecallback
	EventCallbackUrl *string `json:"EventCallbackUrl,omitempty" xml:"EventCallbackUrl,omitempty"`
	// The time when the interactive messaging application was modified. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1698305471
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The retention period of group messages in the application. Valid values:
	//
	// 	- 0 (default): 30 days
	//
	// 	- 1: 90 days
	//
	// 	- 2: 180 days
	//
	// example:
	//
	// 1
	MsgLifeCycle *int32 `json:"MsgLifeCycle,omitempty" xml:"MsgLifeCycle,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9FB68B5B-ED07-18F0-A3CF-083F4E74****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLiveMessageAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveMessageAppResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveMessageAppResponseBody) GetAppId() *string {
	return s.AppId
}

func (s *DescribeLiveMessageAppResponseBody) GetAppKey() *string {
	return s.AppKey
}

func (s *DescribeLiveMessageAppResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *DescribeLiveMessageAppResponseBody) GetAppSign() *string {
	return s.AppSign
}

func (s *DescribeLiveMessageAppResponseBody) GetAuditType() *int32 {
	return s.AuditType
}

func (s *DescribeLiveMessageAppResponseBody) GetAuditUrl() *string {
	return s.AuditUrl
}

func (s *DescribeLiveMessageAppResponseBody) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeLiveMessageAppResponseBody) GetDataCenter() *string {
	return s.DataCenter
}

func (s *DescribeLiveMessageAppResponseBody) GetDisable() *bool {
	return s.Disable
}

func (s *DescribeLiveMessageAppResponseBody) GetEventCallbackUrl() *string {
	return s.EventCallbackUrl
}

func (s *DescribeLiveMessageAppResponseBody) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *DescribeLiveMessageAppResponseBody) GetMsgLifeCycle() *int32 {
	return s.MsgLifeCycle
}

func (s *DescribeLiveMessageAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveMessageAppResponseBody) SetAppId(v string) *DescribeLiveMessageAppResponseBody {
	s.AppId = &v
	return s
}

func (s *DescribeLiveMessageAppResponseBody) SetAppKey(v string) *DescribeLiveMessageAppResponseBody {
	s.AppKey = &v
	return s
}

func (s *DescribeLiveMessageAppResponseBody) SetAppName(v string) *DescribeLiveMessageAppResponseBody {
	s.AppName = &v
	return s
}

func (s *DescribeLiveMessageAppResponseBody) SetAppSign(v string) *DescribeLiveMessageAppResponseBody {
	s.AppSign = &v
	return s
}

func (s *DescribeLiveMessageAppResponseBody) SetAuditType(v int32) *DescribeLiveMessageAppResponseBody {
	s.AuditType = &v
	return s
}

func (s *DescribeLiveMessageAppResponseBody) SetAuditUrl(v string) *DescribeLiveMessageAppResponseBody {
	s.AuditUrl = &v
	return s
}

func (s *DescribeLiveMessageAppResponseBody) SetCreateTime(v int64) *DescribeLiveMessageAppResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeLiveMessageAppResponseBody) SetDataCenter(v string) *DescribeLiveMessageAppResponseBody {
	s.DataCenter = &v
	return s
}

func (s *DescribeLiveMessageAppResponseBody) SetDisable(v bool) *DescribeLiveMessageAppResponseBody {
	s.Disable = &v
	return s
}

func (s *DescribeLiveMessageAppResponseBody) SetEventCallbackUrl(v string) *DescribeLiveMessageAppResponseBody {
	s.EventCallbackUrl = &v
	return s
}

func (s *DescribeLiveMessageAppResponseBody) SetModifyTime(v int64) *DescribeLiveMessageAppResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *DescribeLiveMessageAppResponseBody) SetMsgLifeCycle(v int32) *DescribeLiveMessageAppResponseBody {
	s.MsgLifeCycle = &v
	return s
}

func (s *DescribeLiveMessageAppResponseBody) SetRequestId(v string) *DescribeLiveMessageAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveMessageAppResponseBody) Validate() error {
	return dara.Validate(s)
}

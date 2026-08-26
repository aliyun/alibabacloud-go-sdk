// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendLiveMessageGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *SendLiveMessageGroupRequest
	GetAppId() *string
	SetBody(v string) *SendLiveMessageGroupRequest
	GetBody() *string
	SetDataCenter(v string) *SendLiveMessageGroupRequest
	GetDataCenter() *string
	SetGroupId(v string) *SendLiveMessageGroupRequest
	GetGroupId() *string
	SetMsgTid(v string) *SendLiveMessageGroupRequest
	GetMsgTid() *string
	SetMsgType(v int64) *SendLiveMessageGroupRequest
	GetMsgType() *int64
	SetNoCache(v bool) *SendLiveMessageGroupRequest
	GetNoCache() *bool
	SetNoStorage(v bool) *SendLiveMessageGroupRequest
	GetNoStorage() *bool
	SetSenderId(v string) *SendLiveMessageGroupRequest
	GetSenderId() *string
	SetSenderMetaInfo(v string) *SendLiveMessageGroupRequest
	GetSenderMetaInfo() *string
	SetStaticsIncrease(v int64) *SendLiveMessageGroupRequest
	GetStaticsIncrease() *int64
	SetWeight(v int64) *SendLiveMessageGroupRequest
	GetWeight() *int64
}

type SendLiveMessageGroupRequest struct {
	// The ID of the interactive messaging application that is used to receive the message.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The message body. It can be up to 15 KB in length.
	//
	// example:
	//
	// hello,group
	Body *string `json:"Body,omitempty" xml:"Body,omitempty"`
	// The data center, which must be the same as the data center specified in [CreateLiveMessageApp](https://help.aliyun.com/document_detail/2848162.html). Valid values:
	//
	// - cn-shanghai: Shanghai
	//
	// - ap-southeast-1: Singapore
	//
	// example:
	//
	// cn-shanghai
	DataCenter *string `json:"DataCenter,omitempty" xml:"DataCenter,omitempty"`
	// The ID of the group that receives the message.
	//
	// > Make sure that the specified GroupId is available. Otherwise, a ResourceNotExist error is returned.
	//
	// This parameter is required.
	//
	// example:
	//
	// grouptest
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The unique identifier of the message. This parameter is used to delete the message. The ID can contain only letters and digits and can be up to 64 bytes in length.
	//
	// example:
	//
	// 169830****
	MsgTid *string `json:"MsgTid,omitempty" xml:"MsgTid,omitempty"`
	// The message type. The total number of message types sent within a single group cannot exceed 30.
	//
	// example:
	//
	// 1
	MsgType *int64 `json:"MsgType,omitempty" xml:"MsgType,omitempty"`
	// Specifies whether to disable message caching to the list of the latest messages in the group. Valid values: true and false. Default value: false, which indicates that the message is cached.
	//
	// example:
	//
	// false
	NoCache *bool `json:"NoCache,omitempty" xml:"NoCache,omitempty"`
	// Specifies whether to disable message storage. Valid values: true and false. Default value: false, which indicates that the message is stored for a validity period of 30 days. You can find the message in the response of the ListLiveMessageGroupMessages operation. If you do not need to store the message, set this parameter to true.
	//
	// example:
	//
	// false
	NoStorage *bool `json:"NoStorage,omitempty" xml:"NoStorage,omitempty"`
	// The ID of the user who sends the message. The ID can be up to 64 bytes in length and can contain letters and digits.
	//
	// This parameter is required.
	//
	// example:
	//
	// uid1
	SenderId *string `json:"SenderId,omitempty" xml:"SenderId,omitempty"`
	// The additional information about the user who sends the message. The value can be up to 512 bytes in length.
	//
	// example:
	//
	// uid1meta1
	SenderMetaInfo *string `json:"SenderMetaInfo,omitempty" xml:"SenderMetaInfo,omitempty"`
	// The contribution of the message to the increase in the number of messages of this type. Default value: 1.
	//
	// example:
	//
	// 1
	StaticsIncrease *int64 `json:"StaticsIncrease,omitempty" xml:"StaticsIncrease,omitempty"`
	// The weight of the message. Default value: 1.
	//
	// - For low-priority messages such as likes, you can set the weight to 1.
	//
	// - For regular text messages such as bullet comments, you can set the weight to 5.
	//
	// - For high-priority messages such as red envelopes and gifts, you can set the weight to 1000000.
	//
	// example:
	//
	// 1
	Weight *int64 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s SendLiveMessageGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s SendLiveMessageGroupRequest) GoString() string {
	return s.String()
}

func (s *SendLiveMessageGroupRequest) GetAppId() *string {
	return s.AppId
}

func (s *SendLiveMessageGroupRequest) GetBody() *string {
	return s.Body
}

func (s *SendLiveMessageGroupRequest) GetDataCenter() *string {
	return s.DataCenter
}

func (s *SendLiveMessageGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *SendLiveMessageGroupRequest) GetMsgTid() *string {
	return s.MsgTid
}

func (s *SendLiveMessageGroupRequest) GetMsgType() *int64 {
	return s.MsgType
}

func (s *SendLiveMessageGroupRequest) GetNoCache() *bool {
	return s.NoCache
}

func (s *SendLiveMessageGroupRequest) GetNoStorage() *bool {
	return s.NoStorage
}

func (s *SendLiveMessageGroupRequest) GetSenderId() *string {
	return s.SenderId
}

func (s *SendLiveMessageGroupRequest) GetSenderMetaInfo() *string {
	return s.SenderMetaInfo
}

func (s *SendLiveMessageGroupRequest) GetStaticsIncrease() *int64 {
	return s.StaticsIncrease
}

func (s *SendLiveMessageGroupRequest) GetWeight() *int64 {
	return s.Weight
}

func (s *SendLiveMessageGroupRequest) SetAppId(v string) *SendLiveMessageGroupRequest {
	s.AppId = &v
	return s
}

func (s *SendLiveMessageGroupRequest) SetBody(v string) *SendLiveMessageGroupRequest {
	s.Body = &v
	return s
}

func (s *SendLiveMessageGroupRequest) SetDataCenter(v string) *SendLiveMessageGroupRequest {
	s.DataCenter = &v
	return s
}

func (s *SendLiveMessageGroupRequest) SetGroupId(v string) *SendLiveMessageGroupRequest {
	s.GroupId = &v
	return s
}

func (s *SendLiveMessageGroupRequest) SetMsgTid(v string) *SendLiveMessageGroupRequest {
	s.MsgTid = &v
	return s
}

func (s *SendLiveMessageGroupRequest) SetMsgType(v int64) *SendLiveMessageGroupRequest {
	s.MsgType = &v
	return s
}

func (s *SendLiveMessageGroupRequest) SetNoCache(v bool) *SendLiveMessageGroupRequest {
	s.NoCache = &v
	return s
}

func (s *SendLiveMessageGroupRequest) SetNoStorage(v bool) *SendLiveMessageGroupRequest {
	s.NoStorage = &v
	return s
}

func (s *SendLiveMessageGroupRequest) SetSenderId(v string) *SendLiveMessageGroupRequest {
	s.SenderId = &v
	return s
}

func (s *SendLiveMessageGroupRequest) SetSenderMetaInfo(v string) *SendLiveMessageGroupRequest {
	s.SenderMetaInfo = &v
	return s
}

func (s *SendLiveMessageGroupRequest) SetStaticsIncrease(v int64) *SendLiveMessageGroupRequest {
	s.StaticsIncrease = &v
	return s
}

func (s *SendLiveMessageGroupRequest) SetWeight(v int64) *SendLiveMessageGroupRequest {
	s.Weight = &v
	return s
}

func (s *SendLiveMessageGroupRequest) Validate() error {
	return dara.Validate(s)
}

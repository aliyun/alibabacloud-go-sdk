// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveMessageUserMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteLiveMessageUserMessageRequest
	GetAppId() *string
	SetDataCenter(v string) *DeleteLiveMessageUserMessageRequest
	GetDataCenter() *string
	SetDeleterId(v string) *DeleteLiveMessageUserMessageRequest
	GetDeleterId() *string
	SetDeleterInfo(v string) *DeleteLiveMessageUserMessageRequest
	GetDeleterInfo() *string
	SetMessageId(v string) *DeleteLiveMessageUserMessageRequest
	GetMessageId() *string
	SetReceiverId(v string) *DeleteLiveMessageUserMessageRequest
	GetReceiverId() *string
}

type DeleteLiveMessageUserMessageRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The data center, which must be the same as the data center specified in [CreateLiveMessageApp](https://help.aliyun.com/document_detail/2848162.html). Currently, Shanghai (cn-shanghai) and Singapore (ap-southeast-1) are supported.
	//
	// example:
	//
	// cn-shanghai
	DataCenter *string `json:"DataCenter,omitempty" xml:"DataCenter,omitempty"`
	// The ID of the user who initiates the message deletion. The value consists of uppercase and lowercase letters and digits, and cannot exceed 64 bytes in length. This parameter is required in practice. If not specified, InputInvalid is returned.
	//
	// example:
	//
	// 169830****
	DeleterId *string `json:"DeleterId,omitempty" xml:"DeleterId,omitempty"`
	// The extended information of the user who initiates the message deletion. The value cannot exceed 512 bytes in length.
	//
	// example:
	//
	// username
	DeleterInfo *string `json:"DeleterInfo,omitempty" xml:"DeleterInfo,omitempty"`
	// The ID of the message to be deleted, which corresponds to the MsgTid in the send message operation. The value consists of uppercase and lowercase letters and digits, and cannot exceed 64 bytes in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// **********
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// The ID of the user who receives the delete message notification. The value consists of uppercase and lowercase letters and digits, and cannot exceed 64 bytes in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// 169830****
	ReceiverId *string `json:"ReceiverId,omitempty" xml:"ReceiverId,omitempty"`
}

func (s DeleteLiveMessageUserMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveMessageUserMessageRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveMessageUserMessageRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteLiveMessageUserMessageRequest) GetDataCenter() *string {
	return s.DataCenter
}

func (s *DeleteLiveMessageUserMessageRequest) GetDeleterId() *string {
	return s.DeleterId
}

func (s *DeleteLiveMessageUserMessageRequest) GetDeleterInfo() *string {
	return s.DeleterInfo
}

func (s *DeleteLiveMessageUserMessageRequest) GetMessageId() *string {
	return s.MessageId
}

func (s *DeleteLiveMessageUserMessageRequest) GetReceiverId() *string {
	return s.ReceiverId
}

func (s *DeleteLiveMessageUserMessageRequest) SetAppId(v string) *DeleteLiveMessageUserMessageRequest {
	s.AppId = &v
	return s
}

func (s *DeleteLiveMessageUserMessageRequest) SetDataCenter(v string) *DeleteLiveMessageUserMessageRequest {
	s.DataCenter = &v
	return s
}

func (s *DeleteLiveMessageUserMessageRequest) SetDeleterId(v string) *DeleteLiveMessageUserMessageRequest {
	s.DeleterId = &v
	return s
}

func (s *DeleteLiveMessageUserMessageRequest) SetDeleterInfo(v string) *DeleteLiveMessageUserMessageRequest {
	s.DeleterInfo = &v
	return s
}

func (s *DeleteLiveMessageUserMessageRequest) SetMessageId(v string) *DeleteLiveMessageUserMessageRequest {
	s.MessageId = &v
	return s
}

func (s *DeleteLiveMessageUserMessageRequest) SetReceiverId(v string) *DeleteLiveMessageUserMessageRequest {
	s.ReceiverId = &v
	return s
}

func (s *DeleteLiveMessageUserMessageRequest) Validate() error {
	return dara.Validate(s)
}

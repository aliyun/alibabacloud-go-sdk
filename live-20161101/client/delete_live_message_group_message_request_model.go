// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveMessageGroupMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteLiveMessageGroupMessageRequest
	GetAppId() *string
	SetDataCenter(v string) *DeleteLiveMessageGroupMessageRequest
	GetDataCenter() *string
	SetDeleterId(v string) *DeleteLiveMessageGroupMessageRequest
	GetDeleterId() *string
	SetDeleterInfo(v string) *DeleteLiveMessageGroupMessageRequest
	GetDeleterInfo() *string
	SetGroupId(v string) *DeleteLiveMessageGroupMessageRequest
	GetGroupId() *string
	SetMessageId(v string) *DeleteLiveMessageGroupMessageRequest
	GetMessageId() *string
}

type DeleteLiveMessageGroupMessageRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The data center. It must be the same as the data center that was specified when you called the [CreateLiveMessageApp](https://help.aliyun.com/document_detail/2848162.html) operation to create the interactive messaging application. Valid values: cn-shanghai and ap-southeast-1 (Singapore).
	//
	// example:
	//
	// cn-shanghai
	DataCenter *string `json:"DataCenter,omitempty" xml:"DataCenter,omitempty"`
	// The ID of the user who deletes the message. The ID must be up to 64 bytes in length and can contain letters and digits.
	//
	// example:
	//
	// 169830****
	DeleterId *string `json:"DeleterId,omitempty" xml:"DeleterId,omitempty"`
	// The additional information about the user who performs the deletion operation. The value can be up to 512 bytes in length.
	//
	// example:
	//
	// testname
	DeleterInfo *string `json:"DeleterInfo,omitempty" xml:"DeleterInfo,omitempty"`
	// The group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// grouptest
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the message that you want to delete. Set this parameter to the value of MsgTid that you specified when you called the SendLiveMessageGroup operation. The ID must be up to 64 bytes in length and can contain letters and digits.
	//
	// This parameter is required.
	//
	// example:
	//
	// a74a8fbd3cfe4b2daa8517e4e3******
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
}

func (s DeleteLiveMessageGroupMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveMessageGroupMessageRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveMessageGroupMessageRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteLiveMessageGroupMessageRequest) GetDataCenter() *string {
	return s.DataCenter
}

func (s *DeleteLiveMessageGroupMessageRequest) GetDeleterId() *string {
	return s.DeleterId
}

func (s *DeleteLiveMessageGroupMessageRequest) GetDeleterInfo() *string {
	return s.DeleterInfo
}

func (s *DeleteLiveMessageGroupMessageRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DeleteLiveMessageGroupMessageRequest) GetMessageId() *string {
	return s.MessageId
}

func (s *DeleteLiveMessageGroupMessageRequest) SetAppId(v string) *DeleteLiveMessageGroupMessageRequest {
	s.AppId = &v
	return s
}

func (s *DeleteLiveMessageGroupMessageRequest) SetDataCenter(v string) *DeleteLiveMessageGroupMessageRequest {
	s.DataCenter = &v
	return s
}

func (s *DeleteLiveMessageGroupMessageRequest) SetDeleterId(v string) *DeleteLiveMessageGroupMessageRequest {
	s.DeleterId = &v
	return s
}

func (s *DeleteLiveMessageGroupMessageRequest) SetDeleterInfo(v string) *DeleteLiveMessageGroupMessageRequest {
	s.DeleterInfo = &v
	return s
}

func (s *DeleteLiveMessageGroupMessageRequest) SetGroupId(v string) *DeleteLiveMessageGroupMessageRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteLiveMessageGroupMessageRequest) SetMessageId(v string) *DeleteLiveMessageGroupMessageRequest {
	s.MessageId = &v
	return s
}

func (s *DeleteLiveMessageGroupMessageRequest) Validate() error {
	return dara.Validate(s)
}

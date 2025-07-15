// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLiveMessageUserInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ModifyLiveMessageUserInfoRequest
	GetAppId() *string
	SetDataCenter(v string) *ModifyLiveMessageUserInfoRequest
	GetDataCenter() *string
	SetUserId(v string) *ModifyLiveMessageUserInfoRequest
	GetUserId() *string
	SetUserMetaInfo(v string) *ModifyLiveMessageUserInfoRequest
	GetUserMetaInfo() *string
}

type ModifyLiveMessageUserInfoRequest struct {
	// The ID of the interactive messaging application whose user information you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The data center. It must be the same as the data center that was specified when you called the [CreateLiveMessageApp](https://help.aliyun.com/document_detail/2593195.html) operation to create the interactive messaging application. Valid values: cn-shanghai and ap-southeast-1 (Singapore).
	//
	// example:
	//
	// cn-shanghai
	DataCenter *string `json:"DataCenter,omitempty" xml:"DataCenter,omitempty"`
	// The ID of the user whose information you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// uid2
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The additional information about the user after the modification. The value can be up to 512 bytes in length.
	//
	// example:
	//
	// uid2meta2
	UserMetaInfo *string `json:"UserMetaInfo,omitempty" xml:"UserMetaInfo,omitempty"`
}

func (s ModifyLiveMessageUserInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyLiveMessageUserInfoRequest) GoString() string {
	return s.String()
}

func (s *ModifyLiveMessageUserInfoRequest) GetAppId() *string {
	return s.AppId
}

func (s *ModifyLiveMessageUserInfoRequest) GetDataCenter() *string {
	return s.DataCenter
}

func (s *ModifyLiveMessageUserInfoRequest) GetUserId() *string {
	return s.UserId
}

func (s *ModifyLiveMessageUserInfoRequest) GetUserMetaInfo() *string {
	return s.UserMetaInfo
}

func (s *ModifyLiveMessageUserInfoRequest) SetAppId(v string) *ModifyLiveMessageUserInfoRequest {
	s.AppId = &v
	return s
}

func (s *ModifyLiveMessageUserInfoRequest) SetDataCenter(v string) *ModifyLiveMessageUserInfoRequest {
	s.DataCenter = &v
	return s
}

func (s *ModifyLiveMessageUserInfoRequest) SetUserId(v string) *ModifyLiveMessageUserInfoRequest {
	s.UserId = &v
	return s
}

func (s *ModifyLiveMessageUserInfoRequest) SetUserMetaInfo(v string) *ModifyLiveMessageUserInfoRequest {
	s.UserMetaInfo = &v
	return s
}

func (s *ModifyLiveMessageUserInfoRequest) Validate() error {
	return dara.Validate(s)
}

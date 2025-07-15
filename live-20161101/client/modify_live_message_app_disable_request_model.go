// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLiveMessageAppDisableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ModifyLiveMessageAppDisableRequest
	GetAppId() *string
	SetDataCenter(v string) *ModifyLiveMessageAppDisableRequest
	GetDataCenter() *string
	SetDisable(v bool) *ModifyLiveMessageAppDisableRequest
	GetDisable() *bool
}

type ModifyLiveMessageAppDisableRequest struct {
	// The ID of the interactive messaging application.
	//
	// This parameter is required.
	//
	// example:
	//
	// ab6b5740****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The data center. It must be the same as the data center that was specified when you called the [CreateLiveMessageApp](https://help.aliyun.com/document_detail/2848162.html) operation to create the interactive messaging application. Valid values: cn-shanghai and ap-southeast-1 (Singapore).
	//
	// example:
	//
	// cn-shanghai
	DataCenter *string `json:"DataCenter,omitempty" xml:"DataCenter,omitempty"`
	// Specifies whether to disable the interactive messaging application.
	//
	// example:
	//
	// true
	Disable *bool `json:"Disable,omitempty" xml:"Disable,omitempty"`
}

func (s ModifyLiveMessageAppDisableRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyLiveMessageAppDisableRequest) GoString() string {
	return s.String()
}

func (s *ModifyLiveMessageAppDisableRequest) GetAppId() *string {
	return s.AppId
}

func (s *ModifyLiveMessageAppDisableRequest) GetDataCenter() *string {
	return s.DataCenter
}

func (s *ModifyLiveMessageAppDisableRequest) GetDisable() *bool {
	return s.Disable
}

func (s *ModifyLiveMessageAppDisableRequest) SetAppId(v string) *ModifyLiveMessageAppDisableRequest {
	s.AppId = &v
	return s
}

func (s *ModifyLiveMessageAppDisableRequest) SetDataCenter(v string) *ModifyLiveMessageAppDisableRequest {
	s.DataCenter = &v
	return s
}

func (s *ModifyLiveMessageAppDisableRequest) SetDisable(v bool) *ModifyLiveMessageAppDisableRequest {
	s.Disable = &v
	return s
}

func (s *ModifyLiveMessageAppDisableRequest) Validate() error {
	return dara.Validate(s)
}

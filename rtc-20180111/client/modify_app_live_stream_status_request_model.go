// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppLiveStreamStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ModifyAppLiveStreamStatusRequest
	GetAppId() *string
	SetClientToken(v string) *ModifyAppLiveStreamStatusRequest
	GetClientToken() *string
}

type ModifyAppLiveStreamStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ac7N****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 53200b81-b761-4c10-842a-a0726xxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s ModifyAppLiveStreamStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppLiveStreamStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyAppLiveStreamStatusRequest) GetAppId() *string {
	return s.AppId
}

func (s *ModifyAppLiveStreamStatusRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyAppLiveStreamStatusRequest) SetAppId(v string) *ModifyAppLiveStreamStatusRequest {
	s.AppId = &v
	return s
}

func (s *ModifyAppLiveStreamStatusRequest) SetClientToken(v string) *ModifyAppLiveStreamStatusRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyAppLiveStreamStatusRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v int32) *ModifyAppRequest
	GetAppId() *int32
	SetAppName(v string) *ModifyAppRequest
	GetAppName() *string
	SetDescription(v string) *ModifyAppRequest
	GetDescription() *string
	SetIconUrl(v string) *ModifyAppRequest
	GetIconUrl() *string
}

type ModifyAppRequest struct {
	// The ID of the application.
	//
	// example:
	//
	// 1234
	AppId *int32 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// defaultAppName
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// default description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The URL of the icon.
	//
	// example:
	//
	// https://defaultIcon.png
	IconUrl *string `json:"IconUrl,omitempty" xml:"IconUrl,omitempty"`
}

func (s ModifyAppRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppRequest) GoString() string {
	return s.String()
}

func (s *ModifyAppRequest) GetAppId() *int32 {
	return s.AppId
}

func (s *ModifyAppRequest) GetAppName() *string {
	return s.AppName
}

func (s *ModifyAppRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyAppRequest) GetIconUrl() *string {
	return s.IconUrl
}

func (s *ModifyAppRequest) SetAppId(v int32) *ModifyAppRequest {
	s.AppId = &v
	return s
}

func (s *ModifyAppRequest) SetAppName(v string) *ModifyAppRequest {
	s.AppName = &v
	return s
}

func (s *ModifyAppRequest) SetDescription(v string) *ModifyAppRequest {
	s.Description = &v
	return s
}

func (s *ModifyAppRequest) SetIconUrl(v string) *ModifyAppRequest {
	s.IconUrl = &v
	return s
}

func (s *ModifyAppRequest) Validate() error {
	return dara.Validate(s)
}

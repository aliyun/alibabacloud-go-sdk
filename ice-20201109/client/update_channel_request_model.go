// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessPolicy(v bool) *UpdateChannelRequest
	GetAccessPolicy() *bool
	SetAccessToken(v string) *UpdateChannelRequest
	GetAccessToken() *string
	SetChannelName(v string) *UpdateChannelRequest
	GetChannelName() *string
	SetFillerSourceLocationName(v string) *UpdateChannelRequest
	GetFillerSourceLocationName() *string
	SetFillerSourceName(v string) *UpdateChannelRequest
	GetFillerSourceName() *string
	SetOutPutConfigList(v string) *UpdateChannelRequest
	GetOutPutConfigList() *string
}

type UpdateChannelRequest struct {
	// Specifies whether to enable access control.
	//
	// example:
	//
	// true
	AccessPolicy *bool `json:"AccessPolicy,omitempty" xml:"AccessPolicy,omitempty"`
	// The token for accessing the channel.
	//
	// example:
	//
	// xxxxx
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// The name of the channel.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyChannel
	ChannelName *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	// The source location of the filler slate.
	//
	// example:
	//
	// MySourceLocation
	FillerSourceLocationName *string `json:"FillerSourceLocationName,omitempty" xml:"FillerSourceLocationName,omitempty"`
	// The name of the filler slate.
	//
	// example:
	//
	// MySource
	FillerSourceName *string `json:"FillerSourceName,omitempty" xml:"FillerSourceName,omitempty"`
	// The channel output configurations.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{
	//
	// 	"ManifestName": "manifest-1",
	//
	// 	"Format": "HLS",
	//
	// 	"SourceGroupName": "source-group-1",
	//
	// 	"ManifestSettings": {
	//
	// 		"WindowDuration": 60,
	//
	// 		"AdMarkType": "Daterange"
	//
	// 	}
	//
	// }]
	OutPutConfigList *string `json:"OutPutConfigList,omitempty" xml:"OutPutConfigList,omitempty"`
}

func (s UpdateChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateChannelRequest) GoString() string {
	return s.String()
}

func (s *UpdateChannelRequest) GetAccessPolicy() *bool {
	return s.AccessPolicy
}

func (s *UpdateChannelRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *UpdateChannelRequest) GetChannelName() *string {
	return s.ChannelName
}

func (s *UpdateChannelRequest) GetFillerSourceLocationName() *string {
	return s.FillerSourceLocationName
}

func (s *UpdateChannelRequest) GetFillerSourceName() *string {
	return s.FillerSourceName
}

func (s *UpdateChannelRequest) GetOutPutConfigList() *string {
	return s.OutPutConfigList
}

func (s *UpdateChannelRequest) SetAccessPolicy(v bool) *UpdateChannelRequest {
	s.AccessPolicy = &v
	return s
}

func (s *UpdateChannelRequest) SetAccessToken(v string) *UpdateChannelRequest {
	s.AccessToken = &v
	return s
}

func (s *UpdateChannelRequest) SetChannelName(v string) *UpdateChannelRequest {
	s.ChannelName = &v
	return s
}

func (s *UpdateChannelRequest) SetFillerSourceLocationName(v string) *UpdateChannelRequest {
	s.FillerSourceLocationName = &v
	return s
}

func (s *UpdateChannelRequest) SetFillerSourceName(v string) *UpdateChannelRequest {
	s.FillerSourceName = &v
	return s
}

func (s *UpdateChannelRequest) SetOutPutConfigList(v string) *UpdateChannelRequest {
	s.OutPutConfigList = &v
	return s
}

func (s *UpdateChannelRequest) Validate() error {
	return dara.Validate(s)
}

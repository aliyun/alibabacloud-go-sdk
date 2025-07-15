// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveStreamTranscodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApp(v string) *DeleteLiveStreamTranscodeRequest
	GetApp() *string
	SetDomain(v string) *DeleteLiveStreamTranscodeRequest
	GetDomain() *string
	SetOwnerId(v int64) *DeleteLiveStreamTranscodeRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DeleteLiveStreamTranscodeRequest
	GetSecurityToken() *string
	SetTemplate(v string) *DeleteLiveStreamTranscodeRequest
	GetTemplate() *string
}

type DeleteLiveStreamTranscodeRequest struct {
	// The name of the application to which the live stream belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveApp****
	App *string `json:"App,omitempty" xml:"App,omitempty"`
	// The main streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	Domain        *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The transcoding template ID. Valid values:
	//
	// 	- Standard transcoding:
	//
	//     	- lld: low definition
	//
	//     	- lsd: standard definition
	//
	//     	- lhd: high definition
	//
	//     	- lud : ultra-high definition
	//
	// 	- Narrowband HD™ transcoding:
	//
	//     	- ld: low definition
	//
	//     	- sd: standard definition
	//
	//     	- hd: high definition
	//
	//     	- ud: ultra-high definition
	//
	// 	- Custom transcoding: a custom ID
	//
	// This parameter is required.
	//
	// example:
	//
	// lld
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
}

func (s DeleteLiveStreamTranscodeRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveStreamTranscodeRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveStreamTranscodeRequest) GetApp() *string {
	return s.App
}

func (s *DeleteLiveStreamTranscodeRequest) GetDomain() *string {
	return s.Domain
}

func (s *DeleteLiveStreamTranscodeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteLiveStreamTranscodeRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DeleteLiveStreamTranscodeRequest) GetTemplate() *string {
	return s.Template
}

func (s *DeleteLiveStreamTranscodeRequest) SetApp(v string) *DeleteLiveStreamTranscodeRequest {
	s.App = &v
	return s
}

func (s *DeleteLiveStreamTranscodeRequest) SetDomain(v string) *DeleteLiveStreamTranscodeRequest {
	s.Domain = &v
	return s
}

func (s *DeleteLiveStreamTranscodeRequest) SetOwnerId(v int64) *DeleteLiveStreamTranscodeRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteLiveStreamTranscodeRequest) SetSecurityToken(v string) *DeleteLiveStreamTranscodeRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteLiveStreamTranscodeRequest) SetTemplate(v string) *DeleteLiveStreamTranscodeRequest {
	s.Template = &v
	return s
}

func (s *DeleteLiveStreamTranscodeRequest) Validate() error {
	return dara.Validate(s)
}

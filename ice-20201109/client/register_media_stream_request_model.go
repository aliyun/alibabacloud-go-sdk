// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterMediaStreamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInputURL(v string) *RegisterMediaStreamRequest
	GetInputURL() *string
	SetMediaId(v string) *RegisterMediaStreamRequest
	GetMediaId() *string
	SetStreamTags(v string) *RegisterMediaStreamRequest
	GetStreamTags() *string
	SetUserData(v string) *RegisterMediaStreamRequest
	GetUserData() *string
}

type RegisterMediaStreamRequest struct {
	// The URL of the media asset in another service. The URL is associated with the ID of the media asset in IMS. The URL cannot be modified once registered.
	//
	// Set this parameter to the OSS URL of the media asset. The following formats are supported:
	//
	// http(s)://example-bucket.oss-cn-shanghai.aliyuncs.com/example.mp4
	//
	// oss://example-bucket/example.mp4: In this format, it is considered by default that the region of the OSS bucket in which the media asset resides is the same as the region in which IMS is activated.
	InputURL *string `json:"InputURL,omitempty" xml:"InputURL,omitempty"`
	// The ID of the media asset.
	//
	// example:
	//
	// 5e778ec0027b71ed80a8909598506***
	MediaId    *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	StreamTags *string `json:"StreamTags,omitempty" xml:"StreamTags,omitempty"`
	// The user data.
	//
	// example:
	//
	// {"MessageCallback":{"CallbackURL":"http://test.test.com"}, "Extend":{"localId":"xxx","test":"www"}}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s RegisterMediaStreamRequest) String() string {
	return dara.Prettify(s)
}

func (s RegisterMediaStreamRequest) GoString() string {
	return s.String()
}

func (s *RegisterMediaStreamRequest) GetInputURL() *string {
	return s.InputURL
}

func (s *RegisterMediaStreamRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *RegisterMediaStreamRequest) GetStreamTags() *string {
	return s.StreamTags
}

func (s *RegisterMediaStreamRequest) GetUserData() *string {
	return s.UserData
}

func (s *RegisterMediaStreamRequest) SetInputURL(v string) *RegisterMediaStreamRequest {
	s.InputURL = &v
	return s
}

func (s *RegisterMediaStreamRequest) SetMediaId(v string) *RegisterMediaStreamRequest {
	s.MediaId = &v
	return s
}

func (s *RegisterMediaStreamRequest) SetStreamTags(v string) *RegisterMediaStreamRequest {
	s.StreamTags = &v
	return s
}

func (s *RegisterMediaStreamRequest) SetUserData(v string) *RegisterMediaStreamRequest {
	s.UserData = &v
	return s
}

func (s *RegisterMediaStreamRequest) Validate() error {
	return dara.Validate(s)
}

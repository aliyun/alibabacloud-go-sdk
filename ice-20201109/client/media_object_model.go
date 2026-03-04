// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMediaObject interface {
	dara.Model
	String() string
	GoString() string
	SetMedia(v string) *MediaObject
	GetMedia() *string
	SetType(v string) *MediaObject
	GetType() *string
}

type MediaObject struct {
	// The identifier for the media file.
	//
	// 	- If Type is set to OSS, the value is the URL of the media file. The following formats are supported: oss://... and https://...
	//
	// 	- If Type is set to Media, the value is the ID of the media asset.
	//
	// example:
	//
	// http://bucket.loction.aliyuncs.com/path/to/video.mp4
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The type of media source. Valid values:
	//
	// 	- OSS: an OSS object.
	//
	// 	- Media: a media asset.
	//
	// 	- ExternalURL: a publicly accessible external URL. This is not available for public use.
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s MediaObject) String() string {
	return dara.Prettify(s)
}

func (s MediaObject) GoString() string {
	return s.String()
}

func (s *MediaObject) GetMedia() *string {
	return s.Media
}

func (s *MediaObject) GetType() *string {
	return s.Type
}

func (s *MediaObject) SetMedia(v string) *MediaObject {
	s.Media = &v
	return s
}

func (s *MediaObject) SetType(v string) *MediaObject {
	s.Type = &v
	return s
}

func (s *MediaObject) Validate() error {
	return dara.Validate(s)
}

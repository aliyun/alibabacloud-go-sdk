// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertMediaToSearchLibRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImagesInput(v string) *InsertMediaToSearchLibRequest
	GetImagesInput() *string
	SetInput(v string) *InsertMediaToSearchLibRequest
	GetInput() *string
	SetMediaId(v string) *InsertMediaToSearchLibRequest
	GetMediaId() *string
	SetMediaType(v string) *InsertMediaToSearchLibRequest
	GetMediaType() *string
	SetMsgBody(v string) *InsertMediaToSearchLibRequest
	GetMsgBody() *string
	SetNamespace(v string) *InsertMediaToSearchLibRequest
	GetNamespace() *string
	SetSearchLibName(v string) *InsertMediaToSearchLibRequest
	GetSearchLibName() *string
}

type InsertMediaToSearchLibRequest struct {
	ImagesInput *string `json:"ImagesInput,omitempty" xml:"ImagesInput,omitempty"`
	// The URL of the video, audio, or image file that you want to import to the search library.
	//
	// Note: Make sure that you specify a correct file name and the bucket in which the file resides is in the same region where this operation is called. Otherwise, the file cannot be found or the operation may fail.
	//
	// Specify an Object Storage Service (OSS) URL in the following format: oss://[Bucket name]/[File path]. For example, you can specify oss://[example-bucket-****]/[object_path-****].
	//
	// Specify an HTTP URL in the following format: public endpoint. For example, you can specify http://example-test-\\*\\*\\*\\*.mp4.
	//
	// This parameter is required.
	//
	// example:
	//
	// http://example-test-****.mp4
	Input *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// The ID of the media asset. Each media ID is unique. If you leave this parameter empty, a media ID is automatically generated for this parameter.
	//
	// example:
	//
	// 411bed50018971edb60b0764a0ec6***
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The type of the media asset. Valid values:
	//
	// 	- video (default)
	//
	// 	- image
	//
	// 	- audio
	//
	// example:
	//
	// video
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// The message body.
	//
	// example:
	//
	// {}
	MsgBody   *string `json:"MsgBody,omitempty" xml:"MsgBody,omitempty"`
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The name of the search library. Default value: ims-default-search-lib.
	//
	// example:
	//
	// test1
	SearchLibName *string `json:"SearchLibName,omitempty" xml:"SearchLibName,omitempty"`
}

func (s InsertMediaToSearchLibRequest) String() string {
	return dara.Prettify(s)
}

func (s InsertMediaToSearchLibRequest) GoString() string {
	return s.String()
}

func (s *InsertMediaToSearchLibRequest) GetImagesInput() *string {
	return s.ImagesInput
}

func (s *InsertMediaToSearchLibRequest) GetInput() *string {
	return s.Input
}

func (s *InsertMediaToSearchLibRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *InsertMediaToSearchLibRequest) GetMediaType() *string {
	return s.MediaType
}

func (s *InsertMediaToSearchLibRequest) GetMsgBody() *string {
	return s.MsgBody
}

func (s *InsertMediaToSearchLibRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *InsertMediaToSearchLibRequest) GetSearchLibName() *string {
	return s.SearchLibName
}

func (s *InsertMediaToSearchLibRequest) SetImagesInput(v string) *InsertMediaToSearchLibRequest {
	s.ImagesInput = &v
	return s
}

func (s *InsertMediaToSearchLibRequest) SetInput(v string) *InsertMediaToSearchLibRequest {
	s.Input = &v
	return s
}

func (s *InsertMediaToSearchLibRequest) SetMediaId(v string) *InsertMediaToSearchLibRequest {
	s.MediaId = &v
	return s
}

func (s *InsertMediaToSearchLibRequest) SetMediaType(v string) *InsertMediaToSearchLibRequest {
	s.MediaType = &v
	return s
}

func (s *InsertMediaToSearchLibRequest) SetMsgBody(v string) *InsertMediaToSearchLibRequest {
	s.MsgBody = &v
	return s
}

func (s *InsertMediaToSearchLibRequest) SetNamespace(v string) *InsertMediaToSearchLibRequest {
	s.Namespace = &v
	return s
}

func (s *InsertMediaToSearchLibRequest) SetSearchLibName(v string) *InsertMediaToSearchLibRequest {
	s.SearchLibName = &v
	return s
}

func (s *InsertMediaToSearchLibRequest) Validate() error {
	return dara.Validate(s)
}

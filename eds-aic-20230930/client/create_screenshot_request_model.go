// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScreenshotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAndroidInstanceIdList(v []*string) *CreateScreenshotRequest
	GetAndroidInstanceIdList() []*string
	SetOssBucketName(v string) *CreateScreenshotRequest
	GetOssBucketName() *string
	SetScreenshotId(v string) *CreateScreenshotRequest
	GetScreenshotId() *string
	SetSkipCheckPolicyConfig(v string) *CreateScreenshotRequest
	GetSkipCheckPolicyConfig() *string
}

type CreateScreenshotRequest struct {
	// The list of instance IDs. Batch screenshots are supported.
	//
	// This parameter is required.
	AndroidInstanceIdList []*string `json:"AndroidInstanceIdList,omitempty" xml:"AndroidInstanceIdList,omitempty" type:"Repeated"`
	// The custom OSS bucket. The bucket name must start with the cloudphone-saved-bucket- prefix. The cloud phone instance and the OSS bucket must be in the same region. If you leave this parameter empty, a default bucket named cloudphone-saved-bucket-{RegionOfCloudPhone}-{AliUid} is created.
	//
	// example:
	//
	// cloudphone-saved-bucket-cn-shanghai-default
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// The screenshot ID. The generated screenshot is named "ScreenshotId_AndroidInstanceId.png."
	//
	// 	Notice:
	//
	// The ScreenshotId must be 2 to 128 characters long, beginning with a letter, but cannot start with http\\:// or https\\://. Allowed characters include letters, digits, underscores (_), and hyphens (-).
	//
	// example:
	//
	// image
	ScreenshotId *string `json:"ScreenshotId,omitempty" xml:"ScreenshotId,omitempty"`
	// Specifies whether to skip the screenshot policy check. The default value is false.
	//
	// example:
	//
	// false
	SkipCheckPolicyConfig *string `json:"SkipCheckPolicyConfig,omitempty" xml:"SkipCheckPolicyConfig,omitempty"`
}

func (s CreateScreenshotRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateScreenshotRequest) GoString() string {
	return s.String()
}

func (s *CreateScreenshotRequest) GetAndroidInstanceIdList() []*string {
	return s.AndroidInstanceIdList
}

func (s *CreateScreenshotRequest) GetOssBucketName() *string {
	return s.OssBucketName
}

func (s *CreateScreenshotRequest) GetScreenshotId() *string {
	return s.ScreenshotId
}

func (s *CreateScreenshotRequest) GetSkipCheckPolicyConfig() *string {
	return s.SkipCheckPolicyConfig
}

func (s *CreateScreenshotRequest) SetAndroidInstanceIdList(v []*string) *CreateScreenshotRequest {
	s.AndroidInstanceIdList = v
	return s
}

func (s *CreateScreenshotRequest) SetOssBucketName(v string) *CreateScreenshotRequest {
	s.OssBucketName = &v
	return s
}

func (s *CreateScreenshotRequest) SetScreenshotId(v string) *CreateScreenshotRequest {
	s.ScreenshotId = &v
	return s
}

func (s *CreateScreenshotRequest) SetSkipCheckPolicyConfig(v string) *CreateScreenshotRequest {
	s.SkipCheckPolicyConfig = &v
	return s
}

func (s *CreateScreenshotRequest) Validate() error {
	return dara.Validate(s)
}

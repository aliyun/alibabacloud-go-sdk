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
	SetSkipCheckPolicyConfig(v string) *CreateScreenshotRequest
	GetSkipCheckPolicyConfig() *string
}

type CreateScreenshotRequest struct {
	// The IDs of the cloud phone instances. You can create multiple snapshots simultaneously.
	//
	// This parameter is required.
	AndroidInstanceIdList []*string `json:"AndroidInstanceIdList,omitempty" xml:"AndroidInstanceIdList,omitempty" type:"Repeated"`
	// The name of the OSS bucket. The name must start with "cloudphone-saved-bucket-". The OSS bucket and the cloud phone instance must be in the same region. If you leave this parameter empty, the system will create a default OSS bucket named “cloudphone-saved-bucket-{Region of the cloud phone instance}-{AliUid}.”
	//
	// example:
	//
	// cloudphone-saved-bucket-cn-shanghai-default
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// Specifies whether to bypass the snapshot policy control. Default value: false.
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

func (s *CreateScreenshotRequest) SetSkipCheckPolicyConfig(v string) *CreateScreenshotRequest {
	s.SkipCheckPolicyConfig = &v
	return s
}

func (s *CreateScreenshotRequest) Validate() error {
	return dara.Validate(s)
}

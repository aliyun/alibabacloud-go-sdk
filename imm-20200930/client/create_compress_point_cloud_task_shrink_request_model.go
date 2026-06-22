// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCompressPointCloudTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCompressMethod(v string) *CreateCompressPointCloudTaskShrinkRequest
	GetCompressMethod() *string
	SetCredentialConfigShrink(v string) *CreateCompressPointCloudTaskShrinkRequest
	GetCredentialConfigShrink() *string
	SetKdtreeOptionShrink(v string) *CreateCompressPointCloudTaskShrinkRequest
	GetKdtreeOptionShrink() *string
	SetNotificationShrink(v string) *CreateCompressPointCloudTaskShrinkRequest
	GetNotificationShrink() *string
	SetOctreeOptionShrink(v string) *CreateCompressPointCloudTaskShrinkRequest
	GetOctreeOptionShrink() *string
	SetPointCloudFieldsShrink(v string) *CreateCompressPointCloudTaskShrinkRequest
	GetPointCloudFieldsShrink() *string
	SetPointCloudFileFormat(v string) *CreateCompressPointCloudTaskShrinkRequest
	GetPointCloudFileFormat() *string
	SetProjectName(v string) *CreateCompressPointCloudTaskShrinkRequest
	GetProjectName() *string
	SetSourceURI(v string) *CreateCompressPointCloudTaskShrinkRequest
	GetSourceURI() *string
	SetTagsShrink(v string) *CreateCompressPointCloudTaskShrinkRequest
	GetTagsShrink() *string
	SetTargetURI(v string) *CreateCompressPointCloudTaskShrinkRequest
	GetTargetURI() *string
	SetUserData(v string) *CreateCompressPointCloudTaskShrinkRequest
	GetUserData() *string
}

type CreateCompressPointCloudTaskShrinkRequest struct {
	// The name of the compression algorithm. Valid values:
	//
	// - octree: octree
	//
	// - kdtree: K-d tree
	//
	// This parameter is required.
	//
	// example:
	//
	// octree
	CompressMethod *string `json:"CompressMethod,omitempty" xml:"CompressMethod,omitempty"`
	// **If you do not have special requirements, leave this parameter empty.**
	//
	// The chained authorization configuration. This parameter is not required. For more information, see [Use chained authorization to access other entity resources](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// The parameters for K-d tree compression.
	KdtreeOptionShrink *string `json:"KdtreeOption,omitempty" xml:"KdtreeOption,omitempty"`
	// The notification configuration. For more information, click Notification. For the format of asynchronous notification messages, see [Asynchronous notification message format](https://help.aliyun.com/document_detail/2743997.html).
	//
	// > Intelligent Media Management API callbacks do not support specifying a webhook address. Use MNS instead.
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// The parameters for Octree compression.
	OctreeOptionShrink *string `json:"OctreeOption,omitempty" xml:"OctreeOption,omitempty"`
	// The PCD property fields to compress and the compression order. After compression, the data is decompressed in this order.
	//
	// - If you use Octree compression from the Point Cloud Library (PCL), only ["xyz"] is supported.
	//
	// - If you use K-d tree compression from the Draco library, ["xyz"] or ["xyz", "intensity"] is supported.
	//
	// This parameter is required.
	PointCloudFieldsShrink *string `json:"PointCloudFields,omitempty" xml:"PointCloudFields,omitempty"`
	// The format of the point cloud file. Only the PCD format is supported. The default value is pcd.
	//
	// example:
	//
	// pcd
	PointCloudFileFormat *string `json:"PointCloudFileFormat,omitempty" xml:"PointCloudFileFormat,omitempty"`
	// The project name. For more information, see [Create a project](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The OSS URI of the point cloud file.
	//
	// The URI must be in the format oss\\://${Bucket}/${Object}. ${Bucket} is the name of the OSS bucket in the same region as the project. ${Object} is the full path of the file, including the file name extension.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://test/src/test.pcd
	SourceURI *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	// Custom tags to search for and filter asynchronous tasks.
	//
	// example:
	//
	// {"LabelKey": "Value"}
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The OSS URI of the output file after compression.
	//
	// The URI must be in the format oss\\://${Bucket}/${Object}. ${Bucket} is the name of the OSS bucket in the same region as the project. ${Object} is the full path of the file, including the file name extension.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://test/tgt
	TargetURI *string `json:"TargetURI,omitempty" xml:"TargetURI,omitempty"`
	// Custom information that is returned in the asynchronous notification message. You can use this information to associate notification messages in your system. The maximum length is 2048 bytes.
	//
	// example:
	//
	// {"ID": "user1","Name": "test-user1","Avatar": "http://example.com?id=user1"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateCompressPointCloudTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCompressPointCloudTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateCompressPointCloudTaskShrinkRequest) GetCompressMethod() *string {
	return s.CompressMethod
}

func (s *CreateCompressPointCloudTaskShrinkRequest) GetCredentialConfigShrink() *string {
	return s.CredentialConfigShrink
}

func (s *CreateCompressPointCloudTaskShrinkRequest) GetKdtreeOptionShrink() *string {
	return s.KdtreeOptionShrink
}

func (s *CreateCompressPointCloudTaskShrinkRequest) GetNotificationShrink() *string {
	return s.NotificationShrink
}

func (s *CreateCompressPointCloudTaskShrinkRequest) GetOctreeOptionShrink() *string {
	return s.OctreeOptionShrink
}

func (s *CreateCompressPointCloudTaskShrinkRequest) GetPointCloudFieldsShrink() *string {
	return s.PointCloudFieldsShrink
}

func (s *CreateCompressPointCloudTaskShrinkRequest) GetPointCloudFileFormat() *string {
	return s.PointCloudFileFormat
}

func (s *CreateCompressPointCloudTaskShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateCompressPointCloudTaskShrinkRequest) GetSourceURI() *string {
	return s.SourceURI
}

func (s *CreateCompressPointCloudTaskShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *CreateCompressPointCloudTaskShrinkRequest) GetTargetURI() *string {
	return s.TargetURI
}

func (s *CreateCompressPointCloudTaskShrinkRequest) GetUserData() *string {
	return s.UserData
}

func (s *CreateCompressPointCloudTaskShrinkRequest) SetCompressMethod(v string) *CreateCompressPointCloudTaskShrinkRequest {
	s.CompressMethod = &v
	return s
}

func (s *CreateCompressPointCloudTaskShrinkRequest) SetCredentialConfigShrink(v string) *CreateCompressPointCloudTaskShrinkRequest {
	s.CredentialConfigShrink = &v
	return s
}

func (s *CreateCompressPointCloudTaskShrinkRequest) SetKdtreeOptionShrink(v string) *CreateCompressPointCloudTaskShrinkRequest {
	s.KdtreeOptionShrink = &v
	return s
}

func (s *CreateCompressPointCloudTaskShrinkRequest) SetNotificationShrink(v string) *CreateCompressPointCloudTaskShrinkRequest {
	s.NotificationShrink = &v
	return s
}

func (s *CreateCompressPointCloudTaskShrinkRequest) SetOctreeOptionShrink(v string) *CreateCompressPointCloudTaskShrinkRequest {
	s.OctreeOptionShrink = &v
	return s
}

func (s *CreateCompressPointCloudTaskShrinkRequest) SetPointCloudFieldsShrink(v string) *CreateCompressPointCloudTaskShrinkRequest {
	s.PointCloudFieldsShrink = &v
	return s
}

func (s *CreateCompressPointCloudTaskShrinkRequest) SetPointCloudFileFormat(v string) *CreateCompressPointCloudTaskShrinkRequest {
	s.PointCloudFileFormat = &v
	return s
}

func (s *CreateCompressPointCloudTaskShrinkRequest) SetProjectName(v string) *CreateCompressPointCloudTaskShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateCompressPointCloudTaskShrinkRequest) SetSourceURI(v string) *CreateCompressPointCloudTaskShrinkRequest {
	s.SourceURI = &v
	return s
}

func (s *CreateCompressPointCloudTaskShrinkRequest) SetTagsShrink(v string) *CreateCompressPointCloudTaskShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *CreateCompressPointCloudTaskShrinkRequest) SetTargetURI(v string) *CreateCompressPointCloudTaskShrinkRequest {
	s.TargetURI = &v
	return s
}

func (s *CreateCompressPointCloudTaskShrinkRequest) SetUserData(v string) *CreateCompressPointCloudTaskShrinkRequest {
	s.UserData = &v
	return s
}

func (s *CreateCompressPointCloudTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCompressPointCloudTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCompressMethod(v string) *CreateCompressPointCloudTaskRequest
	GetCompressMethod() *string
	SetCredentialConfig(v *CredentialConfig) *CreateCompressPointCloudTaskRequest
	GetCredentialConfig() *CredentialConfig
	SetKdtreeOption(v *KdtreeOption) *CreateCompressPointCloudTaskRequest
	GetKdtreeOption() *KdtreeOption
	SetNotification(v *Notification) *CreateCompressPointCloudTaskRequest
	GetNotification() *Notification
	SetOctreeOption(v *OctreeOption) *CreateCompressPointCloudTaskRequest
	GetOctreeOption() *OctreeOption
	SetPointCloudFields(v []*string) *CreateCompressPointCloudTaskRequest
	GetPointCloudFields() []*string
	SetPointCloudFileFormat(v string) *CreateCompressPointCloudTaskRequest
	GetPointCloudFileFormat() *string
	SetProjectName(v string) *CreateCompressPointCloudTaskRequest
	GetProjectName() *string
	SetSourceURI(v string) *CreateCompressPointCloudTaskRequest
	GetSourceURI() *string
	SetTags(v map[string]interface{}) *CreateCompressPointCloudTaskRequest
	GetTags() map[string]interface{}
	SetTargetURI(v string) *CreateCompressPointCloudTaskRequest
	GetTargetURI() *string
	SetUserData(v string) *CreateCompressPointCloudTaskRequest
	GetUserData() *string
}

type CreateCompressPointCloudTaskRequest struct {
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
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// The parameters for K-d tree compression.
	KdtreeOption *KdtreeOption `json:"KdtreeOption,omitempty" xml:"KdtreeOption,omitempty"`
	// The notification configuration. For more information, click Notification. For the format of asynchronous notification messages, see [Asynchronous notification message format](https://help.aliyun.com/document_detail/2743997.html).
	//
	// > Intelligent Media Management API callbacks do not support specifying a webhook address. Use MNS instead.
	Notification *Notification `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// The parameters for Octree compression.
	OctreeOption *OctreeOption `json:"OctreeOption,omitempty" xml:"OctreeOption,omitempty"`
	// The PCD property fields to compress and the compression order. After compression, the data is decompressed in this order.
	//
	// - If you use Octree compression from the Point Cloud Library (PCL), only ["xyz"] is supported.
	//
	// - If you use K-d tree compression from the Draco library, ["xyz"] or ["xyz", "intensity"] is supported.
	//
	// This parameter is required.
	PointCloudFields []*string `json:"PointCloudFields,omitempty" xml:"PointCloudFields,omitempty" type:"Repeated"`
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
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
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

func (s CreateCompressPointCloudTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCompressPointCloudTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateCompressPointCloudTaskRequest) GetCompressMethod() *string {
	return s.CompressMethod
}

func (s *CreateCompressPointCloudTaskRequest) GetCredentialConfig() *CredentialConfig {
	return s.CredentialConfig
}

func (s *CreateCompressPointCloudTaskRequest) GetKdtreeOption() *KdtreeOption {
	return s.KdtreeOption
}

func (s *CreateCompressPointCloudTaskRequest) GetNotification() *Notification {
	return s.Notification
}

func (s *CreateCompressPointCloudTaskRequest) GetOctreeOption() *OctreeOption {
	return s.OctreeOption
}

func (s *CreateCompressPointCloudTaskRequest) GetPointCloudFields() []*string {
	return s.PointCloudFields
}

func (s *CreateCompressPointCloudTaskRequest) GetPointCloudFileFormat() *string {
	return s.PointCloudFileFormat
}

func (s *CreateCompressPointCloudTaskRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateCompressPointCloudTaskRequest) GetSourceURI() *string {
	return s.SourceURI
}

func (s *CreateCompressPointCloudTaskRequest) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *CreateCompressPointCloudTaskRequest) GetTargetURI() *string {
	return s.TargetURI
}

func (s *CreateCompressPointCloudTaskRequest) GetUserData() *string {
	return s.UserData
}

func (s *CreateCompressPointCloudTaskRequest) SetCompressMethod(v string) *CreateCompressPointCloudTaskRequest {
	s.CompressMethod = &v
	return s
}

func (s *CreateCompressPointCloudTaskRequest) SetCredentialConfig(v *CredentialConfig) *CreateCompressPointCloudTaskRequest {
	s.CredentialConfig = v
	return s
}

func (s *CreateCompressPointCloudTaskRequest) SetKdtreeOption(v *KdtreeOption) *CreateCompressPointCloudTaskRequest {
	s.KdtreeOption = v
	return s
}

func (s *CreateCompressPointCloudTaskRequest) SetNotification(v *Notification) *CreateCompressPointCloudTaskRequest {
	s.Notification = v
	return s
}

func (s *CreateCompressPointCloudTaskRequest) SetOctreeOption(v *OctreeOption) *CreateCompressPointCloudTaskRequest {
	s.OctreeOption = v
	return s
}

func (s *CreateCompressPointCloudTaskRequest) SetPointCloudFields(v []*string) *CreateCompressPointCloudTaskRequest {
	s.PointCloudFields = v
	return s
}

func (s *CreateCompressPointCloudTaskRequest) SetPointCloudFileFormat(v string) *CreateCompressPointCloudTaskRequest {
	s.PointCloudFileFormat = &v
	return s
}

func (s *CreateCompressPointCloudTaskRequest) SetProjectName(v string) *CreateCompressPointCloudTaskRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateCompressPointCloudTaskRequest) SetSourceURI(v string) *CreateCompressPointCloudTaskRequest {
	s.SourceURI = &v
	return s
}

func (s *CreateCompressPointCloudTaskRequest) SetTags(v map[string]interface{}) *CreateCompressPointCloudTaskRequest {
	s.Tags = v
	return s
}

func (s *CreateCompressPointCloudTaskRequest) SetTargetURI(v string) *CreateCompressPointCloudTaskRequest {
	s.TargetURI = &v
	return s
}

func (s *CreateCompressPointCloudTaskRequest) SetUserData(v string) *CreateCompressPointCloudTaskRequest {
	s.UserData = &v
	return s
}

func (s *CreateCompressPointCloudTaskRequest) Validate() error {
	if s.CredentialConfig != nil {
		if err := s.CredentialConfig.Validate(); err != nil {
			return err
		}
	}
	if s.KdtreeOption != nil {
		if err := s.KdtreeOption.Validate(); err != nil {
			return err
		}
	}
	if s.Notification != nil {
		if err := s.Notification.Validate(); err != nil {
			return err
		}
	}
	if s.OctreeOption != nil {
		if err := s.OctreeOption.Validate(); err != nil {
			return err
		}
	}
	return nil
}

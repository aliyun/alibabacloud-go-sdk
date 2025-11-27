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
	// The compression algorithm. Valid values:
	//
	// 	- octree
	//
	// 	- kdtree
	//
	// This parameter is required.
	//
	// example:
	//
	// octree
	CompressMethod *string `json:"CompressMethod,omitempty" xml:"CompressMethod,omitempty"`
	// **If you have no special requirements, leave this parameter empty.**
	//
	// The configurations of authorization chains. This parameter is optional. For more information, see [Use authorization chains to access resources of other entities](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// The k-d tree compression options.
	KdtreeOption *KdtreeOption `json:"KdtreeOption,omitempty" xml:"KdtreeOption,omitempty"`
	// The notification settings. For information about the asynchronous notification format, see [Asynchronous message examples](https://help.aliyun.com/document_detail/2743997.html).
	//
	// >  The IMM operation does not support a callback URL. We recommend that you use Simple Message Queue (SMQ) to receive notifications.
	Notification *Notification `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// The octree compression options.
	OctreeOption *OctreeOption `json:"OctreeOption,omitempty" xml:"OctreeOption,omitempty"`
	// The PCD property fields and the compression order in which the data is decompressed after the compression is complete.
	//
	// 	- If octree of Point Cloud Library (PCL) is used for compression, ["xyz"] is supported.
	//
	// 	- If Draco k-dimensional (k-d) tree is used for compression, ["xyz"] and ["xyz", "intensity"] are supported.
	//
	// This parameter is required.
	PointCloudFields []*string `json:"PointCloudFields,omitempty" xml:"PointCloudFields,omitempty" type:"Repeated"`
	// The file format. Set the value to the default value: pcd.
	//
	// example:
	//
	// pcd
	PointCloudFileFormat *string `json:"PointCloudFileFormat,omitempty" xml:"PointCloudFileFormat,omitempty"`
	// The name of the project. For more information, see [CreateProject](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The OSS URL of the PCD file.
	//
	// Specify the value in the oss://${Bucket}/${Object} format. `${Bucket}` specifies the name of the OSS bucket that resides in the same region as the current project. `${Object}` specifies the path of the object with the extension included.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://test/src/test.pcd
	SourceURI *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	// The custom tags, which can be used to search for and filter asynchronous tasks.
	//
	// example:
	//
	// {"LabelKey": "Value"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The OSS URL of the output file after compression.
	//
	// Specify the value in the oss://${Bucket}/${Object} format. `${Bucket}` specifies the name of the OSS bucket that resides in the same region as the current project. `${Object}` specifies the path of the object with the extension included.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://test/tgt
	TargetURI *string `json:"TargetURI,omitempty" xml:"TargetURI,omitempty"`
	// The custom data, which is returned in an asynchronous notification and facilitates notification management. The maximum length is 2,048 bytes.
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

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUploadOSSFileJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *CreateUploadOSSFileJobRequest
	GetFileName() *string
	SetFileSource(v string) *CreateUploadOSSFileJobRequest
	GetFileSource() *string
	SetTid(v int64) *CreateUploadOSSFileJobRequest
	GetTid() *int64
	SetUploadTarget(v *CreateUploadOSSFileJobRequestUploadTarget) *CreateUploadOSSFileJobRequest
	GetUploadTarget() *CreateUploadOSSFileJobRequestUploadTarget
}

type CreateUploadOSSFileJobRequest struct {
	// The name of the file.
	//
	// > The file name must end with .txt or .sql. For example, the file name can be text.txt.
	//
	// This parameter is required.
	//
	// example:
	//
	// test.sql
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The purpose of the file upload task. Valid values:
	//
	// 	- **datacorrect**: The file is uploaded to change data.
	//
	// 	- **order_info_attachment**: The file is uploaded as an attachment in a ticket.
	//
	// 	- **big-file**: The file is uploaded to import multiple data records at a time.
	//
	// 	- **sqlreview**: The file is uploaded for SQL review.
	//
	// This parameter is required.
	//
	// example:
	//
	// datacorrect
	FileSource *string `json:"FileSource,omitempty" xml:"FileSource,omitempty"`
	// The ID of the tenant.
	//
	// > To view the ID of the tenant, move the pointer over the profile picture in the upper-right corner of the Data Management (DMS) console. For more information, see the "View information about the current tenant" section of the [Manage DMS tenants](https://help.aliyun.com/document_detail/181330.html) topic.
	//
	// example:
	//
	// -1
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// The information about the OSS file to be uploaded.
	//
	// This parameter is required.
	UploadTarget *CreateUploadOSSFileJobRequestUploadTarget `json:"UploadTarget,omitempty" xml:"UploadTarget,omitempty" type:"Struct"`
}

func (s CreateUploadOSSFileJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUploadOSSFileJobRequest) GoString() string {
	return s.String()
}

func (s *CreateUploadOSSFileJobRequest) GetFileName() *string {
	return s.FileName
}

func (s *CreateUploadOSSFileJobRequest) GetFileSource() *string {
	return s.FileSource
}

func (s *CreateUploadOSSFileJobRequest) GetTid() *int64 {
	return s.Tid
}

func (s *CreateUploadOSSFileJobRequest) GetUploadTarget() *CreateUploadOSSFileJobRequestUploadTarget {
	return s.UploadTarget
}

func (s *CreateUploadOSSFileJobRequest) SetFileName(v string) *CreateUploadOSSFileJobRequest {
	s.FileName = &v
	return s
}

func (s *CreateUploadOSSFileJobRequest) SetFileSource(v string) *CreateUploadOSSFileJobRequest {
	s.FileSource = &v
	return s
}

func (s *CreateUploadOSSFileJobRequest) SetTid(v int64) *CreateUploadOSSFileJobRequest {
	s.Tid = &v
	return s
}

func (s *CreateUploadOSSFileJobRequest) SetUploadTarget(v *CreateUploadOSSFileJobRequestUploadTarget) *CreateUploadOSSFileJobRequest {
	s.UploadTarget = v
	return s
}

func (s *CreateUploadOSSFileJobRequest) Validate() error {
	if s.UploadTarget != nil {
		if err := s.UploadTarget.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateUploadOSSFileJobRequestUploadTarget struct {
	// The name of the OSS bucket.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_bucket
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The endpoint of the OSS bucket.
	//
	// This parameter is required.
	//
	// example:
	//
	// http://oss-cn-hangzhou.aliyuncs.com
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The name of the OSS object.
	//
	// This parameter is required.
	//
	// example:
	//
	// test.sql
	ObjectName *string `json:"ObjectName,omitempty" xml:"ObjectName,omitempty"`
}

func (s CreateUploadOSSFileJobRequestUploadTarget) String() string {
	return dara.Prettify(s)
}

func (s CreateUploadOSSFileJobRequestUploadTarget) GoString() string {
	return s.String()
}

func (s *CreateUploadOSSFileJobRequestUploadTarget) GetBucketName() *string {
	return s.BucketName
}

func (s *CreateUploadOSSFileJobRequestUploadTarget) GetEndpoint() *string {
	return s.Endpoint
}

func (s *CreateUploadOSSFileJobRequestUploadTarget) GetObjectName() *string {
	return s.ObjectName
}

func (s *CreateUploadOSSFileJobRequestUploadTarget) SetBucketName(v string) *CreateUploadOSSFileJobRequestUploadTarget {
	s.BucketName = &v
	return s
}

func (s *CreateUploadOSSFileJobRequestUploadTarget) SetEndpoint(v string) *CreateUploadOSSFileJobRequestUploadTarget {
	s.Endpoint = &v
	return s
}

func (s *CreateUploadOSSFileJobRequestUploadTarget) SetObjectName(v string) *CreateUploadOSSFileJobRequestUploadTarget {
	s.ObjectName = &v
	return s
}

func (s *CreateUploadOSSFileJobRequestUploadTarget) Validate() error {
	return dara.Validate(s)
}

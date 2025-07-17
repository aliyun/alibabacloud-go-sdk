// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUploadOSSFileJobShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *CreateUploadOSSFileJobShrinkRequest
	GetFileName() *string
	SetFileSource(v string) *CreateUploadOSSFileJobShrinkRequest
	GetFileSource() *string
	SetTid(v int64) *CreateUploadOSSFileJobShrinkRequest
	GetTid() *int64
	SetUploadTargetShrink(v string) *CreateUploadOSSFileJobShrinkRequest
	GetUploadTargetShrink() *string
}

type CreateUploadOSSFileJobShrinkRequest struct {
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
	UploadTargetShrink *string `json:"UploadTarget,omitempty" xml:"UploadTarget,omitempty"`
}

func (s CreateUploadOSSFileJobShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUploadOSSFileJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateUploadOSSFileJobShrinkRequest) GetFileName() *string {
	return s.FileName
}

func (s *CreateUploadOSSFileJobShrinkRequest) GetFileSource() *string {
	return s.FileSource
}

func (s *CreateUploadOSSFileJobShrinkRequest) GetTid() *int64 {
	return s.Tid
}

func (s *CreateUploadOSSFileJobShrinkRequest) GetUploadTargetShrink() *string {
	return s.UploadTargetShrink
}

func (s *CreateUploadOSSFileJobShrinkRequest) SetFileName(v string) *CreateUploadOSSFileJobShrinkRequest {
	s.FileName = &v
	return s
}

func (s *CreateUploadOSSFileJobShrinkRequest) SetFileSource(v string) *CreateUploadOSSFileJobShrinkRequest {
	s.FileSource = &v
	return s
}

func (s *CreateUploadOSSFileJobShrinkRequest) SetTid(v int64) *CreateUploadOSSFileJobShrinkRequest {
	s.Tid = &v
	return s
}

func (s *CreateUploadOSSFileJobShrinkRequest) SetUploadTargetShrink(v string) *CreateUploadOSSFileJobShrinkRequest {
	s.UploadTargetShrink = &v
	return s
}

func (s *CreateUploadOSSFileJobShrinkRequest) Validate() error {
	return dara.Validate(s)
}

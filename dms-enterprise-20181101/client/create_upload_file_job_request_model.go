// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUploadFileJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *CreateUploadFileJobRequest
	GetFileName() *string
	SetFileSource(v string) *CreateUploadFileJobRequest
	GetFileSource() *string
	SetTid(v int64) *CreateUploadFileJobRequest
	GetTid() *int64
	SetUploadURL(v string) *CreateUploadFileJobRequest
	GetUploadURL() *string
}

type CreateUploadFileJobRequest struct {
	// The name of the attachment file.
	//
	// >  The file name must end with .txt or .sql. For example, the file name can be test.txt or test.sql.
	//
	// This parameter is required.
	//
	// example:
	//
	// test.txt
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The purpose of the attachment file. Valid values:
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
	// >  You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) operation to query the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// The URL of the attachment file. The URL must be an HTTP URL or an HTTPS URL.
	//
	// >  You can upload the attachment file to an Object Storage Service (OSS) bucket and obtain the URL of the file in the OSS console. For more information, see [Share objects](https://help.aliyun.com/document_detail/195674.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// https://XXX.oss-cn-hangzhou.aliyuncs.com/test.txt
	UploadURL *string `json:"UploadURL,omitempty" xml:"UploadURL,omitempty"`
}

func (s CreateUploadFileJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUploadFileJobRequest) GoString() string {
	return s.String()
}

func (s *CreateUploadFileJobRequest) GetFileName() *string {
	return s.FileName
}

func (s *CreateUploadFileJobRequest) GetFileSource() *string {
	return s.FileSource
}

func (s *CreateUploadFileJobRequest) GetTid() *int64 {
	return s.Tid
}

func (s *CreateUploadFileJobRequest) GetUploadURL() *string {
	return s.UploadURL
}

func (s *CreateUploadFileJobRequest) SetFileName(v string) *CreateUploadFileJobRequest {
	s.FileName = &v
	return s
}

func (s *CreateUploadFileJobRequest) SetFileSource(v string) *CreateUploadFileJobRequest {
	s.FileSource = &v
	return s
}

func (s *CreateUploadFileJobRequest) SetTid(v int64) *CreateUploadFileJobRequest {
	s.Tid = &v
	return s
}

func (s *CreateUploadFileJobRequest) SetUploadURL(v string) *CreateUploadFileJobRequest {
	s.UploadURL = &v
	return s
}

func (s *CreateUploadFileJobRequest) Validate() error {
	return dara.Validate(s)
}

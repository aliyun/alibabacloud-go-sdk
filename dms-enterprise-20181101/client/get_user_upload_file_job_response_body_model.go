// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserUploadFileJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetUserUploadFileJobResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetUserUploadFileJobResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetUserUploadFileJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetUserUploadFileJobResponseBody
	GetSuccess() *bool
	SetUploadFileJobDetail(v *GetUserUploadFileJobResponseBodyUploadFileJobDetail) *GetUserUploadFileJobResponseBody
	GetUploadFileJobDetail() *GetUserUploadFileJobResponseBodyUploadFileJobDetail
}

type GetUserUploadFileJobResponseBody struct {
	// The error code.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BDEFE9F2-B3B4-42D0-83AE-ECF9FC067DCD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The details of the file upload task.
	UploadFileJobDetail *GetUserUploadFileJobResponseBodyUploadFileJobDetail `json:"UploadFileJobDetail,omitempty" xml:"UploadFileJobDetail,omitempty" type:"Struct"`
}

func (s GetUserUploadFileJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserUploadFileJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserUploadFileJobResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetUserUploadFileJobResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetUserUploadFileJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserUploadFileJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetUserUploadFileJobResponseBody) GetUploadFileJobDetail() *GetUserUploadFileJobResponseBodyUploadFileJobDetail {
	return s.UploadFileJobDetail
}

func (s *GetUserUploadFileJobResponseBody) SetErrorCode(v string) *GetUserUploadFileJobResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetUserUploadFileJobResponseBody) SetErrorMessage(v string) *GetUserUploadFileJobResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetUserUploadFileJobResponseBody) SetRequestId(v string) *GetUserUploadFileJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserUploadFileJobResponseBody) SetSuccess(v bool) *GetUserUploadFileJobResponseBody {
	s.Success = &v
	return s
}

func (s *GetUserUploadFileJobResponseBody) SetUploadFileJobDetail(v *GetUserUploadFileJobResponseBodyUploadFileJobDetail) *GetUserUploadFileJobResponseBody {
	s.UploadFileJobDetail = v
	return s
}

func (s *GetUserUploadFileJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetUserUploadFileJobResponseBodyUploadFileJobDetail struct {
	// The key of the file that is returned after the file is uploaded. You can use this key when you upload the file as an attachment in a ticket.
	//
	// example:
	//
	// upload_3c7edea3-e4c3-4403-857d-737043036f69_test.sql
	AttachmentKey *string `json:"AttachmentKey,omitempty" xml:"AttachmentKey,omitempty"`
	// The name of the file.
	//
	// example:
	//
	// test.sql
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The size of the file. Unit: byte.
	//
	// example:
	//
	// 2968269
	FileSize *int64 `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// The purpose of the uploaded file. Valid values:
	//
	// 	- **datacorrect**: The file is uploaded to change data.
	//
	// 	- **order_info_attachment**: The file is uploaded as an attachment in a ticket.
	//
	// 	- **big-file**: The file is uploaded to import multiple data records at a time.
	//
	// 	- **sqlreview**: The file is uploaded for SQL review.
	//
	// example:
	//
	// datacorrect
	FileSource *string `json:"FileSource,omitempty" xml:"FileSource,omitempty"`
	// The key of the file upload task.
	//
	// example:
	//
	// 65254a4c1614235217749100e
	JobKey *string `json:"JobKey,omitempty" xml:"JobKey,omitempty"`
	// The status of the file upload task. Valid values:
	//
	// 	- **INIT**: The file upload task was initialized.
	//
	// 	- **PENDING**: The file upload task waited to be run.
	//
	// 	- **BE_SCHEDULED**: The file upload task waited to be scheduled.
	//
	// 	- **FAIL**: The file upload task failed.
	//
	// 	- **SUCCESS**: The file upload task was successful.
	//
	// 	- **RUNNING**: The file upload task was being run.
	//
	// example:
	//
	// SUCCESS
	JobStatus *string `json:"JobStatus,omitempty" xml:"JobStatus,omitempty"`
	// The information about the status of the file upload task.
	//
	// example:
	//
	// success
	JobStatusDesc *string `json:"JobStatusDesc,omitempty" xml:"JobStatusDesc,omitempty"`
	// The information about the Object Storage Service (OSS) bucket from which the file is uploaded.
	//
	// > This parameter is returned if the value of **UploadType*	- is **OSS**.
	UploadOSSParam *GetUserUploadFileJobResponseBodyUploadFileJobDetailUploadOSSParam `json:"UploadOSSParam,omitempty" xml:"UploadOSSParam,omitempty" type:"Struct"`
	// The method used to upload the file. Valid values:
	//
	// 	- **URL**
	//
	// 	- **OSS**
	//
	// example:
	//
	// URL
	UploadType *string `json:"UploadType,omitempty" xml:"UploadType,omitempty"`
	// The URL of the file.
	//
	// > This parameter is returned if the value of **UploadType*	- is **URL**.
	//
	// example:
	//
	// http://xxxx/test.sql
	UploadURL *string `json:"UploadURL,omitempty" xml:"UploadURL,omitempty"`
	// The size of the uploaded file. Unit: byte.
	//
	// example:
	//
	// 2968269
	UploadedSize *int64 `json:"UploadedSize,omitempty" xml:"UploadedSize,omitempty"`
}

func (s GetUserUploadFileJobResponseBodyUploadFileJobDetail) String() string {
	return dara.Prettify(s)
}

func (s GetUserUploadFileJobResponseBodyUploadFileJobDetail) GoString() string {
	return s.String()
}

func (s *GetUserUploadFileJobResponseBodyUploadFileJobDetail) GetAttachmentKey() *string {
	return s.AttachmentKey
}

func (s *GetUserUploadFileJobResponseBodyUploadFileJobDetail) GetFileName() *string {
	return s.FileName
}

func (s *GetUserUploadFileJobResponseBodyUploadFileJobDetail) GetFileSize() *int64 {
	return s.FileSize
}

func (s *GetUserUploadFileJobResponseBodyUploadFileJobDetail) GetFileSource() *string {
	return s.FileSource
}

func (s *GetUserUploadFileJobResponseBodyUploadFileJobDetail) GetJobKey() *string {
	return s.JobKey
}

func (s *GetUserUploadFileJobResponseBodyUploadFileJobDetail) GetJobStatus() *string {
	return s.JobStatus
}

func (s *GetUserUploadFileJobResponseBodyUploadFileJobDetail) GetJobStatusDesc() *string {
	return s.JobStatusDesc
}

func (s *GetUserUploadFileJobResponseBodyUploadFileJobDetail) GetUploadOSSParam() *GetUserUploadFileJobResponseBodyUploadFileJobDetailUploadOSSParam {
	return s.UploadOSSParam
}

func (s *GetUserUploadFileJobResponseBodyUploadFileJobDetail) GetUploadType() *string {
	return s.UploadType
}

func (s *GetUserUploadFileJobResponseBodyUploadFileJobDetail) GetUploadURL() *string {
	return s.UploadURL
}

func (s *GetUserUploadFileJobResponseBodyUploadFileJobDetail) GetUploadedSize() *int64 {
	return s.UploadedSize
}

func (s *GetUserUploadFileJobResponseBodyUploadFileJobDetail) SetAttachmentKey(v string) *GetUserUploadFileJobResponseBodyUploadFileJobDetail {
	s.AttachmentKey = &v
	return s
}

func (s *GetUserUploadFileJobResponseBodyUploadFileJobDetail) SetFileName(v string) *GetUserUploadFileJobResponseBodyUploadFileJobDetail {
	s.FileName = &v
	return s
}

func (s *GetUserUploadFileJobResponseBodyUploadFileJobDetail) SetFileSize(v int64) *GetUserUploadFileJobResponseBodyUploadFileJobDetail {
	s.FileSize = &v
	return s
}

func (s *GetUserUploadFileJobResponseBodyUploadFileJobDetail) SetFileSource(v string) *GetUserUploadFileJobResponseBodyUploadFileJobDetail {
	s.FileSource = &v
	return s
}

func (s *GetUserUploadFileJobResponseBodyUploadFileJobDetail) SetJobKey(v string) *GetUserUploadFileJobResponseBodyUploadFileJobDetail {
	s.JobKey = &v
	return s
}

func (s *GetUserUploadFileJobResponseBodyUploadFileJobDetail) SetJobStatus(v string) *GetUserUploadFileJobResponseBodyUploadFileJobDetail {
	s.JobStatus = &v
	return s
}

func (s *GetUserUploadFileJobResponseBodyUploadFileJobDetail) SetJobStatusDesc(v string) *GetUserUploadFileJobResponseBodyUploadFileJobDetail {
	s.JobStatusDesc = &v
	return s
}

func (s *GetUserUploadFileJobResponseBodyUploadFileJobDetail) SetUploadOSSParam(v *GetUserUploadFileJobResponseBodyUploadFileJobDetailUploadOSSParam) *GetUserUploadFileJobResponseBodyUploadFileJobDetail {
	s.UploadOSSParam = v
	return s
}

func (s *GetUserUploadFileJobResponseBodyUploadFileJobDetail) SetUploadType(v string) *GetUserUploadFileJobResponseBodyUploadFileJobDetail {
	s.UploadType = &v
	return s
}

func (s *GetUserUploadFileJobResponseBodyUploadFileJobDetail) SetUploadURL(v string) *GetUserUploadFileJobResponseBodyUploadFileJobDetail {
	s.UploadURL = &v
	return s
}

func (s *GetUserUploadFileJobResponseBodyUploadFileJobDetail) SetUploadedSize(v int64) *GetUserUploadFileJobResponseBodyUploadFileJobDetail {
	s.UploadedSize = &v
	return s
}

func (s *GetUserUploadFileJobResponseBodyUploadFileJobDetail) Validate() error {
	return dara.Validate(s)
}

type GetUserUploadFileJobResponseBodyUploadFileJobDetailUploadOSSParam struct {
	// The name of the OSS bucket.
	//
	// example:
	//
	// test_bucket
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The endpoint of the OSS bucket.
	//
	// example:
	//
	// http://oss-cn-hangzhou.aliyuncs.com
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The name of the OSS object.
	//
	// example:
	//
	// test.sql
	ObjectName *string `json:"ObjectName,omitempty" xml:"ObjectName,omitempty"`
}

func (s GetUserUploadFileJobResponseBodyUploadFileJobDetailUploadOSSParam) String() string {
	return dara.Prettify(s)
}

func (s GetUserUploadFileJobResponseBodyUploadFileJobDetailUploadOSSParam) GoString() string {
	return s.String()
}

func (s *GetUserUploadFileJobResponseBodyUploadFileJobDetailUploadOSSParam) GetBucketName() *string {
	return s.BucketName
}

func (s *GetUserUploadFileJobResponseBodyUploadFileJobDetailUploadOSSParam) GetEndpoint() *string {
	return s.Endpoint
}

func (s *GetUserUploadFileJobResponseBodyUploadFileJobDetailUploadOSSParam) GetObjectName() *string {
	return s.ObjectName
}

func (s *GetUserUploadFileJobResponseBodyUploadFileJobDetailUploadOSSParam) SetBucketName(v string) *GetUserUploadFileJobResponseBodyUploadFileJobDetailUploadOSSParam {
	s.BucketName = &v
	return s
}

func (s *GetUserUploadFileJobResponseBodyUploadFileJobDetailUploadOSSParam) SetEndpoint(v string) *GetUserUploadFileJobResponseBodyUploadFileJobDetailUploadOSSParam {
	s.Endpoint = &v
	return s
}

func (s *GetUserUploadFileJobResponseBodyUploadFileJobDetailUploadOSSParam) SetObjectName(v string) *GetUserUploadFileJobResponseBodyUploadFileJobDetailUploadOSSParam {
	s.ObjectName = &v
	return s
}

func (s *GetUserUploadFileJobResponseBodyUploadFileJobDetailUploadOSSParam) Validate() error {
	return dara.Validate(s)
}

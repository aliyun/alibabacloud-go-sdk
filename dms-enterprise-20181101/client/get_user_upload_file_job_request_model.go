// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserUploadFileJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobKey(v string) *GetUserUploadFileJobRequest
	GetJobKey() *string
	SetTid(v int64) *GetUserUploadFileJobRequest
	GetTid() *int64
}

type GetUserUploadFileJobRequest struct {
	// The key of the file upload task. The key is returned when you call the [CreateUploadFileJob](https://help.aliyun.com/document_detail/206059.html) or [CreateUploadOSSFileJob](https://help.aliyun.com/document_detail/206060.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 65254a49100e
	JobKey *string `json:"JobKey,omitempty" xml:"JobKey,omitempty"`
	// The tenant ID.
	//
	// > To view the tenant ID, move the pointer over the profile picture in the upper-right corner of the Data Management (DMS) console. For more information, see [Manage DMS tenants](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// -1
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetUserUploadFileJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserUploadFileJobRequest) GoString() string {
	return s.String()
}

func (s *GetUserUploadFileJobRequest) GetJobKey() *string {
	return s.JobKey
}

func (s *GetUserUploadFileJobRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetUserUploadFileJobRequest) SetJobKey(v string) *GetUserUploadFileJobRequest {
	s.JobKey = &v
	return s
}

func (s *GetUserUploadFileJobRequest) SetTid(v int64) *GetUserUploadFileJobRequest {
	s.Tid = &v
	return s
}

func (s *GetUserUploadFileJobRequest) Validate() error {
	return dara.Validate(s)
}

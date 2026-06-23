// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUploadTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSiteId(v int64) *GetUploadTaskRequest
	GetSiteId() *int64
	SetUploadId(v int64) *GetUploadTaskRequest
	GetUploadId() *int64
}

type GetUploadTaskRequest struct {
	// The site ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// 	Notice: This parameter is required when querying an upload task..
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The file upload task ID, which is assigned when you call the [UploadFile](https://help.aliyun.com/document_detail/2850466.html) operation.
	//
	// 	Notice: This parameter is required and must be a valid task ID returned by the UploadFile operation. If this parameter is not provided or the specified ID does not exist, an InvalidParameters (400) error is returned..
	//
	// example:
	//
	// 159253299357****
	UploadId *int64 `json:"UploadId,omitempty" xml:"UploadId,omitempty"`
}

func (s GetUploadTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUploadTaskRequest) GoString() string {
	return s.String()
}

func (s *GetUploadTaskRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetUploadTaskRequest) GetUploadId() *int64 {
	return s.UploadId
}

func (s *GetUploadTaskRequest) SetSiteId(v int64) *GetUploadTaskRequest {
	s.SiteId = &v
	return s
}

func (s *GetUploadTaskRequest) SetUploadId(v int64) *GetUploadTaskRequest {
	s.UploadId = &v
	return s
}

func (s *GetUploadTaskRequest) Validate() error {
	return dara.Validate(s)
}

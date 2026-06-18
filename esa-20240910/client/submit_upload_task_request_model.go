// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitUploadTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForce(v bool) *SubmitUploadTaskRequest
	GetForce() *bool
	SetSiteId(v int64) *SubmitUploadTaskRequest
	GetSiteId() *int64
	SetUploadId(v int64) *SubmitUploadTaskRequest
	GetUploadId() *int64
}

type SubmitUploadTaskRequest struct {
	// Specifies whether to refresh resources in the corresponding directory if the requested content is different from that on the origin server. Default value: false. This parameter takes effect for a purge task.
	//
	// 	- **true**: purges all resources in the directory.
	//
	// 	- **false**: refresh the changed resources in the directory.
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// The website ID. You can call the [ListSites](~~ListSites~~) operation to obtain the ID.
	//
	// example:
	//
	// 123456789****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The ID of the file upload task, which is generated when you call [UploadTask](~~UploadTask~~).
	//
	// example:
	//
	// 1593805857882113
	UploadId *int64 `json:"UploadId,omitempty" xml:"UploadId,omitempty"`
}

func (s SubmitUploadTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitUploadTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitUploadTaskRequest) GetForce() *bool {
	return s.Force
}

func (s *SubmitUploadTaskRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *SubmitUploadTaskRequest) GetUploadId() *int64 {
	return s.UploadId
}

func (s *SubmitUploadTaskRequest) SetForce(v bool) *SubmitUploadTaskRequest {
	s.Force = &v
	return s
}

func (s *SubmitUploadTaskRequest) SetSiteId(v int64) *SubmitUploadTaskRequest {
	s.SiteId = &v
	return s
}

func (s *SubmitUploadTaskRequest) SetUploadId(v int64) *SubmitUploadTaskRequest {
	s.UploadId = &v
	return s
}

func (s *SubmitUploadTaskRequest) Validate() error {
	return dara.Validate(s)
}

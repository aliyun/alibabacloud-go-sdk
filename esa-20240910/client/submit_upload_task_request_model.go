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
	// Specifies whether to purge resources in the corresponding directory when the back-to-origin content is inconsistent with the origin server resources. Default value: false. This parameter is valid only for purge tasks.
	//
	// - **true**: Purges all resources in the corresponding directory.
	//
	// - **false**: Purges only the changed resources in the corresponding directory.
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// The site ID. You can call the [ListSites](~~ListSites~~) operation to obtain the site ID.
	//
	// example:
	//
	// 123456789****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The file upload task ID, which is generated when you call the [UploadTask](~~UploadTask~~) operation.
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

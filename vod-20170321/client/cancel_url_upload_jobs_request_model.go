// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelUrlUploadJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobIds(v string) *CancelUrlUploadJobsRequest
	GetJobIds() *string
	SetUploadUrls(v string) *CancelUrlUploadJobsRequest
	GetUploadUrls() *string
}

type CancelUrlUploadJobsRequest struct {
	// The list of task IDs. You can obtain the task ID (JobId) from the PlayInfo struct returned by the [GetPlayInfo](https://help.aliyun.com/document_detail/56124.html) operation.
	//
	// - A maximum of 10 IDs are supported.
	//
	// - Separate multiple IDs with commas (,).
	//
	// > You must specify either JobIds or UploadUrls. If both are specified, only JobIds is processed.
	//
	// example:
	//
	// 341c92e6c18dc435ee31253685****,0193d395194a83ad6ee2ef27a5b5****
	JobIds *string `json:"JobIds,omitempty" xml:"JobIds,omitempty"`
	// The list of source video upload URLs. Separate multiple URLs with commas (,). A maximum of 10 URLs are supported.
	//
	// > - URL-encode the URLs before use.
	//
	// > - You must specify either JobIds or UploadUrls. If both are specified, only JobIds is processed.
	UploadUrls *string `json:"UploadUrls,omitempty" xml:"UploadUrls,omitempty"`
}

func (s CancelUrlUploadJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelUrlUploadJobsRequest) GoString() string {
	return s.String()
}

func (s *CancelUrlUploadJobsRequest) GetJobIds() *string {
	return s.JobIds
}

func (s *CancelUrlUploadJobsRequest) GetUploadUrls() *string {
	return s.UploadUrls
}

func (s *CancelUrlUploadJobsRequest) SetJobIds(v string) *CancelUrlUploadJobsRequest {
	s.JobIds = &v
	return s
}

func (s *CancelUrlUploadJobsRequest) SetUploadUrls(v string) *CancelUrlUploadJobsRequest {
	s.UploadUrls = &v
	return s
}

func (s *CancelUrlUploadJobsRequest) Validate() error {
	return dara.Validate(s)
}

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
	// The IDs of the upload jobs. You can obtain the job IDs from PlayInfo in the response to the [GetPlayInfo](https://help.aliyun.com/document_detail/56124.html) operation.
	//
	// 	- You can specify a maximum of 10 IDs.
	//
	// 	- Separate multiple IDs with commas (,).
	//
	// >  You must specify either JobIds or UploadUrls. If you specify both the JobIds and UploadUrls parameters, only the value of the JobIds parameter takes effect.
	//
	// example:
	//
	// 341c92e6c18dc435ee31253685****,0193d395194a83ad6ee2ef27a5b5****
	JobIds *string `json:"JobIds,omitempty" xml:"JobIds,omitempty"`
	// The upload URLs of source video files. Separate multiple URLs with commas (,). You can specify a maximum of 10 URLs.
	//
	// > 	- You must encode the URLs before you use the URLs.
	//
	// > 	- You must specify either JobIds or UploadUrls. If you specify both the JobIds and UploadUrls parameters, only the value of the JobIds parameter takes effect.
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

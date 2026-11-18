// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetURLUploadInfosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobIds(v string) *GetURLUploadInfosRequest
	GetJobIds() *string
	SetUploadURLs(v string) *GetURLUploadInfosRequest
	GetUploadURLs() *string
}

type GetURLUploadInfosRequest struct {
	// The list of upload task IDs (JobIds). It consists of one or more JobIds. A JobId is the value of the JobId response parameter returned when you call the [UploadMediaByURL](https://help.aliyun.com/document_detail/86311.html) operation.
	//
	// - A maximum of 10 IDs are supported.
	//
	// - Separate multiple IDs with commas (,).
	//
	// > You must specify either JobIds or UploadURLs. If both are specified, only JobIds is processed.
	//
	// example:
	//
	// 86c1925fba0****,7afb201e7fa****,2cc4997378****
	JobIds *string `json:"JobIds,omitempty" xml:"JobIds,omitempty"`
	// The list of URLs of the video source files to upload. Separate multiple URLs with commas (,). A maximum of 10 URLs are supported.
	//
	// > - URL-encode the URLs before you use them.
	//
	// > - If the same video URL is uploaded multiple times, specify a single URL for the query.
	//
	// > - You must specify either JobIds or UploadURLs. If both are specified, only JobIds is processed.
	//
	// example:
	//
	// http://****.mp4
	UploadURLs *string `json:"UploadURLs,omitempty" xml:"UploadURLs,omitempty"`
}

func (s GetURLUploadInfosRequest) String() string {
	return dara.Prettify(s)
}

func (s GetURLUploadInfosRequest) GoString() string {
	return s.String()
}

func (s *GetURLUploadInfosRequest) GetJobIds() *string {
	return s.JobIds
}

func (s *GetURLUploadInfosRequest) GetUploadURLs() *string {
	return s.UploadURLs
}

func (s *GetURLUploadInfosRequest) SetJobIds(v string) *GetURLUploadInfosRequest {
	s.JobIds = &v
	return s
}

func (s *GetURLUploadInfosRequest) SetUploadURLs(v string) *GetURLUploadInfosRequest {
	s.UploadURLs = &v
	return s
}

func (s *GetURLUploadInfosRequest) Validate() error {
	return dara.Validate(s)
}

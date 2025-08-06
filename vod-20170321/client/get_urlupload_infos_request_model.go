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
	// The one or more IDs of upload jobs. You can obtain the job IDs in the response parameter JobId of the [UploadMediaByURL](https://help.aliyun.com/document_detail/86311.html) operation.
	//
	// - You can specify a maximum of 10 IDs.
	//
	// - Separate multiple IDs with commas (,).
	//
	// > You must set one of the JobIds and the UploadURLs parameters. If you set both the JobIds and UploadURLs parameters, only the value of the JobIds parameter takes effect.
	//
	// example:
	//
	// 86c1925fba0****,7afb201e7fa****,2cc4997378****
	JobIds *string `json:"JobIds,omitempty" xml:"JobIds,omitempty"`
	// The one or more upload URLs of the source files. Separate multiple URLs with commas (,). You can specify a maximum of 10 URLs.
	//
	// > 	- You must encode the URLs before you use the URLs.
	//
	// > 	- If a media file is uploaded multiple times, pass the URL of the media file to this parameter only once.
	//
	// > 	- You must set one of the JobIds and the UploadURLs parameters. If you set both the JobIds and UploadURLs parameters, only the value of the JobIds parameter takes effect.
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

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUrlUploadInfosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobIds(v string) *GetUrlUploadInfosRequest
	GetJobIds() *string
	SetUploadURLs(v string) *GetUrlUploadInfosRequest
	GetUploadURLs() *string
}

type GetUrlUploadInfosRequest struct {
	// The IDs of the upload jobs. You can specify one or more job IDs. You can obtain the job IDs from the response parameter JobId of the [UploadMediaByURL](https://help.aliyun.com/document_detail/86311.html) operation.
	//
	// 	- You can specify a maximum of 10 job IDs.
	//
	// 	- Separate the job IDs with commas (,).
	//
	// >  You must specify either JobIds or UploadURLs. If you specify both parameters, only the value of JobIds takes effect.
	//
	// example:
	//
	// df2ac80b481346daa1db6a7c40edc7f8
	JobIds *string `json:"JobIds,omitempty" xml:"JobIds,omitempty"`
	// The upload URLs of the source files. You can specify a maximum of 10 URLs. Separate the URLs with commas (,).
	//
	// >
	//
	// 	- The URLs must be encoded.
	//
	// 	- If a media file is uploaded multiple times, we recommend that you specify the URL of the media file only once in this parameter.
	//
	// 	- You must specify either JobIds or UploadURLs. If you specify both parameters, only the value of JobIds takes effect.
	//
	// example:
	//
	// https://media.w3.org/2010/05/sintel/trailer.mp4
	UploadURLs *string `json:"UploadURLs,omitempty" xml:"UploadURLs,omitempty"`
}

func (s GetUrlUploadInfosRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUrlUploadInfosRequest) GoString() string {
	return s.String()
}

func (s *GetUrlUploadInfosRequest) GetJobIds() *string {
	return s.JobIds
}

func (s *GetUrlUploadInfosRequest) GetUploadURLs() *string {
	return s.UploadURLs
}

func (s *GetUrlUploadInfosRequest) SetJobIds(v string) *GetUrlUploadInfosRequest {
	s.JobIds = &v
	return s
}

func (s *GetUrlUploadInfosRequest) SetUploadURLs(v string) *GetUrlUploadInfosRequest {
	s.UploadURLs = &v
	return s
}

func (s *GetUrlUploadInfosRequest) Validate() error {
	return dara.Validate(s)
}

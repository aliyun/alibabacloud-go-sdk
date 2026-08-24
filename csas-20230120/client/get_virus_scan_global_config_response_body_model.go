// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVirusScanGlobalConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetVirusScanGlobalConfigResponseBody
	GetRequestId() *string
	SetUploadFileMaxSize(v int64) *GetVirusScanGlobalConfigResponseBody
	GetUploadFileMaxSize() *int64
	SetUploadFileMaxSpeed(v int64) *GetVirusScanGlobalConfigResponseBody
	GetUploadFileMaxSpeed() *int64
	SetUploadFileSuffixBlacklist(v []*string) *GetVirusScanGlobalConfigResponseBody
	GetUploadFileSuffixBlacklist() []*string
	SetVirusFileUpload(v bool) *GetVirusScanGlobalConfigResponseBody
	GetVirusFileUpload() *bool
}

type GetVirusScanGlobalConfigResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 3D7EC0AF-DB2A-5D9C-90EC-F090A6BAAEA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The maximum size of a single virus file that can be uploaded, in KB. A value of 0 indicates no size limit.
	//
	// example:
	//
	// 10240
	UploadFileMaxSize *int64 `json:"UploadFileMaxSize,omitempty" xml:"UploadFileMaxSize,omitempty"`
	// The maximum upload rate for virus files, in KB/s. A value of 0 indicates no rate limit.
	//
	// example:
	//
	// 1024
	UploadFileMaxSpeed *int64 `json:"UploadFileMaxSpeed,omitempty" xml:"UploadFileMaxSpeed,omitempty"`
	// The collection of file types that are prohibited from being uploaded. Files that match these types are not uploaded even if the upload feature is enabled. An empty list is returned if no file types are configured.
	UploadFileSuffixBlacklist []*string `json:"UploadFileSuffixBlacklist,omitempty" xml:"UploadFileSuffixBlacklist,omitempty" type:"Repeated"`
	// Indicates whether user terminal devices are allowed to upload detected virus files to the cloud for further analysis. Valid values:
	//
	// - **true**: Upload is allowed.
	//
	// - **false**: Upload is not allowed.
	//
	// example:
	//
	// true
	VirusFileUpload *bool `json:"VirusFileUpload,omitempty" xml:"VirusFileUpload,omitempty"`
}

func (s GetVirusScanGlobalConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVirusScanGlobalConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetVirusScanGlobalConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVirusScanGlobalConfigResponseBody) GetUploadFileMaxSize() *int64 {
	return s.UploadFileMaxSize
}

func (s *GetVirusScanGlobalConfigResponseBody) GetUploadFileMaxSpeed() *int64 {
	return s.UploadFileMaxSpeed
}

func (s *GetVirusScanGlobalConfigResponseBody) GetUploadFileSuffixBlacklist() []*string {
	return s.UploadFileSuffixBlacklist
}

func (s *GetVirusScanGlobalConfigResponseBody) GetVirusFileUpload() *bool {
	return s.VirusFileUpload
}

func (s *GetVirusScanGlobalConfigResponseBody) SetRequestId(v string) *GetVirusScanGlobalConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVirusScanGlobalConfigResponseBody) SetUploadFileMaxSize(v int64) *GetVirusScanGlobalConfigResponseBody {
	s.UploadFileMaxSize = &v
	return s
}

func (s *GetVirusScanGlobalConfigResponseBody) SetUploadFileMaxSpeed(v int64) *GetVirusScanGlobalConfigResponseBody {
	s.UploadFileMaxSpeed = &v
	return s
}

func (s *GetVirusScanGlobalConfigResponseBody) SetUploadFileSuffixBlacklist(v []*string) *GetVirusScanGlobalConfigResponseBody {
	s.UploadFileSuffixBlacklist = v
	return s
}

func (s *GetVirusScanGlobalConfigResponseBody) SetVirusFileUpload(v bool) *GetVirusScanGlobalConfigResponseBody {
	s.VirusFileUpload = &v
	return s
}

func (s *GetVirusScanGlobalConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

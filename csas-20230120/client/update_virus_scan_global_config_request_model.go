// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVirusScanGlobalConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUploadFileMaxSize(v int64) *UpdateVirusScanGlobalConfigRequest
	GetUploadFileMaxSize() *int64
	SetUploadFileMaxSpeed(v int64) *UpdateVirusScanGlobalConfigRequest
	GetUploadFileMaxSpeed() *int64
	SetUploadFileSuffixBlacklist(v []*string) *UpdateVirusScanGlobalConfigRequest
	GetUploadFileSuffixBlacklist() []*string
	SetVirusFileUpload(v bool) *UpdateVirusScanGlobalConfigRequest
	GetVirusFileUpload() *bool
}

type UpdateVirusScanGlobalConfigRequest struct {
	// The maximum size of a single virus file that can be uploaded. Unit: KB. Valid values: 0 to 204800. A value of 0 indicates no size limit. Values from 1 to 204800 specify the actual size limit.
	//
	// example:
	//
	// 10240
	UploadFileMaxSize *int64 `json:"UploadFileMaxSize,omitempty" xml:"UploadFileMaxSize,omitempty"`
	// The maximum upload rate for virus files. Unit: KB/s. Valid values: 0 to 102400. A value of 0 indicates no rate limit.
	//
	// example:
	//
	// 1024
	UploadFileMaxSpeed *int64 `json:"UploadFileMaxSpeed,omitempty" xml:"UploadFileMaxSpeed,omitempty"`
	// The collection of file types that are prohibited from being uploaded. Duplicate values are not allowed. Files that match the specified types are not uploaded even if upload is enabled.
	UploadFileSuffixBlacklist []*string `json:"UploadFileSuffixBlacklist,omitempty" xml:"UploadFileSuffixBlacklist,omitempty" type:"Repeated"`
	// Specifies whether user terminal devices are allowed to upload detected virus files to the cloud for further analysis. Valid values:
	//
	// - **true**: Allowed.
	//
	// - **false**: Not allowed.
	//
	// example:
	//
	// true
	VirusFileUpload *bool `json:"VirusFileUpload,omitempty" xml:"VirusFileUpload,omitempty"`
}

func (s UpdateVirusScanGlobalConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateVirusScanGlobalConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateVirusScanGlobalConfigRequest) GetUploadFileMaxSize() *int64 {
	return s.UploadFileMaxSize
}

func (s *UpdateVirusScanGlobalConfigRequest) GetUploadFileMaxSpeed() *int64 {
	return s.UploadFileMaxSpeed
}

func (s *UpdateVirusScanGlobalConfigRequest) GetUploadFileSuffixBlacklist() []*string {
	return s.UploadFileSuffixBlacklist
}

func (s *UpdateVirusScanGlobalConfigRequest) GetVirusFileUpload() *bool {
	return s.VirusFileUpload
}

func (s *UpdateVirusScanGlobalConfigRequest) SetUploadFileMaxSize(v int64) *UpdateVirusScanGlobalConfigRequest {
	s.UploadFileMaxSize = &v
	return s
}

func (s *UpdateVirusScanGlobalConfigRequest) SetUploadFileMaxSpeed(v int64) *UpdateVirusScanGlobalConfigRequest {
	s.UploadFileMaxSpeed = &v
	return s
}

func (s *UpdateVirusScanGlobalConfigRequest) SetUploadFileSuffixBlacklist(v []*string) *UpdateVirusScanGlobalConfigRequest {
	s.UploadFileSuffixBlacklist = v
	return s
}

func (s *UpdateVirusScanGlobalConfigRequest) SetVirusFileUpload(v bool) *UpdateVirusScanGlobalConfigRequest {
	s.VirusFileUpload = &v
	return s
}

func (s *UpdateVirusScanGlobalConfigRequest) Validate() error {
	return dara.Validate(s)
}

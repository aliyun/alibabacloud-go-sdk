// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestoreMediaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMediaIds(v string) *RestoreMediaRequest
	GetMediaIds() *string
	SetRestoreDays(v string) *RestoreMediaRequest
	GetRestoreDays() *string
	SetRestoreTier(v string) *RestoreMediaRequest
	GetRestoreTier() *string
	SetScope(v string) *RestoreMediaRequest
	GetScope() *string
}

type RestoreMediaRequest struct {
	// The media IDs, which are audio or video IDs (VideoId). Separate multiple IDs with commas (,). A maximum of 20 IDs are supported. You can obtain the IDs by using the following methods:
	//
	// - For audio or video files uploaded in the console, log on to the ApsaraVideo VOD console and choose Media Files > Audio/Video to view the audio or video ID.
	//
	// - When you call the CreateUploadVideo operation to obtain the upload URL and credential, the video ID is the value of the VideoId parameter in the response.
	//
	// - After the audio or video file is uploaded, you can call the SearchMedia operation to query the video ID, which is the value of the VideoId parameter in the response.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8bc8e94fe4e55abde85718****,eb186180e989dd56****
	MediaIds *string `json:"MediaIds,omitempty" xml:"MediaIds,omitempty"`
	// The restoration duration. Default value: 1 day. Maximum value for Archive media assets: 7 days. Maximum value for Cold Archive media assets: 365 days.
	//
	// example:
	//
	// 2
	RestoreDays *string `json:"RestoreDays,omitempty" xml:"RestoreDays,omitempty"`
	// The restoration priority. This parameter is required only for Cold Archive media assets. If this parameter is not specified, the default value **Standard*	- is used. Valid values:
	//
	// - **Expedited**: High priority. The restoration is completed within 1 hour.
	//
	// - **Standard*	- (default): Standard priority. The restoration is completed within 2 to 5 hours.
	//
	// - **Bulk**: Batch priority. The restoration is completed within 5 to 12 hours.
	//
	// example:
	//
	// Standard
	RestoreTier *string `json:"RestoreTier,omitempty" xml:"RestoreTier,omitempty"`
	// The scope of the change. If this parameter is not specified, the default value **All*	- is used. Valid values:
	//
	// - **All*	- (default): Applies tiered storage to all resources (source files and transcoded streams) of the media asset.
	//
	// - **SourceFile**: Applies tiered storage only to the video source file of the media asset ID. Resources other than the source file use Standard storage.
	//
	// example:
	//
	// All
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
}

func (s RestoreMediaRequest) String() string {
	return dara.Prettify(s)
}

func (s RestoreMediaRequest) GoString() string {
	return s.String()
}

func (s *RestoreMediaRequest) GetMediaIds() *string {
	return s.MediaIds
}

func (s *RestoreMediaRequest) GetRestoreDays() *string {
	return s.RestoreDays
}

func (s *RestoreMediaRequest) GetRestoreTier() *string {
	return s.RestoreTier
}

func (s *RestoreMediaRequest) GetScope() *string {
	return s.Scope
}

func (s *RestoreMediaRequest) SetMediaIds(v string) *RestoreMediaRequest {
	s.MediaIds = &v
	return s
}

func (s *RestoreMediaRequest) SetRestoreDays(v string) *RestoreMediaRequest {
	s.RestoreDays = &v
	return s
}

func (s *RestoreMediaRequest) SetRestoreTier(v string) *RestoreMediaRequest {
	s.RestoreTier = &v
	return s
}

func (s *RestoreMediaRequest) SetScope(v string) *RestoreMediaRequest {
	s.Scope = &v
	return s
}

func (s *RestoreMediaRequest) Validate() error {
	return dara.Validate(s)
}

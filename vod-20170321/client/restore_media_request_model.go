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
	// The ID of the media asset (VideoId). Separate multiple IDs with commas (,). You can specify a maximum of 20 IDs. You can use one of the following methods to obtain the ID of the media asset:
	//
	// 	- Log on to the ApsaraVideo VOD console. In the left-side navigation pane, choose Media Files > Audio/Video. On the Video and Audio page, view the ID of the media asset. This method is applicable to files that are uploaded by using the ApsaraVideo VOD console.
	//
	// 	- Obtain the value of VideoId from the response to the CreateUploadVideo operation that you call to upload media assets.
	//
	// 	- Obtain the value of VideoId from the response to the SearchMedia operation that you call to query the media ID after the media asset is uploaded.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8bc8e94fe4e55abde85718****,eb186180e989dd56****
	MediaIds *string `json:"MediaIds,omitempty" xml:"MediaIds,omitempty"`
	// The number of days during which media assets remain in the restored state. Default value: 1. The maximum validity period of a restored Archive media asset is 7 days and the maximum validity period of a restored Cold Archive media asset is 365 days.
	//
	// example:
	//
	// 2
	RestoreDays *string `json:"RestoreDays,omitempty" xml:"RestoreDays,omitempty"`
	// The restoration priority. This parameter is required only when you restore a Cold Archive media file. Valid values:
	//
	// 	- **Expedited**: The file is restored within 1 hour.
	//
	// 	- **Standard**: The file is restored within 2 to 5 hours.
	//
	// 	- **Bulk**: The file is restored within 5 to 12 hours.
	//
	// example:
	//
	// Standard
	RestoreTier *string `json:"RestoreTier,omitempty" xml:"RestoreTier,omitempty"`
	// The modification range. Valid values:
	//
	// 	- **All**: restores all resources, including the source files and transcoded streams.
	//
	// 	- **SourceFile**: restores only the source files.
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

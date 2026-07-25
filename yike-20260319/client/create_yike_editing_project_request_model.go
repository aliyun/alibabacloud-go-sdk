// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateYikeEditingProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCoverURL(v string) *CreateYikeEditingProjectRequest
	GetCoverURL() *string
	SetMaterialMaps(v string) *CreateYikeEditingProjectRequest
	GetMaterialMaps() *string
	SetTimeline(v string) *CreateYikeEditingProjectRequest
	GetTimeline() *string
	SetTitle(v string) *CreateYikeEditingProjectRequest
	GetTitle() *string
}

type CreateYikeEditingProjectRequest struct {
	// The cover URL of the cloud editing project.
	//
	// example:
	//
	// https://example.com/example.png
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// The materials associated with the project. Separate multiple materials with commas (,). A maximum of 10 material IDs are supported for each type.
	//
	// example:
	//
	// {"video":"*****2e057304fcd9b145c5cafc*****", "image":"****8021a8d493da643c8acd98*****,*****cb6307a4edea614d8b3f3c*****", "liveStream": "rtmp://domain.com/app/stream", "editingProject": "xxxxx"}
	MaterialMaps *string `json:"MaterialMaps,omitempty" xml:"MaterialMaps,omitempty"`
	// The timeline of the cloud editing project in JSON format.
	//
	// example:
	//
	// {"VideoTracks":[{"VideoTrackClips":[{"MediaId":"****4d7cf14dc7b83b0e801c****"},{"MediaId":"****4d7cf14dc7b83b0e801c****"}]}]}
	Timeline *string `json:"Timeline,omitempty" xml:"Timeline,omitempty"`
	// The title of the cloud editing project.
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CreateYikeEditingProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateYikeEditingProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateYikeEditingProjectRequest) GetCoverURL() *string {
	return s.CoverURL
}

func (s *CreateYikeEditingProjectRequest) GetMaterialMaps() *string {
	return s.MaterialMaps
}

func (s *CreateYikeEditingProjectRequest) GetTimeline() *string {
	return s.Timeline
}

func (s *CreateYikeEditingProjectRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateYikeEditingProjectRequest) SetCoverURL(v string) *CreateYikeEditingProjectRequest {
	s.CoverURL = &v
	return s
}

func (s *CreateYikeEditingProjectRequest) SetMaterialMaps(v string) *CreateYikeEditingProjectRequest {
	s.MaterialMaps = &v
	return s
}

func (s *CreateYikeEditingProjectRequest) SetTimeline(v string) *CreateYikeEditingProjectRequest {
	s.Timeline = &v
	return s
}

func (s *CreateYikeEditingProjectRequest) SetTitle(v string) *CreateYikeEditingProjectRequest {
	s.Title = &v
	return s
}

func (s *CreateYikeEditingProjectRequest) Validate() error {
	return dara.Validate(s)
}
